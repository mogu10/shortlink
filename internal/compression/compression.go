package compression

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type compressWriter struct {
	writer     http.ResponseWriter
	gzipWriter *gzip.Writer
}

func newCompressWriter(writer http.ResponseWriter) *compressWriter {
	return &compressWriter{
		writer:     writer,
		gzipWriter: gzip.NewWriter(writer),
	}
}

type compressReader struct {
	reader     io.ReadCloser
	gripReader *gzip.Reader
}

func (c *compressReader) Close() {
	c.gripReader.Close()
}

func newCompressReader(reader io.ReadCloser) (*compressReader, error) {
	gripReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		reader:     reader,
		gripReader: gripReader,
	}, nil
}

func (c compressWriter) Header() http.Header {
	return c.writer.Header()
}

func (c compressWriter) Write(bytes []byte) (int, error) {
	return c.gzipWriter.Write(bytes)
}

func (c compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.writer.Header().Set("Content-Encoding", "gzip")
	}
	c.writer.WriteHeader(statusCode)
}

func (c *compressWriter) Close() {
	c.gzipWriter.Close()
}

func GzipMV(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		finalWriter := writer

		acceptEncoding := request.Header.Get("Accept-Encoding")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		if supportsGzip {
			compWriter := newCompressWriter(writer)
			finalWriter = compWriter
			defer compWriter.Close()

			request.Header.Add("Accept-Encoding", "gzip")
		}

		contentEncoding := request.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			compReader, err := newCompressReader(request.Body)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			request.Body = compReader.gripReader
			defer compReader.Close()
		}

		h.ServeHTTP(finalWriter, request)
	})
}
