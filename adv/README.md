# Advertisement

**Advertisement(adv)** is a blockchain application built using Cosmos SDK and Tendermint.

## Challenge

### Description
Welcome to the adv, you can post an advertisement on the adv chain, and the more coin you deposit the higher your advertisement ranking will be.
* Attachment: https://github.com/ewagmig/ctf-cosmos
* RPC: http://ip:26657

## Deployment
```
$ git clone https://github.com/ewagmig/ctf-cosmos.git
$ cd adv
$ make all
$ cd script
$ bash start_testnet.sh
```

### Goal
Send a successful `CaptureTheFlag` type transaction.

### Instruction
1. Use the `ctfer` key as the account for testing

>PS: During the competition, we provided ctfer preset built-in accounts rather than the genesis account, then transferred enough coins to those addresses. 


2. Check balance, should be none-zero

```
$ advcli query account $ADDRESS --node $RPC
```

3. Post an advertisement

```
$ advcli tx adv create-advertisement $ID $CONTENT --from $KEY --chain-id mainnet --fees 10ctc --node $RPC
```

### Hint
* The playground website is only used for AD display and TX hash submission (Not a web challenge !!!)
* There are `adv` subcommands as below:
    - create-advertisement Create a new advertisement
    - delete-advertisement Delete an advertisement by ID
    - deposit              Deposit an amount of token on an advertisement
    - withdraw             Withdraw an amount of token from an advertisement
    - ctf                  Capture The Flag
    Try to use all these commands to exploit the chain!

