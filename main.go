package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	version    = "0.1.0"
	apiURL     = "https://www.wolframalpha.com/api/v1/llm-api"
	timeoutSec = 30
)

const (
	exitSuccess      = 0
	exitGeneral      = 1
	exitAuth         = 2
	exitNetwork      = 3
	exitInvalidInput = 4
	exitRateLimit    = 5
)

func main() {
	var (
		markdown    bool
		maxChars    int
		showHelp    bool
		showVersion bool
		showExamples bool
	)

	flag.BoolVar(&markdown, "markdown", false, "rich markdown formatting")
	flag.BoolVar(&markdown, "m", false, "rich markdown formatting (shorthand)")
	flag.IntVar(&maxChars, "max-chars", 0, "response character limit")
	flag.BoolVar(&showHelp, "help", false, "show usage")
	flag.BoolVar(&showHelp, "h", false, "show usage (shorthand)")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version (shorthand)")
	flag.BoolVar(&showExamples, "examples", false, "show examples")
	flag.BoolVar(&showExamples, "e", false, "show examples (shorthand)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: walpha [flags] \"query\"\n\n")
		fmt.Fprintf(os.Stderr, "Query Wolfram Alpha LLM API\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		fmt.Fprintf(os.Stderr, "  -m, --markdown     rich markdown formatting\n")
		fmt.Fprintf(os.Stderr, "  --max-chars N      response character limit (default: 6800)\n")
		fmt.Fprintf(os.Stderr, "  -e, --examples     show usage examples\n")
		fmt.Fprintf(os.Stderr, "  -h, --help         show this help\n")
		fmt.Fprintf(os.Stderr, "  -v, --version      show version\n\n")
		fmt.Fprintf(os.Stderr, "Environment:\n")
		fmt.Fprintf(os.Stderr, "  WOLFRAM_APP_ID     API key (required)\n")
	}

	flag.Parse()

	if showVersion {
		fmt.Printf("walpha %s\n", version)
		os.Exit(exitSuccess)
	}

	if showHelp {
		flag.Usage()
		os.Exit(exitSuccess)
	}

	if showExamples {
		printExamples()
		os.Exit(exitSuccess)
	}

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: query required")
		flag.Usage()
		os.Exit(exitGeneral)
	}

	query := flag.Arg(0)
	if strings.Contains(query, "\n") {
		fmt.Fprintln(os.Stderr, "error: multi-line queries not supported")
		os.Exit(exitGeneral)
	}

	appID := os.Getenv("WOLFRAM_APP_ID")
	if appID == "" {
		fmt.Fprintln(os.Stderr, "error: WOLFRAM_APP_ID not set")
		os.Exit(exitAuth)
	}

	result, exitCode := doQuery(appID, query, maxChars)
	if exitCode != exitSuccess {
		os.Exit(exitCode)
	}

	output := stripWebsiteLink(result)
	if markdown {
		output = formatMarkdown(output)
	}

	fmt.Print(output)
}

func doQuery(appID, query string, maxChars int) (string, int) {
	reqURL := fmt.Sprintf("%s?input=%s", apiURL, url.QueryEscape(query))
	if maxChars > 0 {
		reqURL += fmt.Sprintf("&maxchars=%d", maxChars)
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return "", exitGeneral
	}
	req.Header.Set("Authorization", "Bearer "+appID)

	client := &http.Client{Timeout: timeoutSec * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: network failure\n")
		return "", exitNetwork
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: reading response\n")
		return "", exitNetwork
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return string(body), exitSuccess
	case http.StatusForbidden:
		fmt.Fprintln(os.Stderr, "error: invalid app ID")
		return "", exitAuth
	case http.StatusTooManyRequests:
		fmt.Fprintln(os.Stderr, "error: rate limited, try again later")
		return "", exitRateLimit
	case 501:
		fmt.Fprintln(os.Stderr, "error: query not understood")
		return "", exitInvalidInput
	case http.StatusBadRequest:
		fmt.Fprintln(os.Stderr, "error: bad request")
		return "", exitInvalidInput
	default:
		fmt.Fprintf(os.Stderr, "error: HTTP %d\n", resp.StatusCode)
		return "", exitGeneral
	}
}

func stripWebsiteLink(s string) string {
	lines := strings.Split(s, "\n")
	var result []string
	for _, line := range lines {
		if strings.HasPrefix(line, "Wolfram|Alpha website result") ||
			strings.HasPrefix(line, "https://www.wolframalpha.com/input") {
			continue
		}
		result = append(result, line)
	}
	return strings.TrimRight(strings.Join(result, "\n"), "\n") + "\n"
}

func formatMarkdown(s string) string {
	lines := strings.Split(s, "\n")
	var result []string

	// Section headers: short, no special chars like = / ( )
	sectionRe := regexp.MustCompile(`^([A-Z][A-Za-z\s]+):$`)
	imageRe := regexp.MustCompile(`^image:\s*(https?://\S+)`)

	for _, line := range lines {
		if line == "" {
			result = append(result, line)
			continue
		}

		if matches := sectionRe.FindStringSubmatch(line); matches != nil {
			result = append(result, fmt.Sprintf("## %s", matches[1]))
			continue
		}

		if matches := imageRe.FindStringSubmatch(line); matches != nil {
			result = append(result, fmt.Sprintf("![result](%s)", matches[1]))
			continue
		}

		// Bold key-value patterns like "key | value"
		if strings.Contains(line, " | ") {
			parts := strings.SplitN(line, " | ", 2)
			if len(parts) == 2 && strings.TrimSpace(parts[0]) != "" {
				result = append(result, fmt.Sprintf("**%s** | %s", strings.TrimSpace(parts[0]), parts[1]))
				continue
			}
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

func printExamples() {
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: cannot find executable path")
		os.Exit(exitGeneral)
	}

	exeDir := filepath.Dir(exe)
	paths := []string{
		filepath.Join(exeDir, "docs", "examples.md"),
		filepath.Join(".", "docs", "examples.md"),
	}

	for _, path := range paths {
		content, err := os.ReadFile(path)
		if err == nil {
			fmt.Print(string(content))
			return
		}
	}

	fmt.Fprintln(os.Stderr, "error: examples.md not found")
	os.Exit(exitGeneral)
}
