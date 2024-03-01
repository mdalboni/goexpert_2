package services

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ZipCodeService interface {
	GetZipCode(ctx context.Context, zipCode string) (map[string]any, string, error)
}

type brazilApiZipService struct {
	url string
}

func NewBrazilApiZipCodeService() *brazilApiZipService {
	return &brazilApiZipService{
		url: "https://brasilapi.com.br/api/cep/v1/",
	}
}

func (s *brazilApiZipService) GetZipCode(ctx context.Context, zipCode string) (map[string]any, string, error) {
	url := s.url + zipCode
	result, err := performRequest(ctx, url)
	if err != nil {
		return nil, "", err
	}
	return result, url, nil
}

type viaCepZipService struct {
	url string
}

func NewViaCepZipCodeService() *viaCepZipService {
	return &viaCepZipService{
		url: "http://viacep.com.br/ws/",
	}
}

func (s *viaCepZipService) GetZipCode(ctx context.Context, zipCode string) (map[string]any, string, error) {
	url := s.url + zipCode + "/json"
	result, err := performRequest(ctx, url)
	if err != nil {
		return nil, "", err
	}
	return result, url, nil
}

func performRequest(ctx context.Context, url string) (map[string]any, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return result, nil
}
