package main

import (
	"github.com/ethereum/go-ethereum/common"
	go_listen "go-eth/go-listen"
	go_wallet "go-eth/go-wallet"
	"regexp"
)

func main() {
	go_listen.Checkvalue()
	go_wallet.Wallet()
	//valid := IsValidAddress("0x0BC2336477531cdBdC89d26455e61Ae80B25eEaC")
	//fmt.Println(valid) // true
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}
