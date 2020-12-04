package phash

import (
	"fmt"
	"testing"
)

func TestPhash(t *testing.T) {
	//Phash("https://oss.mokalh.com/copywriting/outerside/images/LxyxCICDkybIOXn3RzzVnYaoRQQDmzU3UeFagWUc.jpeg")
	hash1 := Phash("./images/image1.jpeg")
	hash2 := Phash("./images/image2.jpeg")
	fmt.Println(hash1)
	fmt.Println(hash2)
	distance := HammingDistance(hash1, hash2)
	fmt.Printf("%d,相似度:%f%%", distance, (1-float64(distance)/64)*100)
}
