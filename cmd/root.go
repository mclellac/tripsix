package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mclellac/tripsix/scanner"
	"github.com/urfave/cli"
)

func Run() {
	app := cli.NewApp()
	app.Name = "tripsix"
	app.Usage = "Quick and simple port scanner."
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		scanner.Start()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
