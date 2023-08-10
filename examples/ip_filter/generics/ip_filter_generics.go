package generics

import (
	"encoding/hex"
	"fmt"
	"generics/examples/ip_filter"
	"net/netip"
)

type IpFilterRow interface {
	ip_filter.IPV4FilterRow | ip_filter.IPV6FilterRow
}

type IPFilter[T IpFilterRow] struct {
	version  int
	gm       *ip_filter.Geomap
	rows     []T
	buffer   []byte
	ipsCount int
}

func (f *IPFilter[T]) Process(row interface{}) error {
	ipBytes, err := f.inetNtoA(row)
	if err != nil {
		return err
	}

	if !f.isIPValid(ipBytes) {
		return nil
	}

	isThrottling, err := f.isThrottling(row)
	if err != nil {
		return err
	}

	if isThrottling == 1 {
		f.addRange(row, 2)
	} else {
		f.addRange(row, 1)
		maskBytes, err := f.prefixToMask(row)
		if err != nil {
			return err
		}
		f.ipsCount++
		f.buffer = append(f.buffer, ipBytes...)
		f.buffer = append(f.buffer, maskBytes...)
	}

	return nil
}

func NewIPFilterProcessor[T IpFilterRow](rows []T, gm *ip_filter.Geomap, version int) *IPFilter[T] {
	return &IPFilter[T]{
		rows:     rows,
		version:  version,
		gm:       gm,
		ipsCount: 0,
	}
}

func (f *IPFilter[T]) GetBytes() []byte {
	return f.buffer
}

func (f *IPFilter[T]) GetIPCount() uint32 {
	return uint32(f.ipsCount)
}

func (f *IPFilter[T]) inetNtoA(ipRec interface{}) ([]byte, error) {
	var bytes []byte
	var err error

	switch f.version {
	case ip_filter.IPV4:
		bytes, err = f.inetNtoAV4(ipRec.(ip_filter.IPV4FilterRow).IPMin)
	case ip_filter.IPV6:
		bytes, err = f.inetNtoAV6(ipRec.(ip_filter.IPV6FilterRow).IPMin)
	}

	return bytes, err
}

func (f *IPFilter[T]) inetNtoAV4(ipNumber uint32) ([]byte, error) {
	bytes := make([]byte, 4)

	bytes[3] = byte(ipNumber & 0xFF)
	bytes[2] = byte((ipNumber >> 8) & 0xFF)
	bytes[1] = byte((ipNumber >> 16) & 0xFF)
	bytes[0] = byte((ipNumber >> 24) & 0xFF)

	return bytes, nil
}

func (f *IPFilter[T]) inetNtoAV6(ipString string) ([]byte, error) {
	bytes, err := hex.DecodeString(ipString)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func (f *IPFilter[T]) prefixToMask(row interface{}) ([]byte, error) {
	switch f.version {
	case ip_filter.IPV4:
		return f.prefixToMaskV4(row.(ip_filter.IPV4FilterRow).Mask)
	case ip_filter.IPV6:
		return f.prefixToMaskV6(row.(ip_filter.IPV6FilterRow).Mask)
	}
	return []byte{}, fmt.Errorf("invalid version")
}

func (f *IPFilter[T]) prefixToMaskV4(prefix int) ([]byte, error) {
	return f.inetNtoAV4(0xFFFFFFFF << (32 - prefix))
}

func (f *IPFilter[T]) prefixToMaskV6(prefix int) ([]byte, error) {
	maskBytes := make([]byte, 16)
	maskValue := byte(0xFF << (8 - prefix))

	for i := 0; i < 16; i++ {
		if prefix >= 8 {
			maskBytes[i] = 0xFF
			prefix -= 8
		} else {
			maskBytes[i] = maskValue
			break
		}
	}

	return maskBytes, nil
}

func (f *IPFilter[T]) isIPValid(bytes []byte) bool {
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

func (f *IPFilter[T]) isThrottling(row interface{}) (uint, error) {
	switch f.version {
	case ip_filter.IPV4:
		return row.(ip_filter.IPV4FilterRow).IsThrottling, nil
	case ip_filter.IPV6:
		return row.(ip_filter.IPV6FilterRow).IsThrottling, nil
	}
	return 0, fmt.Errorf("invalid version")
}

func (f *IPFilter[T]) addRange(row interface{}, rowType uint) {
	switch f.version {
	case ip_filter.IPV4:
		v4filterRow := row.(ip_filter.IPV4FilterRow)
		f.gm.AddRange4(v4filterRow.IPMin, v4filterRow.IPMax, rowType)
	case ip_filter.IPV6:
		v6filterRow := row.(ip_filter.IPV6FilterRow)
		f.gm.AddRange6(v6filterRow.IPMin, v6filterRow.IPMax, rowType)
	}
}
