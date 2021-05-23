package services

/**
 * @Description: 加解密接口
 */
type CryptServer interface {
	Encrypt(data interface{}) error
	Decrypt(data interface{}) error
}