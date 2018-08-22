package main

import (
	"qos-agent/agent"
	"fmt"
)

func main() {
	qosiptables := agent.NewQosController()
	qosiptables.InitIptables()

	chans, err :=qosiptables.GetChains("nat")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(chans)

	err1 := qosiptables.InsertRule("mangle", "POSTROUTING", 1,
		"-o", "ens160",
				"-d", "10.222.119.72",
				"-s", "10.222.88.202",
				"-j", "MARK",
				"--set-mark", "1000",)
	if err1 != nil {
		fmt.Println(err.Error())
	}
}


