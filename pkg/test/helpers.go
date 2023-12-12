package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CreateRequest(t *testing.T, method, url string, body interface{}) (*httptest.ResponseRecorder, *http.Request) {
	t.Helper()

	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	r, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	return w, r
}

func UnmarshalResponse(t *testing.T, recorder *httptest.ResponseRecorder, response interface{}) {
	t.Helper()

	err := json.Unmarshal(recorder.Body.Bytes(), response)
	if err != nil {
		t.Fatal(err)
	}
}
