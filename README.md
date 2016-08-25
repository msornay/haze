# Haze

Ethereum command-line wallet for shits and giggles.

## Why ?

Because I like my crypto done locally, and my syncing done remotely. Haze
allows you to interact with a remote ethreum node in a secure way, without
your private keys leaving your local machine.

## How ?

Start & sync a remote Ethereum node, with RPC enabled on localhost :

    /opt/geth/geth --rpc

Open an ssh tunnel from your local machine to the remote node :

    ssh -p 443 -L 8545:localhost:8545 username@example.com -v