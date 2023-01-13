package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/easierway/concurrency_utils"
)

func TestAsynExecutor(t *testing.T) {
	retArry := [5]int{1, 2, 3, 4, 5}
	TimeCosumingOps := func(ctx context.Context) concurrency.TaskResult {
		time.Sleep(time.Second * 1)
		return concurrency.TaskResult{
			&retArry, nil,
		}
	}
	// We can invok time-cosuming firstly.
	retStub := concurrency.AsynExecutor(context.TODO(),
		TimeCosumingOps, 2000)
	// Do something else.
	fmt.Println("Do something else")
	// Get the results while we just need it
	if retStub.GetResult().Err != nil {
		fmt.Println("The error is", retStub.GetResult().Err)
		t.Error(retStub.GetResult().Err)
	}
	ret, _ := retStub.GetResult().Result.(*[5]int)
	fmt.Println("Now, we need the result", *ret)
}
