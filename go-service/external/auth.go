package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/GauravJ3/go-service/config"
)

var (
	AuthToken    string
	CsrfToken    string
	RefreshToken string
	CookieJar    http.CookieJar
)

func LoginAndSetTokens() error {
	loginURL := "http://localhost:5007/api/v1/auth/login"

	credentials := map[string]string{
		"username": config.GetEnv("LOGIN_EMAIL"),
		"password": config.GetEnv("LOGIN_PASSWORD"),
	}
	body, _ := json.Marshal(credentials)

	client := &http.Client{}
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
	CookieJar = jar

	req, _ := http.NewRequest("POST", loginURL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed: %s", string(respBody))
	}

	for _, c := range resp.Cookies() {
		switch c.Name {
		case "accessToken":
			AuthToken = c.Value
		case "csrfToken":
			CsrfToken = c.Value
		case "refreshToken":
			RefreshToken = c.Value
		}
	}

	if AuthToken == "" || CsrfToken == "" || RefreshToken == "" {
		return fmt.Errorf("required tokens missing in login response")
	}

	fmt.Println("✅ Logged in successfully")
	return nil
}
