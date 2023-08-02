## Challenge cosmos-faucet

### Description
Welcome to the cosmos-faucet challenge, here is some info for it:
* Attachment: https://github.com/ignite/cli/ @v0.19.1
* faucet-RPC: http://ip:4500
* Rest-RPC: http://ip:1337

## Deployment
```
$ git clone -b v0.19.1 https://github.com/ignite/cli.git
$ cd cli
$ make install
$ starport chain build
$ starport chain serve --config ../config.yml
```

### Notethe *config.yaml* is attached as below:
```yml
accounts:
  - name: alice
    coins: ["0token", "200000000stake"]
  - name: bob
    coins: ["500token", "100000000stake"]
validator:
  name: alice
  staked: "100000000stake"
client:
  openapi:
    path: "docs/static/openapi.yml"
  vuex:
    path: "vue/src/store"
faucet:
  name: bob
  coins: ["5token", "100000stake"]  
  coins_max: ["12token", "10000]
```

### Goal
Let *bob* get more than **12token** from faucet successfully to exceed the *coins_max*!

### Instruction
1. Setup the chain via *starport* for testing


2. Get token from faucet via *$faucet-RPC* with POST with body:
  ```json
  {
    "address": "BOB_ADDRESS"
  }
  ```

3. Check the balance of *bob* via $Rest-RPC 

### Hint
* Each time *bob* could get *5token* from the faucet and with limit of *12token*, this means only 2 times could be performed towards the faucet.
