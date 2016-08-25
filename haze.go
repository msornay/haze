package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

func unit(pow int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(pow), nil)
}

var (
	Ether = unit(18)
)

type Request struct {
	Id      int      `json:"id"`
	Version string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params,omitempty"`
}

type Response struct {
	Version string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}

func getBalance(addr string) (*big.Int, error) {
	req := Request{
		Version: "2.0",
		Method:  "eth_getBalance",
		Params:  []string{*addrFlag, "latest"},
		Id:      1,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))

	resp, err := http.Post(
		"http://localhost:8545", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ethereum client responded with status %s", resp.Status)
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}


	balance := new(big.Int)
	if _, ok := balance.SetString(response.Result, 0); !ok {
		return nil, fmt.Errorf("invalid base 16 integer: %s", response.Result)
	}

	return balance, nil

}

var addrFlag = flag.String("addr", "0x0000000000000000000000000000000000000000", "")

func main() {

	flag.Parse()

	log.Println(flag.Args())

	weiBalance, err := getBalance(*addrFlag)
	if err != nil {
		log.Fatalf("cannot get balance: %v", err)
	}

	etherBalance := new(big.Rat).Quo(
		new(big.Rat).SetInt(weiBalance),
		new(big.Rat).SetInt(Ether),
	)
	fmt.Println(etherBalance.FloatString(8))
}
