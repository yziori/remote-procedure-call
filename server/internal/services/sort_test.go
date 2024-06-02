package services

import (
	"rpc-server/internal/models"
	"rpc-server/internal/testutils"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name   string
		req    models.Request
		expect models.Response
	}{
		{
			name: "valid input",
			req: models.Request{
				ID: 1,
				Params: []interface{}{
					[]interface{}{"banana", "apple", "cherry"},
				},
			},
			expect: models.Response{
				ID:         1,
				Result:     []interface{}{"apple", "banana", "cherry"},
				ResultType: "[]string",
				Error:      "",
			},
		},
		{
			name: "valid param count",
			req: models.Request{
				ID:     1,
				Params: []interface{}{"banana", "apple"},
			},
			expect: models.Response{
				ID:    1,
				Error: "Invalid parameter count",
			},
		},
		{
			name: "valid param type",
			req: models.Request{
				ID:     1,
				Params: []interface{}{3.14},
			},
			expect: models.Response{
				ID:    1,
				Error: "Invalid parameter type",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Sort(tc.req)
			testutils.CheckResponse(t, tc.name, result, tc.expect)
		})
	}
}
