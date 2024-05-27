package services

import (
	"rpc-server/internal/models"
	"testing"
)

func TestFloor(t *testing.T) {
	tests := []struct {
		name   string
		req    models.Request
		expect models.Response
	}{
		{
			name: "valid input",
			req: models.Request{
				ID:     1,
				Params: []interface{}{3.14},
			},
			expect: models.Response{
				ID:         1,
				Result:     3.0,
				ResultType: "int",
				Error:      "",
			},
		},
		{
			name: "invalid params type",
			req: models.Request{
				ID:     2,
				Params: []interface{}{"invalid"},
			},
			expect: models.Response{
				ID:    2,
				Error: "Invalid params",
			},
		},
		{
			name: "incorrect number of params",
			req: models.Request{
				ID:     3,
				Params: []interface{}{3.14, 2.71},
			},
			expect: models.Response{
				ID:    3,
				Error: "Invalid params",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Floor(tc.req)
			if result.ID != tc.expect.ID || result.ResultType != tc.expect.ResultType || result.Error != tc.expect.Error {
				t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, result, tc.expect)
			}
			if r, ok := result.Result.(float64); ok {
				if r != tc.expect.Result {
					t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, r, tc.expect.Result)
				}
			} else if result.Result != tc.expect.Result {
				t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, result.Result, tc.expect.Result)
			}
		})
	}
}
