package uploadprovider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/maiquocthinh/go-comic/config"
	"io"
	"net/http"
	"strings"
	"time"
)

type dropboxProvider struct {
	appKey              string
	appSecret           string
	refreshToken        string
	accessToken         string
	accessTokenExpireAt time.Time
}

const (
	baseContentUrl = "https://content.dropboxapi.com"
	baseApiUrl     = "https://api.dropboxapi.com"
)

func NewDropboxProvider(cfg *config.DropboxConfig) *dropboxProvider {
	return &dropboxProvider{
		appKey:       cfg.AppKey,
		appSecret:    cfg.AppSecret,
		refreshToken: cfg.RefreshToken,
	}
}

func (provider *dropboxProvider) getAccessToken() string {
	if provider.isLiveAccessToken() {
		return provider.accessToken
	}

	err := provider.refreshAccessToken()
	if err != nil {
		panic(err)
	}

	return provider.accessToken
}

func (provider *dropboxProvider) refreshAccessToken() error {
	client := &http.Client{}
	url := "https://api.dropbox.com/oauth2/token"
	maxRetries := 3

	payload := strings.NewReader(fmt.Sprintf(
		"refresh_token=%v&grant_type=refresh_token&client_id=%v&client_secret=%v",
		provider.refreshToken, provider.appKey, provider.appSecret,
	))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// retry request
	for i := 0; i < maxRetries; i++ {
		res, err := client.Do(req)
		if err != nil {
			continue
		}

		if res.StatusCode == http.StatusOK {
			isSuccess := false

			func() {
				defer res.Body.Close()

				body, err := io.ReadAll(res.Body)
				if err != nil {
					return
				}

				resJson := struct {
					AccessToken string `json:"access_token"`
					TokenType   string `json:"token_type"`
					ExpiresIn   int    `json:"expires_in"`
				}{}

				if err := json.Unmarshal(body, &resJson); err != nil {
					return
				}

				provider.accessToken = resJson.AccessToken
				provider.accessTokenExpireAt = time.Now().Add(time.Duration(resJson.ExpiresIn) * time.Second)

				isSuccess = true
			}()

			if isSuccess {
				return nil
			}
		}

	}

	return errors.New("Refresh token fail.")
}

func (provider *dropboxProvider) isLiveAccessToken() bool {
	if provider.accessTokenExpireAt.Before(time.Now()) {
		return false
	}

	maxReties := 3
	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/2/check/user", baseApiUrl), strings.NewReader(`{}`))
	if err != nil {
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", provider.accessToken))

	// retry request
	for i := 0; i < maxReties; i++ {
		res, err := client.Do(req)
		if err == nil {
			if res.StatusCode != http.StatusOK {
				return true
			} else if res.StatusCode != http.StatusUnauthorized {
				return false
			}
		}
	}

	return false
}

func (provider *dropboxProvider) UploadImage(ctx context.Context, data []byte, path string) error {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/2/files/upload", baseContentUrl), bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", provider.getAccessToken()))
	req.Header.Add("Dropbox-API-Arg", fmt.Sprintf(`{"mode":"overwrite","path":"%s"}`, path))
	req.Header.Add("Content-Type", "application/octet-stream")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("Upload fail.")
	}

	return nil
}

func (provider *dropboxProvider) GetShareLink(ctx context.Context, path string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/2/sharing/create_shared_link", baseApiUrl),
		strings.NewReader(fmt.Sprintf(`{"path": "%v"}`, path)),
	)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", provider.getAccessToken()))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return "", err
		}

		resJson := struct {
			Url string `json:"url"`
		}{}
		if err := json.Unmarshal(body, &resJson); err != nil {
			return "", err
		}
		return resJson.Url, nil

	}

	return "", errors.New("Get share link fail")
}
