package main

import (
	"fmt"
	"golang-yt-microservices/sdk/client"
	"golang-yt-microservices/sdk/client/products"
	"testing"
)

func TestOurClient(t* testing.T) {
	config := client.DefaultTransportConfig().WithHost("localhost:9090")
	testClient := client.NewHTTPClientWithConfig(nil, config)
	params := products.NewListProductsParams()

	products,err := testClient.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", products.GetPayload()[0])
	t.Fail()
}