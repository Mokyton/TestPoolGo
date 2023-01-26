package main

import (
	"errors"
	"ex00/gen/restapi"
	"ex00/gen/restapi/operations"
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"log"
)

var (
	portFlag   = flag.Int("port", 3333, "Port to run this service on")
	candyPrice = make(map[string]int64)
)

type good struct {
	name  string
	price int64
}

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCandyServerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag
	priceList := getPriceList()
	fillMap(priceList)
	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(
		func(params operations.BuyCandyParams) middleware.Responder {
			candy := struct {
				name  string
				count int64
				money int64
			}{
				swag.StringValue(params.Order.CandyType),
				swag.Int64Value(params.Order.CandyCount),
				swag.Int64Value(params.Order.Money),
			}
			fmt.Println(candy)
			_, ok := candyPrice[candy.name]
			if !ok {
				return operations.NewBuyCandyBadRequest().WithPayload(
					&operations.BuyCandyBadRequestBody{Error: fmt.Sprintf("Incorrect input: we don't have %s", candy.name)})
			}

			if candy.count <= 0 || candy.money <= 0 {
				return operations.NewBuyCandyBadRequest().WithPayload(
					&operations.BuyCandyBadRequestBody{Error: "Incorrect input: money or count can't be zero or negative"})
			}

			change, err := calculateChange(candy.count, candyPrice[candy.name], candy.money)
			if err != nil {
				return operations.NewBuyCandyPaymentRequired().WithPayload(
					&operations.BuyCandyPaymentRequiredBody{Error: err.Error()})
			}

			return operations.NewBuyCandyCreated().WithPayload(
				&operations.BuyCandyCreatedBody{Change: change, Thanks: "Thank you!"})
		})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func getPriceList() []good {
	return []good{
		{name: "CE", price: 10},
		{name: "AA", price: 15},
		{name: "NT", price: 17},
		{name: "DE", price: 21},
		{name: "YR", price: 23},
	}
}

func fillMap(priceList []good) {
	for _, v := range priceList {
		candyPrice[v.name] = v.price
	}
}

func calculateChange(count, price, money int64) (int64, error) {
	if money < price*count {
		return 0, errors.New("Not enough money for goods")
	}
	return money - (price * count), nil
}
