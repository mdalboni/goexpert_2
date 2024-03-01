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

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		service := services.NewBrazilApiZipCodeService()
		result, url, err := service.GetZipCode(ctx, "01153000")
		if err != nil {
			ch1 <- err.Error()
			return
		}
		ch1 <- fmt.Sprintf("%s\n%v", url, result)
	}()

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		service := services.NewViaCepZipCodeService()
		result, url, err := service.GetZipCode(ctx, "01153000")
		if err != nil {
			ch2 <- err.Error()
			return
		}
		ch2 <- fmt.Sprintf("%s\n%v", url, result)
	}()

	// Receive data from ch2
	select {
	case res := <-ch1:
		fmt.Println(res)
	case res := <-ch2:
		fmt.Println(res)
	}
}
