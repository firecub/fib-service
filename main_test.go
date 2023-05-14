package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGetNumber(t *testing.T) {
    routerGroup := setUpRouterGroup()
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/fib/number/6", nil)
    routerGroup.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Fatalf("Expected HTTP status OK response.")
    }
}

func TestGetIndex(t *testing.T) {
    routerGroup := setUpRouterGroup()
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/fib/index/6", nil)
    routerGroup.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Fatalf("Expected HTTP status OK response.")
    }
}

