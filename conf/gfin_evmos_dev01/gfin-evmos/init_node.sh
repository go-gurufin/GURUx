#!/bin/bash

# 환경 변수 로딩
. ./env.sh

# jq 설치 여부에 대한 확인
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# 클라이언트 설정
evmosd config chain-id ${CHAINID} --home ${EVMOSD_HOME}
evmosd config keyring-backend ${KEYRING} --home ${EVMOSD_HOME}
evmosd config output json --home ${EVMOSD_HOME}

# Initialize the chain
evmosd init ${MONIKER} --chain-id ${CHAINID} --home ${EVMOSD_HOME}

# adding accounts
evmosd keys add ${KEYNAME} --keyring-backend=${KEYRING} --home ${EVMOSD_HOME} <<! > ${EVMOSD_HOME}/keys.txt
${PASSWORD}
${PASSWORD}
!

# Genesis 계정 추가 및 코인 발행
evmosd add-genesis-account ${KEYNAME} ${AMOUNT}${SYMBOL} --keyring-backend=${KEYRING} --home ${EVMOSD_HOME} <<!
${PASSWORD}
!

# Generate a genesis tx carrying a self delegation
evmosd gentx ${KEYNAME} 100000000000000000000${SYMBOL} \
  --home=${EVMOSD_HOME} \
  --keyring-backend=${KEYRING} \
  --chain-id=${CHAINID} \
  --moniker=${MONIKER} \
  --commission-max-change-rate=0.01 \
  --commission-max-rate=1.0 \
  --commission-rate=0.07 \
  --details="" \
  --security-contact="" \
  --website=${WEBSITE} <<! 
${PASSWORD}
!

# Change parameter token denominations to ${SYMBOL}
cat $EVMOSD_HOME/config/genesis.json | jq --arg symbol "${SYMBOL}" '.app_state["staking"]["params"]["bond_denom"]=$symbol' > $EVMOSD_HOME/config/tmp_genesis.json && mv $EVMOSD_HOME/config/tmp_genesis.json $EVMOSD_HOME/config/genesis.json
cat $EVMOSD_HOME/config/genesis.json | jq --arg symbol "${SYMBOL}" '.app_state["crisis"]["constant_fee"]["denom"]=$symbol' > $EVMOSD_HOME/config/tmp_genesis.json && mv $EVMOSD_HOME/config/tmp_genesis.json $EVMOSD_HOME/config/genesis.json
cat $EVMOSD_HOME/config/genesis.json | jq --arg symbol "${SYMBOL}" '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]=$symbol' > $EVMOSD_HOME/config/tmp_genesis.json && mv $EVMOSD_HOME/config/tmp_genesis.json $EVMOSD_HOME/config/genesis.json
cat $EVMOSD_HOME/config/genesis.json | jq --arg symbol "${SYMBOL}" '.app_state["evm"]["params"]["evm_denom"]=$symbol' > $EVMOSD_HOME/config/tmp_genesis.json && mv $EVMOSD_HOME/config/tmp_genesis.json $EVMOSD_HOME/config/genesis.json
cat $EVMOSD_HOME/config/genesis.json | jq --arg symbol "${SYMBOL}" '.app_state["inflation"]["params"]["mint_denom"]=$symbol' > $EVMOSD_HOME/config/tmp_genesis.json && mv $EVMOSD_HOME/config/tmp_genesis.json $EVMOSD_HOME/config/genesis.json
cat $EVMOSD_HOME/config/genesis.json | jq --arg symbol "${SYMBOL}" '.app_state["claims"]["params"]["claims_denom"]=$symbol' > $EVMOSD_HOME/config/tmp_genesis.json && mv $EVMOSD_HOME/config/tmp_genesis.json $EVMOSD_HOME/config/genesis.json

# Set inflation to false
cat $EVMOSD_HOME/config/genesis.json | jq '.app_state["inflation"]["params"]["enable_inflation"]=false' > $EVMOSD_HOME/config/tmp_genesis.json && mv $EVMOSD_HOME/config/tmp_genesis.json $EVMOSD_HOME/config/genesis.json

# Set gas limit in genesis(블록 가스 총량(Block GasLimit))
#cat $EVMOSD_HOME/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="40000000"' > $EVMOSD_HOME/config/tmp_genesis.json && mv $EVMOSD_HOME/config/tmp_genesis.json $EVMOSD_HOME/config/genesis.json

# if historical queries are needed, set pruning to "nothing". all historic states will be saved"
sed -i 's/pruning = "default"/pruning = "nothing"/' ${EVMOSD_HOME}/config/app.toml

# increase block time (?)
#cat $EVMOSD_HOME/config/genesis.json | jq '.consensus_params["block"]["time_iota_ms"]="30000"' > $EVMOSD_HOME/config/tmp_genesis.json && mv $EVMOSD_HOME/config/tmp_genesis.json $EVMOSD_HOME/config/genesis.json

# disable produce empty block
#sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' ${EVMOSD_HOME}/config/config.toml

# Genesis Transaction 취합
evmosd collect-gentxs --home ${EVMOSD_HOME}

# genesis.json file 오류 확인
evmosd validate-genesis --home ${EVMOSD_HOME}

# 서버 시작
#pm2 start pm2.evmosd.json

# Distribute genesis.json ...
#cp ${EVMOSD_HOME}/config/genesis.json /home/ubuntu
