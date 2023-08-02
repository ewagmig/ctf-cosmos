# You can run all of these commands from your home directory
cd $HOME
rm -rf .advd
rm -rf .advcli

# Initialize the genesis.json file that will help you to bootstrap the network
advd init testing --chain-id=mainnet

# Create a key to hold your validator account
advcli keys add val
advcli keys add ctfer

advcli config indent true
advcli config output json
advcli config trust-node true

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
advd add-genesis-account $(advcli keys show val -a) 1000000000stake
advd add-genesis-account $(advcli keys show ctfer -a) 1000ctc

# Generate the transaction that creates your validator
advd gentx --name val

# Add the generated bonding transaction to the genesis file
advd collect-gentxs

# Now its safe to start `advd`
advd start
