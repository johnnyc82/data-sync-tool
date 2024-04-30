package synctool

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	ct "github.com/ivystreetweb/config-tool/cli"

	"github.com/manifoldco/promptui"
)

func ConnectKinsta(projectName string) error {

	sv, err := ct.GetProjectConfig(projectName)
	if err != nil {
		return err
	}

	envSelect := ""
	env := ""

	prompt := promptui.Select{
		Label: "Select environment",
		Items: []string{"Live", "Staging"},
	}

	_, envSelect, err = prompt.Run()
	if err != nil {
		return err
	}

	if envSelect == "Live" {
		env = sv.KinstaPortLive
	} else {
		env = sv.KinstaPortStaging
	}

	log.Println("Connect to Kinsta site's server...")
	serverName := fmt.Sprintf("%s@%s", sv.KinstaUserName, sv.KinstaIP)
	cmdssh := exec.Command("ssh", "-T", "-i", "~/.ssh/halie_sync", serverName, "-p", env)
	cmdssh.Stdout = os.Stdout
	cmdssh.Stderr = os.Stderr
	err = cmdssh.Run()
	if err != nil {
		return err
	}
	log.Println("Connected via SSH")

	return nil
}
