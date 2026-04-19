package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/daulet/tokenizers"
	"github.com/yalue/onnxruntime_go"
)

const onnxSharedLibraryPath = "/usr/local/lib/libonnxruntime.so"
const modelPath = "/usr/local/share/pguide/all-MiniLM-L6-v2.onnx"
const tokenizerPath = "/usr/local/share/pguide/tokenizer.json"
const maxTokens = 512 // The maximum context window for MiniLM

// PrecomputedRule holds the data we will bake into our final binary
type PrecomputedRule struct {
	Filename  string    `json:"filename"`
	Content   string    `json:"content"`
	Embedding []float32 `json:"embedding"`
}

func main() {
	rulesDir := "."

	onnxruntime_go.SetSharedLibraryPath(onnxSharedLibraryPath)
	if err := onnxruntime_go.InitializeEnvironment(); err != nil {
		fmt.Printf("Fatal: Failed to init ONNX: %v\n", err)
		os.Exit(1)
	}
	defer onnxruntime_go.DestroyEnvironment()

	tk, err := tokenizers.FromFile(tokenizerPath)
	if err != nil {
		fmt.Printf("Fatal: Failed to load tokenizer: %v\n", err)
		os.Exit(1)
	}
	defer tk.Close()

	var rules []PrecomputedRule

	fmt.Println("Scanning directory:", rulesDir)
	err = filepath.WalkDir(rulesDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(d.Name(), ".md") || strings.HasPrefix(d.Name(), "prompt") {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("could not read %s: %v", path, err)
		}
		content := string(contentBytes)

		fmt.Printf("Embedding: %s...\n", d.Name())
		embedding, err := getEmbedding(content, tk)
		if err != nil {
			fmt.Printf("Warning: Failed to embed %s: %v\n", path, err)
			return nil
		}

		rules = append(rules, PrecomputedRule{
			Filename:  filepath.Base(path),
			Content:   content,
			Embedding: embedding,
		})

		return nil
	})

	if err != nil {
		fmt.Printf("Failed during directory walk: %v\n", err)
		os.Exit(1)
	}

	outputPath := filepath.Join(os.ExpandEnv("$HOME"), ".local", "share", "pguide")
	err = os.MkdirAll(outputPath, 0755)
	if err != nil {
		fmt.Printf("Failed to create output path: %v\n", err)
		os.Exit(1)
	}

	outputFilePath := filepath.Join(outputPath, "embedded_rules.json")
	file, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Failed to create output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(rules); err != nil {
		fmt.Printf("Failed to write JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nSuccess! Generated embeddings for %d files and saved to %s\n", len(rules), outputFilePath)
}

// getEmbedding handles tokenization, truncation, and ONNX execution
func getEmbedding(text string, tk *tokenizers.Tokenizer) ([]float32, error) {
	tokens, _ := tk.Encode(text, false)

	if len(tokens) > maxTokens {
		tokens = tokens[:maxTokens]
	}

	inputIDs := make([]int64, len(tokens))
	attentionMask := make([]int64, len(tokens))
	tokenTypeIDs := make([]int64, len(tokens))

	for i, t := range tokens {
		inputIDs[i] = int64(t)
		attentionMask[i] = 1
		tokenTypeIDs[i] = 0
	}

	tensorShape := []int64{1, int64(len(inputIDs))}

	inputTensor, _ := onnxruntime_go.NewTensor(tensorShape, inputIDs)
	defer inputTensor.Destroy()

	maskTensor, _ := onnxruntime_go.NewTensor(tensorShape, attentionMask)
	defer maskTensor.Destroy()

	typeTensor, _ := onnxruntime_go.NewTensor(tensorShape, tokenTypeIDs)
	defer typeTensor.Destroy()

	outputData := make([]float32, 384)
	outputTensor, _ := onnxruntime_go.NewTensor([]int64{1, 384}, outputData)
	defer outputTensor.Destroy()

	session, err := onnxruntime_go.NewAdvancedSession(
		modelPath,
		[]string{"input_ids", "attention_mask", "token_type_ids"},
		[]string{"sentence_embedding"},
		[]onnxruntime_go.ArbitraryTensor{inputTensor, maskTensor, typeTensor},
		[]onnxruntime_go.ArbitraryTensor{outputTensor},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Destroy()

	if err := session.Run(); err != nil {
		return nil, fmt.Errorf("model execution failed: %w", err)
	}

	return outputData, nil
}
