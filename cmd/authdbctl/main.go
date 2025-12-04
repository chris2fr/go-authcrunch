// Copyright 2022 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/greenpau/versioned"

	"github.com/BurntSushi/toml"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	app        *versioned.PackageManager
	appVersion string
	gitBranch  string
	gitCommit  string
	buildUser  string
	buildDate  string
	sh         *cli.App
)

func init() {

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	bundle.MustLoadMessageFile("assets/locale/active.fr.toml")
	localizer := i18n.NewLocalizer(bundle, language.English.String(), language.French.String()) // Initialize localizer
	Messageappdescription := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "app-description",
			Other: "AuthDB management client",
		},
	})
	Messageconfig := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "config",
			Other: "Sets `PATH` to configuration file",
		},
	})
	Messagetokenpath := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "token-path",
			Other: "Sets `PATH` to token file",
		},
	})
	Messageformat := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "format",
			Other: "Sets `NAME` of the output format",
		},
	})
	Messagedebug := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "debug",
			Other: "Enabled debug logging",
		},
	})
	Messageconnect := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "connect",
			Other: "connect to auth portal and obtain access token",
		},
	})
	Messagemetadata := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "metadata",
			Other: "fetch metadata",
		},
	})
	Messageadd := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "add",
			Other: "add database objects",
		},
	})
	Messagelist := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "list",
			Other: "list database objects",
		},
	})
	// i18n.MustLoadTranslationFile("./locale/en-us.all.json")
	// T, _ := i18n.Tfunc("en-US")

	app = versioned.NewPackageManager("authdbctl")
	app.Description = Messageappdescription
	app.Documentation = "https://github.com/greenpau/go-authcrunch/"
	app.SetVersion(appVersion, "1.1.7")
	app.SetGitBranch(gitBranch, "main")
	app.SetGitCommit(gitCommit, "v1.1.6-1-g25b3ec7")
	app.SetBuildUser(buildUser, "")
	app.SetBuildDate(buildDate, "")

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(os.Stdout, "%s\n", app.Banner())
	}

	sh = cli.NewApp()
	sh.Name = app.Name
	sh.Version = app.Version
	sh.Usage = app.Description
	sh.Description = app.Documentation
	sh.HideHelp = false
	sh.HideVersion = false
	sh.Flags = append(sh.Flags, &cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Usage:       Messageconfig,
		Value:       `~/.config/authdbctl/config.yaml`,
		DefaultText: `~/.config/authdbctl/config.yaml`,
		EnvVars:     []string{"AUTHDBCTL_CONFIG_PATH"},
	})
	sh.Flags = append(sh.Flags, &cli.StringFlag{
		Name:        "token-path",
		Usage:       Messagetokenpath,
		Value:       `~/.config/authdbctl/token.jwt`,
		DefaultText: `~/.config/authdbctl/token.jwt`,
		EnvVars:     []string{"AUTHDBCTL_TOKEN_PATH"},
	})
	sh.Flags = append(sh.Flags, &cli.StringFlag{
		Name:        "format",
		Usage:       Messageformat,
		Value:       `json`,
		DefaultText: `json`,
		EnvVars:     []string{"AUTHDBCTL_OUTPUT_FORMAT"},
	})
	sh.Flags = append(sh.Flags, &cli.BoolFlag{
		Name:  "debug",
		Usage: Messagedebug,
	})
	sh.Commands = []*cli.Command{
		{
			Name:   "connect",
			Usage:  Messageconnect,
			Action: connect,
		},
		{
			Name:   "metadata",
			Usage:  Messagemetadata,
			Action: metadata,
		},
		{
			Name:        "add",
			Usage:       Messageadd,
			Subcommands: addSubcmd,
		},
		{
			Name:        "list",
			Usage:       Messagelist,
			Subcommands: listSubcmd,
		},
	}
}

func main() {
	err := sh.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
