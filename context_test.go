package golangcontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {

	Background := context.Background()
	fmt.Println(Background)

	Todo := context.TODO()
	fmt.Println(Todo)

}
