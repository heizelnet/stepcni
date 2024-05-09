package bridge

import (
	"fmt"
	"os/exec"

	"k8s.io/klog/v2"
)

/*
	# create a new bridge if doesn't exist
	brctl addbr cni0
	brctl addif cni0 eth1
	ip link set cni0 up
*/

const mtu = 1500

func CreateBridge() {
	cmd := exec.Command("ip", "addr")

	output, err := cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return
	}

	fmt.Println(output)
}
