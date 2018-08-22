package main

import (
	"qos-agent/agent"
	"fmt"
)

func main() {
	qosiptables := agent.NewQosIptables()
	qosiptables.InitIptables()

	chans, err :=qosiptables.GetChans("nat")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(chans)
}
