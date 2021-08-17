package xdp_dns_mod

import (
	"fmt"
	"github.com/cilium/ebpf"
	"xdp_dns_mod/xdpobject"
)

var objs xdpobject.XdpDNSObjects

type CacheIns interface {
	Update(key, val interface{}, flags ebpf.MapUpdateFlags) error
	Put(key, val interface{}) error
	Delete(key interface{}) error
	Lookup(key, valueOut interface{}) error
	LookupAndDelete(key, valueOut interface{}) error
}

type MapCache struct {
	Qname CacheIns
	Count CacheIns
}

func GetXDPObject() (*MapCache, error) {
	if &objs == nil{
		return nil, ObjNotLoadError
	}
	m := &MapCache{Count: objs.FrameCount}
	return m, nil
}

func init() {
	if err := xdpobject.LoadXdpDNSObjects(&objs, nil); err != nil{
		fmt.Println("load xdp dns object error.")
		return
	}
}


