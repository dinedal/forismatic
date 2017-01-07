package forismatic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	SenderName  string `json:"senderName"`
	SenderLink  string `json:"senderLink"`
	QuoteLink   string `json:"quoteLink"`
}

func GetQuote() (Quote, error) {
	a := Quote{}
	resp, httpErr := http.Get("http://api.forismatic.com/api/1.0/?format=json&method=getQuote&lang=en")
	if httpErr != nil {
		return a, httpErr
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return a, readErr
	}

	fixedBody := bytes.Replace(body, []byte("\\'"), []byte("'"), -1)

	jsonErr := json.Unmarshal(fixedBody, &a)
	if jsonErr != nil {
		return a, jsonErr
	}
	return a, nil
}
