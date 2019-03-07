package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

func main() {
	enc := md5.New()
	enc.Write([]byte("testfssdfdsfdsfsdfsfsdfdsfsfsdfds"))
	log.Println(len(hex.EncodeToString(enc.Sum(nil))))
	// library.StartServer(":8000", pendaftaran.ApiRoute())
	// log.Printf("miaw")
}
