package main

import (
	_ "embed"
	"log"

	st "github.com/ivystreetweb/halie-sync-tool/synctool"

	"github.com/leaanthony/clir"
)

//go:embed version.txt
var version string

func main() {

	var fromPath, toPath, projectName string
	var pull, push bool
	var live, staging, local bool
	var data, assets, masterplan, media, all bool

	// Halie Sync CLI created
	cli := clir.NewCli("hsync", "Halie Sync Tool", version)

	//
	// capture flags
	//
	cli.StringFlag("from", "location data is coming from", &fromPath).
		StringFlag("to", "location data is going to", &toPath).
		BoolFlag("pull", "data is being pulled from remote server", &pull).
		BoolFlag("push", "data is being pushed to remote server", &push).
		BoolFlag("live", "sync with production environment", &live).
		BoolFlag("staging", "sync with staging environment", &staging).
		BoolFlag("local", "sync with local environment", &local).
		BoolFlag("data", "sync halie stock & configuration data", &data).
		BoolFlag("assets", "sync halie assets (images/pdfs etc.)", &assets).
		BoolFlag("masterplan", "sync halie masterplan", &masterplan).
		BoolFlag("media", "sync wordpress media files (does not sync db)", &media).
		BoolFlag("all", "sync all halie data & assets", &all).
		StringFlag("projectname", "project to sync with", &projectName)

	// init command
	initCmd := cli.NewSubCommand("init", "Test SSH connection to selected Kinsta site")
	initCmd.Action(func() error {
		return st.ConnectKinsta(projectName)
	})

	// sync command
	syncCmd := cli.NewSubCommandInheritFlags("sync", "Sync data and assets for halie sites")
	syncCmd.Action(func() error {
		// capture the flags provided
		sf := &st.SyncFlags{
			ProjectName: projectName,
			FromPath:    fromPath,
			ToPath:      toPath,
			Pull:        pull,
			Push:        push,
			Live:        live,
			Staging:     staging,
			Local:       local,
			Data:        data,
			Assets:      assets,
			Masterplan:  masterplan,
			Media:       media,
			All:         all,
		}
		// check that the flags are sensible
		err := sf.Check()
		if err != nil {
			return err
		}

		// run the sync command with these flags
		return st.Sync(sf)
	})

	err := cli.Run()
	if err != nil {
		log.Fatal(err)
	}
}
