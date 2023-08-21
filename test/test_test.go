package test_test

import (
	"testing"

	"github.com/alexshin/httpcache"
	"github.com/alexshin/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
