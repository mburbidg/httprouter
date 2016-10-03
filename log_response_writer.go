package httprouter

import (
	"net/http"
)

type LogResponseWriter struct {
	method string
	path   string
	status int
	size   int
	http.ResponseWriter
}

func NewLogResponseWriter(method string, path string, res http.ResponseWriter) *LogResponseWriter {
	// Default the status code to 200
	return &LogResponseWriter{method, path, 200, 6, res}
}

func (w *LogResponseWriter) Method() string {
	return w.method
}

func (w *LogResponseWriter) Path() string {
	return w.path
}

// Status provides an easy way to retrieve the status code
func (w *LogResponseWriter) Status() int {
	return w.status
}

// Size provides an easy way to retrieve the response size in bytes
func (w *LogResponseWriter) Size() int {
	return w.size
}

// Header returns & satisfies the http.ResponseWriter interface
func (w *LogResponseWriter) Header() http.Header {

	return w.ResponseWriter.Header()
}

// Write satisfies the http.ResponseWriter interface and
// captures data written, in bytes
func (w *LogResponseWriter) Write(data []byte) (int, error) {

	written, err := w.ResponseWriter.Write(data)
	w.size += written

	return written, err
}

// WriteHeader satisfies the http.ResponseWriter interface and
// allows us to cach the status code
func (w *LogResponseWriter) WriteHeader(statusCode int) {

	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
