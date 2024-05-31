package services

import (
	"rpc-server/internal/models"
	"rpc-server/internal/testutils"
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
			testutils.CheckResponse(t, tc.name, result, tc.expect)
		})
	}
}
