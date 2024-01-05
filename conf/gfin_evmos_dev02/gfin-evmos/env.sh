#!/bin/bash

CHAINID="gurufin_8282-1"  # {identifier}_{EIP155}-{version}
KEYRING="file"
MONIKER="gfin_val01"
KEYNAME="gfin_val01"
PASSWORD="gurufin830!@"
WEBSITE="https://www.gurufin.io"
SYMBOL="gurux"
AMOUNT=100000000000000000000000000000  # supply * (10 ** 18) 개
EVMOSD_HOME=/root/evm_home

MONIKER2="gfin_val02"
KEYNAME2="gfin_val02"

# SEED 노드 정보
HOST="gfin-gurux01"
IP="172.31.51.101"
VALID="c19b485da8a4d7c3aeedbab7118da2130d9a717d"  # evmosd tendermint show-node-id --home=/root/evm_home


