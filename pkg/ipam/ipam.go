package ipam

import (
	"fmt"
	"net"
	"os/exec"

	"github.com/heizelnet/stepcni/pkg/ipam/datastore"

	"k8s.io/klog/v2"
)

/*
 1. Allocate IP

 2. Manage IP with file (lock with mutex when accesing)

 3. Other Utilizing IP address functions
*/

/*
	#cacluating ip address
	ip netns exec $nsname ip link set $CNI_IFNAME up
	ip netns exec $nsname ip addr add $ip/24 dev $CNI_IFNAME
	ip netns exec $nsname ip route add default via $podcidr_gw dev $CNI_IFNAME
*/

type IPAM struct {
	subnet    *net.IPNet
	gateway   net.IP
	datastore *datastore.DataStore
}

func NewIPAM() (*IPAM, error) {

	return nil, nil
}

func (ipam *IPAM) AllocateIP() (net.IP, error) {
	ipam.datastore.

	return nil, nil
}

/*
func SetIPAM(conf *config.CNIConf, containerID string, ifname string) (*IPAM, error) {

	//ip netns exec $nsname ip link set $CNI_IFNAME up
	cmd := exec.Command("ip", "netns", "exec", containerID, "ip", "link", "set", ifname, "up")
	_, err := cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return nil, err
	}

	//ip netns exec $nsname ip addr add $ip/24 dev $CNI_IFNAME
	cmd = exec.Command("ip", "netns", "exec", containerID, "ip", "addr", "add", conf.Subnet, "dev", ifname)
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return nil, err
	}

	//ip netns exec $nsname ip route add default via $podcidr_gw dev $CNI_IFNAME
	cmd = exec.Command("ip", "netns", "exec", containerID, "ip", "route", "add", "default", "via", conf.PodCidr, "dev", ifname)
	_, err = cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return nil, err
	}

	_, ipnet, err := net.ParseCIDR(conf.Subnet)
	if err != nil {
		return nil, err
	}

	im := &IPAM{
		subnet: ipnet,
	}

	im.gateway, err = im.Ne

}
*/

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
