package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mclellac/tripsix/scanner"
	"github.com/urfave/cli"
	"golang.org/x/sync/semaphore"
)

func Run() {
	app := cli.NewApp()
	app.Name = "tripsix"
	app.Usage = "Quick, simple & dirty port scanner."
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")

		p := &scanner.PortScanner{
			IP:   "127.0.0.1",
			Lock: semaphore.NewWeighted(scanner.Ulimit()),
		}
		p.Start(1, 65535, 500*time.Millisecond)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
