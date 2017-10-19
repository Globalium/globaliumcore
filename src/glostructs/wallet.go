package glostructs

import(
	"crypto/rand"
    "crypto/rsa"
    "crypto/x509"
	"encoding/pem"
)

type Wallet struct{
	PublicKey [64]byte 	//Public key and account id of the wallet
	Amount uint64		//Amount of money holded by the account, unsigned int of 8 digits as minimum
}

func(wallet *Wallet) New(passphrase string)(bool, error){
	pk, err := wallet.createPrivateKey(2048)
	if err != nil {
		return false, err
	}
	pem := wallet.getPemFromPK(pk)
	block, err := wallet.EncryptPEMBlock(pem,passphrase)
	if err != nil {
		return false, err
	}
	
	return true,nil
}

func(wallet *Wallet) createPrivateKey(bits int) (*rsa.PrivateKey, error){
    pk, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, err
	}
	return pk,nil
}

func(wallet *Wallet) getPemFromPK(pk *rsa.PrivateKey) (*pem.Block){
    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(pk),
	}
	return block
}

func(wallet *Wallet) EncryptPEMBlock(block *pem.Block, passphrase string) (*pem.Block, error){
	block, err := x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(passphrase), x509.PEMCipherAES256)
	if err != nil {
		return nil, err
	}
	return block,nil
}

func(wallet *Wallet) DecryptPEMBlock(block *pem.Block, passphrase string) ([]byte, error){
	der, err := x509.DecryptPEMBlock(block, []byte(passphrase))
	if err != nil {
		return nil, err
	}
	return der,nil
}

func(wallet *Wallet) getPKFromPem(der []byte) (*rsa.PrivateKey, error){
	pk, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil{
		return nil, err
	}
	return pk, nil
}
//pem.EncodeToMemory(block), nil