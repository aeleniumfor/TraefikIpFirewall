package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"net/http"
	"os"
)

type Rule struct {
	BlackList []string `yaml:"black_list"`
}

func loadRule(filename string) (Rule, error) {
	data, _ := os.ReadFile(filename)
	var rule Rule
	yaml.Unmarshal(data, &rule)
	return rule, nil
}

func IP_Black_List_Handler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.Header.Get("X-Forwarded-For")
	ip := net.ParseIP(clientIP)
	if ip == nil {
		// X-Forwarded-Forが IPアドレス形式ではない場合を考慮
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Println("clientIP: ", clientIP)
	rule, err := loadRule("rule.yaml")
	if err != nil {
		// ruleファイルが失敗 == 認証サービスで機能していない
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	if clientIP != "" {
		for _, black := range rule.BlackList {
			// 表現しやすいCIDR形式を採用 ...
			// 1つだけブロックしたい場合は <ip>/32
			// 範囲のIPをブロックしたい場合は <ip>/<任意の値>
			_, cidr, _ := net.ParseCIDR(black)
			if cidr != nil {
				if cidr.Contains(ip) {
					w.WriteHeader(http.StatusUnauthorized)
				}
			}
		}
	} else {
		// ClientIP が空 == 認証できない
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	http.HandleFunc("/", IP_Black_List_Handler)
	log.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}
