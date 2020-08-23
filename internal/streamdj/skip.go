package streamdj

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	skipPath = "request_skip"
)

func SkipTrack(baseURL, channelID, apiKey string) error {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	skipURL := fmt.Sprintf("%v/%v/%v/%v", baseURL, skipPath, channelID, apiKey)
	req, _ := http.NewRequest(http.MethodGet, skipURL, nil)
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("skip error")
	}

	log.Printf("status code %v", resp.StatusCode)
	return nil
}
