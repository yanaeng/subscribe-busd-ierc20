# Subscribing and Reading IERC-20 Event Logs

## Online

[Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book).

## Quick Overview Setup

1. Installing Solidity on Mac

brew install node
npm install -g solc


2. Installing solc-select to manages installing and setting solc compiler to match your .sol versions

pip3 install solc-select
solc --version
solc-select install <version>
solc-select use <version>


3. Installing abigen tool

Download and install abigen via [Go Ethereum](https://geth.ethereum.org/downloads/).

or

go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools


4. Compiling solidity into go

solc --abi Foo.sol -o .
solc --bin Foo.sol -o .
abigen --bin=Foo.bin --abi=Foo.abi --pkg=bar --out=Foo.go


## Usage

When you’re ready, run with

go run .