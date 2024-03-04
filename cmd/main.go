package main

import (
	"context"
	"fmt"
	"goexpert_2/public/services"
	"time"
)

func main() {
	// Create two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	go brasilAPIRequest(ch1)

	go viaCepRequest(ch2)

	// Receive data from ch2
	select {
	case res := <-ch1:
		fmt.Println(res)
	case res := <-ch2:
		fmt.Println(res)
	}
}

func brasilAPIRequest(ch1 chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for {

		service := services.NewBrazilApiZipCodeService()
		result, url, _ := service.GetZipCode(ctx, "01153000")
		select {
		case <-ctx.Done():
			ch1 <- "Request timeout"
		default:
			if result != nil {
				ch1 <- fmt.Sprintf("%s\n%v", url, result)
				return
			}
		}
	}

}

func viaCepRequest(ch2 chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for {

		service := services.NewViaCepZipCodeService()
		result, url, _ := service.GetZipCode(ctx, "01153000")
		select {
		case <-ctx.Done():
			ch2 <- "Request timeout"
		default:
			if result != nil {
				ch2 <- fmt.Sprintf("%s\n%v", url, result)
				return
			}
		}
	}
}
