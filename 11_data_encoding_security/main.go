package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

// 📂 11_data_encoding_security
// 🖥️ BASE64 ENCODING & DECODING


func main() {
	// ENCODING
	data:="Learning about Base64 ENCODING & DECODING"
	encoded:=base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("ENCODED:",encoded)

	// DECODING
	encodedStr:="TGVhcm5pbmcgYWJvdXQgQmFzZTY0IEVOQ09ESU5HICYgREVDT0RJTkc="
	decodedStr,err:=base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		log.Fatal("ERROR: ",err.Error())
	}

	if string(decodedStr) != data{
		log.Fatal("⚠️ ERR : DECODED STR. DOESN'T MATCH ENCODED DATA!")
	}

	fmt.Println("DECODED:",string(decodedStr))

}

// $ go run main.go
// ENCODED: TGVhcm5pbmcgYWJvdXQgQmFzZTY0IEVOQ09ESU5HICYgREVDT0RJTkc=
// DECODED: Learning about Base64 ENCODING & DECODING
