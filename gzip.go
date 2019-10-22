package runtime

import (
	//"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	//"reflect"
)

func GZipConsumer() Consumer {
	return ConsumerFunc(func(reader io.Reader, data interface{}) error {
		if reader == nil {
			return errors.New("GZipConsumer requires reader") // early exit
		}

		gzipReader, err := gzip.NewReader(reader)
		if err != nil {
			return err
		}
		writer, ok := data.(io.Writer)
		if !ok {
			return errors.New("data type must be io.Writer")
		}
		gzipWriter := gzip.NewWriter(writer)

		//buf := new(bytes.Buffer)
		//_, err = buf.ReadFrom(reader)
		//if err != nil {
		//	return err
		//}
		//b := buf.Bytes()

		_, err = gzipReader.Read(data.([]byte))
		if err != nil {
			return err
		}
		if _, err = io.Copy(gzipWriter, gzipReader); err != nil {
			return err
		}

		gzipWriter.Close()
		gzipReader.Close()

		return fmt.Errorf("Test error")
	})
}
