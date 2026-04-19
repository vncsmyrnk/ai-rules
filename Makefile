.PHONY: dependencies embeddings build

dependencies:
	curl -L --parallel \
		https://huggingface.co/onnx-models/all-MiniLM-L6-v2-onnx/resolve/main/model.onnx -o /tmp/all-MiniLM-L6-v2.onnx \
		https://huggingface.co/onnx-models/all-MiniLM-L6-v2-onnx/resolve/main/tokenizer.json -o /tmp/tokenizer.json \
		https://github.com/daulet/tokenizers/releases/download/v1.27.0/libtokenizers.linux-amd64.tar.gz -o /tmp/libtokenizers.linux-amd64.tar.gz \
		https://github.com/microsoft/onnxruntime/releases/download/v1.24.4/onnxruntime-linux-x64-1.24.4.tgz -o /tmp/onnxruntime-linux-x64-1.24.4.tgz
	mkdir -p /tmp/libtokenizers; tar -xzvf /tmp/libtokenizers.linux-amd64.tar.gz -C /tmp/libtokenizers
	mkdir -p /tmp/onnxruntime; tar -xzvf /tmp/onnxruntime-linux-x64-1.24.4.tgz -C /tmp/onnxruntime
	sudo install -Dm655 /tmp/all-MiniLM-L6-v2.onnx /usr/local/share/pguide/all-MiniLM-L6-v2.onnx
	sudo install -Dm655 /tmp/tokenizer.json /usr/local/share/pguide/tokenizer.json
	sudo install -Dm655 /tmp/libtokenizers/libtokenizers.a /usr/local/lib/libtokenizers.a
	sudo install -Dm655 /tmp/onnxruntime/onnxruntime-linux-x64-1.24.4/lib/libonnxruntime.so /usr/local/lib/libonnxruntime.so

embeddings:
	CGO_ENABLED=1 go run ./cmd/embeddings/main.go
	@cp "$$HOME/.local/share/pguide/embedded_rules.json" ./cmd/pguide/

build:
	CGO_ENABLED=1 go build -o ./dist/pguide ./cmd/pguide/main.go
