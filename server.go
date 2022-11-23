package main

import (
	"github.com/conekta/conekta-go"
	"github.com/conekta/conekta-go/order"
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func createCheckout() (error, string) {
	conekta.APIKey = "key_y2TV3l4GLgqOwc6POUfFGP1"

	customerParams := &conekta.CustomerParams{
		ID: "cus_2shKUY6Rvwg1nSEiq",
	}

	lineItemParams := &conekta.LineItemsParams{
		Name:      "Naranjas Robadas",
		UnitPrice: 10000,
		Quantity:  2,
	}

	orderCheckoutParams := conekta.OrderCheckoutParams{
		Type:                  "HostedPayment",
		ExpiresAt:             1669323410,
		AllowedPaymentMethods: []string{"cash", "card", "bank_transfer"},
		SuccessUrl:            "https://www.mysite.com/payment/confirmation",
		FailureUrl:            "https://www.mysite.com/payment/failure",
	}

	orderParams := &conekta.OrderParams{}
	orderParams.Currency = "MXN"
	orderParams.CustomerInfo = customerParams
	orderParams.LineItems = append(orderParams.LineItems, lineItemParams)
	orderParams.Checkout = &orderCheckoutParams

	ord, err := order.Create(orderParams)
	if err != nil {
		return err, ""
	}
	return nil, ord.Checkout.Url
}

type Response struct {
	Url string `json:"url"`
}

func main() {
	e := echo.New()
	e.POST("/checkout", func(c echo.Context) error {
		err, url := createCheckout()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, Response{Url: url})
	})
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8080"))
}
