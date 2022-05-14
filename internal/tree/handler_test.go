package tree

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTreeHandler_calculateMaxPathSum(t *testing.T) {

	type args struct {
		testBody []byte
	}
	type results struct {
		statusCode int
		result     string
	}
	tests := []struct {
		name string
		args args
		want results
	}{
		{"WithValidInputShouldReturnResult", args{testBody1}, results{200, "18"}},
		{"WithInputNotToBeBindedShouldReturn400", args{testBody2}, results{400, `"Request body cannot be binded"`}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := getResponseForTesting(tt.args.testBody)
			assert.Equal(t, tt.want.statusCode, rr.Code)
			assert.Equal(t, tt.want.result, rr.Body.String())
		})
	}
}

// getResponseForTesting creates http request and record the response for testing purposes
func getResponseForTesting(testBody []byte) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", "http://127.0.0.1:8090/api/v1/binary-tree/max-path-sum", bytes.NewBuffer(testBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	c.Request = req

	calculateMaxPathSum(c)
	return rr
}

var (
	// testBody1 is an example of valid input
	testBody1 = []byte(`{
"tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "3", "left": "6", "right": "7", "value": 3},
      {"id": "7", "left": null, "right": null, "value": 7},
      {"id": "6", "left": null, "right": null, "value": 6},
      {"id": "2", "left": "4", "right": "5", "value": 2},
      {"id": "5", "left": null, "right": null, "value": 5},
      {"id": "4", "left": null, "right": null, "value": 4}
],
"root": "1" }}
`)
	// testBody1 is an example of an input that cannot be binded due to left and right properties are given as integers
	testBody2 = []byte(`{
"tree": {
    "nodes": [
      {"id": "1", "left": 2, "right": 3, "value": 1},
      {"id": "3", "left": 6, "right": 7, "value": 3},
      {"id": "7", "left": null, "right": null, "value": 7},
      {"id": "6", "left": null, "right": null, "value": 6},
      {"id": "2", "left": 4, "right": 5, "value": 2},
      {"id": "5", "left": null, "right": null, "value": 5},
      {"id": "4", "left": null, "right": null, "value": 4}
],
"root": "1" }}
`)
)
