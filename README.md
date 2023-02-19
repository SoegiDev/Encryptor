# Encryption and Decrytion Password

##
##
Features
    - Create Encryption (AES)
    - Create Decryption (AES)
##
// Install Encryption
##
First get `Encryption`
##
bash
##
go get github.com/SoegiDec/Encryptor
##
// Initiate Encryption
##
SampleSecret_32Bit := "mypasswordstrong32bitpasswordgoh"
##
SampleSecret_24Bit := "mypasswordstrong24bitpas"
##
SampleSecret_16Bit := "pass16bitstrong!"
##
encryption :=  Encryption.New(SampleSecret_32Bit)

// Get Encryption
##
Password := "PasswordString"
##
encryptPassword, _ := encryption.EncryptPassword(Password)
##
// Get Decryption
##
decryptPassword, _ := encryption.DecryptPassword(encryptPassword)
