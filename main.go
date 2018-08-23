package main

import (

	"qos-agent/agent"
	"fmt"
)



func main() {
	qos := agent.NewQosController()
	/*
	err := qos.InsertRule("10.222.88.202","10.222.119.72", "1000")
	if err != nil {
		fmt.Println(err.Error())
	}

	chans, err :=qos.GetChains("nat")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(chans)
	*/
	err := qos.DeleteRule("10.222.88.202","10.222.119.72", "1000")
	if err != nil {
		fmt.Println(err.Error())
	}
}


