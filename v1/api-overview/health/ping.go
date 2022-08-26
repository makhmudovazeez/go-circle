package health

import (
	"net/http"
)

var (
	url string = "https://api-sandbox.circle.com/ping"
)

func PING() (res *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return res, err
	}

	req.Header.Add("Accept", "application/json")

	res, err = http.DefaultClient.Do(req)

	return res, err
}
