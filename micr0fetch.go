package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var iconchoice, colorchoice string
	flag.StringVar(&iconchoice, "icon", "", "override icon (Arch, Ubuntu, Manjaro, MacOs, Linux)")
	flag.StringVar(&colorchoice, "color", "", "override color (Red, Green, Yellow, Blue, Purple, Cyan, Grey, White)")

	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	// Detect Operating System
	cmd := exec.Command("uname", "-s")
	osname, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var isMacOs bool

	if strings.Contains(strings.ToLower(string(osname)), "darwin") {
		isMacOs = true
	}

	// Get kernel version (works on Both Mac and Linux)

	cmd = exec.Command("uname", "-r")
	kerneldata, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	kernel := strings.ReplaceAll(string(kerneldata), "\n", "")

	var operatingsys, architecture, host, uptime string

	if !isMacOs {
		// Get Operating system, Architecture, Hostname, and Uptime (Linux only)

		txtcmd := "hostnamectl | grep \"Operating System\""
		cmd = exec.Command("bash", "-c", txtcmd)
		operatingsysdata, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		operatingsys = strings.ReplaceAll(string(operatingsysdata), "Operating System: ", "")
		operatingsys = strings.ReplaceAll(operatingsys, "\n", "")
		operatingsys = strings.TrimSpace(operatingsys)

		txtcmd = "hostnamectl | grep \"Architecture\""
		cmd = exec.Command("bash", "-c", txtcmd)
		architecturedata, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		architecture = strings.ReplaceAll(string(architecturedata), "Architecture: ", "")
		architecture = strings.ReplaceAll(architecture, "\n", "")
		architecture = strings.TrimSpace(architecture)

		txtcmd = "hostnamectl | grep \"Static hostname\""
		cmd = exec.Command("bash", "-c", txtcmd)

		hostdata, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		host = strings.ReplaceAll(string(hostdata), "Static hostname: ", "")
		host = strings.ReplaceAll(host, "\n", "")
		host = strings.TrimSpace(host)

		cmd = exec.Command("uptime", "-p")

		updata, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		uptime = strings.ReplaceAll(string(updata), "up ", "")
		uptime = strings.ReplaceAll(uptime, "\n", "")
		uptime = strings.TrimSpace(uptime)
	} else {
		// Get Operating system, Architecture, Hostname, and Uptime (Mac only)

		operatingsys = "macOS"

		cmd = exec.Command("uname", "-m")
		architecturedata, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		architecture = strings.ReplaceAll(string(architecturedata), "\n", "")

		cmd = exec.Command("uname", "-n")
		hostdata, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		host = strings.ReplaceAll(string(hostdata), "\n", "")
		host = strings.TrimSpace(host)

		cmd = exec.Command("uptime")
		updata, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		uptime = strings.Split(string(updata), "up")[1]
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
	cmd = exec.Command("id", "-u", "-n")
	userdata, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user := strings.ReplaceAll(string(userdata), "\n", "")
	user = strings.TrimSpace(user)

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
