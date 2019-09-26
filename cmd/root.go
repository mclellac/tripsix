package cmd

import (
	"fmt"
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
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:        "scan",
			Usage:       "Scan a host",
			Description: "Scans an IP or Hostname for open ports\n\nEXAMPLE:\n   $ tripsix scan 127.0.0.1",
			ArgsUsage:   "[\"IP\"] or [\"host.domain.tld\"]",
			Action: func(c *cli.Context) error {
				if len(c.Args()) != 1 {
					fmt.Println("You might want to double check your command there.")
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
		},
	}
	app.Run(os.Args)
}

// Run get this party started...
//func Run() {
//	tasks := []string{"scan", "ports", "help", "localhost", "127.0.0.1"}
//
//	app := cli.NewApp()
//	app.Name = "tripsix"
//	app.Usage = "Quick, simple & dirty port scanner."
//	app.Version = "0.0.1"
//	app.EnableBashCompletion = true
//
//	app := &cli.App{
//		Commands: []*cli.Command{
//			{
//				Name:        "scan",
//				Usage:       "Scan a host",
//				Description: "Scans an IP or Hostname for open ports\n\nEXAMPLE:\n$ tripsix scan 127.0.0.1",
//				ArgsUsage:   "TARGET",
//
//				Action: func(c *cli.Context) error {
//					if len(c.Args()) != 1 {
//						fmt.Println("You might want to double check your command there, chummer.")
//						return nil
//					}
//
//					ip := c.Args().Get(0)
//
//					p := &scanner.PortScanner{
//						IP:   ip,
//						Lock: semaphore.NewWeighted(scanner.Ulimit()),
//					}
//
//					p.Start(1, 65535, 500*time.Millisecond)
//
//					return nil
//				},
//				BashComplete: func(c *cli.Context) {
//					// This will complete if no args are passed
//					if c.NArg() > 0 {
//						return
//					}
//					for _, t := range tasks {
//						fmt.Println(t)
//					}
//				},
//			},
//		},
//	}
//	app.Run(os.Args)
//
//}
//
