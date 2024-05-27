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
				Method: "floor",
				Params: []interface{}{3.14},
			},
			expect: models.Response{
				Result:     3,
				ResultType: "int",
				ID:         1,
			},
		},
	}

	for _, tc := range tests {
		t.Run((tc.name), func(t *testing.T) {
			result := Floor(tc.req)
			if result.ID != tc.expect.ID ||
				result.Result != tc.expect.Result ||
				result.ResultType != tc.expect.ResultType ||
				result.Error != tc.expect.Error {
				t.Errorf("unexpected result for test %s: got %v, want %v", tc.name, result, tc.expect)
			}
		})
	}
}
