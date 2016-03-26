package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/denkhaus/irspamd/command"
	"github.com/denkhaus/tcgl/applog"
)

func main() {
	app := cli.NewApp()
	app.Name = "irspamd"
	app.Version = "1.0 by pbatkov"
	app.Usage = "A command line app that scans your IMAP mail for spam."
	app.Flags = []cli.Flag{
                        cli.StringFlag{
                                Name: "host, H",
                                Value: "imap.cm.hc.ru",
                                Usage: "Host to connect to.",
                        },
                        cli.StringFlag{
                                Name: "port, P",
                                Value: "993",
                                Usage: "Port number to connect to.",
                        },
                        cli.StringFlag{
                                Name: "user, u",
                                Value: "",
                                Usage: "Your username at host",
                        },
			cli.StringFlag{
                                Name: "pass, p",
                                Value: "",
                                Usage: "Your IMAP password. For security reasons prefer IMAP_PASSWORD='yourpassword'",
				EnvVar: "IMAP_PASSWORD",
                        },
                        cli.BoolFlag{
                                Name: "reset, r",
                                Usage: "Clear database before run",
                        },
	}

	if cmdr, err := command.NewCommander(app); err != nil {
		applog.Errorf("Startup error:: %s", err.Error())
		return
	} else {
		cmdr.Run(os.Args)
	}
}
