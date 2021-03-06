package main

import (
	"bufio"
	// "errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func osMain() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func main() {
	dnssec := flag.Bool("dnssec", false, "Request DNSSEC records")
	port := flag.String("port", "53", "Set the query port")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [name ...]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Println("dnssec: ", *dnssec)
	fmt.Println("port: ", *port)
	flag.PrintDefaults()

	cmd := exec.Command("/bin/ls", "-l")
	buf1, err := cmd.Output()
	if err != nil {
		return
	}
	fmt.Println("output is : %s ", string(buf1))

	r1, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	b, err := ioutil.ReadAll(r1.Body)
	r1.Body.Close()
	if err == nil {
		fmt.Printf("%s", string(b))
	}

	return

	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}
}
