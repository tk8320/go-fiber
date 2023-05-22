package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go-fiber/routes"
	"log"
	"net/http/httptest"
	"testing"
)

var app *fiber.App

func init() {
	app = fiber.New()
	routes.AddRoutes(app)
	// Truncate the tables for easy testing
	err := routes.CTX.TruncateBlog()
	if err != nil {
		log.Fatalln("Unable to truncate the table. Initialization failed")
	}
}

// Unit test for creating a blog
func TestCreate(t *testing.T) {

	requestBody, _ := json.Marshal(map[string]string{
		"title":       "This is the test blog",
		"description": "This is the test blog",
		"body":        "this is the test blog"})

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "Create a blog post",
			route:        "/api/blog-post",
			expectedCode: 201,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest("POST", test.route, bytes.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		if err != nil {
			fmt.Println(err)
		}
		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

// unit test for get
func TestGet(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get HTTP status 200",
			route:        "/api/blog-post",
			expectedCode: 200,
		},
		{
			description:  "Get single Data",
			route:        "/api/blog-post/1",
			expectedCode: 200,
		},
		{
			description:  "Get non existing Data",
			route:        "/api/blog-post/2",
			expectedCode: 404,
		},
	}

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)
		//fmt.Println(req)
		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, err := app.Test(req, -1)
		if err != nil {
			fmt.Println(err)
		}
		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

// Unit test for update
func TestUpdate(t *testing.T) {

	requestBody, _ := json.Marshal(map[string]string{
		"title": "This is updated one"})

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "Update an existing blog",
			route:        "/api/blog-post/1",
			expectedCode: 200,
		},
		{
			description:  "Update a non-existing blog",
			route:        "/api/blog-post/2",
			expectedCode: 404,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest("PATCH", test.route, bytes.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		if err != nil {
			fmt.Println(err)
		}
		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

// Unit Test for Delete
func TestDelete(t *testing.T) {

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "Update an existing blog",
			route:        "/api/blog-post/1",
			expectedCode: 200,
		},
		{
			description:  "Update a non-existing blog",
			route:        "/api/blog-post/2",
			expectedCode: 404,
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest("DELETE", test.route, nil)

		resp, err := app.Test(req, -1)
		if err != nil {
			fmt.Println(err)
		}
		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
