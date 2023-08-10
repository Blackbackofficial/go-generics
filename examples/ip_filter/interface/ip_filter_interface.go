package _interface

import (
	"encoding/hex"
	"fmt"
	"generics/examples/ip_filter"
	"net/netip"
)

type IpFilterRow interface {
	GetIPMin() interface{}
	GetIPMax() interface{}
	GetMask() int
	GetIsThrottling() uint
}

type IPFilter struct {
	version  int
	gm       *ip_filter.Geomap
	rows     []IpFilterRow
	buffer   []byte
	ipsCount int
}

func (f *IPFilter) Process(row IpFilterRow) error {
	ipBytes, err := f.inetNtoA(row.GetIPMin())
	if err != nil {
		return err
	}

	if !f.isIPValid(ipBytes) {
		return nil
	}

	isThrottling := row.GetIsThrottling()

	if isThrottling == 1 {
		f.addRange(row, 2)
	} else {
		f.addRange(row, 1)
		maskBytes, err := f.prefixToMask(row.GetMask())
		if err != nil {
			return err
		}
		f.ipsCount++
		f.buffer = append(f.buffer, ipBytes...)
		f.buffer = append(f.buffer, maskBytes...)
	}

	return nil
}

func NewIPFilterProcessor(rows []IpFilterRow, gm *ip_filter.Geomap, version int) *IPFilter {
	return &IPFilter{
		rows:     rows,
		version:  version,
		gm:       gm,
		ipsCount: 0,
	}
}

func (f *IPFilter) GetBytes() []byte {
	return f.buffer
}

func (f *IPFilter) GetIPCount() uint32 {
	return uint32(f.ipsCount)
}

func (f *IPFilter) inetNtoA(ip interface{}) ([]byte, error) {
	var bytes []byte
	var err error

	switch f.version {
	case ip_filter.IPV4:
		bytes, err = f.inetNtoAV4(ip.(uint32))
	case ip_filter.IPV6:
		bytes, err = f.inetNtoAV6(ip.(string))
	}

	return bytes, err
}

func (f *IPFilter) inetNtoAV4(ipNumber uint32) ([]byte, error) {
	bytes := make([]byte, 4)

	bytes[3] = byte(ipNumber & 0xFF)
	bytes[2] = byte((ipNumber >> 8) & 0xFF)
	bytes[1] = byte((ipNumber >> 16) & 0xFF)
	bytes[0] = byte((ipNumber >> 24) & 0xFF)

	return bytes, nil
}

func (f *IPFilter) inetNtoAV6(ipString string) ([]byte, error) {
	bytes, err := hex.DecodeString(ipString)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func (f *IPFilter) prefixToMask(mask int) ([]byte, error) {
	switch f.version {
	case ip_filter.IPV4:
		return f.prefixToMaskV4(mask)
	case ip_filter.IPV6:
		return f.prefixToMaskV6(mask)
	}
	return []byte{}, fmt.Errorf("invalid version")
}

func (f *IPFilter) prefixToMaskV4(prefix int) ([]byte, error) {
	return f.inetNtoAV4(0xFFFFFFFF << (32 - prefix))
}

func (f *IPFilter) prefixToMaskV6(prefix int) ([]byte, error) {
	maskBytes := make([]byte, 16)

	for i := 0; i < 16; i++ {
		if prefix >= 8 {
			maskBytes[i] = 0xFF
			prefix -= 8
		} else if prefix > 0 {
			maskBytes[i] = byte(0xFF << (8 - prefix))
			prefix = 0
		} else {
			maskBytes[i] = 0
		}
	}

	return maskBytes, nil
}

func (f *IPFilter) isIPValid(bytes []byte) bool {
	var addr netip.Addr

	switch f.version {
	case ip_filter.IPV4:
		var ipV4 [4]byte
		copy(ipV4[:], bytes)
		addr = netip.AddrFrom4(ipV4)
	case ip_filter.IPV6:
		var ipV6 [16]byte
		copy(ipV6[:], bytes)
		addr = netip.AddrFrom16(ipV6)
	}

	return addr.IsValid()
}

func (f *IPFilter) addRange(row IpFilterRow, rowType uint) {
	switch f.version {
	case ip_filter.IPV4:
		f.gm.AddRange4(row.GetIPMin().(uint32), row.GetIPMax().(uint32), rowType)
	case ip_filter.IPV6:
		f.gm.AddRange6(row.GetIPMin().(string), row.GetIPMax().(string), rowType)
	}
}
