package tools

import (
	"fmt"
	"os/exec"

	"google.golang.org/genai"
)

var CheckoutInput = &genai.Schema{
	Type: genai.TypeObject,
	Properties: map[string]*genai.Schema{
		"name": {
			Type:        genai.TypeString,
			Description: "The branch name or the new feature name where git has to checkout",
		},
	},
	Required: []string{"name"},
}

var CheckoutDefination = &genai.FunctionDeclaration{
	Name:        "checkout",
	Description: "Used to create new features/branches.Creates a new branch if no branch exists with given name or Switch to a Branch with given name. If no name is provided it will give an error",
	Parameters:  CheckoutInput,
}

func Checkout(input *genai.FunctionCall) (string, error) {
	name, ok := input.Args["name"].(string)
	if !ok || name == "" {
		return "Error! Check the name provided", fmt.Errorf("error : %v", ok)
	}

	cmd := exec.Command("git", "checkout -b", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	return string(output), nil
}
