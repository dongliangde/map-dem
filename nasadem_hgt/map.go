package hgt

import "io"

type ElevationFileStorer interface {
	Get(dms string) (ElevationPointGetter, error)
	Close()
}

type ElevationPointGetter interface {
	Get(pt Point) (int32, error)
}

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ElevationPointGetterCloser interface {
	ElevationPointGetter
	io.Closer
}
