package main

import (
	"bytes"
	"fmt"
	"golang.org/x/image/webp"
	"image"
	"io/ioutil"
	"log"
	"math"
)

func main() {
	data, err := ioutil.ReadFile("mapbox-dem/6845.webp")
	if err != nil {
		log.Println(err)
	}
	// 珠峰经纬度
	//lng := 86.92498
	//lat := 27.98812
	// 香炉峰经纬度
	//lng := 116.1785
	//lat := 39.99215
	lng := 85.77917
	lat := 28.3525
	// 图层
	zoom := 14
	// 获取瓦片地址
	x := ((lng + 180.0) / 360.0) * math.Exp2(float64(zoom))
	y := (1 - math.Asinh(math.Tan(lat*math.Pi/180))/math.Pi) * math.Exp2(float64(zoom-1))
	// 可以直接访问mapbox官网获取webp数据
	// 计算rgba坐标
	pixelIntX := int(math.Floor(math.Mod(x*514, 514)))
	pixelIntY := int(math.Floor(math.Mod(y*514, 514)))
	// 解析图片
	img, err := webp.Decode(bytes.NewBuffer(data))
	// 读取图片RGBA值
	colorRGB := img.(image.Image).At(pixelIntX, pixelIntY)
	r, g, b, _ := colorRGB.RGBA()
	height := -10000 + float64(int(uint8(r))*256*256+int(uint8(g))*256+int(uint8(b)))*0.1
	fmt.Println(height)
}
