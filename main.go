package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hbagdi/go-unsplash/unsplash"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

const (
	baseURL string = "https://images.unsplash.com"
	imgName string = "background.png"
)

var (
	collectionIDs []int = []int{880012}
)

func main() {
	env := os.Getenv("GO_ENV")
	if "" == env {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	key := os.Getenv("key")
	if "" == key {
		log.Fatal("value of key empty")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "Client-ID " + key},
	)
	client := oauth2.NewClient(oauth2.NoContext, ts)
	unsplash := unsplash.New(client)
	randomID := getRandomPhotoId(unsplash)
	downloadUrl := baseURL + getDownloadURL(unsplash, randomID)
	downloadDirectory := "."
	if len(os.Args) > 1 {
		downloadDirectory = os.Args[1]
	}
	if err := downloadFile(downloadUrl, downloadDirectory+"/"+imgName); err != nil {
		log.Fatal(err)
	}
}

func getDownloadURL(client *unsplash.Unsplash, id *string) string {
	url, _, err := client.Photos.DownloadLink(*id)
	if err != nil {
		log.Fatal(err)
	}
	return url.URL.Path
}

func getRandomPhotoId(client *unsplash.Unsplash) *string {
	opts := &unsplash.RandomPhotoOpt{
		Orientation:   "landscape",
		CollectionIDs: collectionIDs,
	}
	randomPhoto, _, err := client.Photos.Random(opts)
	if err != nil {
		log.Fatal(err)
	}
	return (*randomPhoto)[0].ID
}

func downloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received response code: " + string(response.StatusCode))
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = io.Copy(file, response.Body); err != nil {
		return err
	}
	return nil
}
