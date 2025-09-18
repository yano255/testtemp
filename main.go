package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DiscordWebhook struct {
	Content string `json:"content"`
}

func main() {
	ip, err := getGlobalIP()
	if err != nil {
		fmt.Printf("aaip: %v\n", err)
		return
	}

	webhookURL := "https://discord.com/api/webhooks/1417023842215395409/FvHZ_bW2VH3OKPk_vcPjn6TYA9D6fJ6muI1eEDUrdK5gkXt3FGZkxqqajE_JWxpkOr_5"

	message := fmt.Sprintf("IP: %s\nMessage: By GitHub Actions Ubuntu machine!!!", ip)

	err = sendDiscordWebhook(webhookURL, message)
	if err != nil {
		fmt.Printf("失敗: %v\n", err)
		return
	}

	fmt.Println("送付完了")
}

func getGlobalIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

func sendDiscordWebhook(webhookURL, message string) error {
	webhookData := DiscordWebhook{
		Content: message,
	}

	jsonData, err := json.Marshal(webhookData)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("エラーwebhook: %s", resp.Status)
	}

	return nil
}