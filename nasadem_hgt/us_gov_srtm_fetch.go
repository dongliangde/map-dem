package hgt

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const nasa30mSrtmURL = "https://e4ftl01.cr.usgs.gov/MEASURES/SRTMGL1.003/2000.02.11/%s.SRTMGL1.hgt.zip"

type nasa30mFile struct {
	authorization  string
	destinationDir string
}

func (n *nasa30mFile) getFile(dms string) (*os.File, error) {
	f, err := os.Open(fmt.Sprintf(".%s/%s.hgt", n.destinationDir, dms))
	if err == nil {
		return f, err
	}

	if err = n.download(dms); err != nil {
		return nil, err
	}

	if err = n.unzip(dms); err != nil {
		return nil, err
	}

	return os.Open(fmt.Sprintf(".%s/%s.hgt", n.destinationDir, dms))
}

func (n *nasa30mFile) download(dms string) error {
	out, err := os.Create(fmt.Sprintf(".%s/%s.hgt.zip", n.destinationDir, dms))
	defer out.Close()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(nasa30mSrtmURL, dms), nil)
	req.Header.Set("Authorization", n.authorization)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil
}

func (n *nasa30mFile) unzip(dms string) error {
	zipFilename := fmt.Sprintf(".%s/%s.hgt.zip", n.destinationDir, dms)
	zipReader, err := zip.OpenReader(zipFilename)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// Archive is supposed to have just a single hgt file
	if len(zipReader.Reader.File) != 1 {
		return errors.New("wrong hgt zip formats")
	}

	file := zipReader.Reader.File[0]
	if file.FileInfo().IsDir() {
		return errors.New("wrong hgt zip format")
	}

	zippedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	extractedFilePath := filepath.Join(
		n.destinationDir,
		file.Name,
	)

	outputFile, err := os.OpenFile(
		"."+extractedFilePath,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		file.Mode(),
	)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if _, err = io.Copy(outputFile, zippedFile); err != nil {
		return err
	}

	return nil
}
