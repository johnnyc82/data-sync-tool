package synctool

import (
	"fmt"
	"os"
	"os/exec"

	ct "github.com/exampleowner/config-tool/cli"
)

var dataPathTemplate = "web/app/themes/%s/resources/data"
var assetsPathTemplate = "web/app/themes/%s/resources/lots"
var mpPathTemplate = "web/app/themes/%s/resources/masterplans"
var mediaPathTemplate = "web/app/uploads"

func Sync(flags *SyncFlags) error {

	sv, err := ct.GetProjectConfig(flags.ProjectName)
	if err != nil {
		return err
	}

	//
	// handle sync to local if config not set
	//
	if flags.Local && sv.LocalPath == "" {
		return fmt.Errorf("cannot sync to local as location not set.")
	}

	var commands []*exec.Cmd
	if flags.Pull {
		commands, err = createPullCommands(flags, sv)
	} else {
		commands, err = createPushCommands(flags, sv)
	}
	if err != nil {
		return err
	}

	for _, cmd := range commands {
		c := cmd
		// log.Println("running command:", c)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err = c.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

// assembles the set of requested pull commands
func createPullCommands(flags *SyncFlags, config ct.Project) ([]*exec.Cmd, error) {

	// commands to execute
	pullCommands := []*exec.Cmd{}

	//
	// if user has set both from & to paths, then we just implement the one
	// command as data type flags are not relevent
	//
	if flags.FromPath != "" && flags.ToPath != "" {
		userCommand, err := pullCommand(flags, config, "")
		if err != nil {
			return nil, err
		}
		pullCommands = append(pullCommands, userCommand)
		return pullCommands, nil
	}

	//
	// build the commands for the data types requested
	//
	if flags.Data {
		dataCommand, err := pullCommand(flags, config, fmt.Sprintf(dataPathTemplate, config.ThemeName))
		if err != nil {
			return nil, err
		}
		pullCommands = append(pullCommands, dataCommand)
	}

	if flags.Assets {
		assetsCommand, err := pullCommand(flags, config, fmt.Sprintf(assetsPathTemplate, config.ThemeName))
		if err != nil {
			return nil, err
		}
		pullCommands = append(pullCommands, assetsCommand)
	}

	if flags.Masterplan {
		mpCommand, err := pullCommand(flags, config, fmt.Sprintf(mpPathTemplate, config.ThemeName))
		if err != nil {
			return nil, err
		}
		pullCommands = append(pullCommands, mpCommand)
	}

	if flags.Media {
		mediaCommand, err := pullCommand(flags, config, mediaPathTemplate)
		if err != nil {
			return nil, err
		}
		pullCommands = append(pullCommands, mediaCommand)
	}

	return pullCommands, nil
}

// builds each required pull command
func pullCommand(flags *SyncFlags, config ct.Project, path string) (*exec.Cmd, error) {

	var cmd *exec.Cmd

	// build ssh section of command, empty string for Local environment
	sshParams := ""
	if flags.Live {
		sshParams = fmt.Sprintf("ssh -i ~/.ssh/Data_sync -p %s", config.KinstaPortLive)
	}
	if flags.Staging {
		sshParams = fmt.Sprintf("ssh -i ~/.ssh/Data_sync -p %s", config.KinstaPortStaging)
	}

	// build remote server section of command
	remoteServerPath := fmt.Sprintf("%s@%s:public", config.KinstaUserName, config.KinstaIP)
	// for local 'server' is just a filepath
	if flags.Local {
		remoteServerPath = config.LocalPath
	}

	// build the complete command params
	var fromPath, toPath string
	fromPath = fmt.Sprintf("%s/%s/", remoteServerPath, path) // default to the path provided
	// override if set via flags
	if flags.FromPath != "" {
		fromPath = fmt.Sprintf("%s/%s/", remoteServerPath, flags.FromPath)
	}
	toPath = path // default to path provided, but at our current location
	// override toPath if set via flags
	if flags.ToPath != "" {
		toPath = flags.ToPath
	}

	// ensure the toPath exists locally (rsync won't create nested paths by default)
	err := os.MkdirAll(toPath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// create the command
	cmd = exec.Command("rsync", "-urave", sshParams, fromPath, toPath)

	return cmd, nil
}

// assembles the set of requested push commands
func createPushCommands(flags *SyncFlags, config ct.Project) ([]*exec.Cmd, error) {

	// commands to execute
	pushCommands := []*exec.Cmd{}

	//
	// if user has set both from & to paths, then we just implement the one
	// command as data type flags are not relevent
	//
	if flags.FromPath != "" && flags.ToPath != "" {
		userCommand, err := pushCommand(flags, config, "")
		if err != nil {
			return nil, err
		}
		pushCommands = append(pushCommands, userCommand)
		return pushCommands, nil
	}

	//
	// build the commands for the data types requested
	//
	if flags.Data {
		dataCommand, err := pushCommand(flags, config, fmt.Sprintf(dataPathTemplate, config.ThemeName))
		if err != nil {
			return nil, err
		}
		pushCommands = append(pushCommands, dataCommand)
	}

	if flags.Assets {
		assetsCommand, err := pushCommand(flags, config, fmt.Sprintf(assetsPathTemplate, config.ThemeName))
		if err != nil {
			return nil, err
		}
		pushCommands = append(pushCommands, assetsCommand)
	}

	if flags.Masterplan {
		mpCommand, err := pushCommand(flags, config, fmt.Sprintf(mpPathTemplate, config.ThemeName))
		if err != nil {
			return nil, err
		}
		pushCommands = append(pushCommands, mpCommand)
	}

	if flags.Media {
		mediaCommand, err := pushCommand(flags, config, mediaPathTemplate)
		if err != nil {
			return nil, err
		}
		pushCommands = append(pushCommands, mediaCommand)
	}

	return pushCommands, nil
}

// builds each required push command
func pushCommand(flags *SyncFlags, config ct.Project, path string) (*exec.Cmd, error) {

	var cmd *exec.Cmd

	// build ssh section of command, empty string for Local environment
	sshParams := ""
	if flags.Live {
		sshParams = fmt.Sprintf("ssh -i ~/.ssh/Data_sync -p %s", config.KinstaPortLive)
	}
	if flags.Staging {
		sshParams = fmt.Sprintf("ssh -i ~/.ssh/Data_sync -p %s", config.KinstaPortStaging)
	}

	// build remote server section of command
	remoteServerPath := fmt.Sprintf("%s@%s:public", config.KinstaUserName, config.KinstaIP)
	// for local 'server' is just a filepath
	if flags.Local {
		remoteServerPath = config.LocalPath
	}

	// build the complete command params
	var fromPath, toPath string
	fromPath = fmt.Sprintf("%s/", path) // default to path provided, but at our current location
	// override if set via flags
	if flags.FromPath != "" {
		fromPath = fmt.Sprintf("%s/", flags.FromPath)
	}
	toPath = fmt.Sprintf("%s/%s", remoteServerPath, path) // default to the path provided
	// override toPath if set via flags
	if flags.ToPath != "" {
		fromPath = fmt.Sprintf("%s/%s", remoteServerPath, flags.ToPath)
	}

	// create the command
	cmd = exec.Command("rsync", "-urave", sshParams, fromPath, toPath)

	return cmd, nil

}
