package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/cdproto/network"
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
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

var text string
	err := chromedp.Run(ctx, network.Enable(), chromedp.Navigate("https://www.albertsons.com/shop/search-results.html?q=egg"),chromedp.TextContent(`span.sr-only`, &text, chromedp.NodeVisible),)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
fmt.Println(text)

}

