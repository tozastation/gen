package main

import (
	"errors"
	"fmt"
	"github.com/tozastation/gogoenv/Initialize"
	"github.com/tozastation/gogoenv/generics"
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	// N0_ARGUMENTS_COMMAND is has not arguments your command
	N0ArgumentsCommand = ""
	InitializeCommand = "init"
	CreateServiceCommand = "create"
	CheckGoPath = "gopath"
)

func main() {
	app := cli.NewApp()
	app.Name = "gogoenv"
	app.Usage = "gogoenv is Create Boiler Template"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		command := c.Args().Get(0)
		switch command {
		case N0ArgumentsCommand:
			return errors.New("[Enough Arguments]: please type one argument at least")
		case InitializeCommand:
			subCommand := c.Args().Get(1)
			if subCommand == "" {
				return errors.New("[Enough Arguments]: init + {Your Application Name}")
			}
			Initialize.Initialize(subCommand)
			return nil
		case CreateServiceCommand:
			domainName := c.Args().Get(1)
			return generics.GenerateGenerics(domainName)
		case CheckGoPath:
			fmt.Printf("Your GOPATH is %s", os.Getenv("GOPATH"))
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
