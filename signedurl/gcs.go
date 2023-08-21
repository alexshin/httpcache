package signedurl

import (
	"net/http"
)

var gcsSignUrlServiceHeadersSet = map[string]bool{
	"X-Goog-Algorithm":     true,
	"X-Goog-Credential":    true,
	"X-Goog-Date":          true,
	"X-Goog-Expires":       true,
	"X-Goog-SignedHeaders": true,
	"X-Goog-Signature":     true,
}

func GcsGenerateCacheKey(req *http.Request) string {
	url := req.URL
	values := url.Query()
	for key, _ := range gcsSignUrlServiceHeadersSet {
		values.Del(key)
	}
	url.RawQuery = values.Encode()

	if req.Method == http.MethodGet {
		return url.String()
	} else {
		return req.Method + " " + url.String()
	}
}

func GcsAuthorizeCache(req *http.Request, c *http.Client) bool {
	if req.Method != http.MethodGet {
		return true
	}

	req2 := req.Clone(req.Context())

	req2.Method = http.MethodHead

	resp, err := c.Do(req2)
	if err != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}
