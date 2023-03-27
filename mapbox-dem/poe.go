package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
)

// 计算瓦片行列号
func latLng2Tile(lat, lng float64, zoom int) (int, int) {
	n := 1 << uint(zoom)
	x := int((lng + 180.0) / 360.0 * float64(n))
	y := int((1.0 - (math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0)) / math.Pi)) / 2.0 * float64(n))
	return x, y
}

// 获取指定瓦片的数据并保存为 PNG 文件
func getTile(x, y, zoom int, accessToken string) error {
	url := fmt.Sprintf("https://api.mapbox.com/raster/v1/mapbox.mapbox-terrain-dem-v1/%d/%d/%d.webp%s", zoom, x, y, accessToken)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Referer", "https://docs.mapbox.com/")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if string(body) == "{\"message\":\"Forbidden\"}" || string(body) == "Internal Server Error" {
		fmt.Printf("z:%d,x:%d,y:%d tile not found \n", zoom, x, y)

	} else if len(body) == 0 {
		fmt.Printf("z:%d,x:%d,y:%d body size 0\n", zoom, x, y)

	} else if string(body) == "Tile does not exist" || string(body) == "Out of bounds" {
		fmt.Printf("z:%d,x:%d,y:%d backImg error\n", zoom, x, y)
	} else {
		fileName := fmt.Sprintf("%d-%d-%d.webp", zoom, x, y)
		out, err := os.Create(fileName)
		defer func() {
			_ = out.Close()
		}()
		if err != nil {
			fmt.Printf("%s create file error %s\n", fileName, err)
		} else {
			_, err = out.Write(body)
			if err != nil {
				fmt.Printf("save file error %s\n", err)
			}
		}
	}
	return nil
}

func main() {
	//minLat, minLng, maxLat, maxLng := 37.7045, -122.5149, 37.8324, -122.3569
	minLat, minLng, maxLat, maxLng := 22.1650202, 120.196396717016, 25.391554, 121.7846595
	minZoom, maxZoom := 1, 10
	accessToken := "?sku=101bds94IWmPl&access_token=pk.eyJ1IjoiZXhhbXBsZXMiLCJhIjoiY2p0MG01MXRqMW45cjQzb2R6b2ptc3J4MSJ9.zA2W0IkI0c6KaAhJfk9bWg"

	for zoom := minZoom; zoom <= maxZoom; zoom++ {
		x1, y1 := latLng2Tile(minLat, minLng, zoom)
		x2, y2 := latLng2Tile(maxLat, maxLng, zoom)

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				err := getTile(x, y, zoom, accessToken)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
