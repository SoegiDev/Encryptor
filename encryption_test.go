package encrypts

import (
	"testing"
	"fmt"
)

// function to test if the original
// go program is working
func TestEncrypt(test *testing.T){
	SECRET_KEY :="abc&1*~#^2^#s0^=)^^7%b34"
	Text:="Fajar Soegi 1901"
	Expected := "LSFN4hoMcPcEZcazInDUUA=="
	encText, err := Encrypt(Text, SECRET_KEY)
	  if err != nil {
	  fmt.Println("error decrypting your encrypted text: ", err)
	  }
	if encText == Expected{
		fmt.Println("Encryption Working Encyption =  "+ encText)
	}else{
		fmt.Println("Not Working")
	}
	  
	 
  
}