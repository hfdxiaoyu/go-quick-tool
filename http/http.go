package http

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"time"
)

func Get(uri string, heads map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return []byte{}, err
	}
	if len(heads) != 0 {
		for s, s2 := range heads {
			req.Header.Add(s, s2)
		}
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	do, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer do.Body.Close()
	body, err := io.ReadAll(do.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func Post(url string, heads map[string]string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return []byte{}, err
	}
	if len(heads) != 0 {
		for s, s2 := range heads {
			req.Header.Add(s, s2)
		}
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	do, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer do.Body.Close()
	body, err := io.ReadAll(do.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func GetByProxy(uri, proxy string, heads map[string]string) ([]byte, error) {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return []byte{}, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return []byte{}, err
	}
	if len(heads) != 0 {
		for s, s2 := range heads {
			req.Header.Add(s, s2)
		}
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}

	do, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer do.Body.Close()
	body, err := io.ReadAll(do.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func PostByProxy(uri, proxy string, heads map[string]string, data []byte) ([]byte, error) {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return []byte{}, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	if err != nil {
		return []byte{}, err
	}
	if len(heads) != 0 {
		for s, s2 := range heads {
			req.Header.Add(s, s2)
		}
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	do, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer do.Body.Close()
	body, err := io.ReadAll(do.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
