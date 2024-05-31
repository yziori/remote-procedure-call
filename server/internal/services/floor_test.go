package services

import (
	"rpc-server/internal/models"
	"rpc-server/internal/testutils"
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
			testutils.CheckResponse(t, tc.name, result, tc.expect)
		})
	}
}
