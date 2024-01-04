package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type config struct {
	Filepath  string `yaml:"filepath"`
	TenantId  string `yaml:"tenantId"`
	AppId     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
	Scope     string `yaml:"scope"`
	TimeRange string `yaml:"timerange"`
}

var appConfig config

func main() {

	// Define the flags
	endpoint := flag.String("endpoint", "", "API endpoint URL eg. /users")
	isPOST := flag.Bool("post", false, "Use POST Method")
	postData := flag.String("data", "", "POST data in JSON format")

	flag.Parse()

	if *endpoint == "" {
		fmt.Println("timerange and endpoint are required flags")
		flag.Usage()
		os.Exit(1)
	}

	configData, err := os.ReadFile("config.yaml")
	handleError(err)

	err = yaml.Unmarshal(configData, &appConfig)
	handleError(err)

	err = configVerify(appConfig)
	if err != nil {
		log.Println(err.Error())
		panic("Missing informaiton in the config.yaml file")
	}

	logs := fetchlogs(*endpoint, *postData, *isPOST)
	fmt.Println(logs)

}

func configVerify(c config) error {
	if &c.AppId == nil {
		return errors.New("Error, App ID was not found ")
	} else if &c.AppSecret == nil {
		return errors.New("Error, App Secret was not found ")
	} else if &c.TenantId == nil {
		return errors.New("Error, Tenant ID was not found ")
	}
	return nil
}

func fetchlogs(endpoint string, data string, isPOST bool) string {

	token := "Bearer " + getToken()

	req := &http.Request{}

	if isPOST {
		req, _ = http.NewRequest("POST", endpoint, strings.NewReader(data))

	} else {
		req, _ = http.NewRequest("GET", endpoint, nil)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
		log.Println("Request Status Code", resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("Error fetching logs: " + string(body))
		os.Exit(1)

	}
	handleError(err)

	logs, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(logs)

}

func getToken() string {

	type Response struct {
		Access_token string `json:"access_token"`
	}

	tenantId := appConfig.TenantId
	appId := appConfig.AppId
	appSecret := appConfig.AppSecret
	scope := appConfig.Scope
	oAuthUri := fmt.Sprintf("%s%s%s", "https://login.microsoftonline.com/", tenantId, "/oauth2/v2.0/token")
	data := url.Values{
		"scope":         {scope},
		"client_id":     {appId},
		"client_secret": {appSecret},
		"grant_type":    {"client_credentials"},
	}
	req, err := http.NewRequest("POST", oAuthUri, strings.NewReader(data.Encode()))
	if err != nil {

		log.Println("Error creating token: ", err)

	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending the request", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		log.Println("Failed Request:")
		log.Println(string(body))
		fmt.Println("Token requests status code", resp.StatusCode)
		os.Exit(1)
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return string(result.Access_token)
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
