package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

// Helper function to run command and return trimmed output string
func runCmd(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// Helper function to extract value from hostnamectl output
func extractHostnameCtlValue(field string) (string, error) {
	txtcmd := fmt.Sprintf("hostnamectl | grep \"%s\"", field)
	data, err := runCmd("bash", "-c", txtcmd)
	if err != nil {
		return "", err
	}
	// Replace field name and remove leading and trailing white spaces
	return strings.TrimSpace(strings.ReplaceAll(data, field+":", "")), nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
	}
}

func run() error {
	var iconchoice, colorchoice string
	flag.StringVar(&iconchoice, "icon", "", "override icon (Arch, Ubuntu, Manjaro, MacOs, Linux)")
	flag.StringVar(&colorchoice, "color", "", "override color (Red, Green, Yellow, Blue, Purple, Cyan, Grey, White)")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	// Detect Operating System
	osname, err := runCmd("uname", "-s")
	if err != nil {
		return err
	}

	var isMacOs bool

	if strings.Contains(strings.ToLower(osname), "darwin") {
		isMacOs = true
	}

	// Get kernel version (works on Both Mac and Linux)
	kernel, err := runCmd("uname", "-r")
	if err != nil {
		return err
	}

	var operatingsys, architecture, host, uptime string

	if !isMacOs {
		// Get Operating system, Architecture, Hostname, and Uptime (Linux only)
		operatingsys, err = extractHostnameCtlValue("Operating System")
		if err != nil {
			return err
		}

		architecture, err = extractHostnameCtlValue("Architecture")
		if err != nil {
			return err
		}

		host, err = extractHostnameCtlValue("Static hostname")
		if err != nil {
			return err
		}

		uptime, err = runCmd("uptime", "-p")
		if err != nil {
			return err
		}

		uptime = strings.ReplaceAll(uptime, "up ", "")
	} else {
		// Get Operating system, Architecture, Hostname, and Uptime (Mac only)
		operatingsys = "macOS"

		architecture, err = runCmd("uname", "-m")
		if err != nil {
			return err
		}

		host, err = runCmd("uname", "-n")
		if err != nil {
			return err
		}

		uptime, err = runCmd("uptime")
		if err != nil {
			return err
		}

		uptime = strings.Split(uptime, "up")[1]
		uptimesplit := strings.Split(uptime, ",")

		extra := strings.Join(strings.Split(uptimesplit[0], "")[1:], "")
		if strings.Contains(uptimesplit[1], "hrs") {
			hours := strings.Split(strings.ReplaceAll(uptimesplit[1], " ", ""), "hrs")[0]
			uptime = string(extra + ", " + hours + " hours")
		} else {
			uptimesplit = strings.Split(uptimesplit[len(uptimesplit)-3], ":")
			minutes := uptimesplit[1]
			uptimesplit = strings.Split(uptimesplit[0], " ")
			hours := uptimesplit[len(uptimesplit)-1]

			uptime = strings.ReplaceAll(string(extra+", "+hours+" hours, "+minutes+" minutes"), " 0", " ")
		}

	}

	// Get Active user (Both Mac and Linux)
	user, err := runCmd("id", "-u", "-n")
	if err != nil {
		return err
	}

	colorReset := "\033[0m"

	var color string

	if operatingsys == "amogos" || colorchoice == "amogus" {
		colors := [8]string{"red", "green", "yellow", "blue", "purple", "cyan", "gray", "white"}
		color = getColor(colors[rand.Intn(len(colors))])
	} else if colorchoice != "" {
		color = getColor(colorchoice)
	} else {
		color = getColor(operatingsys)
	}

	var iconSplit []string

	if iconchoice != "" {
		iconSplit = strings.Split(getIcon(iconchoice, color), "\n")
	} else {
		iconSplit = strings.Split(getIcon(operatingsys, color), "\n")
	}

	fmt.Println(color + iconSplit[1] + colorReset + "  " + color + string(user) + colorReset + "@" + color + string(host) + colorReset)
	fmt.Println(color + iconSplit[2] + "  " + "os     " + colorReset + string(operatingsys) + " " + string(architecture))
	fmt.Println(color + iconSplit[3] + "  " + "kernel " + colorReset + string(kernel))
	fmt.Println(color + iconSplit[4] + "  " + "uptime " + colorReset + string(uptime))

	return nil
}

func getIcon(distro string, color string) string {
	distrosplit := strings.Split(distro, " ")
	switch strings.ToLower(distrosplit[0]) {
	case "arch":
		return `
   /\   
  /\ \  
 /   -\ 
/__/\__\`
	case "ubuntu":
		return `
 ,-O 
O(_))
 ` + "`" + `-O 
     `
	case "manjaro":
		return `
 _ _ _ 
|  _| |
| | | |
|_|_|_|`
	case "macos":
		return `
 _` + "\033[32mQ" + color + `_ 
/   (
\___/
     
`
	case "fedora":
		return `
  (‾)
 _|_ 
  |  
(_)  `
	case "debian":
		return `
 __ 
( c)
 \. 
    `
	case "gentoo":
		return `
/‾‾‾‾‾\
\ o   /
/    / 
‾‾‾‾‾  `
	case "chromeos":
		return `
 ` + "\033[31m,-,_" + color + `
` + "\033[32m\\" + "\033[34m(O)" + "\033[33m)" + color + `
 ` + "\033[32m`-" + "\033[33m/" + color + ` 
`
	case "amogos":
		return `
    
  ඞ 
    
    `

	case "raspbian":
		return `
 ` + "\033[32m\\/ " + color + `
()()
 () 
    `
	case "opensuse":
		return `
,__ 
‾  o\
_ \_,
‾‾‾ 
`
	case "slackware":
		return `
 ╔═╗
 ╚═╗
|╚═╝
` + "`" + `‾‾‾
`
	}
	return `
  .-. 
  oo| 
 /` + "`" + `'\ 
(\_;/)`

}

func getColor(distro string) string {
	distrosplit := strings.Split(distro, " ")
	switch strings.ToLower(distrosplit[0]) {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "purple":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "gray":
		return "\033[90m"
	case "white":
		return "\033[37m"
	case "arch":
		return "\033[36m"
	case "ubuntu":
		return "\033[31m"
	case "manjaro":
		return "\033[32m"
	case "macos":
		return "\033[31m"
	case "fedora":
		return "\033[34m"
	case "debian":
		return "\033[31m"
	case "gentoo":
		return "\033[35m"
	case "chromeos":
		return "\033[32m"
	case "raspbian":
		return "\033[31m"
	case "opensuse":
		return "\033[32m"
	case "slackware":
		return "\033[37m"
	}
	return "\033[33m"
}
