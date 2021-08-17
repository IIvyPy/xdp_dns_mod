package xdp_dns_mod

import (
	"errors"
	"xdp_dns_mod/internal"
)

const (
	XDP_FLAGS_SKB_MODE  = 1 << 1
	XDP_FLAGS_DRV_MODE  = 1 << 2
	XDP_FLAGS_HW_MODE   = 1 << 3
	XDP_FLAGS_REPLACE   = 1 << 4
	XDP_FLAGS_MODES = XDP_FLAGS_SKB_MODE | XDP_FLAGS_DRV_MODE | XDP_FLAGS_HW_MODE
)

var ObjNotLoadError = errors.New("object of XDP DNS is not load ever")

func AttachXDPFD(netCard string) error {
	if &objs == nil {
		return ObjNotLoadError
	}
	// should be objs.XdpDNSPrograms.FD().
	return internal.AttachXDPWithFlags(netCard, -1, XDP_FLAGS_MODES)
}

func DetachXDPFD(netCard string) error {
	if &objs == nil{
		return ObjNotLoadError
	}
	return internal.DetachXDPWithFlags(netCard, XDP_FLAGS_MODES)
}

