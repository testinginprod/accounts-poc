
CHAINID="test"
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

accountsd config keyring-backend "test"
accountsd config chain-id "$CHAINID"

accountsd keys add funder
accountsd keys add beneficiary
accountsd keys add account1
accountsd keys add recoverer

accountsd init "test" --chain-id=$CHAINID

cat $HOME/.accounts/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="utest"' > $HOME/.accounts/config/tmp_genesis.json && mv $HOME/.accounts/config/tmp_genesis.json $HOME/.accounts/config/genesis.json
cat $HOME/.accounts/config/genesis.json | jq '.app_state["staking"]["params"]["unbonding_time"]="300s"' > $HOME/.accounts/config/tmp_genesis.json && mv $HOME/.accounts/config/tmp_genesis.json $HOME/.accounts/config/genesis.json
cat $HOME/.accounts/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="utest"' > $HOME/.accounts/config/tmp_genesis.json && mv $HOME/.accounts/config/tmp_genesis.json $HOME/.accounts/config/genesis.json
cat $HOME/.accounts/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="utest"' > $HOME/.accounts/config/tmp_genesis.json && mv $HOME/.accounts/config/tmp_genesis.json $HOME/.accounts/config/genesis.json
cat $HOME/.accounts/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="utest"' > $HOME/.accounts/config/tmp_genesis.json && mv $HOME/.accounts/config/tmp_genesis.json $HOME/.accounts/config/genesis.json
cat $HOME/.accounts/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="utest"' > $HOME/.accounts/config/tmp_genesis.json && mv $HOME/.accounts/config/tmp_genesis.json $HOME/.accounts/config/genesis.json
cat $HOME/.accounts/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.accounts/config/tmp_genesis.json && mv $HOME/.accounts/config/tmp_genesis.json $HOME/.accounts/config/genesis.json

accountsd add-genesis-account funder 1000000000000000utest --keyring-backend="test"
accountsd add-genesis-account beneficiary 1000000000utest --keyring-backend="test"
accountsd gentx funder 500000000000000utest --keyring-backend="test" --chain-id $CHAINID
accountsd collect-gentxs
accountsd validate-genesis
sed -i 's/127.0.0.1:26657/0.0.0.0:26657/g' $HOME/.accounts/config/config.toml
accountsd start

