package veth

import (
	"fmt"
	"os/exec"

	"k8s.io/klog/v2"
)

/*
	# veth up
	host_ifname="veth$n" # n=1,2,3,...
	ip link add $CNI_IFNAME type veth peer name $host_ifname
	ip link set $host_ifname up

	ip link set $host_ifname master cni0 # connect veth1 to bridge
	ip link set $CNI_IFNAME netns $nsname #mov eth0 to pod ns
*/

func SetupVeth() {
	cmd := exec.Command("ip", "addr")

	output, err := cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return
	}

	fmt.Println(output)
}
