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
		networkInterface: globalInterface,
	}
}

func (this *QosIptables) InitIptables() error{
	err := errors.New("")
	this.qosIptables, err = iptables.NewWithProtocol(iptables.ProtocolIPv4)
	if err != nil {
		return errors.New(fmt.Sprintf("NewQosIptables() is failed. \n" + "   NewWithProtocol(): %s", err.Error()))
	}

	list_chan, err := this.qosIptables.ListChains("filter")
	if err != nil {
		return errors.New(fmt.Sprintf("NewQosIptables() is faild. \n" + "ListChains(): %s", err.Error()))
	}
	fmt.Println(list_chan)
}

