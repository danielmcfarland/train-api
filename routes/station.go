package routes

import (
	"fmt"
	"github.com/danielmcfarland/train-api/services"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func (h *Handler) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(fmt.Sprintf("Get Station: %s", id))

	response := fetchUrl(id)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func fetchUrl(stationId string) *http.Response {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("https://tiger-api.worldline.global/services/%v", stationId)
	headers := services.SignAWSRequest(stationId)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	header := make(http.Header)

	for k, v := range headers {
		header[k] = []string{v}
	}
	header["Access-Control-Allow-Origin"] = []string{"*"}
	header["Accept"] = []string{"application/json"}
	header["Sec-Fetch-Site"] = []string{"same-site"}
	header["Access-Control-Allow-Headers"] = []string{"Content-Type,Access-Control-Allow-Origin,Authorization,X-Amz-Date,X-Api-Key,X-Amz-Security-Token"}
	header["Accept-Language"] = []string{"en-GB,en;q=0.9"}
	header["Sec-Fetch-Mode"] = []string{"cors"}
	header["Accept-Encoding"] = []string{"gzip, deflate, br"}
	header["Origin"] = []string{"https://tiger.worldline.global"}
	header["User-Agent"] = []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.3 Safari/605.1.15"}
	header["Referer"] = []string{"https://tiger.worldline.global/"}
	header["Sec-Fetch-Dest"] = []string{"empty"}
	header["Priority"] = []string{"u=3, i"}

	req.Header = header

	res, err := client.Do(req)
	if err != nil {
		log.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	log.Println(fmt.Sprintf("Response Status Code: %v", res.StatusCode))

	return res
}
