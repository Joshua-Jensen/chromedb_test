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

	var request network.EventRequestWillBeSentExtraInfo

	err := chromedp.Run(ctx, network.Enable(), chromedp.Navigate("https://www.albertsons.com/shop/search-results.html?q=egg"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	listen_for_network_event(ctx)

}

func listen_for_network_event(ctx context.Context) {

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSentExtraInfo:
			if len(ev.Headers) > 0 {
				fmt.Println(ev.Headers)
			}
		}
	})
}
