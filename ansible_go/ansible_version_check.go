package main

import (
	"flag"
	"fmt"
	"os/exec"
)


func ansibleVersion()string {
	// execute ansible version 
	cmd := exec.Command("ansible", "--version")

	// run command
	if output, err := cmd.Output(); err != nil {
		fmt.Println("Error!", err)
	} else {
		fmt.Printf("Ansible Version: %s\n", output)
	}
	return ""
}

// Priv Deploy

func privDeploy() string {
	// execute ansible version 
	cmd := exec.Command("ansible-playbook", "--version")

	// run command
	if output, err := cmd.Output(); err != nil {
		fmt.Println("Error!", err)
	} else {
		fmt.Printf("Ansible Version: %s\n", output)
	}
	return ""

}

func main() {
	var (
		checkVersion bool 
		priv 		 bool 

	)

	{
		flag.BoolVar(&checkVersion, "CheckVersion", true,  "Check Ansible version")
		flag.BoolVar(&priv, "priv", true, "check priv")
		flag.Parse()
	}
	if checkVersion {
		ansibleVersion()
		return
	}

	if priv {
		privDeploy()
		return
	}
	
	flag.Usage()
}
