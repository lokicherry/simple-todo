package main

import (
	"testing"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
	"fmt"
	"strings"
)

const apiUrl = "http://localhost:12345"
const resource = "/todos"

func TestGetTodo(t *testing.T) {
	urlStr := formTodoApiUrl()
	resp := get(urlStr)

	checkResponseCode(t, http.StatusOK, resp.StatusCode)
}

func TestGetTodoById(t *testing.T) {
	urlStr := formTodoApiUrl()
	urlStrWithId2 := urlStr + "/4"
	resp := get(urlStrWithId2)

	checkResponseCode(t, http.StatusOK, resp.StatusCode)

	m := getJsonResponse(resp)
	if m["name"] != "Go to Office" {
		t.Errorf("Expected task name to be 'Go to Office'. Got '%s'", m["name"])
	}
}

func TestDeleteTodoWithId1(t *testing.T) {
	urlStr := formTodoApiUrl()
	urlStrWithId1 := urlStr + "/1"
	resp := delete(urlStrWithId1)

	checkResponseCode(t, http.StatusOK, resp.StatusCode)
}

func TestCreateTodo(t *testing.T) {
	data := url.Values{}
	data.Set("name", "Wake up Arjun!")
	urlStr := formTodoApiUrl()
	resp := post(urlStr, data)

	checkResponseCode(t, http.StatusOK, resp.StatusCode)

	m := getJsonResponse(resp)
	if m["name"] != "Wake up Arjun!" {
		t.Errorf("Expected task name to be 'Wake up Arjun!'. Got '%s'", m["name"])
	}
}

func getJsonResponse(resp *http.Response) map[string]interface{} {
	var m map[string]interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &m)
	return m
}

func post(urlStr string, data url.Values) (*http.Response) {
	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, _ := client.Do(r)
	return resp
}

func get(urlStr string) (*http.Response) {
	client := &http.Client{}
	r, _ := http.NewRequest("GET", urlStr, nil)
	resp, _ := client.Do(r)
	return resp
}

func delete(urlStr string) (*http.Response) {
	client := &http.Client{}
	r, _ := http.NewRequest("DELETE", urlStr, nil)
	resp, _ := client.Do(r)
	return resp
}

func formTodoApiUrl() string {
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()
	fmt.Println(urlStr)
	return urlStr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
