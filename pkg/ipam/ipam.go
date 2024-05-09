package ipam

import (
	"fmt"
	"os/exec"
)

/*
 1. Allocate IP

 2. Manage IP with file (lock with mutex when accesing)
*/

func AllocateIP() {
	cmd := exec.Command("ip", "addr")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(output)
}

func GetSubnetFromFile() {
	cmd := exec.Command("ip", "addr")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(output)
}
