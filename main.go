package main

import "qos-agent/agent"

func main() {
	qosiptables := agent.NewQosIptables()
	qosiptables.InitIptables()
}
