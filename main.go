package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

var build = "1" // build number set at compile time

func main() {
	app := cli.NewApp()
	app.Name = "Drone-Sonar-Plugin"
	app.Usage = "Drone plugin to integrate with SonarQube."
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:   "key",
			Usage:  "project key",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "name",
			Usage:  "project name",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "host",
			Usage:  "SonarQube host",
			EnvVar: "PLUGIN_SONAR_HOST",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "SonarQube token",
			EnvVar: "PLUGIN_SONAR_TOKEN",
		},

		// advanced parameters
		cli.StringFlag{
			Name:   "branch",
			Usage:  "Branch name",
			Value:  "",
			EnvVar: "PLUGIN_BRANCH_NAME",
		},
		cli.StringFlag{
			Name:   "target",
			Usage:  "Target branch name",
			Value:  "",
			EnvVar: "PLUGIN_BRANCH_TARGET",
		},
		cli.StringFlag{
			Name:   "ver",
			Usage:  "Project version",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "timeout",
			Usage:  "Web request timeout",
			Value:  "60",
			EnvVar: "PLUGIN_TIMEOUT",
		},
		cli.StringFlag{
			Name:   "sources",
			Usage:  "analysis sources",
			Value:  ".",
			EnvVar: "PLUGIN_SOURCES",
		},
		cli.StringFlag{
			Name:   "inclusions",
			Usage:  "code inclusions",
			EnvVar: "PLUGIN_INCLUSIONS",
		},
		cli.StringFlag{
			Name:   "exclusions",
			Usage:  "code exclusions",
			EnvVar: "PLUGIN_EXCLUSIONS",
		},
		cli.StringFlag{
			Name:   "level",
			Usage:  "log level",
			Value:  "INFO",
			EnvVar: "PLUGIN_LEVEL",
		},
		cli.StringFlag{
			Name:   "showProfiling",
			Usage:  "showProfiling during analysis",
			Value:  "false",
			EnvVar: "PLUGIN_SHOWPROFILING",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) {
	plugin := Plugin{
		Config: Config{
			Key:   c.String("key"),
			Name:  c.String("name"),
			Host:  c.String("host"),
			Token: c.String("token"),

			Branch:        c.String("branch"),
			Target:        c.String("target"),
			Version:       c.String("ver"),
			Timeout:       c.String("timeout"),
			Sources:       c.String("sources"),
			Inclusions:    c.String("inclusions"),
			Exclusions:    c.String("exclusions"),
			Level:         c.String("level"),
			showProfiling: c.String("showProfiling"),
		},
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
