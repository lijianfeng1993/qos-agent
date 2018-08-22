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
}
