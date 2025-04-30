package tools

import (
	
	"fmt"
	"os/exec"
	

	"google.golang.org/genai"
)


var CommintChangesInput = &genai.Schema{
	Type: genai.TypeObject,
	Properties: map[string]*genai.Schema{
		"message": {
			Type:        genai.TypeString,
			Description: "message while committing",
		},
	},
}

var CommitChangesDefination = &genai.FunctionDeclaration{
	Name:	"commitChanges",
	Description: "Commits all the staged files with relevant message. If no message is provided try to figure what changes are done and based on that provide message to commit",
	Parameters:  CommintChangesInput,
}




func CommitChanges(input *genai.FunctionCall) (string, error) {
	message, ok := input.Args["message"].(string)

	// If no message provided, auto-generate it
	if !ok || message == "" {
		// Step 1: Get staged diff
		diffCmd := exec.Command("git", "diff", "--cached")
		diffOutput, err := diffCmd.CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("failed to get diff: %v", err)
		}

		// Step 2: Call genai to summarize changes into a commit message
		diffText := string(diffOutput)
		if diffText == "" {
			return "", fmt.Errorf("no staged changes found to commit")
		}

		// Step 3: Generate a commit message using genai
		prompt := fmt.Sprintf("Write a clear, concise Git commit message summarizing the following changes:\n\n%s", diffText)
		functionCall := &genai.FunctionCall{
			Args: map[string]interface{}{
				"message": prompt,
			},
		}
		resp, err := GenerateCommitMessage(functionCall) // You must define this or use genai SDK accordingly
		if err != nil {
			return "", fmt.Errorf("AI generation failed: %v", err)
		}

		message = resp // or resp.Choices[0].Message.Content depending on your SDK
	}

	// Step 4: Run the git commit command
	cmd := exec.Command("git", "commit", "-m", message)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error committing: %v\n%s", err, string(output))
	}

	return string(output), nil
}





