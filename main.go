package main 

import (
 "fmt"
 "net/url"
 "bufio"
 "os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		ToEncodeValue := scanner.Text()
		encodedValue := url.QueryEscape(ToEncodeValue)
		fmt.Println(encodedValue)
	}
}


