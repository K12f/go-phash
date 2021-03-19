package phash

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"strings"
)

func httpGetOriImage(path string) (io.ReadCloser, error) {
	rep, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	return rep.Body, err
}

// https://zh.wikipedia.org/zh/YUV
func Matrix64(path string) (matrix64 [64][64]float64) {
	var src image.Image
	var err error
	// 是否是远程图片
	if strings.Contains(path, "http") {
		file, err := httpGetOriImage(path)
		if err != nil {
			log.Fatalf("failed to read image: %v", err)
		}
		src, err = imaging.Decode(file)
	} else {
		src, err = imaging.Open(path, imaging.AutoOrientation(true))
	}
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	//缩小尺寸 为了后边的步骤计算简单些
	src = imaging.Resize(src, 64, 64, imaging.Lanczos)

	//for w := 0; w <= src.Bounds().Max.X; w++ {
	//    for h := 0; h <= src.Bounds().Max.Y; h++ {
	//        fmt.Print(src.At(w, h))
	//    }
	//    fmt.Println("")
	//}

	gray := imaging.Grayscale(src)

	width := gray.Rect.Max.X
	height := gray.Rect.Max.Y
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			pointColor := gray.At(w, h)
			r := pointColor.(color.NRGBA).R
			g := pointColor.(color.NRGBA).G
			b := pointColor.(color.NRGBA).B

			// RGB 转化 YUV 的公式（经过 PAL制式 CRT伽玛校正）如下：
			// Y = 0.299R’ + 0.587G’ + 0.114B'
			pixel := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			//fmt.Println(r)
			matrix64[h][w] = pixel
		}
	}
	return
}
