package main

/*
	// Todo
	1. Add detail error handling with using PluginMainFuncsWithError..
*/
import (
	"runtime"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/heizelnet/stepcni/pkg/config"
	klog "k8s.io/klog/v2"
)

func init() {
	// this ensures that main runs only on main thread (thread group leader).
	// since namespace ops (unshare, setns) are done for a single thread, we
	// must ensure that the goroutine does not jump from OS thread to thread
	runtime.LockOSThread()

	klog.InitFlags(nil)
}

func main() {
	skel.PluginMainFuncs(skel.CNIFuncs{
		Add:   cmdAdd,
		Check: cmdCheck,
		Del:   cmdDel,
	}, version.All, "CNI stepcni v0.4.0")
}

func cmdAdd(args *skel.CmdArgs) error {

	//log for ADD request
	klog.Infof("[+] cmdAdd details: containerID = %s, netNs = %s, ifName = %s, path = %s, stdin = %s",
		args.ContainerID,
		args.Netns,
		args.IfName,
		args.Path,
		string(args.StdinData),
	)

	conf, err := config.LoadCNIConfig(args.StdinData)
	if err != nil {
		klog.Error("[-] LoadCNIConfig Error!")
		return err
	}

	klog.Infof("[+] DataDir Path : %s", conf.DataDir)
	/*
		(Done) 1. LoadCNIConfig

		2. Open ip Store File (modify ip stored file)

		3. ipam.NewIPAM

		4. gateway = ipam.gateway

		5. ipam.AllocateIP

		6. CreateBridge

		7. bridge.SetupVeth
	*/

	return nil
}

// TODO
func cmdCheck(args *skel.CmdArgs) error {
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	//log for DEL request
	klog.Infof("[+] cmdAdd details: containerID = %s, netNs = %s, ifName = %s, args = %s, path = %s, stdin = %s",
		args.ContainerID,
		args.Netns,
		args.IfName,
		args.Args,
		args.Path,
		string(args.StdinData),
	)

	conf, err := config.LoadCNIConfig(args.StdinData)
	if err != nil {
		klog.Error("[-] LoadCNIConfig Error!")
		return err
	}

	klog.Infof("[+] DataDir Path : %s", conf.DataDir)
	/*
		1. LoadCNIConfig

		2. NewStore (modify ip stored file)

		3. ipam.NewIPAM

		4. ReleaseIP

		5. netns.GetNS

		6. bridge.DelVeth
	*/
	return nil
}

/*
 error return
 	if err != nil {
		return nil, err
	}
*/
