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
	return &QosController{
	}
}

func (this *QosController) InitIptables() (err error){

	this.qosIptables, err = iptables.NewWithProtocol(iptables.ProtocolIPv4)
	if err != nil {
		return  errors.New(fmt.Sprintf("NewQosIptables() is failed. \n" + "   NewWithProtocol(): %s", err.Error()))
	}

	return nil
}

/*
func (this *QosController) InitTc() (err error) {

}
*/

func (this *QosController) GetChains(table string) (chans []string, err error) {
	chans, err = this.qosIptables.ListChains(table)
	if err != nil {
		chains := []string {""}
		return chains,errors.New(fmt.Sprintf("ListChans(): %s", err.Error()))
	}
	return chans, nil
}

func (this *QosController) InsertRule(table string, chain string, pos int, rulespec ...string) error{
	//fmt.Println(rulespec)   ---->  [-o ens160 -d 10.222.119.72 -s 10.222.88.202 -j MARK --set-mark 1000]
	//fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(rulespec), "[]"), " ", ",", -1))
	// -----> -o,ens160,-d,10.222.119.72,-s,10.222.88.202,-j,MARK,--set-mark,1000
	

	//err := this.qosIptables.Insert(table, chain, pos, strings.Replace(strings.Trim(fmt.Sprint(rulespec), "[]"), " ", ",", -1),)
	err := this.qosIptables.Insert("mangle", "POSTROUTING", 1, "-o", "ens160", "-d", "10.222.119.72", "-s", "10.222.88.202", "-j", "MARK", "--set-mark", "1000",)
	if err != nil {
		return errors.New(fmt.Sprintf("InsertRule(): %s \n", err.Error()))
	}
	return nil
}







