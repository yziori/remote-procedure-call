package services

import (
	"rpc-server/internal/models"
	"testing"
)

func TestNroot(t *testing.T) {
	tests := []struct {
		name   string
		req    models.Request
		expect models.Response
	}{
		{
			name: "valid input",
			req: models.Request{
				ID:     1,
				Params: []interface{}{2.0, 16.0},
			},
			expect: models.Response{
				ID:         1,
				Result:     4.0,
				ResultType: "float64",
				Error:      "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Nroot(tc.req)
			if result.ID != tc.expect.ID || result.ResultType != tc.expect.ResultType || result.Error != tc.expect.Error {
				t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, result, tc.expect)
			}
		})
	}
}
