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
		{
			name: "valid parameter type",
			req: models.Request{
				ID:     1,
				Params: []interface{}{2.0, "invalid"},
			},
			expect: models.Response{
				ID:    1,
				Error: "Invalid Parameter type",
			},
		},
		{
			name: "valid parameter count",
			req: models.Request{
				ID:     1,
				Params: []interface{}{3.0},
			},
			expect: models.Response{
				ID:    1,
				Error: "Invalid parameter count",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Nroot(tc.req)
			if result.ID != tc.expect.ID || result.ResultType != tc.expect.ResultType || result.Error != tc.expect.Error {
				t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, result, tc.expect)
			}
			// float64かどうか確認し、一致する場合は値を比較する
			if r, ok := result.Result.(float64); ok {
				if r != tc.expect.Result {
					t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, r, tc.expect)
				}
			} else if result.Result != tc.expect.Result {
				t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, result.Result, tc.expect.Result)
			}
		})
	}
}
