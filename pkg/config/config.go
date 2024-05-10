package config

import (
	"encoding/json"
	"fmt"
	"os"

	klog "k8s.io/klog/v2"
)

/*
	/etc/cni/net.d/20-step-cni.conf
	{
		cniVersion : 0.4.0
		name : stepcni
		type : stepcni
		DataDir : ~ pathname
		podcidr : 10.240.0.0/24
	}
*/

const (
	cniConfigPath  = "/etc/cni/net.d/20-stepcni.conf"
	SubnetFilePath = "/run/stepcni/subnet.json"
	DefaultBridge  = "cni0"
)

type Defaultconf struct {
	Name       string `json:"name"`
	CNIVersion string `json:"cniVersion"`
	Type       string `json:"type"`
	PodCidr    string `json:"podcidr"`
}

type SubnetConf struct {
	Subnet string `json:"subnet"`
	Bridge string `json:"bridge"`
}

type CNIConf struct {
	Defaultconf
	SubnetConf
}

func LoadDefaultConfig(stdin []byte) (*Defaultconf, error) {
	conf := Defaultconf{}

	/*
		for debugging, print entire data
	*/
	klog.Infof("%s", stdin)

	if err := json.Unmarshal(stdin, &conf); err != nil {
		klog.Error("[-] Fail to load default CNI config!")
		return nil, err
	}

	return &conf, nil
}

func LoadSubnetConfig() (*SubnetConf, error) {
	data, err := os.ReadFile(SubnetFilePath)
	if err != nil {
		klog.Error("[-] Fail to load Subnet config!")
		return nil, err
	}

	conf := &SubnetConf{}
	if err := json.Unmarshal(data, conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func LoadCNIConfig(stdin []byte) (*CNIConf, error) {
	defaultconf, err := LoadDefaultConfig(stdin)
	if err != nil {
		return nil, err
	}

	subnetConf, err := LoadSubnetConfig()
	if err != nil {
		return nil, err
	}
	return &CNIConf{Defaultconf: *defaultconf, SubnetConf: *subnetConf}, nil
}

func StoreSubnetConfig(conf *SubnetConf) error {
	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	return os.WriteFile(SubnetFilePath, data, 0644)
}

func StoreCNIPluginConfig() error {
	fd, err := os.OpenFile(cniConfigPath,
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	if err != nil {
		klog.Errorf("[-] Open CNI config file error")
		return err
	}

	defer fd.Close()

	if _, err = fd.Write([]byte(fmt.Sprintf(cniConfTemplate, node.Spec.PodCIDR))); err != nil {
		klog.Errorf("[-]  Write CNI config file error")
		return err
	}

	return nil
}
