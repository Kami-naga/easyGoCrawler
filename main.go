package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		//check & get correct charset decoding approach
		e := determineEncoding(resp.Body)
		utf8Reader := transform.NewReader(resp.Body,
			e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", all)
	} else {
		fmt.Println("response is not OK, status code: ", resp.StatusCode)
	}
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(nil)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
