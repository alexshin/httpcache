package signedurl

import (
	"net/http"
	"net/http/httptest"
	url2 "net/url"
	"testing"
)

func TestGcsGenerateBasicCacheKey(t *testing.T) {
	req := &http.Request{
		Method: http.MethodGet,
		URL:    &url2.URL{RawQuery: "foo=bar&baz=qux", Scheme: "http", Host: "localhost"},
	}

	key1 := GcsGenerateCacheKey(req)
	if key1 != "http://localhost?baz=qux&foo=bar" {
		t.Fatalf("key1 = %s", key1)
	}

	req.Method = http.MethodPost
	key2 := GcsGenerateCacheKey(req)
	if key2 != "POST http://localhost?baz=qux&foo=bar" {
		t.Fatalf("key1 = %s", key1)
	}
}

func TestGcsGenerateSignedCacheKey(t *testing.T) {
	req := &http.Request{
		Method: http.MethodGet,
		URL:    &url2.URL{RawQuery: "X-Goog-Algorithm=xxx&X-Goog-Signature=yyy", Scheme: "http", Host: "localhost"},
	}

	key1 := GcsGenerateCacheKey(req)
	if key1 != "http://localhost" {
		t.Fatalf("key1 = %s", key1)
	}

	req.Method = http.MethodPost
	req.URL.RawQuery = req.URL.RawQuery + "&foo=bar"
	key2 := GcsGenerateCacheKey(req)
	if key2 != "POST http://localhost?foo=bar" {
		t.Fatalf("key2 = %s", key2)
	}
}

func TestGcsAuthorizeCache(t *testing.T) {
	mux := http.NewServeMux()

	isGet := false
	isHead := false

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			isGet = true
		}
		if r.Method == http.MethodHead {
			isHead = true
		}
		w.WriteHeader(http.StatusOK)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	req := &http.Request{Method: http.MethodGet, URL: &url2.URL{Scheme: "http", Host: ts.Listener.Addr().String()}}

	client := http.DefaultClient
	client.Do(req)
	res2 := GcsAuthorizeCache(req, client)

	if !isGet {
		t.Fatalf("isGet should be true")
	}

	if !isHead {
		t.Fatalf("isHead should be true")
	}

	if !res2 {
		t.Fatalf("GcsAuthorizeCache should return true")
	}
}
