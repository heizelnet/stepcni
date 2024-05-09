package ipam

import (
	"fmt"
	"net"
	"os/exec"

	"k8s.io/klog/v2"
)

/*
 1. Allocate IP

 2. Manage IP with file (lock with mutex when accesing)

 3. Other Utilizing IP address functions
*/

type IPAM struct {
	subnet  *net.IPNet
	gateway net.IP
}

func AllocateIP() {
	cmd := exec.Command("ip", "addr")

	output, err := cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return
	}

	fmt.Println(output)
}

func GetSubnetFromFile() {
	cmd := exec.Command("ip", "addr")

	output, err := cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return
	}

	fmt.Println(output)
}
