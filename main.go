package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mrturkmencom/youtubeto/client"
)

const (
	LIST_OF_PLAYLIST_URLS     = "./resource/playlist-list.csv"
	OLD_LIST_OF_PLAYLIST_URLS = "./resource/old-playlist-list.csv"
)

func main() {
	cli := client.New()

	var oldFile, err = os.OpenFile(OLD_LIST_OF_PLAYLIST_URLS, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer oldFile.Close()
	csvfile, err := os.Open(LIST_OF_PLAYLIST_URLS)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	writer := csv.NewWriter(oldFile)
	defer writer.Flush()
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if err := cli.YoutubeDL.DownloadWithOutputName(record[0], record[1]); err != nil {
			panic(err)
		}

		if err := cli.Tar.CompressWithPIGZ(fmt.Sprintf("%s.tar.gz", record[0]), record[0]); err != nil {
			panic(err)
		}
		writer.Write([]string{record[0], record[1]})
	}

	if err := os.Remove("./resource/playlist-list.csv"); err != nil {
		panic(err)
	}

	createFile("./resource/playlist-list.csv")

}

func createFile(path string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return
		}
		defer file.Close()
	}
}
