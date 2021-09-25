/**
	* hope it helps u 										*
	* By Nikita Vtorushin<n.vtorushin@inbox.ru> 			*
	* @nikitavoryet 										*
	* GoLang search admin panel (phpmyadmin or any			*
**/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var admins, _ = readLines("adminpath.txt")
var (
	domain   string
	protocol string
)

func header1() {
	print("coded by @nikitavoryet")
	fmt.Println("\n")
}

func header2() {
	fmt.Println(strings.Repeat("-", 74))
	fmt.Print("|   ")
	fmt.Print(" status    ")
	fmt.Print("|                          ")
	fmt.Print("path url")
	fmt.Println("                       |")
	fmt.Println(strings.Repeat("-", 74))
}

func main() {
	header1()
	fmt.Print("enter domain >> ")
	fmt.Scan(&domain)
	if !checkProtocol(domain) {
		fmt.Println("choose the protocol:\nhttps - 1\nhttp - 2")
		fmt.Print(">> ")
		fmt.Scan(&protocol)

		if protocol == "1" {
			protocol = "https://"
		} else {
			protocol = "http://"
		}

		domain = protocol + domain
	}

	header2()
	now := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	_, err := os.Create(now + "_output.txt")
	if err != nil {
		panic(err)
	}
	fs, err := os.OpenFile(now+"_output.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	for _, i := range admins {
		_domain := domain + "/" + i
		resp, _ := http.Get(_domain)
		if resp.StatusCode != 404 {
			fmt.Print("      OK       |     ")
			if _, err := fs.WriteString(_domain + "\n"); err != nil {
				panic(err)
			}
		} else {
			fmt.Print("    failed     ")
			fmt.Print("|     ")
		}
		fmt.Println(domain + "/" + i)
	}
	fmt.Println(strings.Repeat("-", 74))
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func checkProtocol(domain string) bool {
	var re = regexp.MustCompile(`http|https`)
	if !re.MatchString(domain) {
		return false
	} else {
		return true
	}
}
