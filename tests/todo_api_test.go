package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

type todoResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func TestTodoCRUDWorkflow(t *testing.T) {
	baseURL := strings.TrimRight(os.Getenv("TODO_API_BASE_URL"), "/")
	if baseURL == "" {
		t.Skip("set TODO_API_BASE_URL to run API integration tests")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	taskName := fmt.Sprintf("integration test %d", time.Now().UnixNano())
	var taskID string
	deleted := false
	defer func() {
		if taskID == "" || deleted {
			return
		}
		response, err := requestJSON(client, http.MethodDelete, baseURL+"/v1/todos/"+taskID, nil)
		if err == nil {
			response.Body.Close()
		}
	}()

	response, err := requestJSON(client, http.MethodPost, baseURL+"/v1/todos", map[string]string{"name": taskName})
	if err != nil {
		t.Fatalf("create TODO request failed: %v", err)
	}
	defer response.Body.Close()
	requireStatus(t, response, http.StatusCreated)

	var created todoResponse
	if err := json.NewDecoder(response.Body).Decode(&created); err != nil {
		t.Fatalf("decode create response: %v", err)
	}
	if created.ID == "" || created.Name != taskName || created.Status != "pending" {
		t.Fatalf("unexpected created TODO: %+v", created)
	}
	taskID = created.ID

	response, err = requestJSON(client, http.MethodGet, baseURL+"/v1/todos/"+taskID, nil)
	if err != nil {
		t.Fatalf("get TODO request failed: %v", err)
	}
	defer response.Body.Close()
	requireStatus(t, response, http.StatusOK)

	var fetched todoResponse
	if err := json.NewDecoder(response.Body).Decode(&fetched); err != nil {
		t.Fatalf("decode get response: %v", err)
	}
	if fetched.ID != taskID || fetched.Name != taskName || fetched.Status != "pending" {
		t.Fatalf("unexpected fetched TODO: %+v", fetched)
	}

	response, err = requestJSON(client, http.MethodPatch, baseURL+"/v1/todos/"+taskID, map[string]string{"status": "in_progress"})
	if err != nil {
		t.Fatalf("update TODO request failed: %v", err)
	}
	defer response.Body.Close()
	requireStatus(t, response, http.StatusOK)

	var updated todoResponse
	if err := json.NewDecoder(response.Body).Decode(&updated); err != nil {
		t.Fatalf("decode update response: %v", err)
	}
	if updated.ID != taskID || updated.Status != "in_progress" {
		t.Fatalf("unexpected updated TODO: %+v", updated)
	}

	response, err = requestJSON(client, http.MethodDelete, baseURL+"/v1/todos/"+taskID, nil)
	if err != nil {
		t.Fatalf("delete TODO request failed: %v", err)
	}
	defer response.Body.Close()
	requireStatus(t, response, http.StatusNoContent)
	deleted = true
}

func requestJSON(client *http.Client, method, url string, payload interface{}) (*http.Response, error) {
	var body io.Reader
	if payload != nil {
		encoded, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("encode request body: %w", err)
		}
		body = bytes.NewReader(encoded)
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	return client.Do(request)
}

func requireStatus(t *testing.T, response *http.Response, expected int) {
	t.Helper()
	if response.StatusCode == expected {
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("expected HTTP %d, got %d; reading response body failed: %v", expected, response.StatusCode, err)
	}
	t.Fatalf("expected HTTP %d, got %d; response: %s", expected, response.StatusCode, strings.TrimSpace(string(body)))
}
