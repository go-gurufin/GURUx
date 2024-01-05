#!/bin/bash

. ../env.sh

if [ $# != 2 ]
then
  echo "Usage: ${0} [address] [amount(${SYMBOL}})]"
  exit 1
fi
amt=`echo "${2} * 1000000000000000000" | bc -l | awk -F . '{print $1}'`
fees="0${SYMBOL}"

evmosd tx bank send ${KEYNAME} ${1} ${amt}${SYMBOL} --fees $fees -b block -y --home=${EVMOSD_HOME} --keyring-backend $KEYRING << !
${PASSWORD}
!

#Usage:
#  evmosd tx bank send [from_key_or_address] [to_address] [amount] [flags]

#Flags:
#  -a, --account-number uint   The account number of the signing account (offline mode only)
#      --dry-run               ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it
#      --fee-account string    Fee account pays fees for the transaction instead of deducting from the signer
#      --gas string            gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
#      --generate-only         Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
#  -h, --help                  help for send
#      --keyring-dir string    The client Keyring directory; if omitted, the default 'home' directory will be used
#      --ledger                Use a connected Ledger device
#      --note string           Note to add a description to the transaction (previously --memo)
#      --offline               Offline mode (does not allow any online functionality
#  -o, --output string         Output format (text|json) (default "json")
#  -s, --sequence uint         The sequence number of the signing account (offline mode only)
#      --sign-mode string      Choose sign mode (direct|amino-json), this is an advanced feature
#      --timeout-height uint   Set a block timeout height to prevent the tx from being committed past a certain height
#  -y, --yes                   Skip tx broadcasting prompt confirmation
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
#      --trace                    print out full stack trace on errors
