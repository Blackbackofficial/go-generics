package ip_filter

import (
	"generics/examples/easy/ip_filter"
	"generics/examples/easy/ip_filter/generics"
	"testing"
)

func BenchmarkIPFilterGenericsIPV6(b *testing.B) {
	rows := make([]ip_filter.IPV6FilterRow, 10000)

	for i := range rows {
		rows[i].IPMin = "20010db885a3000000008a2e03707334" // equivalent to 2001:0db8:85a3:0000:0000:8a2e:0370:7334
		rows[i].IPMax = "20010db885a3000000008a2e03707335" // equivalent to 2001:0db8:85a3:0000:0000:8a2e:0370:7335
		rows[i].Mask = 64
		rows[i].IsThrottling = 1
	}

	filter := generics.NewIPFilterProcessor(rows, &ip_filter.Geomap{}, ip_filter.IPV6)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, row := range rows {
			err := filter.Process(row)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}

func BenchmarkIPFilterProcessGenericsIPV4(b *testing.B) {
	rows := make([]ip_filter.IPV4FilterRow, 10000)

	for i := range rows {
		rows[i].IPMin = 3232235776 // equivalent to 192.168.1.0
		rows[i].IPMax = 3232235777 // equivalent to 192.168.1.1
		rows[i].Mask = 24
		rows[i].IsThrottling = 1
	}

	filter := generics.NewIPFilterProcessor(rows, &ip_filter.Geomap{}, ip_filter.IPV4)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, row := range rows {
			err := filter.Process(row)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
