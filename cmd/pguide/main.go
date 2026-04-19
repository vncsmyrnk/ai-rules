package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/daulet/tokenizers"
	"github.com/yalue/onnxruntime_go"
)

//go:embed embedded_rules.json
var embeddedRulesJSON []byte

//go:embed prompt.md
var promptTemplateText string

const onnxSharedLibraryPath = "/usr/local/lib/libonnxruntime.so"
const modelPath = "/usr/local/share/pguide/all-MiniLM-L6-v2.onnx"
const tokenizerPath = "/usr/local/share/pguide/tokenizer.json"
const maxTokens = 512

type PrecomputedRule struct {
	Filename  string    `json:"filename"`
	Content   string    `json:"content"`
	Embedding []float32 `json:"embedding"`
	Score     float32   `json:"-"`
}

type PromptData struct {
	Task  string
	Rules string
}

func main() {
	taskPtr := flag.String("task", "", "The user task and search query")
	flag.Parse()

	if *taskPtr == "" {
		fmt.Fprintln(os.Stderr, "Error: Please provide a task using the -task flag.")
		os.Exit(1)
	}

	var rules []PrecomputedRule
	if err := json.Unmarshal(embeddedRulesJSON, &rules); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: failed to load embedded rules: %v\n", err)
		os.Exit(1)
	}

	onnxruntime_go.SetSharedLibraryPath(onnxSharedLibraryPath)
	if err := onnxruntime_go.InitializeEnvironment(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: Failed to init ONNX: %v\n", err)
		os.Exit(1)
	}
	defer onnxruntime_go.DestroyEnvironment()

	tk, err := tokenizers.FromFile(tokenizerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: Failed to load tokenizer: %v\n", err)
		os.Exit(1)
	}
	defer tk.Close()

	queryEmbedding, err := getEmbedding(*taskPtr, tk)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to embed task: %v\n", err)
		os.Exit(1)
	}

	for i := range rules {
		rules[i].Score = calculateCosineSimilarity(queryEmbedding, rules[i].Embedding)
	}

	sort.Slice(rules, func(i, j int) bool {
		return rules[i].Score > rules[j].Score
	})

	var ruleContents strings.Builder
	for i, rule := range rules {
		if i >= 2 { // Take only the top 2 matches to save context window
			break
		}
		fmt.Fprint(&ruleContents, "\n")
		ruleContents.WriteString(rule.Content)
		ruleContents.WriteString("\n\n")
	}

	tmpl, err := template.New("prompt").Parse(promptTemplateText)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: failed to parse template: %v\n", err)
		os.Exit(1)
	}

	data := PromptData{
		Task:  *taskPtr,
		Rules: ruleContents.String(),
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: failed to execute template: %v\n", err)
		os.Exit(1)
	}
}

// getEmbedding handles tokenization and ONNX execution for the user prompt
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

	outputData := make([]float32, 384) // MiniLM produces 384-dimensional vectors
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

// calculateCosineSimilarity measures the distance between two vectors
func calculateCosineSimilarity(a, b []float32) float32 {
	var dotProduct, normA, normB float32
	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	if normA == 0 || normB == 0 {
		return 0.0
	}
	return dotProduct / float32(math.Sqrt(float64(normA))*math.Sqrt(float64(normB)))
}
