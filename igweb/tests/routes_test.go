package tests

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var testHost string = "http://localhost:8080"

func checkRoute(t *testing.T, route string, expectedToken string) {

	testURL := testHost + route
	response, err := http.Get(testURL)
	if err != nil {
		t.Errorf("Could not connect to URL: %s. Failed with error: %s", testURL, err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("Could not read response body. Failed with error: %s", err)
		}
		if strings.Contains(string(contents), expectedToken) == false {
			t.Errorf("Could not find expected string token: \"%s\", in response body for URL: %s", expectedToken, testURL)
		}
	}
}

func TestServerSideRoutes(t *testing.T) {

	routes := []string{"", "/", "/index", "/products", "/product-detail/swiss-army-knife", "/about", "/contact", "/shopping-cart"}
	tokenMap := map[string]string{"": "IGWEB", "/": "IGWEB", "/index": "IGWEB", "/products": "Add To Cart", "/product-detail/swiss-army-knife": "Swiss Army Knife", "/about": "Molly", "/contact": "Enter your message for us here"}

	for _, r := range routes {
		checkRoute(t, r, tokenMap[r])
	}
}