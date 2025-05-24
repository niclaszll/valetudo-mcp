package tools

import (
	"fmt"
	"valetudo-mcp/client"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

var BasicControlDefinition = ToolDefinition[BasicControlArgs]{
	Name:        "basic_control",
	Description: "Control basic robot functions (start, stop, pause, home)",
}

type BasicControlArgs struct {
	Action string `json:"action" jsonschema:"required,description=Action to perform (start, stop, pause, home)"`
}

func BasicControl(valetudoClient *client.ValetudoClient, args BasicControlArgs) (*mcp_golang.ToolResponse, error) {
	validActions := map[string]bool{
		"start": true,
		"stop":  true,
		"pause": true,
		"home":  true,
	}

	if !validActions[args.Action] {
		return nil, fmt.Errorf("invalid action: %s. Valid actions are: start, stop, pause, home", args.Action)
	}

	payload := map[string]string{
		"action": args.Action,
	}

	responseBody, err := valetudoClient.MakeRequest("PUT", "/api/v2/robot/capabilities/BasicControlCapability", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to execute basic control action: %v", err)
	}

	var response string
	if string(responseBody) == "OK" {
		response = fmt.Sprintf("Successfully executed %s action", args.Action)
	} else {
		response = fmt.Sprintf("Unexpected response: %s", string(responseBody))
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(response)), nil
}
