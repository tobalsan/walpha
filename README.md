# walpha

CLI tool to query [Wolfram Alpha LLM API](https://products.wolframalpha.com/llm-api/documentation).

## Setup

1. Get an API key from [Wolfram Alpha Developer Portal](https://developer.wolframalpha.com/)

2. Set environment variable:
```bash
export WOLFRAM_APP_ID="your-app-id"
```

3. Install:
```bash
# From source
git clone https://github.com/thinhle19/walpha && cd walpha
make link

# Or via go install
go install github.com/thinhle19/walpha@latest
```

## Usage

```bash
walpha "query"
walpha --markdown "query"    # Rich markdown formatting
walpha --max-chars 500 "query"
walpha --examples            # Show usage examples
```

## Examples

```bash
walpha "integrate x^2 sin(x)"
walpha "150 lbs to kg"
walpha "speed of light"
walpha "population of France"
walpha --markdown "10 densest metals"
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Auth error (missing/invalid API key) |
| 3 | Network error |
| 4 | Invalid input |
| 5 | Rate limited |
