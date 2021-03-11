package anko

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func TestRunThen(t *testing.T) {
	type Each struct {
		Name           string
		Code           string
		ExpectedResult interface{}
		ExpectedError  error
	}
	for _, each := range []*Each{
		{Name: "Replace data with a return interface version", ExpectedError: nil, ExpectedResult: map[interface{}]interface{}{"T": (1)}, Code: `data = map[interface{}]interface{}{"T": 1}`},
		{Name: "Replace data with a return string version", ExpectedError: nil, ExpectedResult: map[string]interface{}{"T": (1)}, Code: `data = map[string]interface{}{"T": 1}`},
	} {
		t.Run(each.Name, func(t *testing.T) {
			ctx, cf := context.WithTimeout(context.Background(), 1*time.Minute)
			defer cf()
			actualResult, err := Run(ctx, each.Code)
			if err != each.ExpectedError {
				t.Errorf("Error with error. Expected %v got %v", each.ExpectedError, err)
			}
			if diff := cmp.Diff(actualResult, each.ExpectedResult); diff != "" {
				t.Errorf("Data doesn't match: %v", diff)
			}
		})
	}

}
