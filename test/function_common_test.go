package main

import (
	"base/internal/utils"
	"testing"
	"github.com/stretchr/testify/assert"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 09:27
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func TestGenHash(t *testing.T) {
	plainText := "vccorpadtechadmicrobigdataplatform"
	hash256 := "0faec427ba144f820c2f4faceb7393dd3cbe953b3dea48d53e011a0586b3fadd"
	assert.Equal(t, utils.GenHash(plainText), hash256)
}

func BenchmarkGenHash(b *testing.B)  {
	for test := 0; test < b.N; test++ {
		var plainText string  = "vccorpadtechadmicrobigdataplatform"
		plainText += string(test)
		utils.GenHash(plainText)
	}
}