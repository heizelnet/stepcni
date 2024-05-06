package config

import (
	"encoding/json"
	"os"

	klog "k8s.io/klog/v2"
)

const (
	SubnetFilePath = "/run/stepcni/subnet.json"
	DefaultBridge  = "cni0"
)

type Defaultconf struct {
	Name       string `json:"name"`
	CNIVersion string `json:"cniVersion"`
	Type       string `json:"type"`
	DataDir    string `json:"dataDir"`
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

func StoreSubnetConfig(conf *SubnetConf) error {
	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	return os.WriteFile(SubnetFilePath, data, 0644)
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
