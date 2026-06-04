# go-bitshares

Go client library for the BitShares blockchain. Connect to a full node over WebSocket, or to `cli_wallet` over HTTP RPC. Serialize and sign transactions without cgo.

**Maintained fork:** [github.com/pi314x/go-bitshares](https://github.com/pi314x/go-bitshares) — use this repo for updates, fixes, and new operations (HTLC, liquidity pools, SametFund, credit offers, tickets, and more).

**Upstream:** [github.com/denkhaus/go-bitshares](https://github.com/denkhaus/go-bitshares) — original library; `go get github.com/denkhaus/bitshares` without a `replace` still resolves there and will not include fork changes.

Imports stay `github.com/denkhaus/bitshares/...` for compatibility; point the module at pi314x with `replace` (see below).

Examples: [examples](/examples). Tests: [tests](/tests).

## Install

Clone or browse the fork:

```bash
git clone https://github.com/pi314x/go-bitshares.git
```

In a consuming project, **replace** the module path with pi314x, then resolve dependencies (the fork’s `go.mod` still declares `module github.com/denkhaus/bitshares`):

```bash
go mod edit -replace=github.com/denkhaus/bitshares=github.com/pi314x/go-bitshares@master
go get github.com/denkhaus/bitshares@master
go mod tidy
```

`go mod tidy` writes a pseudo-version (not the literal string `master`) into `go.mod`. Pin a commit for reproducible builds:

```bash
go mod edit -replace=github.com/denkhaus/bitshares=github.com/pi314x/go-bitshares@f99cfca
go get github.com/denkhaus/bitshares@f99cfca
go mod tidy
```

Local checkout during development:

```go
replace github.com/denkhaus/bitshares => ../go-bitshares
```

Install dev dependencies:

```bash
make init
```

## Code generation

Types and operations use [ffjson](https://github.com/pquerna/ffjson) for fast JSON marshaling. After changing types or operations, regenerate helpers:

```bash
make generate
```

If generation fails, regenerate from scratch:

```bash
make generate_new
```

Object ID helpers are generated from [types/gen.go](/types/gen.go) via `go generate`.

## Operation samples

The [gen](/gen) package fetches real operations from the chain and writes sample code under `gen/samples/`. Those samples are injected automatically when running operation tests. Rebuild samples with:

```bash
make opsamples
```

## Testing

Operation tests use blockchain samples from [gen](/gen). API tests expect a reachable public node (or your own). Wallet tests expect a local `cli_wallet` RPC endpoint.

```bash
make test_operations
make test_api
```

Long-running block serialize/deserialize comparison:

```bash
make test_blocks
```

For wallet and operation tests against a local chain, see [bitshares-docker](https://github.com/denkhaus/bitshares-docker).

## Examples

| Example | Description |
|---------|-------------|
| [examples/transfer](/examples/transfer) | Sign and broadcast a transfer on testnet |
| [examples/limitorder](/examples/limitorder) | Create a limit order (requires `BTS_TEST_ACCOUNT`, `BTS_TEST_WIF`) |
| [examples/memo](/examples/memo) | Memo encryption helpers |
| [examples/history](/examples/history) | Account history queries |

## Usage

Default public nodes (XBTS): mainnet `wss://node.xbts.io/ws`, testnet `wss://testnet.xbts.io/ws`.

### WebSocket API (full node)

```go
package main

import (
	"log"

	"github.com/denkhaus/bitshares"
	"github.com/denkhaus/bitshares/types"
)

func main() {
	api := bitshares.NewWebsocketAPI("wss://node.xbts.io/ws")
	if err := api.Connect(); err != nil {
		log.Fatal(err)
	}

	api.OnError(func(err error) {
		log.Fatal(err)
	})

	userID := types.NewAccountID("1.2.253")
	assetBTS := types.NewAssetID("1.3.0")

	balances, err := api.GetAccountBalances(userID, assetBTS)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("balances: %v", balances)
}
```

### Wallet API (`cli_wallet` RPC)

```go
api := bitshares.NewWalletAPI("http://localhost:8095")

if err := api.Connect(); err != nil {
	log.Fatal(err)
}
```

### Auto-selected node (latency tester)

For long-lived apps, `NewWebsocketAPIWithAutoEndpoint` probes public nodes and picks a low-latency endpoint. The URL passed in is used until the first probe completes.

```go
api, err := bitshares.NewWebsocketAPIWithAutoEndpoint("wss://node.xbts.io/ws")
if err != nil {
	log.Fatal(err)
}

if err := api.Connect(); err != nil {
	log.Fatal(err)
}

api.OnError(func(err error) {
	log.Fatal(err)
})
```

For one-off scripts, prefer `NewWebsocketAPI` to avoid probe startup delay.

### Testnet

Use the testnet endpoint and set chain config before signing (see [examples/transfer](/examples/transfer)):

```go
config.SetCurrent(config.ChainIDTest)
api := bitshares.NewWebsocketAPI("wss://testnet.xbts.io/ws")
```

## Implemented operations

Operations registered in `types.OperationMap` support binary serialize/deserialize. Items marked *(virtual)* are chain-emitted, not user-broadcast. Only a subset have round-trip tests against chain samples in [tests/operation_test.go](/tests/operation_test.go); the checklist below means “implemented in code,” not “fully tested.”

### Core (0–46)

- [x] OperationTypeTransfer
- [x] OperationTypeLimitOrderCreate
- [x] OperationTypeLimitOrderCancel
- [x] OperationTypeCallOrderUpdate
- [x] OperationTypeFillOrder *(virtual)*
- [x] OperationTypeAccountCreate
- [x] OperationTypeAccountUpdate
- [x] OperationTypeAccountWhitelist
- [x] OperationTypeAccountUpgrade
- [x] OperationTypeAccountTransfer
- [x] OperationTypeAssetCreate
- [x] OperationTypeAssetUpdate
- [x] OperationTypeAssetUpdateBitasset
- [x] OperationTypeAssetUpdateFeedProducers
- [x] OperationTypeAssetIssue
- [x] OperationTypeAssetReserve
- [x] OperationTypeAssetFundFeePool
- [x] OperationTypeAssetSettle
- [x] OperationTypeAssetGlobalSettle
- [x] OperationTypeAssetPublishFeed
- [x] OperationTypeWitnessCreate
- [x] OperationTypeWitnessUpdate
- [x] OperationTypeProposalCreate
- [x] OperationTypeProposalUpdate
- [x] OperationTypeProposalDelete
- [x] OperationTypeWithdrawPermissionCreate
- [x] OperationTypeWithdrawPermissionUpdate
- [x] OperationTypeWithdrawPermissionClaim
- [x] OperationTypeWithdrawPermissionDelete
- [x] OperationTypeCommitteeMemberCreate
- [x] OperationTypeCommitteeMemberUpdate
- [x] OperationTypeCommitteeMemberUpdateGlobalParameters
- [x] OperationTypeVestingBalanceCreate
- [x] OperationTypeVestingBalanceWithdraw
- [x] OperationTypeWorkerCreate
- [x] OperationTypeCustom
- [x] OperationTypeAssert
- [x] OperationTypeBalanceClaim
- [x] OperationTypeOverrideTransfer
- [x] OperationTypeTransferToBlind
- [x] OperationTypeBlindTransfer
- [x] OperationTypeTransferFromBlind
- [x] OperationTypeAssetSettleCancel
- [x] OperationTypeAssetClaimFees
- [x] OperationTypeFBADistribute
- [x] OperationTypeBidCollateral
- [x] OperationTypeExecuteBid

### BitShares core v3+ (47+)

- [x] OperationTypeAssetClaimPool
- [x] OperationTypeAssetUpdateIssuer
- [x] OperationTypeHTLCCreate
- [x] OperationTypeHTLCRedeem
- [x] OperationTypeHTLCRedeemed *(virtual)*
- [x] OperationTypeHTLCExtend
- [x] OperationTypeHTLCRefund *(virtual)*
- [x] OperationTypeCustomAuthorityCreate
- [x] OperationTypeCustomAuthorityUpdate
- [x] OperationTypeCustomAuthorityDelete
- [x] OperationTypeTicketCreate
- [x] OperationTypeTicketUpdate
- [x] OperationTypeLiquidityPoolCreate
- [x] OperationTypeLiquidityPoolDelete
- [x] OperationTypeLiquidityPoolDeposit
- [x] OperationTypeLiquidityPoolWithdraw
- [x] OperationTypeLiquidityPoolExchange
- [x] OperationTypeLiquidityPoolUpdate
- [x] OperationTypeSametFundCreate
- [x] OperationTypeSametFundDelete
- [x] OperationTypeSametFundUpdate
- [x] OperationTypeSametFundBorrow
- [x] OperationTypeSametFundRepay
- [x] OperationTypeCreditOfferCreate
- [x] OperationTypeCreditOfferDelete
- [x] OperationTypeCreditOfferUpdate
- [x] OperationTypeCreditOfferAccept
- [x] OperationTypeCreditDealRepay
- [x] OperationTypeCreditDealExpired *(virtual)*
- [x] OperationTypeCreditDealUpdate
- [x] OperationTypeLimitOrderUpdate

## Contributing

Open issues and pull requests on [pi314x/go-bitshares](https://github.com/pi314x/go-bitshares). New operations, samples in [gen](/gen), and tests in [tests](/tests) are welcome. Use at your own risk on mainnet.
