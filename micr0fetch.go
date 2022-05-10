package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("uname", "-r")
	kerneldata, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	kernel := strings.ReplaceAll(string(kerneldata), "\n", "")

	txtcmd := "hostnamectl | grep \"Operating System\""
	cmd = exec.Command("bash", "-c", txtcmd)
	operatingsysdata, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	operatingsys := strings.ReplaceAll(string(operatingsysdata), "Operating System: ", "")
	operatingsys = strings.ReplaceAll(operatingsys, "\n", "")

	txtcmd = "hostnamectl | grep \"Architecture\""
	cmd = exec.Command("bash", "-c", txtcmd)
	architecturedata, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	architecture := strings.ReplaceAll(string(architecturedata), "Architecture: ", "")
	architecture = strings.ReplaceAll(architecture, "\n", "")
	architecture = strings.ReplaceAll(architecture, " ", "")

	txtcmd = "hostnamectl | grep \"Static hostname\""
	cmd = exec.Command("bash", "-c", txtcmd)

	hostdata, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	host := strings.ReplaceAll(string(hostdata), "Static hostname: ", "")
	host = strings.ReplaceAll(host, "\n", "")
	host = strings.ReplaceAll(host, " ", "")

	cmd = exec.Command("id", "-u", "-n")

	userdata, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user := strings.ReplaceAll(string(userdata), "\n", "")

	cmd = exec.Command("uptime", "-p")

	updata, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	uptime := strings.ReplaceAll(string(updata), "up ", "")
	uptime = strings.ReplaceAll(uptime, "\n", "")

	iconSplit := strings.Split(getIcon(operatingsys), "\n")

	fmt.Println(iconSplit[1] + "  " + string(user) + "@" + string(host))
	fmt.Println(iconSplit[2] + "  " + "os     " + string(operatingsys) + " " + string(architecture))
	fmt.Println(iconSplit[3] + "  " + "kernel " + string(kernel))
	fmt.Println(iconSplit[4] + "  " + "uptime " + string(uptime))
}

func getIcon(distro string) string {
	switch distro {
	case "Arch Linux":
		return `
   /\   
  /\ \  
 /   -\ 
/__/\__\`
	}
	return `
  .-. 
  oo| 
 /` + "`" + `'\ 
(\_;/)`
}
