package agent

import (
	"github.com/coreos/go-iptables/iptables"
	"errors"
	"fmt"
)

type QosIptables struct {
	qosIptables *iptables.IPTables
	networkInterface string
}

func NewQosIptables() *QosIptables {
	return &QosIptables{
		networkInterface: "ens160",
	}
}

func (this *QosIptables) InitIptables() (err error){

	this.qosIptables, err = iptables.NewWithProtocol(iptables.ProtocolIPv4)
	if err != nil {
		return  errors.New(fmt.Sprintf("NewQosIptables() is failed. \n" + "   NewWithProtocol(): %s", err.Error()))
	}

	return nil
}

func (this *QosIptables) GetChans(table string) (chans []string, err error) {
	chans, err = this.qosIptables.ListChains(table)
	if err != nil {
		chans := []string {""}
		return chans,errors.New(fmt.Sprintf("ListChans(): %s", err.Error()))
	}
	return chans, nil
}






