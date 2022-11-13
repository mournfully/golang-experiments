package urlshort

import (
	"fmt"
	"net/http"
)

/*
	MapHandler will return an http.HandlerFunc that will attempt to
	- map any paths (localhost:8080/<path>) to corresponding URL (https://<url>)
	- if provided path =/= in map, then do fallback http.Handler
*/

// Introduction to HTTP with Go : Our first microservice - YouTube -- https://www.youtube.com/watch?v=MKkokYpGyTU
// docker-development-youtube-series/golang/introduction/part-3.http at master · marcel-dempers/docker-development-youtube-series -- https://github.com/marcel-dempers/docker-development-youtube-series/tree/master/golang/introduction/part-3.http
// http package - net/http - Go Packages -- https://pkg.go.dev/net/http
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Golang Maps/HashMap/HashTable - ULTIMATE Golang Tutorial! - YouTube -- https://www.youtube.com/watch?v=92Q8n3LlMOY
		url, ok := pathsToUrls[path]
		/*
			fmt.Println("found path ")
			fmt.Println("path " + path)
			fmt.Println("url " + url)

			Looked at solution, and thought I saw redirect or something. Searched offical docs and found two methods. Tried this the one below, but nothing happened even though it was a 'success'.
			http.RedirectHandler(url, http.StatusMovedPermanently)
			So, I just thought 'fuck it' and tried the other one, and it worked.
		*/
		if ok {
			// HTTP response status codes - HTTP | MDN -- https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fmt.Println("didn't find path ")
			fmt.Println(fallback)
			fallback.ServeHTTP(w, r)
		}
	}
}
