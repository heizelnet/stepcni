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

func CreateBridge(gateway string, brname string) error {

	//check already has cni0 interface
	cmd := exec.Command("ip", "link", "show", brname)
	_, err := cmd.Output()
	if err != nil {
		klog.Error("[-] Already has cni0!")
		return err
	}

	//add bridge
	cmd = exec.Command("ip", "link", "add", "name", brname, "type", "bridge")
	_, err = cmd.Output()
	if err != nil {
		klog.Errorf("[-] Fail to add bridge %s!", brname)
		return err
	}

	//interface setup
	cmd = exec.Command("ip", "link", brname, "up")
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to bridge")
		return err
	}

	//add address
	cmd = exec.Command("ip", "addr", "add", fmt.Sprintf("%s/24", gateway), "dev", brname)
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to add interface to bridge!")
		return err
	}

	return nil
}
