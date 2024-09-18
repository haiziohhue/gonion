# GONION

## Use
```go
package example

import (
	"context"
	"github.com/haiziohhue/gonion"
)

// create a decorator, use next continue
func example1(ctx context.Context) (context.Context, error) {
	ctx, err := gonion.Next(ctx)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

// use it
func main() {
	srv := gonion.BlankHand()
	ctx, err := srv.
		At(example1).
		At(example2).
		At(example3).
		Exe(exe1).
		Exec(context.TODO())
}
```
## Install
```bash
go get github.com/haiziohhue/gonion
```
## License
MIT
