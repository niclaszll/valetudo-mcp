# Valetudo MCP

A Go-based MCP server for interacting with Valetudo-powered robot vacuums via the Valetudo HTTP API. The server provides a set of tools to query robot state, info, capabilities, and control log level, exposing them as MCP tools.

## Features

- Get the current state of the robot (attributes, map)
- Get basic robot information (manufacturer, model, implementation)
- Get the list of capabilities supported by the robot
- Set the Valetudo log level (trace, debug, info, warn, error)
- Control basic robot functions (start, stop, pause, home)
- Locate the robot (plays a sound to help find it)
- Zone cleaning (clean specific areas of your home)

## Requirements

- Go 1.24+
- A robot vacuum running [Valetudo](https://valetudo.cloud/) (with accessible HTTP API)

## Setup

1. Clone this repository:
   ```sh
   git clone https://github.com/niclaszll/valetudo-mcp.git
   cd valetudo-mcp
   ```
2. Set the `VALETUDO_URL` environment variable to your Valetudo instance base URL:
   ```sh
   export VALETUDO_URL="http://<valetudo-ip>:<port>"
   ```
3. Build and run:
   ```sh
   go build -o valetudo-mcp
   ./valetudo-mcp
   ```
4. Integrate it with your MCP compatible hosts of choice, like Claude Desktop or Cursor.

## Usage

The MCP server will start and register the following tools:

- `get_robot_state`: Get the current state of the robot
- `get_robot_info`: Get basic robot information
- `get_robot_capabilities`: Get the list of capabilities supported by the robot
- `set_log_level`: Set the Valetudo log level
- `basic_control`: Control basic robot functions
- `locate`: Play a sound to help locate the robot
- `zone_clean`: Clean specific zones in your home

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Attribution

This project uses the [Valetudo](https://valetudo.cloud/) API exposed by your robot vacuum. Valetudo is an open-source project that enables local control of robot vacuums. All credit for the API and its development goes to the Valetudo contributors.
