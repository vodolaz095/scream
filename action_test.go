package scream

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestStartingServer(t *testing.T) {
	Cfg.Address = ":3000"
	Cfg.Key = "lalala"
	t.Parallel()
	go func() {
		err := StartServer()
		if err != nil {
			t.Error(err)
		}
	}()

}

func TestGetIndexPage(t *testing.T) {
	t.Parallel()
	time.Sleep(300 * time.Millisecond)
	resp, err := http.Get("http://localhost:3000/")
	if err != nil {
		t.Error(err)
	}

	if resp.Header["Content-Type"][0] != "text/html; charset=utf-8" {
		t.Error("not text/html content type")
	}
	if resp.StatusCode != 200 {
		t.Error("not 200 response code")
	}
	return
}

func TestSendNotificationSuccessJSON(t *testing.T) {
	t.Parallel()
	//	time.Sleep(1 * time.Second)
	n := notification{
		Key:     "lalala",
		Type:    "notification",
		Message: "hello",
	}
	client := &http.Client{}
	body, _ := json.Marshal(n)
	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		t.Error("not 201 response code")
	}
}

func TestSendNotificationFailWrongJSONKey(t *testing.T) {
	t.Parallel()
	//	time.Sleep(1 * time.Second)
	n := notification{
		Key:     "suka_padla",
		Type:    "notification",
		Message: "hello",
	}
	client := &http.Client{}
	body, _ := json.Marshal(n)
	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 403 {
		t.Error("not 403 response code")
	}
}

func TestSendNotificationFailJSONKeyEmpty(t *testing.T) {
	t.Parallel()
	//	time.Sleep(1 * time.Second)
	n := notification{
		Key:     "",
		Type:    "notification",
		Message: "hello",
	}
	client := &http.Client{}
	body, _ := json.Marshal(n)
	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 403 {
		t.Error("not 403 response code")
	}
}

func TestSendNotificationFailJSONTypeEmpty(t *testing.T) {
	t.Parallel()
	//	time.Sleep(1 * time.Second)
	n := notification{
		Key:     "lalala",
		Type:    "",
		Message: "hello",
	}
	client := &http.Client{}
	body, _ := json.Marshal(n)
	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Error("not 400 response code")
	}

}

func TestSendNotificationFailJSONMessageEmpty(t *testing.T) {
	t.Parallel()
	//	time.Sleep(1 * time.Second)
	n := notification{
		Key:     "lalala",
		Type:    "notification",
		Message: "",
	}
	client := &http.Client{}
	body, _ := json.Marshal(n)
	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Error("not 400 response code")
	}

}

func TestSendNotificationFailWrongJSONSyntax(t *testing.T) {
	t.Parallel()
	//	time.Sleep(1 * time.Second)
	client := &http.Client{}
	body := []byte("some shit, not json")
	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Error("not 400 response code")
	}
}

func TestSendNotificationSuccessForm(t *testing.T) {
	t.Parallel()
	//	time.Sleep(1 * time.Second)

	data := url.Values{}
	data.Add("key", "lalala")
	data.Add("type", "notification")
	data.Add("message", "message")

	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		t.Error("not 201 response code")
	}
}

func TestSendNotificationFailWrongFormKey(t *testing.T) {
	t.Parallel()
	//time.Sleep(1 * time.Second)

	data := url.Values{}
	data.Set("key", "foo")
	data.Add("type", "notification")
	data.Add("message", "message")

	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 403 {
		t.Error("not 403 response code")
	}
}

func TestSendNotificationFailFormEmptyKey(t *testing.T) {
	t.Parallel()
	//	time.Sleep(1 * time.Second)

	data := url.Values{}
	data.Add("key", "")
	data.Add("type", "notification")
	data.Add("message", "message")

	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 403 {
		t.Error("not 403 response code")
	}
}

func TestSendNotificationFailFormEmptyType(t *testing.T) {
	t.Parallel()
	//	time.Sleep(1 * time.Second)

	data := url.Values{}
	data.Add("key", "lalala")
	data.Add("type", "")
	data.Add("message", "message")

	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Error("not 400 response code")
	}
}

func TestSendNotificationFailFormEmptyMessage(t *testing.T) {
	t.Parallel()
	//time.Sleep(1 * time.Second)

	data := url.Values{}
	data.Add("key", "lalala")
	data.Add("type", "notification")
	data.Add("message", "")

	req, _ := http.NewRequest("POST", "http://localhost:3000/", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Error("not 400 response code")
	}
}
