package httpr

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// Gzip golang explanation - https://www.youtube.com/watch?v=oSDtwcvU2-o

type gzipResponseWriter struct {
	http.ResponseWriter
	GzipWriter io.Writer
}

func (g gzipResponseWriter) Write(data []byte) (int, error) {
	return g.GzipWriter.Write(data)
}

func acceptsGzip(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
}

func serveGzip(w http.ResponseWriter, r *http.Request, f http.Handler) {
	w.Header().Add("Content-Encoding", "gzip")
	gzw := gzip.NewWriter(w)
	defer gzw.Close()

	gzipRW := gzipResponseWriter{
		ResponseWriter: w,
		GzipWriter:     gzw,
	}

	f.ServeHTTP(gzipRW, r)
}
