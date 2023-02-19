# Encryption and Decrytion Password

Features
- Create Encryption (AES)
- Create Decryption (AES)

# Install
First get `Encryption`
```bash
go get github.com/SoegiDev/Encryptor
```
# Usage

SampleSecret_32Bit := "mypasswordstrong32bitpasswordgoh"

SampleSecret_24Bit := "mypasswordstrong24bitpas"

SampleSecret_16Bit := "pass16bitstrong!"

initiate Encryption
```
encryption :=  Encryption.New(SampleSecret_32Bit)
```
Get EncryptPassword
```
Password := "PasswordString"
encryptPassword, _ := encryption.EncryptPassword(Password)
```
Get DecryptPassword
```
decryptPassword, _ := encryption.DecryptPassword(encryptPassword)
```
