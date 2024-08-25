package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// AlertManagerWebhook 结构体用于映射从 Alertmanager 接收到的 JSON 数据
type AlertManagerWebhook struct {
	Receiver          string      `json:"receiver"`
	Status            string      `json:"status"`
	Alerts            []Alert     `json:"alerts"`
	GroupLabels       Labels      `json:"groupLabels"`
	CommonLabels      Labels      `json:"commonLabels"`
	CommonAnnotations Annotations `json:"commonAnnotations"`
	ExternalURL       string      `json:"externalURL"`
	Version           string      `json:"version"`
	GroupKey          string      `json:"groupKey"`
}

type Alert struct {
	Status       string      `json:"status"`
	Labels       Labels      `json:"labels"`
	Annotations  Annotations `json:"annotations"`
	StartsAt     string      `json:"startsAt"`
	EndsAt       string      `json:"endsAt"`
	GeneratorURL string      `json:"generatorURL"`
	Fingerprint  string      `json:"fingerprint"`
}

type Labels struct {
	AlertName string `json:"alertname"`
	Instance  string `json:"instance"`
	Severity  string `json:"severity"`
}

type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}
func alertHandler(w http.ResponseWriter, r *http.Request) {
	// 读取请求体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusBadRequest)
		return
	}

	// 打印原始的 JSON 字符串
	fmt.Println("Received alert:", string(body))

	// 解析 JSON 数据
	var webhook AlertManagerWebhook
	err = json.Unmarshal(body, &webhook)
	if err != nil {
		http.Error(w, "Could not parse JSON", http.StatusBadRequest)
		return
	}

	// 处理解析后的数据
	for _, alert := range webhook.Alerts {
		fmt.Printf("Alerddt: %s\n", alert.Labels.AlertName)
		fmt.Printf("Insccdtance: %s\n", alert.Labels.Instance)
		fmt.Printf("Sevccderity: %s\n", alert.Labels.Severity)
		fmt.Printf("Sumccdmary: %s\n", alert.Annotations.Summary)
		fmt.Printf("Desccdcription: %s\n", alert.Annotations.Description)
		fmt.Printf("Staccdtus: %s\n", webhook.Status)
	}

	// 返回成功响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Alssssert received"}`))
}

func main() {
	http.HandleFunc("/alert", alertHandler)
	http.HandleFunc("/", rootHandler)
	log.Println("55555ffff555555Starting server on :18080")
	log.Fatal(http.ListenAndServe(":18080", nil))
}
