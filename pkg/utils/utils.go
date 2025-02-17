package utils

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// RemoveCodeFences removes ```html from the start and ``` from the end of the input string.
func LintCodeFences(input *string, language string) *string {
	codeFenceStart := fmt.Sprintf("```%v", language)
	const codeFenceEnd = "```"

	// Trim the starting "```html"
	*input = strings.TrimPrefix(*input, codeFenceStart)

	// Trim any leading/trailing whitespace or newlines to better detect the ending code fence
	*input = strings.TrimSpace(*input)

	// Trim the ending "```"
	*input = strings.TrimSuffix(*input, codeFenceEnd)

	// Trim excess whitespace again
	trimmedInput := strings.TrimSpace(*input)

	return &trimmedInput
}

// WatchAndBuildCSS runs the build once and watches CSS changes concurrently
func WatchAndBuildCSS() {
	RunBuildCSS()

	// Run watch command in a goroutine so it doesn't block
	go RunWatchCSS()

	// Keep the main function alive
	select {}
}

func RunBuildCSS() {
	fmt.Println("Building CSS...")
	cmd := exec.Command("npm", "run", "build")

	// Pipe output to terminal
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	fmt.Println("Building CSS... ")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Build failed: %v", err)
	}
}

func RunWatchCSS() {
	cmd := exec.Command("npm", "run", "watch")

	// Pipe output to terminal
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	fmt.Println("Starting CSS watch ")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Watch failed: %v", err)
	}
}

// hashFile computes the SHA-256 hash of a file
func HashFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(content)
	return fmt.Sprintf("%x", hash), nil
}
