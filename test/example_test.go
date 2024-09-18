package test

import (
	"context"
	"fmt"
	"gonion"
	"testing"
)

func example1(ctx context.Context) (context.Context, error) {
	fmt.Println("[1] start")
	fmt.Println("[1] next")
	ctx, err := gonion.Next(ctx)
	if err != nil {
		return nil, err
	}
	port, ok := ctx.Value("network_port").(int)
	if !ok {
		return ctx, fmt.Errorf("[1] network_port no find")
	}
	fmt.Println("[1] get port by ctx: ", port)
	fmt.Println("[1] done")
	return ctx, nil
}
func example2(ctx context.Context) (context.Context, error) {
	fmt.Println("[2] start")
	fmt.Println("[2] set ctx network_port")
	ctx = context.WithValue(ctx, "network_port", 8888)
	fmt.Println("[2] next")
	ctx, err := gonion.Next(ctx)
	if err != nil {
		return ctx, err
	}
	fmt.Println("[2] done")
	return ctx, nil
}
func example3(ctx context.Context) (context.Context, error) {
	fmt.Println("[3] next")
	ctx, err := gonion.Next(ctx)
	if err != nil {
		return ctx, err
	}
	fmt.Println("[3] start")
	fmt.Println("[3] get ctx network_port")
	port, ok := ctx.Value("network_port").(int)
	if !ok {
		return ctx, fmt.Errorf("[3] network_port no find")
	}
	fmt.Println("[3] get port by ctx: ", port)
	fmt.Println("[3] done")
	return ctx, nil
}

func TestOnionFunction(t *testing.T) {
	exe1 := func() gonion.SrvHandler {
		//define your cfg...
		str := "hello world"
		return func(ctx context.Context) (context.Context, error) {
			fmt.Println(str)
			return ctx, nil
		}
	}
	srv := gonion.BlankHand()
	ctx, err := srv.
		At(example1).
		At(example2).
		At(example3).
		Exe(exe1()).
		Exec(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("ctx: %v", ctx)
}
