package main

import (
	"github.com/dfuse-io/dfuse-eosio/cmd/dfuseeos/cli"

	_ "github.com/streamingfast/dauth/ratelimiter/null"
)

var version = "dev"
var commit = ""

func init() {
	cli.RootCmd.Version = version + "-" + commit
}

func main() {
	cli.Main()
}
