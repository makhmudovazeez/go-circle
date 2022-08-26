package wallets

import (
	"fmt"
	"github.com/makhmudovazeez/go-circle/v1/response"
	"io/ioutil"
	"net/http"
)

func GetWallet(token, walletId string) (*response.CircleResponse, error) {
	url := baseUrl + "/" + walletId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
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
