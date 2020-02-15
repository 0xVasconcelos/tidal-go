package main

import (
	"io/ioutil"

	"github.com/lucaslg26/tidal-go/pkg/tidal"
)

func main() {

	enc, _ := ioutil.ReadFile("input.flac")

	MyMedia := &tidal.Media{
		Media:         enc,
		EncryptionKey: []byte(""),
	}

	MyMedia.GetKey()
	MyMedia.DecryptMedia()

	err := ioutil.WriteFile("output.flac", MyMedia.Media, 0644)

	if err != nil {
		panic(err)
	}

}
