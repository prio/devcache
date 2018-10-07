package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Storing in cache directory")
	err := os.Mkdir("cache", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.RawQuery[4:]
		fmt.Printf("Processing: %s\n", url)
		fname := fmt.Sprintf("cache/%X", md5.Sum([]byte(url)))

		if _, err := os.Stat(fname); err != nil {
			body, err := getPage(url, r.Header)
			if err != nil {
				log.Fatal(err)
			}
			// write file contents to cache dir
			if err := ioutil.WriteFile(fname, body, 0644); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("Found it: %s\n", fname)
		}

		content, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(content)
	})

	fmt.Println("Listening on localhost:4321...")
	fmt.Println("Append 'http://localhost:4321?url=' to an API call to cache it")
	http.ListenAndServe(":4321", nil)
}

func getPage(url string, header http.Header) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.Header = header
	// Keep it simple, don't want to deal with gzip etc.
	req.Header["Accept-Encoding"] = []string{"identity"}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
