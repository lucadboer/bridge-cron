package theta

import (
	"encoding/json"
	"io"
	"net/http"
)

type Theta struct{}

type Price struct {
	ThetaPrice string `json:"theta_price"`
	TfuelPrice string `json:"tfuel_price"`
	RplayPrice string `json:"RPLAY_price"`
}

// https://thetascan.io/document/
// https://thetascan.io/contracts/?data=0x3da3d8cde7b12cd2cbb688e2655bcacd8946399d&index=1800
func (theta Theta) GetPrice() (*Price, error) {
	resp, err := http.Get("http://www.thetascan.io/api/price/")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var prices Price
	err = json.Unmarshal(body, &prices)
	if err != nil {
		return nil, err
	}
	return &prices, nil
}
