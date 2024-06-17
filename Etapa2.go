// Author: Gilberto Luis Koerbes Junior
// Este código faz parte de 3 / 3 etapas para troca de mensagens.
// Nesta etapa, recebemos uma mensagem, deciframos a mesma, invertemos o conteúdo e criframos novamente
// Principais variveis de entrada: key (calculada na etapa anterior) - ciphertext (mensagem cifrada)
// #Doc: ESTE CÓDIGO É BASEADO NA DOCUMENTAÇÃO OFICIAL GOLANG.
// CODE EXAMPLE https://pkg.go.dev/crypto/cipher#NewCBCDecrypter
// CODE EXAMPLE https://pkg.go.dev/crypto/cipher#NewCBCEncrypter
// go version go1.18.1 linux/amd64
// para executar sem compilar => go run nome_arquivo

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func addPKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	fmt.Println("Padding => ", padding, padText)
	return append(data, padText...)
}

func main() {

	/////////////////////////////////////////////////////////////////////
	/////////////////////// D E C R Y P T ///////////////////////////////
	/////////////////////////////////////////////////////////////////////
	key, _ := hex.DecodeString("663fdd7206e22c085b683c89f72b14b3")

	//MENSAGEM 1 RECEBIDA
	//ciphertext, _ := hex.DecodeString("9752C89280BC97E0E66CE51688E76C5F7FB10ADF452EE59DDBF336499EFDCB84F7D3117E24AA1F0EC2D044EB0E520CCACA0A4747E4D04F271F3F51C7A34FEFFE57931116C9715823DDB927ECA6913F530C08FDE0AFC361098A80BA21FD2F3F151E12ACF82DD2EE87B000C132A35FC7EC")

	//MENSAGEM 2 RECEBIDA
	ciphertext, _ := hex.DecodeString("937F26B2BBA44A656ED831B2B19049C71AA2A0AA2B64CDC19321C52327531BD8C9D90699243B72841B4F0E6C351F2D9DF2F6014E363D037B70820C0EBA383DE437A635ABE9A06CFDDA17360C8C5E08B673A7614ACDAA6A8D93F981A15084436280F69102C5D47CE01E57F2330A76C357D67133A33C0AD876BCC3905DECE7066E11CFB2A6C3ABFA641DC883E7398909B960E84C75FBBC6C487A458EA1AA4F1969FC600B6786336571A7349C70DF9C4443DB407534CE1396D45C96A85584D2B2BC8937F10AF2A276D924F69EBC9150E438D503172274E3AF693002920DB9515FA961A85D93FE565157088E59A24F8756E1067DEF497CE55573DE2414161AC7FFD22B6572D8A09FBC40A27C48AF34EC12A4536D4764F9D04335F8BE80D661F5D837E676FB704F4CD6F229EEAFD4C2CD996F495E330AD3D30926B007D52CF4532D50")

	//Block, em Golang, define um objeto BlockMode com base na chave, que define os atributos como quantidade de rounds e tamanho de chave
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]        //pegar os primeiros bytes que representam IV
	ciphertext = ciphertext[aes.BlockSize:] //pegar os demais bytes que representam mensagem

	fmt.Println("aes.BlockSize", aes.BlockSize)
	fmt.Println("iv", iv)
	fmt.Println("Mensagem cifrada recebida", ciphertext)
	fmt.Println("Length Mensagem recebida", len(ciphertext))

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	//Aqui criamos o modo de operação, passando o objeto Block(informações da chave, tamanho de chave e rounds) e o IV
	mode := cipher.NewCBCDecrypter(block, iv)

	deciphertext := ciphertext
	mode.CryptBlocks(deciphertext, ciphertext) //com o modo definido - parametro1: valor onde a mensagem decifrada é gravada, parametro2: mensagem cifrada

	fmt.Println("%s\n Mensagem decifrada recebida => ", string(deciphertext[:]))
	fmt.Println("Length Mensagem decifrada recebida => ", len(deciphertext))

	/////////////////////////////////////////////////////////////////////
	/////////////////////// E N C R Y P T ///////////////////////////////
	/////////////////////////////////////////////////////////////////////
	message := string(deciphertext[:])
	// Inverter a mensagem em texto plano
	r_message := []rune(message)
	for i, j := 0, len(r_message)-1; i < j; i, j = i+1, j-1 {
		r_message[i], r_message[j] = r_message[j], r_message[i]
	}

	reverse_message := []byte(string(r_message))
	fmt.Println("\\n Mensagem Invertida", (string(reverse_message)))
	fmt.Println("\nlen reverse_message", len(string(reverse_message)))

	reverse_message_with_padding := addPKCS7Padding(reverse_message, aes.BlockSize)
	fmt.Println("\n reverse_message_with_padding", reverse_message_with_padding)
	if len(string(reverse_message_with_padding))%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	//generate new IV random
	iv = make([]byte, 16)
	_, err = rand.Read(iv)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("iv = make([]byte, 16)", iv)

	ciphertext_to_send := make([]byte, aes.BlockSize+len(reverse_message_with_padding))
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode = cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext_to_send[aes.BlockSize:], reverse_message_with_padding)

	//Add IV in ciphertext_to_send
	for i := 0; i < aes.BlockSize; i++ {
		ciphertext_to_send[i] = iv[i]

	}

	fmt.Println("ciphertext_to_send reverse_message = ", hex.EncodeToString(ciphertext_to_send))

}
