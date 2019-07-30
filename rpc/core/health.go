package core

import (
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	rpctypes "github.com/tendermint/tendermint/rpc/lib/types"
)

// @Summary Get node health
// @Description Get node health. Returns empty result (200 OK) on success, no response in case of an error
// @Produce json
// @Success 200 {object} github.com/tendermint/tendermint/rpc/core/types.ResultHealth "result will be empty"
// @Router /health [get]
func Health(ctx *rpctypes.Context) (*ctypes.ResultHealth, error) {
	return &ctypes.ResultHealth{}, nil
}
