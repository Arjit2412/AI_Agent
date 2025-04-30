package tools

// import (
// 	"fmt"
// 	// "os/exec"
// 	"os"
// 	"context"

// 	"google.golang.org/genai"
// )

// var GenerateCommitMessageInput = &genai.Schema{
// 	Type: genai.TypeObject,
// 	Properties: map[string]*genai.Schema{
// 		"message": {
// 			Type:        genai.TypeString,
// 			Description: "Message that has to be given as commit message",
// 		},
// 	},
// }

// var GenerateMessageDefination = &genai.FunctionDeclaration{
// 	Name:	"generateCommitMessage",
// 	Description: "Whenever the code has to do commit returns a relevant message",
// 	Parameters:  GenerateCommitMessageInput,
// }

// // func GenerateCommitMessage(input *genai.FunctionCall) (string, error) {
// 	// cmd := exec.Command("git", "diff", "--cached")
// 	// output, err := cmd.CombinedOutput()
	
// 	// if err != nil {
// 	// 	return "", fmt.Errorf("failed to get diff: %v", err)
// 	// }

// 	// fmt.Print(output)
// 	// // Return the diff back to Gemini so it can suggest a commit message
// 	// return string(output), nil

// // }


