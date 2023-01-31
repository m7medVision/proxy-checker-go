package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type Colors struct {
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	White  string
	Reset  string
}

var color = Colors{
	Red:    "\033[31m",
	Green:  "\033[32m",
	Yellow: "\033[33m",
	Blue:   "\033[34m",
	Purple: "\033[35m",
	Cyan:   "\033[36m",
	White:  "\033[37m",
	Reset:  "\033[0m",
}
var (
	good int
	bad  int
)

func main() {
	fmt.Print(color.Red)
	fmt.Println(" ███████████                                                 █████████  █████                        █████                        ")
	fmt.Println("░░███░░░░░███                                               ███░░░░░███░░███                        ░░███                         ")
	fmt.Println(" ░███    ░███ ████████   ██████  █████ █████ █████ ████    ███     ░░░  ░███████    ██████   ██████  ░███ █████  ██████  ████████ ")
	fmt.Println(" ░██████████ ░░███░░███ ███░░███░░███ ░░███ ░░███ ░███    ░███          ░███░░███  ███░░███ ███░░███ ░███░░███  ███░░███░░███░░███")
	fmt.Println(" ░███░░░░░░   ░███ ░░░ ░███ ░███ ░░░█████░   ░███ ░███    ░███          ░███ ░███ ░███████ ░███ ░░░  ░██████░  ░███████  ░███ ░░░ ")
	fmt.Println(" ░███         ░███     ░███ ░███  ███░░░███  ░███ ░███    ░░███     ███ ░███ ░███ ░███░░░  ░███  ███ ░███░░███ ░███░░░   ░███     ")
	fmt.Println(" █████        █████    ░░██████  █████ █████ ░░███████     ░░█████████  ████ █████░░██████ ░░██████  ████ █████░░██████  █████    ")
	fmt.Println("░░░░░        ░░░░░      ░░░░░░  ░░░░░ ░░░░░   ░░░░░███      ░░░░░░░░░  ░░░░ ░░░░░  ░░░░░░   ░░░░░░  ░░░░ ░░░░░  ░░░░░░  ░░░░░     ")
	fmt.Println("                                              ███ ░███                                                                            ")
	fmt.Println("                                             ░░██████                                                                             ")
	fmt.Println("                                              ░░░░░░                                                                              ")
	fmt.Println(color.Red + "Made By: " + color.Reset + "@majhcc")
	var filename string
	var Timeout int
	var target string
	var threads int
	var mode int
	fmt.Print(color.Blue + "Enter File Name " + color.Yellow + "[default proxies.txt]: " + color.Reset)
	fmt.Scanln(&filename)
	if filename == "" {
		filename = "proxies.txt"
	}
	proxies, err := openProxiesFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(color.Blue + "Enter Timeout " + color.Yellow + "[default 10]: " + color.Reset)
	fmt.Scanln(&Timeout)
	if Timeout == 0 {
		Timeout = 10
	}
	fmt.Print(color.Red + "Enter Target " + color.Yellow + "[default http://google.com]: " + color.Reset)
	fmt.Scanln(&target)
	if target == "" {
		target = "http://google.com"
	}
	fmt.Print(color.Blue + "Which Mode? " + color.Red + "(1) POWER[NOT RECOMMENDED] " + color.Green + " (2) NORMAL [default NORMAL]: " + color.Reset)
	fmt.Scanln(&mode)
	if mode == 0 {
		mode = 2
	}
	x := len(proxies)
	if mode == 1 {
		go func() {
			for {
				fmt.Print(color.Green + "Good proxy: " + color.Reset + fmt.Sprint(good) + color.Reset + " " + color.Red + "Bad proxy: " + color.Reset + fmt.Sprint(bad) + " Checked proxies:" + fmt.Sprint(bad+good) + "/" + fmt.Sprint(x) + "\r")
			}
		}()
		wg := sync.WaitGroup{}
		fmt.Println(color.Red+"Proxies: "+color.Reset, x)
		for _, proxy := range proxies {
			wg.Add(1)
			go func(proxy string) {
				defer wg.Done()
				runner(proxy, Timeout, target)
			}(proxy)
		}
		wg.Wait()
		// fmt.Println(color.Red + "Good: " + color.Reset + strconv.Itoa(good) + color.Red + " Bad: " + color.Reset + strconv.Itoa(bad))
	} else if mode == 2 {
		fmt.Print(color.Blue + "Enter Threads " + color.Yellow + "[default 100]: " + color.Reset)
		fmt.Scanln(&threads)
		if threads == 0 {
			threads = 100
		}
		go func() {
			for {
				fmt.Print(color.Green + "Good proxy: " + color.Reset + fmt.Sprint(good) + color.Reset + " " + color.Red + "Bad proxy: " + color.Reset + fmt.Sprint(bad) + " Checked proxies:" + fmt.Sprint(bad+good) + "/" + fmt.Sprint(x) + "\r")
			}
		}()
		queue := make(chan bool, threads)
		fmt.Println(color.Red+"Proxies: "+color.Reset, x)
		for _, proxy := range proxies {
			queue <- true
			go func(proxy string) {
				defer func() { <-queue }()
				runner(proxy, Timeout, target)
			}(proxy)
		}
		for i := 0; i < cap(queue); i++ {
			queue <- true
		}
	}
}

func checkProxy(proxyUrl string, timeout int, proxType string, target string) bool {
	proxy, _ := url.Parse(proxType + "://" + proxyUrl)

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(timeout) * time.Second,
	}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return false
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func openProxiesFile(filename string) ([]string, error) {
	if file, err := os.Open(filename); err == nil {
		defer file.Close()

		var proxies []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			proxies = append(proxies, scanner.Text())
		}

		return proxies, nil
	} else {
		return []string{}, err
	}
}

func runner(proxy string, Timeout int, target string) {
	if !checkProxy(proxy, Timeout, "http", target) && !checkProxy(proxy, Timeout, "https", target) {
		bad++
	} else {
		good++
		f, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.Write([]byte(proxy + "\n"))
		f.Close()
	}
}
