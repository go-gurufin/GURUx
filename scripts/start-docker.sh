#!/bin/bash

KEY="dev0"
CHAINID="gurux_9000-1"
MONIKER="mymoniker"
DATA_DIR="/home"

echo "create and add new keys"
./guruxd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Gurux with moniker=$MONIKER and chain-id=$CHAINID"
./guruxd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./guruxd add-genesis-account \
"$(./guruxd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000agurux,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./guruxd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./guruxd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./guruxd validate-genesis --home $DATA_DIR

echo "starting gurux node $i in background ..."
./guruxd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown 

echo "started gurux node"
tail -f /dev/null