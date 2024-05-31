package services

import (
	"rpc-server/internal/models"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name   string
		req    models.Request
		expect models.Response
	}{
		{
			name: "valid input",
			req: models.Request{
				ID:     1,
				Params: []interface{}{"あいうえお"},
			},
			expect: models.Response{
				ID:         1,
				Result:     "おえういあ",
				ResultType: "string",
				Error:      "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Reverse(tc.req)
			if result.ID != tc.expect.ID || result.ResultType != tc.expect.ResultType || result.Error != tc.expect.Error {
				t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, result, tc.expect)
			}
			if r, ok := result.Result.(string); ok {
				if r != tc.expect.Result {
					t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, r, tc.expect)
				}
			} else if result.Result != tc.expect.Result {
				t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, result.Result, tc.expect.Result)
			}
		})
	}
}
