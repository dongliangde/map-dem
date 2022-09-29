package hgt

import (
	"fmt"
	"log"
	"os"
)

type hgtFileGetter interface {
	getFile(dms string) (*os.File, error)
}

type memcacheElevationStorage struct {
	elevationGettersClosersCache map[string]ElevationPointGetterCloser
	fileGetter                   hgtFileGetter
}

// NewMemcacheNasaElevationFileStorage instanciates a new in memory object holding ElevationPointGetterCloser objects
func NewMemcacheNasaElevationFileStorage(destinationDir, authorization string) *memcacheElevationStorage {
	fileGetter := &nasa30mFile{authorization, destinationDir}

	return &memcacheElevationStorage{
		elevationGettersClosersCache: make(map[string]ElevationPointGetterCloser),
		fileGetter:                   fileGetter,
	}
}

func (m *memcacheElevationStorage) Get(dms string) (ElevationPointGetter, error) {
	if g, ok := m.elevationGettersClosersCache[dms]; ok {
		return g, nil
	}

	log.Println(fmt.Sprintf("Getting file: %s", dms))
	f, err := m.fileGetter.getFile(dms)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Done")

	g := NewStrm1(f)
	m.elevationGettersClosersCache[dms] = g

	return g, nil
}

func (m *memcacheElevationStorage) Close() {
	for _, egc := range m.elevationGettersClosersCache {
		egc.Close()
	}
}
