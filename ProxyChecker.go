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
var counter int

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
	fmt.Println(color.Red + "Made By: " + color.Reset + "Mohammed AlJahwari")
	var filename string
	var Timeout int
	var target string
	fmt.Print("Enter FileName: ")
	fmt.Scanln(&filename)
	fmt.Print("Enter Timeout: ")
	fmt.Scanln(&Timeout)
	fmt.Print("Enter Target: ")
	fmt.Scanln(&target)
	proxies, err := openProxiesFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(color.Red+"Proxies: "+color.Reset, len(proxies))
	var wg sync.WaitGroup
	for _, proxy := range proxies {
		wg.Add(1)
		go runner(proxy, Timeout, target, &wg)
	}
	wg.Wait()
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
	req, _ := http.NewRequest("GET", target, nil)
	resp, _ := client.Do(req)
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
func runner(proxy string, Timeout int, target string, wg *sync.WaitGroup) {
	defer wg.Done()
	if !checkProxy(proxy, Timeout, "http", target) && !checkProxy(proxy, Timeout, "https", target) {
		counter++
		fmt.Println(color.Red + "Proxy is not working" + color.Reset + " " + proxy)
	} else {
		counter++
		fmt.Println(color.Green + "Proxy is working" + color.Reset + " " + proxy)
		f, _ := os.OpenFile("workingProxies.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.Write([]byte(proxy + "\n"))
		f.Close()
	}
	fmt.Print(color.Yellow + "Checking proxy " + color.Reset + fmt.Sprint(counter) + color.Reset + "\r")
}
