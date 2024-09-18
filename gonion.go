package gonion

import (
	"context"
	"fmt"
)

type Srv struct {
	Dec *Decorator
}
type SrvHandler func(ctx context.Context) (context.Context, error)
type Decorator struct {
	Fun  SrvHandler
	Next *Decorator
	Stat int
}

func BlankHand() Srv {
	return Srv{
		Dec: &Decorator{
			Next: nil,
			Stat: ONION_WAIT,
			Fun:  nil,
		},
	}
}
func mounted(dec *Decorator, targetDec *Decorator) {
	if dec.Fun == nil && dec.Next == nil {
		dec.Fun = targetDec.Fun
	} else if dec.Next != nil {
		mounted(dec.Next, targetDec)
	} else {
		dec.Next = targetDec
	}
}

// At 装饰
func (srv *Srv) At(Handler SrvHandler) *Srv {
	//使用递归将处理器放置到内部
	dec := &Decorator{
		Fun:  Handler,
		Next: nil,
		Stat: ONION_WAIT,
	}
	mounted(srv.Dec, dec)
	return srv
}
func (srv *Srv) Exe(Handler SrvHandler) *Srv {
	dec := &Decorator{
		Fun:  Handler,
		Next: nil,
		Stat: ONION_WAIT,
	}
	mounted(srv.Dec, dec)
	return srv
}

func (srv *Srv) Exec(ctx context.Context) (context.Context, error) {
	ctx = context.WithValue(ctx, DEC_CTX, srv.Dec)
	return Next(ctx)
}
func Next(ctx context.Context) (context.Context, error) {
	var err error
	dec, ok := ctx.Value(DEC_CTX).(*Decorator)
	if !ok {
		dec.Stat = ONION_REJECT
		ctx = context.WithValue(ctx, DEC_CTX, dec)
		return ctx, fmt.Errorf("type is error")
	}

	dec.Stat = ONION_PENDING
	ctx = context.WithValue(ctx, DEC_CTX, dec.Next)
	ctx, err = dec.Fun(ctx)

	if dec.Next == nil {
		dec.Stat = ONION_RESLOVE
		ctx = context.WithValue(ctx, DEC_CTX, dec)
		return ctx, nil
	}

	return ctx, err
}
