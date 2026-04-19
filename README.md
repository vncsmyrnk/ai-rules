# AI rules

This project provides a structured set of guidelines, principles, and workflows to generate standardized prompts for AI coding agents. Its goal is to improve the quality of AI-generated code, enforce best practices, and minimize hallucinations.

It provides an executable that, given a user prompt, it auto injects the necessary guidelines using a cosine similarity search to filter only the most relevant guidelines. The output should be the final input for the LLM call.

Output template:

```markdown
You are a software engineer. Complete the user's task while strictly adhering to the architectural rules provided below.

<architectural_rules>
{{ .Rules }}
</architectural_rules>

<user_task>
{{ .Task }}
</user_task>
```

## Structure

The repository is organized into four main categories, all utilizing YAML frontmatter tags for easy filtering and discovery:

- **`principles/`**: Detailed software engineering principles, design patterns, and architectural guidelines. Instructs agents on readability and performance trade-offs.
- **`workflows/`**: Specific task prompts and step-by-step instructions that agents can execute.
- **`languages/`**: Language-specific best practices, common code smells, battle-tested frameworks, and links to official documentation.

## Install

To be documented soon.
