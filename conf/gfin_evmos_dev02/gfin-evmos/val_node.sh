#!/bin/bash

# 환경 변수 로딩
. ./env.sh

# 클라이언트 설정
evmosd config chain-id ${CHAINID} --home ${EVMOSD_HOME}
evmosd config keyring-backend ${KEYRING} --home ${EVMOSD_HOME}
evmosd config output json --home ${EVMOSD_HOME}

# Initialize the chain
evmosd init ${MONIKER2} --chain-id ${CHAINID} --home ${EVMOSD_HOME}

# adding accounts
evmosd keys add ${KEYNAME2} --keyring-backend=${KEYRING} --home ${EVMOSD_HOME} <<! > ${EVMOSD_HOME}/keys.txt
${PASSWORD}
${PASSWORD}
!

# genesis.json file copy
scp ubuntu@${HOST}:./genesis.json ${EVMOSD_HOME}/config

# vi /root/evm_home/config/config.toml
# seeds= <노드1 ID>@<노드1 내부사설IP>:26656
# evmosd tendermint show-node-id --home=/root/evm_home

SEEDS="\"${VALID}@${IP}:26656\""
mv ${EVMOSD_HOME}/config/config.toml ${EVMOSD_HOME}/config/config.toml.bak
awk -v val=$SEEDS '/^seeds/{gsub($3, val)};{print}' ${EVMOSD_HOME}/config/config.toml.bak > ${EVMOSD_HOME}/config/config.toml

# 설정 파일 파라미터 수정
sed -i "s/^pruning = \"default\"/pruning = \"nothing\"/" ${EVMOSD_HOME}/config/app.toml
sed -i "s/^max-open-connections = 1000/max-open-connections = 10000/" ${EVMOSD_HOME}/config/app.toml
sed -i "s/^enable = false/enable = true/" ${EVMOSD_HOME}/config/app.toml
sed -i "s/^swagger = false/swagger = true/" ${EVMOSD_HOME}/config/app.toml

sed -i "s/^persistent_peers = \"\"/persistent_peers = $SEEDS/" ${EVMOSD_HOME}/config/config.toml
sed -i "s/^addr_book_strict = true/addr_book_strict = false/" ${EVMOSD_HOME}/config/config.toml
sed -i "s/^max_packet_msg_payload_size = 1024/max_packet_msg_payload_size = 10240/" ${EVMOSD_HOME}/config/config.toml
sed -i "s/^grpc_max_open_connections = 900/grpc_max_open_connections = 10000/" ${EVMOSD_HOME}/config/config.toml
sed -i "s/^max_open_connections = 900/max_open_connections = 10000/" ${EVMOSD_HOME}/config/config.toml
sed -i "s/^size = 5000/size = 10000/" ${EVMOSD_HOME}/config/config.toml
sed -i "s/^max_txs_bytes = 1073741824/max_txs_bytes = 3221225472/" ${EVMOSD_HOME}/config/config.toml
sed -i "s/^timeout_commit = \"5s\"/timeout_commit = \"3s\"/" ${EVMOSD_HOME}/config/config.toml
sed -i "s/^prometheus = false/prometheus = true/" ${EVMOSD_HOME}/config/config.toml

# 서버 시작
#pm2 start pm2.evmosd.json



