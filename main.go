package main

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

// GreetingInput represents the greeting operation request.
type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

type HumaOutput[R any] struct {
	Status int
	Body   Response[R]
}

type Response[R any] struct {
	Msg  string `json:"msg" doc:"detailed error message"`
	Code int    `json:"code" doc:"detailed business error code"`
	Data R      `json:"data" doc:"business data"`
}

func ParseParam(c context.Context, input *GreetingInput) (*string, error) {
	return &input.Name, nil
}
func HumaHandlerAdaptor[T0 any, T1 any, R any](parseParam func(c context.Context, input *T0) (*T1, error),
	processRequest func(c context.Context, input *T1) (*R, error)) func(ctx context.Context, input *T0) (*HumaOutput[R], error) {
	return func(ctx context.Context, input0 *T0) (*HumaOutput[R], error) {
		input1, _ := parseParam(ctx, input0)

		r, _ := processRequest(ctx, input1)
		return &HumaOutput[R]{
			Status: http.StatusOK,
			Body: Response[R]{
				Code: http.StatusOK,
				Msg:  "",
				Data: *r,
			},
		}, nil

	}
}

func main() {
	// Create a new router & API
	router := gin.New()
	api := humagin.New(router, huma.DefaultConfig("My API", "1.0.0"))

	// Register GET /greeting/{name}
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Summary:     "Get a greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
	},
		HumaHandlerAdaptor(ParseParam, LegacyHandler),
	)

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}
