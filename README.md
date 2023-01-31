# Proxy Checker
![Screenshot](./screenshot.png){: .center-image }

## Download Go
To run this code, you need to have Go installed on your machine. You can download it from the official website: https://golang.org/dl/

## Building the code
Clone the repository to your local machine
`git clone https://github.com/majhcc/proxy-checker-go.git`

Navigate to the project directory
`cd proxy-checker-go`

Build the code
`go build ProxyChecker.go`

Run the executable file
`./proxy-checker.exe`
## Usage
The program will prompt you to enter the filename of the proxy list, the timeout for each request, and the target website to test the proxies against.

It will then test each proxy and display the results with a colored output and it will save it inside `output.txt`. The green color indicates a successful connection, while red indicates a failed one.

## Contribution
If you want to contribute to this project, feel free to open a pull request or issue.

## License
This tool is under MIT License.