package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
)

func OpenImg(sPath string) image.Image {
	f, _ := os.Open(sPath)
	defer f.Close()
	img, fmtName, _ := image.Decode(f)
	fmt.Printf("%v\n", fmtName)
	return img
}

func WriteImage(w http.ResponseWriter, img *image.Image) {
	basee64 := Image2Base64(img)
	w.Write(basee64)
}

func Image2Base64(img *image.Image) (ans []byte) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("unable to encode image.")
	}
	ans = make([]byte, base64.StdEncoding.EncodedLen(len(buffer.Bytes())))
	base64.StdEncoding.Encode(ans, buffer.Bytes())
	return ans
}
