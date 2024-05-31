package services

import (
	"rpc-server/internal/models"
	"rpc-server/internal/testutils"
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
		{
			name: "valid parameter type",
			req: models.Request{
				ID:     1,
				Params: []interface{}{1.3},
			},
			expect: models.Response{
				ID:    1,
				Error: "Invalid parameter type",
			},
		},
		{
			name: "valid parameter count",
			req: models.Request{
				ID:     1,
				Params: []interface{}{"あいうえお", "アイウエオ"},
			},
			expect: models.Response{
				ID:    1,
				Error: "Invalid parameter count",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Reverse(tc.req)
			testutils.CheckResponse(t, tc.name, result, tc.expect)
		})
	}
}
