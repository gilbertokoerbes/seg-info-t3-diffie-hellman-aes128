package main

// Author: Gilberto Luis Koerbes Junior
// Este código faz parte de 1 / 3 etapas para troca de mensagens.
// Nesta etapa, geramos uma chave publica para troca de chaves com DiffieHelmman

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
)

func main() {
	p := "B10B8F96 A080E01D DE92DE5E AE5D54EC 52C99FBC FB06A3C6" +
		"9A6A9DCA 52D23B61 6073E286 75A23D18 9838EF1E 2EE652C0" +
		"13ECB4AE A9061123 24975C3C D49B83BF ACCBDD7D 90C4BD70" +
		"98488E9C 219A7372 4EFFD6FA E5644738 FAA31A4F F55BCCC0" +
		"A151AF5F 0DC8B4BD 45BF37DF 365C1A65 E68CFDA7 6D4DA708" +
		"DF1FB2BC 2E4A4371"

	g := "A4D1CBD5 C3FD3412 6765A442 EFB99905 F8104DD2 58AC507F" +
		"D6406CFF 14266D31 266FEA1E 5C41564B 777E690F 5504F213" +
		"160217B4 B01B886A 5E91547F 9E2749F4 D7FBD7D3 B9A92EE1" +
		"909D0D22 63F80A76 A6A24C08 7A091F53 1DBF0A01 69B6A28A" +
		"D662A4D1 8E73AFA3 2D779D59 18D08BC8 858F4DCE F97C2A24" +
		"855E6EEB 22B3B2E5"

	//Replace de espacos e tabs
	sampleRegexp := regexp.MustCompile(`\s|/\n`)
	p = sampleRegexp.ReplaceAllString(p, "")
	g = sampleRegexp.ReplaceAllString(g, "")
	//fmt.Println("\n p = ", p)
	//fmt.Println("\n g = ", g)

	// Parse hexadecimal string to big.Int
	p_decimalNumber := new(big.Int)
	p_decimalNumber, success := p_decimalNumber.SetString(p, 16)

	g_decimalNumber := new(big.Int)
	g_decimalNumber, success = g_decimalNumber.SetString(g, 16)

	if !success {
		fmt.Println("Error: Invalid hexadecimal string")
		return
	}

	//Random com 30 digitos
	base := big.NewInt(10)
	exp := big.NewInt(30)
	max := new(big.Int)
	max.Exp(base, exp, nil).Sub(max, big.NewInt(1)) // 10^30

	a, err := rand.Int(rand.Reader, max) // gerar um random com até 10^30
	if err != nil {
		//error handling
	}
	fmt.Println("\n a = ", a)

	A := new(big.Int)
	A.Exp(g_decimalNumber, a, p_decimalNumber) // A = (g^a mod p)

	fmt.Println("\n A = ", A)

	// Convert decimal to hexadecimal
	hex_A := fmt.Sprintf("%X", A)

	fmt.Println("\n hex_A = ", hex_A) //Valor de troca de chave a ser mandado para a outra parte

}
