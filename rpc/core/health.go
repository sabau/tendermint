package core

import (
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	rpctypes "github.com/tendermint/tendermint/rpc/lib/types"
)

// Get node health. Returns empty result (200 OK) on success, no response - in
// case of an error.
//
// ```shell
// curl 'localhost:26657/health'
// ```
//
// ```go
// client := client.NewHTTP("tcp://0.0.0.0:26657", "/websocket")
// err := client.Start()
// if err != nil {
//   // handle error
// }
// defer client.Stop()
// result, err := client.Health()
// ```
//
// > The above command returns JSON structured like this:
//
// ```json
// {
// 	"error": "",
// 	"result": {},
// 	"id": "",
// 	"jsonrpc": "2.0"
// }
// ```

// swagger:route GET /health health Health
//
// Get node health.
//
// Returns empty result (200 OK) on success, no response - in
// case of an error.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: ResponseResultHealth
//       500: ResponseResultHealth
func Health(ctx *rpctypes.Context) (*ctypes.ResultHealth, error) {
	return &ctypes.ResultHealth{}, nil
}
