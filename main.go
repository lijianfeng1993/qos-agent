package main

import (
	"qos-agent/agent"
	"fmt"
)

func main() {
	qosiptables := agent.NewQosIptables()
	listchan := qosiptables.InitIptables()
	fmt.Println(listchan)
}
