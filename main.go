package main

import (
	"fmt"
	"github.com/paramah/ledo/app/cmd"
	"github.com/paramah/ledo/app/modules/compose"
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	version = "dev"
)

func GetCurrentVersion() string {
	return version
}

func main() {
	app := cli.NewApp()
	compose.CheckDockerComposeVersion()
	app.Name = "ledo"
	app.Usage = "docker-compose and docker workflow improvement tool"
	app.Description = appDescription
	app.CustomAppHelpTemplate = helpTemplate
	app.Version = GetCurrentVersion()
	app.Commands = []*cli.Command{
		&cmd.CmdInit,
		&cmd.CmdDocker,
		&cmd.CmdImage,
		&cmd.CmdMode,
		&cmd.CmdAutocomplete,
		// &CmdSelfupdate,
	}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		// app.Run already exits for errors implementing ErrorCoder,
		// so we only handle generic errors with code 1 here.
		fmt.Fprintf(app.ErrWriter, "Error: %v\n", err)
		os.Exit(1)
	}
}

var appDescription = `LeadDocker (ledo) is a simple tool for improve docker and docker-compose workflow in your project.
What you can do with this tool:
=> create and manage docker-compose workflow in a project
=> build docker image for project (automatic fqn and docker registry)
=> login to the docker registry (with AWS ECR support)

LeadDocker is helpful in a continuous methodologies. 
If you want use it as a docker service, try dind image: https://hub.docker.com/r/paramah/dind

Enjoy (-_-)
`

var helpTemplate = bold(`
 _                _ ___          _
| |   ___ __ _ __| |   \ ___  __| |_____ _ _
| |__/ -_) _' / _' | |) / _ \/ _| / / -_) '_|
|____\___\__,_\__,_|___/\___/\__|_\_\___|_|    {{.Version}} 

{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}`) + `

 USAGE
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .Commands}} command [subcommand] [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Description}}

 DESCRIPTION
   {{.Description | nindent 3 | trim}}{{end}}{{if .VisibleCommands}}

 COMMANDS{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
   {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

 OPTIONS
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}

 EXAMPLES
   ledo init                       # init ledo in your project
   ledo docker ps                  # print list of docker containers

  
 ABOUT
   Written & maintained by Aleksander "paramah" Cynarski

   More info about ledo on https://leaddocker.tech
   `+bold(`
   Thanks for:
    StreamSage Team        https://streamsage.io
    Jazzy Innovations Team https://jazzy.pro
`)+"\n"

func bold(t string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", t)
}

var CmdSelfupdate = cli.Command{
	Name: "selfupdate",
	Aliases: []string{"update"},
	Category: "SETUP",
	Usage: "Self update Ledo",
	Action: runSelfupdate,
}

var updater = &selfupdate.Updater{
	CurrentVersion: GetCurrentVersion(),
	ApiURL:         "http://updates.yourdomain.com/",
	BinURL:         "http://updates.yourdomain.com/",
	DiffURL:        "http://updates.yourdomain.com/",
	Dir:            "update/",
	CmdName:        "ledo",
}

func runSelfupdate(ctx *cli.Context) error {
	log.Printf("check and update Ledo binary")
	updater.BackgroundRun()

	return nil
}