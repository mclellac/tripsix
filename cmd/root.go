package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mclellac/tripsix/scanner"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/semaphore"
)

// Run get this party started...
func Run() {
	tasks := []string{"scan", "ports", "help", "localhost", "127.0.0.1"}

	//app := cli.NewApp()
	app.Name = "tripsix"
	app.Usage = "Quick, simple & dirty port scanner."
	app.Version = "0.0.1"
	app.EnableBashCompletion = true


	app := &cli.App{
		Commands: []*cli.Command{
		{
			Name:        "scan",
			Usage:       "Scan a host",
			Description: "Scans an IP or Hostname for open ports\n\nEXAMPLE:\n   $ tripsix scan 127.0.0.1",
			ArgsUsage:   "TARGET",
			Action: func(c *cli.Context) error {
				if len(c.Args()) != 1 {
					fmt.Println("You might want to double check your command there, chummer.")
					return nil
				}

				ip := c.Args().Get(0)

				p := &scanner.PortScanner{
					IP:   ip,
					Lock: semaphore.NewWeighted(scanner.Ulimit()),
				}

				p.Start(1, 65535, 500*time.Millisecond)

				return nil
			},
			BashComplete: func(c *cli.Context) {
				// This will complete if no args are passed
				if c.NArg() > 0 {
					return
				}
				for _, t := range tasks {
					fmt.Println(t)
				}
			},
		},
		{
			Name:        "ports",
			Usage:       "Port range",
			Description: "The range of ports (from first to last) to scan\n\nEXAMPLE\n$ tripsix scan localhost range 1 65535",
			//			ArgsUsage:   "[\"First Port\" \"Last port\"]",
			//			Action: func(c *cli.Context) error {
			//				if len(c.Args()) != 2 {
			//					fmt.Println("You'll need to specify a start and end to this range, cowboy.")
			//					return nil
			//				}
			//
			//				// TODO: Complete port range
			//				first := c.Args().Get(0)
			//				Last := c.Args().Get(1)
			//
			//				p := &scanner.PortScanner{
			//					IP:   ip,
			//					Lock: semaphore.NewWeighted(scanner.Ulimit()),
			//				}
			//
			//				p.Start(1, 65535, 500*time.Millisecond)
			//				return nil
			//			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
