package main

// Author: Gilberto Luis Koerbes Junior
// Este código faz parte de 2 / 3 etapas para troca de mensagens.
// Nesta etapa, recebemos uma chave publica e calculamos um valor V, comum as partes
// go version go1.18.1 linux/amd64
// para executar sem compilar => go run nome_arquivo

import (
	"crypto/sha256"
	"encoding/hex"
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

	B := "62090E85206D8D6E74CBFCFEF3970B61774102CC599FFFC1B6F3DB99D272011408FDE67B10AB35B35DF6FA12045BBE7AE9CC05FFD3C8A0FC5223A3C7E1A0D1996AD5031765861E547768164A3E998D3A7327031CA0E79C28C46D88EF72D307A5417A8F9C69B72DA7CEB11C9E3D3F47D0476CC6266A5C18C513CBFB64E5300085"

	a := "f3dc96767f89cf3b5cdb1c62"

	//Replace de espacos e tabs
	sampleRegexp := regexp.MustCompile(`\s|/\n`)
	p = sampleRegexp.ReplaceAllString(p, "")

	fmt.Println("\n p = ", p)
	fmt.Println("\n B = ", B)

	p_decimalNumber := new(big.Int)
	p_decimalNumber, success := p_decimalNumber.SetString(p, 16)

	B_decimalNumber := new(big.Int)
	B_decimalNumber, success = B_decimalNumber.SetString(B, 16)

	a_decimalNumber := new(big.Int)
	a_decimalNumber, success = a_decimalNumber.SetString(a, 16)
	fmt.Println(a_decimalNumber)
	if !success {
		fmt.Println("Error: Invalid hexadecimal string")
		return
	}

	V := new(big.Int)
	V.Exp(B_decimalNumber, a_decimalNumber, p_decimalNumber) // V = (B^a mod p)

	V_bytes := V.Bytes()

	//calculando o hash de V
	S := sha256.Sum256([]byte(V_bytes))
	fmt.Println("\n S = ", S)

	// Converter o S hash do tipo Bytes para hexadecimal string
	hex_S := hex.EncodeToString(S[:])

	fmt.Println("\n hex_S = ", hex_S)
	fmt.Println("\n hex_S[:128 bits] = ", hex_S[:32])

}
