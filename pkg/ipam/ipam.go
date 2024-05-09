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

/*
	#cacluating ip address
	ip netns exec $nsname ip link set $CNI_IFRAME up
	ip netns exec $nsname ip addr add $ip/24 dev $CNI_IFNAME
	ip netns exec $nsname ip route add default via $podcidr_gw dev $CNI_IFNAME
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
