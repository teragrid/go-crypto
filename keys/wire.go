package keys

import (
	amino "github.com/teragrid/go-amino"
	crypto "github.com/teragrid/go-crypto"
)

var cdc = amino.NewCodec()

func init() {
	crypto.RegisterAmino(cdc)
}
