package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	APP_KEY    = "eddyciy"
	APP_SECRET = "go-programming-tour-book"
)

type AccessToken struct {
	Token string `json:"token"`
}

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth",
		APP_KEY, APP_SECRET)
	body, err := a.HttpGet(ctx, url)
	if err != nil {
		return "", err
	}

	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	body, err := a.HttpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *API) HttpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
