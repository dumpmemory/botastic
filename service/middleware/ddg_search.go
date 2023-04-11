package middleware

import (
	"context"
	"fmt"

	"github.com/pandodao/botastic/core"
	"github.com/pandodao/botastic/internal/ddg"
)

type duckDuckGoSearch struct{}

func (m *duckDuckGoSearch) Name() string {
	return core.MiddlewareDuckduckgoSearch
}

type duckDuckGoSearchOptions struct {
	*generalOptions
	Limit int `json:"limit"`
}

func (m *duckDuckGoSearch) ValidateOptions(gopts *generalOptions, opts map[string]any) (any, error) {
	options := &duckDuckGoSearchOptions{
		generalOptions: gopts,
		Limit:          3,
	}

	if val, ok := opts["limit"]; ok {
		v, ok := val.(float64)
		if !ok {
			return nil, fmt.Errorf("limit is not a number: %v", val)
		}

		if v <= 0 || float64(int(v)) != v {
			return nil, fmt.Errorf("limit is not a positive integer: %v", v)
		}
		options.Limit = int(v)
	}

	return options, nil
}

func (m *duckDuckGoSearch) Process(ctx context.Context, opts any, input string) (string, error) {
	options := opts.(*duckDuckGoSearchOptions)
	r, err := ddg.WebSearch(ctx, input, options.Limit)
	if err != nil {
		return "", err
	}
	result, err := r.Text()
	if err != nil {
		return "", err
	}
	return result, nil
}