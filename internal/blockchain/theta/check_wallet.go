package theta

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var lastKnownBalances map[string]string = make(map[string]string)

func CheckWalletChange(walletAddress string) bool {
	url := "http://www.thetascan.io/api/balance/?address=" + walletAddress

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching balance:", err)
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return false
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("error decoding response json:", err)
		return false
	}

	fmt.Println(result)

	currentBalance, ok := result["balance"].(string)
	if !ok {
		fmt.Println("balance not found or is not a string", result)
		return false
	}

	lastBalance, exists := lastKnownBalances[walletAddress]
	if !exists || lastBalance != currentBalance {
		lastKnownBalances[walletAddress] = currentBalance
		return true
	}

	return false
}