// init client from clientPool.
// client is httpClient of tendermint

package helper

import (
	"context"
	"fmt"
	"github.com/tendermint/tendermint/types"
	"testing"

	"github.com/irisnet/irishub-sync/logger"
)

func TestInitClientPool(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]int, 6, 6)
	for index, value := range a {
		b[index] = value
	}
	b[3] = 4
}

func TestGetClient(t *testing.T) {
	client := GetClient()
	fmt.Println("====1======")
	defer func() {
		fmt.Println("====3======")
		if err := recover(); err != nil {
			logger.Debug("debug=======================recover=======================debug")
		}
	}()
	_, err := client.Status()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("====4======")

	txEventsCh := make(chan interface{})
	client.Start()
	client.Subscribe(context.Background(), "xxx", types.EventQueryValidatorSetUpdates, txEventsCh)
	for e := range txEventsCh {
		edt := e.(types.EventDataValidatorSetUpdates)
		fmt.Println(edt.ValidatorUpdates)
	}
}
