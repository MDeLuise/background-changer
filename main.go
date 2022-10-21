package main

import (
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/hbagdi/go-unsplash/unsplash"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

const (
	baseURL string = "https://images.unsplash.com"
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

	imgName := flag.String("name", strconv.FormatInt(time.Now().Unix(), 10)+".png", "Name of the downloaded image")
	downloadDirectory := flag.String("directory", ".", "Directory used to download the images")
	collectionIDsString := flag.String("collections", "880012", "Collection IDs to use as download filter")
	removeOldImages := flag.Bool("clean", true, "Remove old downloaded images")
	flag.Parse()
	collectionIDs := stringToArrayOfInt(collectionIDsString)

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "Client-ID " + key},
	)
	client := oauth2.NewClient(oauth2.NoContext, ts)
	unsplash := unsplash.New(client)
	randomID := getRandomPhotoId(unsplash, collectionIDs)
	downloadUrl := baseURL + getDownloadURL(unsplash, randomID)
	if *removeOldImages {
		removeAllImageInDirectory(*downloadDirectory)
	}
	if err := downloadFile(downloadUrl, *downloadDirectory+"/"+*imgName); err != nil {
		log.Fatal(err)
	}
}

func stringToArrayOfInt(stringArray *string) *[]int {
	result := make([]int, 1)
	for _, elStr := range strings.Split(*stringArray, " ") {
		if elInt, err := strconv.Atoi(elStr); err != nil {
			log.Fatal(err)
		} else {
			result = append(result, elInt)
		}
	}
	return &result
}

func getDownloadURL(client *unsplash.Unsplash, id *string) string {
	url, _, err := client.Photos.DownloadLink(*id)
	if err != nil {
		log.Fatal(err)
	}
	return url.URL.Path
}

func getRandomPhotoId(client *unsplash.Unsplash, collectionIDs *[]int) *string {
	topics := make([]string, 1)
	topics = append(topics, "bo8jQKTaE0Y")
	opts := &unsplash.RandomPhotoOpt{
		Orientation:   "landscape",
		CollectionIDs: *collectionIDs,
		TopicIDs:      topics,
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

func removeAllImageInDirectory(dirPath string) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{dirPath, d.Name()}...))
	}
	return nil
}
