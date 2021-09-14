# Loadout

A hobby Golang project that is a primitive loadtesting tool.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Example Usage](#example usage)

## Installation

Grab the loadout binary from the bin/ directory, or compile it from source yourself.

## Usage

Usage:
```
  loadout [command]

Available Commands:

  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  send        The command that fires of requests.

Flags:
      --config string   config file (default is $HOME/.loadout.yaml)
  -h, --help            help for loadout
  -t, --toggle          Help message for toggle

Use "loadout [command] --help" for more information about a command.
```
## Example Usage
```
./loadout send --count 10 --target www.google.com --uri / --port 443
```
