#!/bin/bash

# 환경 변수 로딩
. ./env.sh

# 밸리데이터 트랜잭션 
evmosd tx staking create-validator \
	--home=${EVMOSD_HOME} \
  --keyring-backend=${KEYRING} \
	--amount=100000000000000000000${SYMBOL} \
	--pubkey=$(evmosd tendermint show-validator --home=${EVMOSD_HOME}) \
	--node-id=$(evmosd tendermint show-node-id --home=${EVMOSD_HOME}) \
  --chain-id=${CHAINID} \
  --moniker=${MONIKER2} \
  --commission-max-change-rate=0.01 \
  --commission-max-rate=1.0 \
  --commission-rate=0.07 \
	--min-self-delegation="1" \
	--from=${KEYNAME2} \
  --yes \
  --website=${WEBSITE} <<! 
${PASSWORD}
!

#create new validator initialized with a self-delegation to it
#Usage:
#  evmosd tx staking create-validator [flags]

#Flags:
#  -a, --account-number uint                 The account number of the signing account (offline mode only)
#      --amount string                       Amount of coins to bond
#      --commission-max-change-rate string   The maximum commission change rate percentage (per day)
#      --commission-max-rate string          The maximum commission rate percentage
#      --commission-rate string              The initial commission rate percentage
#      --details string                      The validator's (optional) details
#      --dry-run                             ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
#      --fee-account string                  Fee account pays fees for the transaction instead of deducting from the signer
#      --gas string                          gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
#      --generate-only                       Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
#  -h, --help                                help for create-validator
#      --identity string                     The optional identity signature (ex. UPort or Keybase)
#      --ip string                           The node's public IP. It takes effect only when used in combination with --generate-only
#      --keyring-dir string                  The client Keyring directory; if omitted, the default 'home' directory will be used
#      --ledger                              Use a connected Ledger device
#      --min-self-delegation string          The minimum self delegation required on the validator
#      --moniker string                      The validator's name
#      --node-id string                      The node's ID
#      --note string                         Note to add a description to the transaction (previously --memo)
#      --offline                             Offline mode (does not allow any online functionality
#  -o, --output string                       Output format (text|json) (default "json")
#      --pubkey string                       The validator's Protobuf JSON encoded public key
#      --security-contact string             The validator's (optional) security contact email
#  -s, --sequence uint                       The sequence number of the signing account (offline mode only)
#      --sign-mode string                    Choose sign mode (direct|amino-json), this is an advanced feature
#      --timeout-height uint                 Set a block timeout height to prevent the tx from being committed past a certain height
#      --website string                      The validator's (optional) website
#  -y, --yes                                 Skip tx broadcasting prompt confirmation
#
#Global Flags:
#  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
#      --chain-id string          The network chain ID
#      --fees string              Fees to pay along with transaction; eg: 10aphoton
#      --from string              Name or address of private key with which to sign
#      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
#      --gas-prices string        Gas prices to determine the transaction fee (e.g. 10aphoton)
#      --home string              directory for config and data (default "/root/.evmosd")
#      --keyring-backend string   Select keyring's backend (default "os")
#      --log_format string        The logging format (json|plain) (default "plain")
#      --log_level string         The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
#      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")



