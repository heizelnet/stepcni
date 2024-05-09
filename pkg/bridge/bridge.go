package bridge

import (
	"fmt"
	"os/exec"

	"k8s.io/klog/v2"
)

func CreateBridge() {
	cmd := exec.Command("ip", "addr")

	output, err := cmd.Output()
	if err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return
	}

	fmt.Println(output)
}
