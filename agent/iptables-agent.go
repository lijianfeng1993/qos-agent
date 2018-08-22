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

func InitIptables() (*QosIptables,error){
	network := "ens160"
	qosIptables, err := iptables.NewWithProtocol(iptables.ProtocolIPv4)
	if err != nil {
		return _, errors.New(fmt.Sprintf("InitIptables() is failed. \n" + "   NewWithProtocol(): %s", err.Error()))
	}
	return &QosIptables{networkInterface: network, qosIptables: qosIptables,}, nil
}




/*
func (ipt *QosIptables) GetChans(table string) (chans []string, err error) {
	chans, err = this.qosIptables.ListChains(table)
	if err != nil {
		chans := []string {""}
		return chans,errors.New(fmt.Sprintf("ListChans(): %s", err.Error()))
	}
	return chans, nil
}
*/

//func (this *QosIptables) InsertRule() error {
//	err = this.InsertRule()
//}




