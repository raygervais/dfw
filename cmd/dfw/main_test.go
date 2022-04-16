package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

const (
	baseUrl = "http://127.0.0.1:8080/v1"
)

func TestApiCRUD(t *testing.T) {
	testCases := []struct {
		desc     string
		endpoint string
		method   string
		status   int
		body     string
		comment  string
	}{
		{
			desc:     "Valid GET on /rooms",
			endpoint: "/rooms",
			method:   "GET",
			status:   200,
			body:     "",
		},
		{
			desc:     "Valid POST to /rooms",
			endpoint: "/rooms",
			method:   "POST",
			status:   201,
			body:     `{"hostEmail":"test@test.com","timeToLive":60}`,
		},
		{
			desc:     "Valid GET on /rooms after POST",
			endpoint: "/rooms",
			method:   "GET",
			status:   200,
			body:     "",
			comment:  "Returned response should contain a single room from the previous POST",
		},
		{
			desc:     "Valid POST to /rooms with an ID provided for lookup in later tests",
			endpoint: "/rooms",
			method:   "POST",
			status:   201,
			body:     `{"hostEmail":"testing123@test.com","timeToLive":60,"id":"123"}`,
		},
		{
			desc:     "Valid GET on /rooms/123",
			endpoint: "/rooms/123",
			method:   "GET",
			status:   200,
			body:     "",
			comment:  "Returned response should contain the room from the previous POST",
		},
		{
			desc:     "Add a room with an invalid JSON body",
			endpoint: "/rooms",
			method:   "POST",
			status:   400,
			body:     `{"hostEmail":"ABC","timeToLive":"ABC"}`,
		},
		{
			desc:     "Add a prompt to a room with ID 123",
			endpoint: "/rooms/123/prompts",
			method:   "POST",
			status:   201,
			body:     `{"text":"Who is the Batman?", "active":true, "id":"123"}`,
		},
		{
			desc:     "Add a prompt to a room that doesn't exist",
			endpoint: "/rooms/456/prompts",
			method:   "POST",
			status:   400,
			body:     `{"text":"Who is the Batman?"}`,
		},
		{
			desc:     "Add a comment to a room with ID 123",
			endpoint: "/rooms/123/comments",
			method:   "POST",
			status:   201,
			body:     `{"comment":"Hello, world!"}`,
		},
		{
			desc:     "Add a comment to a room that doesn't exist",
			endpoint: "/rooms/456/comments",
			method:   "POST",
			status:   400,
			body:     `{"comment":"Hello, world!"}`,
		},
		{
			desc:     "Add a comment to a room with an invalid JSON body",
			endpoint: "/rooms/123/comments",
			method:   "POST",
			status:   400,
			body:     `{"comment":true,"id":123"}`,
		},
		{
			desc:     "Create a room which expires -1 minutes ago which should fail",
			endpoint: "/rooms",
			method:   "POST",
			status:   400,
			body:     `{"hostEmail":"test@example.com","id":"456","timeToLive":-1}`,
		},

		// {
		// 	desc:     "Valid DELETE on /rooms/123",
		// 	endpoint: "/rooms/123",
		// 	method:   "DELETE",
		// 	status:   200,
		// 	body:     "",
		// },
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			resp, err := SendRequest(tC.method, baseUrl+tC.endpoint, tC.body)
			if err != nil {
				t.Errorf("Error sending request: %s", err)
				t.Fail()
			}
			if resp.StatusCode != tC.status {
				t.Errorf("Expected status code: %d, got: %d", tC.status, resp.StatusCode)
				t.Fail()

				// Print Body
				body, _ := ioutil.ReadAll(resp.Body)
				t.Logf("Body: %s", body)
			}
		})
	}
}

func SendRequest(method, endpoint, body string) (*http.Response, error) {
	req, err := http.NewRequest(method, endpoint, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}
