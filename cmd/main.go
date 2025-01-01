package main

import (
	"asterix-parser/asterix"
	"asterix-parser/asterix/cat240"
	"encoding/hex"
	"fmt"
)

func main() {

	sentence := "f0030fe7c80000020000ae720c810ca00000051403e21935000402ec0002ecbb30303030303030303030302e2e2e30303030303030303030303030303030303030303030302e2e2e2e2e2e2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c292d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2a2a2a2a2a2928282a2a2a2a2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2b2b2b2b2b2b2b2b2b2b2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c1d24242424242424242a2a2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2a2a2a2a2a2a2a2a2a2a2a2a2a2a27272727272729292929292929292a2a2a2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b28282828282828282828282828282828282828282828282828282a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a28282828282828282828282828282828282727262929292929292929292929292929292929292929292929292929292828282828282828281f2626262627272727272727272727272727272727272727272727272424242424242323232323232323232323232327272729292929292929292929292929292929292929292929292828282828282828282828282828282828282828282829292929292929292929292929292929292929292929292727272828282828282828282828282828282828282828282828282828252525252525252525252525252525252525252525252525252626262626262626262626262626262626262626262626282828282828282828282828282828282828282828282828282828282828282828232323232222222222222626262626262626262626262626262626262626262626252525252525252525252525252626262626262626262626262626262626262626262626212222222225252525252525252525252525252525252525252525252323232424242424242424242424242424262626262626262632f104"

	// Convert hex to byte using encoding/hex
	message, err := hex.DecodeString(sentence)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	asterixCat240 := asterix.Cat240{}
	cat240.AsterixUnpackMessage(&asterixCat240, &message)

	fmt.Printf("Cat240: %x", asterixCat240)
}
