httpcache
=========

[![Build Status](https://travis-ci.org/gregjones/httpcache.svg?branch=master)](https://travis-ci.org/gregjones/httpcache) [![GoDoc](https://godoc.org/github.com/gregjones/httpcache?status.svg)](https://godoc.org/github.com/gregjones/httpcache)

Package httpcache provides a http.RoundTripper implementation that works as a mostly [RFC 7234](https://tools.ietf.org/html/rfc7234) compliant cache for http responses.

It is only suitable for use as a 'private' cache (i.e. for a web-browser or an API-client and not for a shared proxy).

This project isn't actively maintained; it works for what I, and seemingly others, want to do with it, and I consider it "done". That said, if you find any issues, please open a Pull Request and I will try to review it. Any changes now that change the public API won't be considered.

Cache Backends
--------------

- The built-in 'memory' cache stores responses in an in-memory map.
- [`github.com/gregjones/httpcache/diskcache`](https://github.com/gregjones/httpcache/tree/master/diskcache) provides a filesystem-backed cache using the [diskv](https://github.com/peterbourgon/diskv) library.
- [`github.com/gregjones/httpcache/memcache`](https://github.com/gregjones/httpcache/tree/master/memcache) provides memcache implementations, for both App Engine and 'normal' memcache servers.
- [`sourcegraph.com/sourcegraph/s3cache`](https://sourcegraph.com/github.com/sourcegraph/s3cache) uses Amazon S3 for storage.
- [`github.com/gregjones/httpcache/leveldbcache`](https://github.com/gregjones/httpcache/tree/master/leveldbcache) provides a filesystem-backed cache using [leveldb](https://github.com/syndtr/goleveldb/leveldb).
- [`github.com/die-net/lrucache`](https://github.com/die-net/lrucache) provides an in-memory cache that will evict least-recently used entries.
- [`github.com/die-net/lrucache/twotier`](https://github.com/die-net/lrucache/tree/master/twotier) allows caches to be combined, for example to use lrucache above with a persistent disk-cache.
- [`github.com/birkelund/boltdbcache`](https://github.com/birkelund/boltdbcache) provides a BoltDB implementation (based on the [bbolt](https://github.com/coreos/bbolt) fork).

If you implement any other backend and wish it to be linked here, please send a PR editing this file.

Signed URLs
-----------

Usually there is a big problem related to caching signed URLs. The problem is that the URL is signed with a timestamp, 
so the URL is different every time. This means that the cache will never hit, and the request will always be sent 
to the origin server.

This package provides a solution to this problem. The solution is to use a `httpcache.CacheConfig` to 
provide additional behavior to use cache.

This solution allow you to pass:

- `CacheKeyFn func(req *http.Request) string` - A function that returns a key for the cache. Then you can remove all
  redundant information from the request, and return a key that will be used to store the response in the cache.
- `AuthorizeCacheFn func(req *http.Request, c *http.Client) bool` - function that authorizes original request to be 
  cached. This function should return `true` if the request should be cached, and `false` otherwise.

Example for GCS you can find in [example/gcs](example/gcs) folder. AuthorizeCacheFn is implemented to do HEAD request
to remote to check that client has access to the resource. After this you can get cached response by cache-key with
no any specific URL-Signature parameters.

License
-------

-	[MIT License](LICENSE.txt)
