package glostructs

import(
	"fmt"
	"os"
	"bufio"
	"crypto/rand"
    "crypto/rsa"
    "crypto/x509"
	"encoding/pem"
	"encoding/json"
	"encoding/base64"
	"crypto/sha256"
)

type Wallet struct{
	Address [32]byte	//Address and id account of the wallet
	PublicKey rsa.PublicKey	//Public key 
	Amount uint64	//Amount of money holded by the account, unsigned int of 8 digits as minimum
}

func(wallet *Wallet) New(passphrase string, bits int)(*pem.Block, error){
	pk, err := wallet.createPrivateKey(bits)
	if err != nil {
		return nil, err
	}
	wallet.PublicKey = pk.PublicKey
	pubjson,err := json.Marshal(pk.PublicKey)
	wallet.Address = sha256.Sum256(pubjson)
	pem := wallet.getPemFromPK(pk)
	block, err := wallet.EncryptPEMBlock(pem,passphrase)
	if err != nil {
		return nil, err
	}
	return block,nil
}

func TestNewWallet(){
	in := bufio.NewReader(os.Stdin)
	var wallet Wallet
	fmt.Println("\n**************************")
	fmt.Println("IMPORTANT INFO:")
	fmt.Println("YOUR PASSPHRASE SHOULD BE REMEMBERED, IT ALLOWS YOU TO CERTIFY EVERY TRANSACTION")
	fmt.Println("AT LAST WE WILL GIVE YOU A PRIVATE KEY, IT SHOULD BE SAVED TO CERTIFY YOUR OWNERSHIP ON WALLET LOGIN")
	fmt.Println("**************************")
	
	fmt.Println("\nInput your passphrase (no minimum security implemented yet): ")
	passphrase,err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	passphrase = passphrase[:len(passphrase)-1]
	
	block, err := wallet.New(passphrase,4096)
	if err != nil{
		panic(err)
	}
	f, err := os.Create("extra/pemblock.txt")
	if err != nil{
		panic(err)
	}
	blockjson,err := json.Marshal(block)
	if err != nil{
		panic(err)
	}
	f.WriteString(string(blockjson))

	f2, err := os.Create("extra/pemblockbase64.txt")
	if err != nil{
		panic(err)
	}
	pembase64 := base64.StdEncoding.EncodeToString(blockjson)
	f2.WriteString(pembase64)
	defer f.Close()

	fmt.Println("\nYour wallet is:")
	fmt.Println(wallet)

	fmt.Println("\nYour private key is:")
	fmt.Println(pembase64)
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