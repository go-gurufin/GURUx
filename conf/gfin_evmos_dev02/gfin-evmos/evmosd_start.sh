#!/bin/bash

. ./env.sh

evmosd start \
  --home ${EVMOSD_HOME} \
  --p2p.laddr "tcp://0.0.0.0:26656" \
  --rpc.laddr "tcp://127.0.0.1:26657" \
  --json-rpc.address "0.0.0.0:8545" \
  --json-rpc.ws-address "127.0.0.1:8546" \
  --json-rpc.api "eth,txpool,net,web3" \
  --keyring-backend ${KEYRING} \
  --minimum-gas-prices=0${SYMBOL} \
  --log_level "info" \
  --pruning "nothing" \
  --trace << !
${PASSWORD}
!

#Usage:
#  evmosd start [flags]

#Flags:
#      --abci string                                     specify abci transport (socket | grpc) (default "socket")
#      --address string                                  Listen address (default "tcp://0.0.0.0:26658")
#      --consensus.create_empty_blocks                   set this to false to only produce blocks when there are txs or when the AppHash changes (default true)
#      --consensus.create_empty_blocks_interval string   the possible interval between empty blocks (default "0s")
#      --consensus.double_sign_check_height int          how many blocks to look back to check existence of the node's consensus votes before joining consensus
#      --cpu-profile string                              Enable CPU profiling and write to the provided file
#      --db_backend string                               database backend: goleveldb | cleveldb | boltdb | rocksdb | badgerdb (default "goleveldb")
#      --db_dir string                                   database directory (default "data")
#      --evm.tracer string                               the EVM tracer type to collect execution traces from the EVM transaction execution (json|struct|access_list|markdown)
#      --fast_sync                                       fast blockchain syncing (default true)
#      --genesis_hash bytesHex                           optional SHA-256 hash of the genesis file
#      --grpc-web.address string                         The gRPC-Web server address to listen on (default "0.0.0.0:9091")
#      --grpc-web.enable                                 Define if the gRPC-Web server should be enabled. (Note: gRPC must also be enabled.) (default true)
#      --grpc.address string                             the gRPC server address to listen on (default "0.0.0.0:9090")
#      --grpc.enable                                     Define if the gRPC server should be enabled (default true)
#      --halt-height uint                                Block height at which to gracefully halt the chain and shutdown the node
#      --halt-time uint                                  Minimum block time (in Unix seconds) at which to gracefully halt the chain and shutdown the node
#  -h, --help                                            help for start
#      --inter-block-cache                               Enable inter-block caching (default true)
#      --inv-check-period uint                           Assert registered invariants every N blocks
#      --json-rpc.address string                         the JSON-RPC server address to listen on (default "0.0.0.0:8545")
#      --json-rpc.api strings                            Defines a list of JSON-RPC namespaces that should be enabled (default [eth,net,web3])
#      --json-rpc.enable                                 Define if the gRPC server should be enabled (default true)
#      --json-rpc.evm-timeout duration                   Sets a timeout used for eth_call (0=infinite) (default 5s)
#      --json-rpc.filter-cap int32                       Sets the global cap for total number of filters that can be created (default 200)
#      --json-rpc.gas-cap uint                           Sets a cap on gas that can be used in eth_call/estimateGas unit is aphoton (0=infinite) (default 25000000)
#      --json-rpc.txfee-cap float                        Sets a cap on transaction fee that can be sent via the RPC APIs (1 = default 1 photon) (default 1)
#      --json-rpc.ws-address string                      the JSON-RPC WS server address to listen on (default "0.0.0.0:8546")
#      --min-retain-blocks uint                          Minimum block height offset during ABCI commit to prune Tendermint blocks
#      --minimum-gas-prices string                       Minimum gas prices to accept for transactions; Any fee in a tx must meet this minimum (e.g. 0.01photon;0.0001stake)
#      --moniker string                                  node name (default "gfin-evm-dev01")
#      --p2p.laddr string                                node listen address. (0.0.0.0:0 means any interface, any port) (default "tcp://0.0.0.0:26656")
#      --p2p.persistent_peers string                     comma-delimited ID@host:port persistent peers
#      --p2p.pex                                         enable/disable Peer-Exchange (default true)
#      --p2p.private_peer_ids string                     comma-delimited private peer IDs
#      --p2p.seed_mode                                   enable/disable seed mode
#      --p2p.seeds string                                comma-delimited ID@host:port seed nodes
#      --p2p.unconditional_peer_ids string               comma-delimited IDs of unconditional peers
#      --p2p.upnp                                        enable/disable UPNP port forwarding
#      --priv_validator_laddr string                     socket address to listen on for connections from external priv_validator process
#      --proxy_app string                                proxy app address, or one of: 'kvstore', 'persistent_kvstore', 'counter', 'e2e' or 'noop' for localtesting. (default "tcp://127.0.0.1:26658")
#      --pruning string                                  Pruning strategy (default|nothing|everything|custom) (default "default")
#      --pruning-interval uint                           Height interval at which pruned heights are removed from disk (ignored if pruning is not 'custom')
#      --pruning-keep-every uint                         Offset heights to keep on disk after 'keep-every' (ignored if pruning is not 'custom')
#      --pruning-keep-recent uint                        Number of recent heights to keep on disk (ignored if pruning is not 'custom')
#      --rpc.grpc_laddr string                           GRPC listen address (BroadcastTx only). Port required
#      --rpc.laddr string                                RPC listen address. Port required (default "tcp://127.0.0.1:26657")
#      --rpc.pprof_laddr string                          pprof listen address (https://golang.org/pkg/net/http/pprof)
#      --rpc.unsafe                                      enabled unsafe rpc methods
#      --state-sync.snapshot-interval uint               State sync snapshot interval
#      --state-sync.snapshot-keep-recent uint32          State sync snapshot to keep (default 2)
#      --tls.certificate-path string                     the cert.pem file path for the server TLS configuration
#      --tls.key-path string                             the key.pem file path for the server TLS configuration
#      --trace-store string                              Enable KVStore tracing to an output file
#      --transport string                                Transport protocol: socket, grpc (default "socket")
#      --unsafe-skip-upgrades ints                       Skip a set of upgrade heights to continue the old binary
#      --with-tendermint                                 Run abci app embedded in-process with tendermint (default true)
#      --x-crisis-skip-assert-invariants                 Skip x/crisis invariants check on startup

#Global Flags:
#  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "sync")
#      --chain-id string          Specify Chain ID for sending Tx (default "testnet")
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


