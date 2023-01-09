package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/easierway/concurrency_utils"
)

func TestAsynExecutor(t *testing.T) {
	TimeCosumingOps := func(ctx context.Context) concurrency.TaskResult {
		time.Sleep(time.Second * 1)
		return concurrency.TaskResult{
			"Hello Asynchrony", nil,
		}
	}
	// We can invok time-cosuming firstly.
	retStub := concurrency.AsynExecutor(context.TODO(),
		TimeCosumingOps, 2000)
	// Do something else.
	fmt.Println("Do something else")
	// Get the results while we just need it
	fmt.Println("Now, we need the result", retStub.GetResult().Result)
}
