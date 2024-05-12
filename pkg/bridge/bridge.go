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

	//brctl addbr cni0
	cmd = exec.Command("ip", "link", "add", "name", brname, "type", "bridge")
	_, err = cmd.Output()
	if err != nil {
		klog.Errorf("[-] Fail to add bridge %s!", brname)
		return err
	}

	//brctl addif cni0 eth1
	cmd = exec.Command("ip", "addr", "add", fmt.Sprintf("%s/24", gateway), "dev", brname)
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to add interface to bridge!")
		return err
	}

	//ip link set cni0 up
	cmd = exec.Command("ip", "link", brname, "up")
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to bridge")
		return err
	}

	return nil
}

/*
	# veth up
	host_ifname="veth$n" # n=1,2,3,...
	ip link add $CNI_IFNAME type veth peer name $host_ifname
	ip link set $host_ifname up

	ip link set $host_ifname master cni0 # connect veth1 to bridge
	ip link set $CNI_IFNAME netns $nsname #mov eth0 to pod ns
*/

func SetupVeth(netns string, bridge string, ifname string, containerID string) error {

	//add netns
	cmd := exec.Command("ip", "netns", "add", netns)
	_, err := cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return err
	}

	//check interface existence and add veth pair
	veth := ""
	for i := 1; ; i++ {
		cmd = exec.Command("ip", "link", "show", fmt.Sprintf("veth%d", i))
		_, err = cmd.Output()
		if err != nil {
			klog.Infof("[+] Create veth%d interface!", i)
			veth = fmt.Sprintf("veth%d", i)
			break
		}
	}

	cmd = exec.Command("ip", "link", "add", ifname, "type", "veth", "peer", "name", veth)
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return err
	}

	//set veth interface up
	cmd = exec.Command("ip", "link", "set", veth, "up")
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return err
	}

	//connect veth to bridge
	cmd = exec.Command("ip", "link", "set", veth, "master", bridge)
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return err
	}

	cmd = exec.Command("ln", "-sfT", netns, fmt.Sprintf("/var/run/%s", containerID))
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return err
	}

	cmd = exec.Command("ip", "link", "set", ifname, "netns", containerID)
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return err
	}

	return nil
}
