package crawler

import (
	"fmt"
	"github.com/grokify/html-strip-tags-go"
	"io/ioutil"
	"net/http"
	"strings"
)

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}

func main() {
	resp, err := http.Get("http://www.reuters.com")  //crawling a specific URL here
	if err != nil{
		fmt.Println("http transport error is:", err)
	}

	body, err := ioutil.ReadAll(resp.Body)    //Reading the body from response
	if err != nil {
		fmt.Println("read error is:", err)
		}
	re := strip.StripTags(string(body))		//Removing html Tags from body
	words := wordCount(re)					//Counting words in the string
	for k,v:= range words{
		fmt.Printf("%s : %d \n",k,v)			// Printing the data line by line
	}

}