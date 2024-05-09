package veth

import (
	"fmt"
	"os/exec"
)

func SetupVeth() {
	cmd := exec.Command("ip", "addr")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(output)
}
