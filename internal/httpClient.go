package httpClient

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "time"
)

type HttpClient struct {
    client *http.Client
}

func NewHttpClient() *HttpClient {
    return &HttpClient{
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

func (h *HttpClient) Get(url string, headers map[string]string) ([]byte, error) {
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    for key, value := range headers {
        req.Header.Add(key, value)
    }

    resp, err := h.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}

func (h *HttpClient) Post(url string, headers map[string]string, body interface{}) ([]byte, error) {
    requestBody, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }

    for key, value := range headers {
        req.Header.Add(key, value)
    }

    req.Header.Add("Content-Type", "application/json")

    resp, err := h.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}

func (h *HttpClient) Put(url string, headers map[string]string, body interface{}) ([]byte, error) {
    requestBody, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }

    for key, value := range headers {
        req.Header.Add(key, value)
    }

    req.Header.Add("Content-Type", "application/json")

    resp, err := h.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}

func (h *HttpClient) Delete(url string, headers map[string]string) ([]byte, error) {
    req, err := http.NewRequest(http.MethodDelete, url, nil)
    if err != nil {
        return nil, err
    }

    for key, value := range headers {
        req.Header.Add(key, value)
    }

    resp, err := h.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return io.ReadAll(resp.Body)
}
