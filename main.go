package main

import (
	"qos-agent/agent"
	"fmt"
)

func main() {
	//qosiptables, err := agent.InitIptables()
	qosiptables, err := agent.InitIptables()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(qosiptables)


		/*
	chans, err :=qosiptables.GetChans("nat")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(chans)
	*/


}
