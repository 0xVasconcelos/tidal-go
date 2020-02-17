package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/lucaslg26/tidal-go/pkg/tidal"
)

func main() {

	uid, _ := strconv.Atoi(os.Getenv("USER_ID"))

	Tidal := &tidal.Client{
		Token:        "4zx46pyr9o8qZNRw",
		Username:     os.Getenv("TIDAL_USERNAME"),
		Password:     os.Getenv("TIDAL_PASSWORD"),
		SoundQuality: "LOSSLESS",
		User: &tidal.User{
			UserID:      uid,
			SessionID:   os.Getenv("SESSION_ID"),
			CountryCode: "BR",
		},
	}

	//Tidal.Login()

	fmt.Printf("%v\n", Tidal.User)

	// search, _ := Tidal.Search("Garden Of The Titans")

	// fmt.Println("Artists:")
	// for _, v := range search.Artists.Items {
	// 	fmt.Printf("%s\n", v.Name)
	// }

	// fmt.Println("Playlists:")
	// for _, v := range search.Playlists.Items {
	// 	fmt.Printf("%s\n", v.Title)
	// }

	// fmt.Println("Albums:")
	// for _, v := range search.Albums.Items {
	// 	fmt.Printf("%s\n", v.Title)
	// }

	// fmt.Println("Tracks:")
	// for _, v := range search.Tracks.Items {
	// 	fmt.Printf("%s\n", v.Title)
	// }

	track, err := Tidal.GetStream(&tidal.Track{ID: 1267842})

	if err != nil {
		fmt.Println(err)
	}

	media, err := Tidal.GetMedia(track)

	if err != nil {
		fmt.Println(err)
	}

	media.GetKey()

	media.DecryptMedia()

	err = ioutil.WriteFile("media/output.flac", media.Media, 0644)

	if err != nil {
		panic(err)
	}

}
