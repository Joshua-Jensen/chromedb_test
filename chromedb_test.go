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

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", false),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("window-size", "50,400"),
	)

	alloc_ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(alloc_ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	err := chromedp.Run(ctx)
	if err != nil {
		panic(err)
	}

	network, err := listen_for_network_event(ctx)
	if err != nil {
		println(err)
	}

	write_network(network)
	if err != nil {
		println(err)
	}
}

func listen_for_network_event(ctx context.Context) {
chromedp.ListenTarget(ctx, func(ev interface{}) {
	switch ev := ev.(type){

	}
	}
})
}
