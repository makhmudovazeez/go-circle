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

type manCurrency string

const (
	CURRENCY_USD manCurrency = "USD"
	CURRENCY_EUR manCurrency = "EUR"
	CURRENCY_BTC manCurrency = "BTC"
	CURRENCY_ETH manCurrency = "ETH"
)

type manChain string

const (
	CHAIN_ALGO  manChain = "ALGO"
	CHAIN_AVAX  manChain = "AVAX"
	CHAIN_BTC   manChain = "BTC"
	CHAIN_ETH   manChain = "ETH"
	CHAIN_FLOW  manChain = "FLOW"
	CHAIN_HBAR  manChain = "HBAR"
	CHAIN_MATIC manChain = "MATIC"
	CHAIN_SOL   manChain = "SOL"
	CHAIN_TRX   manChain = "TRX"
	CHAIN_XLM   manChain = "XLM"
)

func CreateBlockchainAddress(token, walletId string, currency manCurrency, chain manChain) (*response.CircleResponse, error) {
	url := baseUrl + "/" + walletId + "/addresses"

	body := struct {
		IdempotencyKey string      `json:"idempotencyKey"`
		Currency       manCurrency `json:"currency"`
		Chain          manChain    `json:"chain"`
	}{
		IdempotencyKey: uuid.NewString(),
		Currency:       currency,
		Chain:          chain,
	}

	b, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(b))

	req, _ := http.NewRequest("POST", url, payload)

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
