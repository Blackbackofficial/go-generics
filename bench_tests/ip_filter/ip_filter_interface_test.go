package ip_filter

import (
	"generics/examples/ip_filter"
	_interface "generics/examples/ip_filter/interface"
	"testing"
)

func BenchmarkProcessInterfaceIPv6(b *testing.B) {
	row := ip_filter.IPV6FilterRow{IPMin: "2001:0db8:85a3:0000:0000:8a2e:0370:7334", IPMax: "2001:0db8:85a3:0000:0000:8a2e:0370:7335", Mask: 64, IsThrottling: 1}
	ipFilter := _interface.NewIPFilterProcessor([]_interface.IpFilterRow{row}, &ip_filter.Geomap{}, ip_filter.IPV6)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = ipFilter.Process(&row)
	}
}

func BenchmarkProcessInterfaceIPv4(b *testing.B) {
	row := ip_filter.IPV4FilterRow{
		IPMin:        3232235776,
		IPMax:        3232235777,
		Mask:         24,
		IsThrottling: 1,
	}
	ipFilter := _interface.NewIPFilterProcessor([]_interface.IpFilterRow{row}, &ip_filter.Geomap{}, ip_filter.IPV4)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = ipFilter.Process(&row)
	}
}
