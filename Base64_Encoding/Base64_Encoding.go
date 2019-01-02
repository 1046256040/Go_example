package main

import (
	"encoding/base64"
	"fmt"
)

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encoder(src []byte) string {
	return coder.EncodeToString(src)
}

func base64Decoder(src string) ([]byte, error) {
	return coder.DecodeString(src)
}

func main() {
	data := "wo shi ni ba ba !"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	tEnc := base64Encoder([]byte(data))
	fmt.Println(tEnc)

	tDec, _ := base64Decoder(tEnc)
	fmt.Println(string(tDec))
}
