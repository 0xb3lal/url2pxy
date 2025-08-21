package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

func printLogo() {
	fmt.Println("            _ ____                  ")
	fmt.Println(" _   _ _ __| |___ \\ _ ____  ___   _ ")
	fmt.Println("| | | | '__| | __) | '_ \\ \\/ / | | |")
	fmt.Println("| |_| | |  | |/ __/| |_) >  <| |_| |")
	fmt.Println(" \\__,_|_|  |_|_____| .__/_/\\_\\\\__, |")
	fmt.Println("                   |_|        |___/ ")
	fmt.Println("")
}

func makeRequest(targetURL string, proxyAddr string, wg *sync.WaitGroup) {
	defer wg.Done()

	proxyURL, err := url.Parse("http://" + proxyAddr)
	if err != nil {
		fmt.Printf("\033[31m[!] Invalid proxy: %s\033[0m\n", err)
		return
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		fmt.Printf("\033[31m[-] Failed: %s - Error: %s\033[0m\n", targetURL, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("\033[32m[+] Success: %s - Status Code: %d\033[0m\n", targetURL, resp.StatusCode)
	} else {
		fmt.Printf("\033[33m[!] Error: %s - Status Code: %d\033[0m\n", targetURL, resp.StatusCode)
	}
}

func main() {
	printLogo()

	listPath := flag.String("l", "", "Path to the list of URLs ")
	proxyAddr := flag.String("p", "", "Proxy address")
	threads := flag.Int("t", 10, "Number of threads to use (default 10)")
	flag.Parse()

	if *listPath == "" || *proxyAddr == "" {
		fmt.Println("Usage:")
		fmt.Println("  url2pxy -p [Proxy] -l [urls_list] -t [Threads]")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -p    Proxy address (host:port)")
		fmt.Println("  -l    Path to the list of URLs")
		fmt.Println("  -t    Number of threads to use (default: 10)")
		return
	}

	file, err := os.Open(*listPath)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	sem := make(chan struct{}, *threads)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		if url == "" {
			continue
		}

		sem <- struct{}{}
		wg.Add(1)

		go func(u string) {
			defer func() { <-sem }()
			makeRequest(u, *proxyAddr, &wg)
		}(url)
	}

	wg.Wait()
}
