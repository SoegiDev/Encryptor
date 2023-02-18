package encrypts

import (
	"testing"
	"fmt"
)

// function to test if the original
// go program is working
func TestDecrypt(test *testing.T){
  
	SECRET_KEY :="abc&1*~#^2^#s0^=)^^7%b34"
	Expected:="Fajar Soegi 1901"
	Text := "LSFN4hoMcPcEZcazInDUUA=="
	decText, err := Decrypt(Text, SECRET_KEY)
	if err != nil {
	  fmt.Println("error decrypting your encrypted text: ", err)
	  }
	if decText == Expected{
		fmt.Println("Decryption Working Encyption =  "+ decText)
	}else{
		fmt.Println("Not Working")
	}	 
  
}