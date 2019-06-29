
package tx

import (
	"testing"
	"strings"
)

func Test(t *testing.T) {

	//	basic tx example from https://bitcoin.org/en/developer-reference#raw-transaction-format

	inputs := []TransactionInput{}

	inputTX := []byte{0x7b,0x1e,0xab,0xe0,0x20,0x9b,0x1f,0xe7,0x94,0x12,0x45,0x75,0xef,0x80,0x70,0x57,0xc7,0x7a,0xda,0x21,0x38,0xae,0x4f,0xa8,0xd6,0xc4,0xde,0x03,0x98,0xa1,0x4f,0x3f}
	scripts := []byte{0x48,    0x30,0x45,0x02,0x21,0x00,0x89,0x49,0xf0,0xcb,0x40,0x00,0x94,0xad,0x2b,0x5e,0xb3,0x99,0xd5,0x9d,0x01,0xc1,0x4d,0x73,0xd8,0xfe,0x6e,0x96,0xdf,0x1a,0x71,0x50,0xde,0xb3,0x88,0xab,0x89,0x35,0x02,0x20,0x79,0x65,0x60,0x90,0xd7,0xf6,0xba,0xc4,0xc9,0xa9,0x4e,0x0a,0xad,0x31,0x1a,0x42,0x68,0xe0,0x82,0xa7,0x25,0xf8,0xae,0xae,0x05,0x73,0xfb,0x12,0xff,0x86,0x6a,0x5f,0x01}
	inputs = append(inputs, createInput(inputTX, 0, scripts))


	outputs := []OutputFormat{}
	scriptFormat := []byte{0x76, 0xa9, 0x14, 0xcb,0xc2, 0x0a, 0x76, 0x64, 0xf2, 0xf6, 0x9e, 0x53,0x55,0xaa,0x42,0x70,0x45,0xbc,0x15,0xe7,0xc6,0xc7,0x72,		0x88, 0xac}
	outputs = append(outputs, createOutput(4999990000, scriptFormat))

	tx := construct(inputs, outputs)
	if !(strings.Compare(tx, "01000000017b1eabe0209b1fe794124575ef807057c77ada2138ae4fa8d6c4de0398a14f3f00000000494830450221008949f0cb400094ad2b5eb399d59d01c14d73d8fe6e96df1a7150deb388ab8935022079656090d7f6bac4c9a94e0aad311a4268e082a725f8aeae0573fb12ff866a5f01ffffffff01f0ca052a010000001976a914cbc20a7664f2f69e5355aa427045bc15e7c6c77288ac00000000") == 0){
		t.Errorf("something is wrong with tx generator")
	}

}