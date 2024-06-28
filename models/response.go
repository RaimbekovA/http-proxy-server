package models

import "net/http"

type ProxyResponse struct {
    ID      int64           `json:"id"`
    Status  int             `json:"status"`
    Headers http.Header     `json:"headers"`
    Length  int             `json:"length"`
}
