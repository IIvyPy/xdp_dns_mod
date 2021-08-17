package internal

import "xdp_dns_mod/internal/unix"

func AttachXDPWithFlags(netCard string, fd, flags int) error{
	link, err := unix.NetLinkByName(netCard)
	if err != nil {
		return err
	}

	err = unix.NetLinkSetXdpFdWithFlags(link, fd, flags)
	if err != nil {
		return err
	}

	return nil
}

func DetachXDPWithFlags(netCard string, flags int) error{
	link, err := unix.NetLinkByName(netCard)
	if err != nil {
		return err
	}

	err = unix.NetLinkSetXdpFdWithFlags(link, -1, flags)
	if err != nil {
		return err
	}

	return nil
}
