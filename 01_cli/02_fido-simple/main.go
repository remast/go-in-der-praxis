package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	urls := []string{
		"https://go.dev/dl/go1.22.2.src.tar.gz",
		"https://go.dev/dl/go1.22.2.darwin-amd64.tar.gz",
		"https://go.dev/dl/go1.22.2.darwin-amd64.pkg",
		"https://go.dev/dl/go1.22.2.windows-amd64.zip",
	}

	ctx := context.Background()
	downloadURLs(ctx, urls)
}

func downloadURLs(ctx context.Context, urls []string) {
	// 1. Channel für Download Jobs
	jobs := make(chan int, len(urls))

	// 2. Go Routinen für Downloads starten
	for _, url := range urls {
		go func() {
			downloadURL(ctx, url)
			jobs <- 1
		}()
	}

	// 3. Warten auf Ende aller Go Routinen
	for range len(urls) {
		<-jobs
	}
}

func downloadURL(ctx context.Context, url string) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Erro", err)
		return
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		os.Remove(filename)
		log.Println("Erro", err)
	}
}
