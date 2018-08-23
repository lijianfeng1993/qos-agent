
package agent

import (
	"github.com/coreos/go-iptables/iptables"
	"errors"
	"fmt"
)

type QosController struct {
	qosIptables *iptables.IPTables
	qosTc string
}

func NewQosController() *QosController {
	qosIptables, _ := iptables.NewWithProtocol(iptables.ProtocolIPv4)
	qoscon := &QosController{
		qosIptables: qosIptables,
		qosTc: "tc",
	}
	return qoscon
}

func (this *QosController) GetChains(table string) (chans []string, err error) {
	chans, err = this.qosIptables.ListChains(table)
	if err != nil {
		chains := []string {""}
		return chains,errors.New(fmt.Sprintf("ListChans(): %s", err.Error()))
	}
	return chans, nil
}

func (this *QosController) InsertRule(srcIp, dstIp, mark string) error{
	//fmt.Println(rulespec)   ---->  [-o ens160 -d 10.222.119.72 -s 10.222.88.202 -j MARK --set-mark 1000]
	//fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(rulespec), "[]"), " ", ",", -1))
	// -----> -o,ens160,-d,10.222.119.72,-s,10.222.88.202,-j,MARK,--set-mark,1000
	//fmt.Println()
	//err := this.qosIptables.Insert(table, chain, pos, strings.Replace(strings.Trim(fmt.Sprint(rulespec), "[]"), " ", ",", -1),)

	err := this.qosIptables.Insert(
		"mangle",
		"POSTROUTING",
		1,
		"-o",
		"ens160",
		"-d", fmt.Sprintf("%s", dstIp),
		"-s", fmt.Sprintf("%s", srcIp),
		"-j", "MARK",
		"--set-mark", fmt.Sprintf("%s", mark),)

	if err != nil {
		return errors.New(fmt.Sprintf("InsertRule(): %s \n", err.Error()))
	}

	return nil
}








