package services

import (
	"rpc-server/internal/models"
	"rpc-server/internal/testutils"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		name   string
		req    models.Request
		expect models.Response
	}{
		{
			name: "valid input",
			req: models.Request{
				ID:     1,
				Params: []interface{}{"あいうえお", "いあえうお"},
			},
			expect: models.Response{
				ID:         1,
				Result:     true,
				ResultType: "bool",
				Error:      "",
			},
		},
		{
			name: "アナグラムにならないケース",
			req: models.Request{
				ID:     1,
				Params: []interface{}{"あいうえお", "いあえうおか"},
			},
			expect: models.Response{
				ID:         1,
				Result:     false,
				ResultType: "bool",
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
				Error: "Invalid parameter type",
			},
		},
		{
			name: "valid parameter count",
			req: models.Request{
				ID:     1,
				Params: []interface{}{"テスト文字列"},
			},
			expect: models.Response{
				ID:    1,
				Error: "Invalid parameter count",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ValidAnagram(tc.req)
			testutils.CheckResponse(t, tc.name, result, tc.expect)
		})
	}
}
