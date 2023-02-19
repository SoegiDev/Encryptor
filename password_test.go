package Encryption_test

import (
	"fmt"
	"os"
	"testing"

	Encryption "github.com/SoegiDev/Encryptor"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestPassword(m *testing.T) {

	secretKey32Bit := "asuperstrong32bitpasswordgohere!"
	//secretKey24Bit := "asuperstrong24bitpasswor"
	//secretKey16Bit := "asuperstrong16bi"
	setPassword := Encryption.New(secretKey32Bit)
	encryptPassword, _ := setPassword.EncryptPassword("Fajarsoegi")
	decryptPassword, _ := setPassword.DecryptPassword(encryptPassword)
	if decryptPassword == "Fajarsoegi" {
		fmt.Println("Password Encryption and Decryption is Working Password Is =  " + decryptPassword)
	} else {
		fmt.Println("Not Working")
	}

}
