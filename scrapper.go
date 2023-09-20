package main

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

type Property struct {
	address      string
	postCode     string
	price        string
	propertyType string
}

func main() {

	opts := append(
		// select all the elements after the third element
		chromedp.DefaultExecAllocatorOptions[3:],
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	)

	parentCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(parentCtx)
	defer cancel()

	var address, postCode, price, propertyType string

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.propertypal.com/120-moylagh-road-omagh/780845"),
		chromedp.Text("ckZmwy", &address, chromedp.NodeVisible),
		chromedp.Text("drredi", &postCode, chromedp.NodeVisible),
		chromedp.Text("doLooj", &propertyType, chromedp.NodeVisible),
		chromedp.Text("pp-property-price-bold", &price, chromedp.NodeVisible),
	)

	if err != nil {
		panic(err)
	}

	property := Property{
		address:      address,
		postCode:     postCode,
		price:        price,
		propertyType: propertyType,
	}

	fmt.Println(property)

}
