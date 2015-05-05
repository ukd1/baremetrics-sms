/**
 * Baremetrics CLI Client
 *
 * I'm a go newb - so sorry about this in advance. I'd love
 * feedback on what I could do better though. :-)
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/jsonq"
	"github.com/sfreiberg/gotwilio"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const endpoint_url = string("https://dashboard.baremetrics.com/")
const version = string("0.0.2")

func main() {
	twilio := gotwilio.NewTwilioClient(os.Getenv("TWILIO_SID"), os.Getenv("TWILIO_TOKEN"))
	jsonstring := fetch_http("stats/mrr/dashboard.json", os.Getenv("BAREMETRICS_COOKIE"))

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonstring))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	mrr, _ := jq.String("current_human")
	change, _ := jq.String("trend", "percent")
	class, _ := jq.String("trend", "class")
	class = strings.Split(class, " ")[0]

	message := "Current MRR: " + mrr + ", " + class + " by " + change + " over last 30 days."

	numbers := strings.Split(os.Getenv("PHONE_NUMBERS"), ",")
	for _, to := range numbers {
		twilio.SendSMS(os.Getenv("TWILIO_FROM"), to, message, "", "")
	}
}

func fetch_http(url string, cookie string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", endpoint_url+url, nil)

	req.Header.Set("Cookie", cookie)
	req.Header.Add("User-Agent", "BaremetricsSms/"+version+" (https://twitter.com/rhs)")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println("Could not fetch JSON:")
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		content, _ := ioutil.ReadAll(resp.Body)
		return string(content[:])
	} else {
		fmt.Println("Error: " + resp.Status)
		os.Exit(1)
		return ""
	}
}
