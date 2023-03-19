// MADE BY XIRAS3C
// ALL RIGHT RESERVED
// SKID IT = GAY

package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "net/http"
    "net/http/httputil"
)

func main() {
    // Adresse du proxy HTTP
    proxyAddr := "proxy.example.com:8080"

    dialer := &net.Dialer{
        Timeout:   30 * time.Second,
        KeepAlive: 30 * time.Second,
        Proxy:     http.ProxyURL(&url.URL{Host: proxyAddr}),
    }

    client := &http.Client{
        Transport: &http.Transport{
            Proxy: http.ProxyURL(&url.URL{Host: proxyAddr}),
            Dial:  dialer.Dial,
        },
    }

    // Adresse du serveur distant
    serverAddr := "server.example.com:80"

    proxy := httputil.ReverseProxy{
        Director: func(req *http.Request) {
            req.URL.Scheme = "http"
            req.URL.Host = serverAddr
        },
        Transport: &http.Transport{
            Proxy: http.ProxyURL(&url.URL{Host: proxyAddr}),
            Dial:  dialer.Dial,
        },
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        proxy.ServeHTTP(w, r)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
