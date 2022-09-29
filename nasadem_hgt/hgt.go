package hgt

import (
	"fmt"
	"log"
	"math"
)

// NewHgt instanciates a new HGT object with a file fetcher from US government
func NewHgt(efs ElevationFileStorer) *hgt {
	return &hgt{
		elevationFileStorage: efs,
	}
}

type hgt struct {
	elevationFileStorage ElevationFileStorer
}

func (h *hgt) Get(points []Point) (ptElevations []int32, err error) {
	for _, pt := range points {
		dms := getDMSFromPoint(pt)
		elevationGetter, err := h.elevationFileStorage.Get(dms)
		if err != nil {
			return nil, err
		}

		ele, err := elevationGetter.Get(pt)
		if err != nil {
			if err == errorWrongElevation {
				log.Panicln("Could not grade (wrong elevation). Probably water, will use 0 instead")
			} else {
				return nil, err
			}
		}
		ptElevations = append(ptElevations, ele)
	}
	return
}

// getDMSFromPoint extract the DMS format (e.g. N09E011) from point
func getDMSFromPoint(pt Point) string {
	latPfx := "N"
	if pt.Lat < 0 {
		latPfx = "S"
	}

	lngPfx := "E"
	if pt.Lng < 0 {
		lngPfx = "W"
	}

	return fmt.Sprintf("%s%02d%s%03d", latPfx, int8(math.Floor(pt.Lat)), lngPfx, int8(math.Floor(pt.Lng)))
}
