package ip_filter

import "fmt"

const (
	IPV4 = 4
	IPV6 = 6
)

type IPV6FilterRow struct {
	IPMin        string
	IPMax        string
	Mask         int
	IsThrottling uint
}

type IPV4FilterRow struct {
	IPMin        uint32
	IPMax        uint32
	Mask         int
	IsThrottling uint
}

type Geomap struct {
	Range string
}

func (g *Geomap) AddRange4(min, max uint32, rType uint) {
	g.Range = fmt.Sprintf("IPv4 Range: Min: %v, Max: %v, Type: %v\n", min, max, rType)
}

func (g *Geomap) AddRange6(min, max string, rType uint) {
	g.Range = fmt.Sprintf("IPv6 Range: Min: %v, Max: %v, Type: %v\n", min, max, rType)

}

// --------- interface ---------

func (row IPV4FilterRow) GetIPMin() interface{} {
	return row.IPMin
}

func (row IPV4FilterRow) GetIPMax() interface{} {
	return row.IPMax
}

func (row IPV4FilterRow) GetMask() int {
	return row.Mask
}

func (row IPV4FilterRow) GetIsThrottling() uint {
	return row.IsThrottling
}

func (row IPV6FilterRow) GetIPMin() interface{} {
	return row.IPMin
}

func (row IPV6FilterRow) GetIPMax() interface{} {
	return row.IPMax
}

func (row IPV6FilterRow) GetMask() int {
	return row.Mask
}

func (row IPV6FilterRow) GetIsThrottling() uint {
	return row.IsThrottling
}
