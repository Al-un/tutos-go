package intermediate

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

// RunNetworkCli https://tutorialedge.net/golang/building-a-cli-in-go/
func RunNetworkCli() {
	// ===== v1
	// err := cli.NewApp().Run(os.Args)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ===== v2
	app := cli.NewApp()
	app.Name = "Website lookup CLI"
	app.Usage = "Learning how to CLI in go :)"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "url",
			Value: "tutorialedge.net",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the nameservers for a url",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				// a simple lookup function
				ns, err := net.LookupNS(c.String("url"))
				if err != nil {
					return err
				}
				// we log the results to our console
				// using a trusty fmt.Println statement
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular url",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("url"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		}, {
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular url",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("url"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cname)
				return nil
			},
		}, {
			Name:  "mx",
			Usage: "Looks up the MX records for a particular url",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("url"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}

	// start our application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
