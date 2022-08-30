package fetcher

import (
	"bufio"
	"io"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		buf = append(buf, buf[:n]...)
	}
	return buf, nil
}
