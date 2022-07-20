package sclient

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

var DefaultClient = &Client{http.DefaultClient, nil, "", ""}

func Get(ctx context.Context, url string, headers ...map[string]string) (*http.Response, []byte, error) {
	return DefaultClient.Get(ctx, url, headers...)
}

func Post(ctx context.Context, url string, contentType string, body io.Reader, headers ...map[string]string) (*http.Response, []byte, error) {
	return DefaultClient.Post(ctx, url, contentType, body, headers...)
}

func PostForm(ctx context.Context, url string, data url.Values, headers ...map[string]string) (*http.Response, []byte, error) {
	return DefaultClient.PostForm(ctx, url, data, headers...)
}
