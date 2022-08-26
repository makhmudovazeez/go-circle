package wallets

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/makhmudovazeez/go-circle/v1/response"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateWallet(token string, description ...string) (*response.CircleResponse, error) {
	body := struct {
		IdempotencyKey string `json:"idempotencyKey"`
		Description    string `json:"description"`
	}{}

	body.IdempotencyKey = uuid.NewString()

	for _, value := range description {
		body.Description = body.Description + ". " + value
		if len(body.Description) > 255 {
			body.Description = body.Description[:255]
		}
	}

	b, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(b))

	req, _ := http.NewRequest("POST", baseUrl, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	mResp := &response.CircleResponse{
		StatusCode: res.StatusCode,
		Body:       string(resBody),
	}

	return mResp, nil
}
