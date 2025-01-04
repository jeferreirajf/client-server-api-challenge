package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ResponseDto struct {
	Quotation struct {
		ID           string `json:"id"`
		CurrencyFrom string `json:"currency_from"`
		ValueFrom    string `json:"value_from"`
		CurrencyTo   string `json:"currency_to"`
		ValueTo      string `json:"value_to"`
		CreatedAt    string `json:"created_at"`
	} `json:"quotation"`
}

func main() {

	url := "http://localhost:8080/cotacao"

	responseDto, err := getQuotation(url)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%v\n", *responseDto)

	filePath := "cotacao.txt"

	saveToFile(filePath, responseDto)
}

func getQuotation(url string) (*ResponseDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond) // 300 milliseconds
	defer cancel()

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		panic(err)
	}

	request = request.WithContext(ctx)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("Status: %v\nMessage: %s", response.Status, body)
		return nil, fmt.Errorf("%s", errorMsg)
	}

	responseDto := ResponseDto{}

	err = json.Unmarshal(body, &responseDto)

	if err != nil {
		panic(err)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
		return nil, ctx.Err()
	default:
		fmt.Println("Success")
		return &responseDto, nil
	}
}

func saveToFile(filePath string, response *ResponseDto) {

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	n, err := file.WriteString(fmt.Sprintf("Dólar: %#v\n", response.Quotation.ValueTo))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Escrevemos %v bytes!\n", n)
}
