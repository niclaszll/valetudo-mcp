package tools

import (
	"fmt"
	"valetudo-mcp/client"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

var SetLogLevelDefinition = ToolDefinition{
	Name:        "set_log_level",
	Description: "Set the Valetudo log level",
}

type SetLogLevelArgs struct {
	Level string `json:"level" jsonschema:"required,description=Log level (trace, debug, info, warn, error)"`
}

func SetLogLevel(valetudoClient *client.ValetudoClient, args SetLogLevelArgs) (*mcp_golang.ToolResponse, error) {
	validLevels := map[string]bool{
		"trace": true,
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}

	if !validLevels[args.Level] {
		return nil, fmt.Errorf("invalid log level: %s. Valid levels are: trace, debug, info, warn, error", args.Level)
	}

	payload := map[string]string{
		"level": args.Level,
	}

	responseBody, err := valetudoClient.MakeRequest("PUT", "/api/v2/valetudo/log/level", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to set log level: %v", err)
	}

	var response string
	if string(responseBody) == "OK" || len(responseBody) == 0 {
		response = fmt.Sprintf("Log level successfully set to: %s", args.Level)
	} else {
		response = fmt.Sprintf("Unexpected response: %s", string(responseBody))
	}

	return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(response)), nil
}
