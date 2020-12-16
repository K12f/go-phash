package phash

import (
	"fmt"
	"math"
)

//缩小尺寸 为了后边的步骤计算简单些
//简化色彩 将图片转化成灰度图像，进一步简化计算量
//计算DCT 计算图片的DCT变换，得到64*64的DCT系数矩阵。
//缩小DCT 虽然DCT的结果是64*64大小的矩阵，但我们只要保留左上角的64*64的矩阵，这部分呈现了图片中的最低频率。
//计算平均值 如同均值哈希一样，计算DCT的均值。
//计算hash值 根据64*64的DCT矩阵，设置0或1的64位的hash值，大于等于DCT均值的设为”1”，小于DCT均值的设为“0”。组合在一起，就构成了一个64位的整数，这就是这张图片的指纹

func TransMa() [64][64]float64 {
	var ma [64][64]float64

	var l = len(ma)
	var a float64
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == 0 {
				a = math.Sqrt(float64(1) / float64(l))
			} else {
				a = math.Sqrt(float64(2) / float64(l))
			}
			ma[i][j] = a * math.Cos(math.Pi*float64(i)*(float64(j)+0.5)/float64(l))
		}
	}

	return ma
}

func DCT(dctMa [64][64]float64, dctMap [64][64]float64) [64][64]float64 {
	var l = len(dctMa)
	var t = 0.0
	var dctMapTemp [64][64]float64

	// 相当于A*I
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			t = 0
			for k := 0; k < l; k++ {
				t += dctMa[i][k] * dctMap[k][j]
			}
			dctMapTemp[i][j] = t
		}
	}
	// 相当于（A*I）后再*A‘
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			t = 0
			for k := 0; k < l; k++ {
				t += dctMapTemp[i][k] * dctMa[j][k]
			}
			dctMap[i][j] = math.Round(t)
		}
	}
	return dctMap
}

func average(dctMap [8][8]float64) float64 {
	var l = len(dctMap)
	var sum = 0.0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			sum += dctMap[i][j]
		}
	}
	return sum / (float64(l) * float64(l))
}

func compare(average float64, dctMap [64][64]float64) (hash []int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if dctMap[i][j] >= average {
				hash = append(hash, 1)
			} else {
				hash = append(hash, 0)
			}
		}
	}
	return
}

func Phash(path string) (hash []int) {
	pixel := Matrix64(path)

	fmt.Println(pixel)
	ma := TransMa()

	dctMap := DCT(ma, pixel)

	//fmt.Println(dctMap)
	var dctMap8 [8][8]float64

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			dctMap8[i][j] = dctMap[i][j]
		}
	}

	average := average(dctMap8)
	hash = compare(average, dctMap)
	return hash
}
