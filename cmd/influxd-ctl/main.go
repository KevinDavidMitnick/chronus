package main

import (
	"github.com/KevinDavidMitnick/chronus/cmd/influxd-ctl/command"
)

func main() {
	command := command.NewCommand()
	command.Execute()
}
