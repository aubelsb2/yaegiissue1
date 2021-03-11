package anko

import (
	"context"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func Run(ctx context.Context, then string) (interface{}, error) {
	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)

	_, err := i.EvalWithContext(ctx, `
var data interface{} = nil
`)
	if err != nil {
		return nil, err
	}

	_, err = i.EvalWithContext(ctx, then)
	if err != nil {
		return nil, err
	}

	res, err := i.Eval("data")
	if err != nil {
		return nil, err
	}

	return res.Interface(), nil
}
