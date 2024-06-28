package handlers

import (
    "encoding/json"
    "net/http"
    "net/http/httputil"
    "sync"
    "time"
    "http-proxy-server/models"
)

var store sync.Map

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
    var req models.ProxyRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    proxyReq, err := http.NewRequest(req.Method, req.URL, nil)
    if err != nil {
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
        return
    }

    for key, value := range req.Headers {
        proxyReq.Header.Set(key, value)
    }

    client := &http.Client{}
    resp, err := client.Do(proxyReq)
    if err != nil {
        http.Error(w, "Failed to send request", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    responseDump, _ := httputil.DumpResponse(resp, true)
    responseLength := len(responseDump)
    responseID := time.Now().UnixNano()

    proxyResp := models.ProxyResponse{
        ID:      responseID,
        Status:  resp.StatusCode,
        Headers: resp.Header,
        Length:  responseLength,
    }

    store.Store(responseID, proxyResp)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(proxyResp)
}
