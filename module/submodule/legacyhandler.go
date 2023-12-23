package submodule

import "context"

type BizData struct {
	Greeting string
	Name     string
}

func LegacyHandler(ctx context.Context, name *string) (*BizData, error) {
	return &BizData{
		Greeting: "Hello world ",
		Name:     *name,
	}, nil
}
