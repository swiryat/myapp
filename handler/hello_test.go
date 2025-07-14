package handler

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHello(t *testing.T) {
    req := httptest.NewRequest("GET", "/?name=Гоша", nil)
    w := httptest.NewRecorder()

    Hello(w, req)

    res := w.Result()
    body := strings.TrimSpace(w.Body.String())

    if res.StatusCode != 200 {
        t.Errorf("ожидался статус 200, получен %d", res.StatusCode)
    }

    if !strings.Contains(body, "Гоша") {
        t.Errorf("ожидалась строка с именем, получено: %s", body)
    }
}
 
