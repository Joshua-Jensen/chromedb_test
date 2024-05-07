package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func write_network(m string) {
	f, err := os.Create("network.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	l, err := f.WriteString(m)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println("wrote to file ", f.Name(), "network requests: ", l)
}

func main() {
ctx, cancel := chromedp.NewContext(
	context.Background(),
	chromedp.WithLogf(log.Printf),
	chromedp.Headless(false),
)

}


