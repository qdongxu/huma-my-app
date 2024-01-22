package submodulex

import "context"

type Generic[T any] struct {
	Data T
}

type BizData struct {
	Greeting string
	Name     string
}

func BizHandler(ctx context.Context, name *string) (*Generic[BizData], error) {
	return &Generic[BizData]{
		Data: BizData{
			Greeting: "Hello world ",
			Name:     *name,
		},
	}, nil
}

type OrderData struct {
	Count int
	Price int
	Me    string
}

func OrderHandler(ctx context.Context, name *string) (*Generic[OrderData], error) {
	return &Generic[OrderData]{
		Data: OrderData{
			Count: 1,
			Price: 1,
		},
	}, nil
}
