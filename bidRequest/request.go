package request

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:8080/read"
	body, err := getResponse(url)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func getResponse(url string) ([]byte, error) {
	if len(url) == 0 {
		return nil, errors.New("Invalid URL")
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	code := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)

	if err == nil && code != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf("Server status error: %v", http.StatusText(code))
	}

	return body, nil
}
