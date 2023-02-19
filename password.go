package Encryption

type AlgorithmKey struct {
	SecretKey string
	ByteKey   []byte
}

var encryption *AlgorithmKey

func New(AesKey string) *AlgorithmKey {
	GenerateKey := []byte(AesKey)
	encryption = &AlgorithmKey{
		SecretKey: AesKey,
		ByteKey:   GenerateKey}
	return encryption
}

func (a *AlgorithmKey) EncryptPassword(password string) (encoded string, err error) {

	getEncrypt, err := Encrypt_AES(a.ByteKey, password)
	if err != nil {
		return
	}

	return getEncrypt, err

}
func (a *AlgorithmKey) DecryptPassword(encryptPassword string) (encoded string, err error) {

	getEncrypt, err := Decrypt_AES(a.ByteKey, encryptPassword)
	if err != nil {
		return
	}

	return getEncrypt, err

}

func Resolve() *AlgorithmKey {
	return encryption
}
