package testutils

import (
	"reflect"
	"rpc-server/internal/models"
	"testing"
)

func CheckResponse(t *testing.T, name string, result, expect models.Response) {
	if result.ID != expect.ID || result.ResultType != expect.ResultType || result.Error != expect.Error {
		t.Errorf("unexpected result for test %s: got %v, want %v", name, result, expect)
	}
	if !reflect.DeepEqual(result.Result, expect.Result) {
		t.Errorf("unexpected result for test %s: got %v, want %v", name, result.Result, expect.Result)
	}
}
