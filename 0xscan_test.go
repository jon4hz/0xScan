package xscan

import (
	"context"
	"testing"
)

func TestGetErc20TokenTransfers(t *testing.T) {
	client := NewClient(BscOpts, "")

	_, _, err := client.GetErc20TokenTransfers(context.Background(), "0xaf5fFcce7CA5115e3dEeD9Be8B7a66199121cE0D", nil)
	if err != nil {
		t.Errorf("GetErc20TokenTransfers error: %v", err)
	}
}
