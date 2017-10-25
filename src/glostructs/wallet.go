package glostructs

import (
	"math/big"
	"crypto/sha512"
	"golang.org/x/crypto/ripemd160"
	"crypto/sha256"
	"encoding/hex"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

type WalletECDSA struct {
	IdAccount     string
	PublicID	ecdsa.PublicKey //public key of Account
	Amount uint64 //amount of wallet
}

func CreateIDAccount(X *big.Int, Y *big.Int) string {
	
	//dos numeros de la clave publica se les pasa un sha256
	h1 := sha256.New()
	h2 := sha256.New()
		
	h1.Write([]byte(X.String()))
	h2.Write([]byte(Y.String()))
	
	//doble sha para seguridad criptográfica
	h1.Write([]byte(h1.Sum(nil)))
	h2.Write([]byte(h2.Sum(nil)))
	
	//se juntan los sha codificados en hexadecimal
	var strAux string = hex.EncodeToString(h1.Sum(nil)) + hex.EncodeToString(h2.Sum(nil))
		
	//se le hace un ripemd160
	r := ripemd160.New()
	r.Write([]byte(strAux))
		
	return hex.EncodeToString(r.Sum(nil))
}

func (w *WalletECDSA) New() {

	//creamos las claves
	priAndPub, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	//validate generate keys of account
	if err != nil {
		panic(err)
	}

	w.PublicID = priAndPub.PublicKey
	

	//ese es el id de la wallet, único
	w.IdAccount = CreateIDAccount(priAndPub.X, priAndPub.Y)

	//guardamos la public key
	w.Save(priAndPub.PublicKey)

	//imprimimos la clave privada por consola para que se la guarde el cliente
	fmt.Println("Your private Key:")
	fmt.Println(priAndPub.D.String())
}

func (w *WalletECDSA) Save(k ecdsa.PublicKey) {
	
	//guardamos los datos en el archivo siguiente
	f, err := os.Create("extra/accounts.txt")

	//miramos errores
	if err != nil {
		panic(err)
	}

	//formato de guardado
	f.WriteString(w.IdAccount + "\n")
	f.WriteString(k.X.String() + "\n")
	f.WriteString(k.Y.String() + "\n")
	f.WriteString(fmt.Sprint(w.Amount) + "\n")
	
	//calculamos el sha correspondiente para añadirlo al bloque
	var s = w.IdAccount + k.X.String() + k.Y.String() + fmt.Sprint(w.Amount)
	sha := sha512.New()
	sha.Write([]byte(s))
	
	//insertamos
	f.WriteString(hex.EncodeToString(sha.Sum(nil)) + "\n")

	//ATENCION
	//no se tiene en cuenta el sha anterior, 
	//habría que sumarle el sha anterior para asegurar la cadena de bloques de cuentas

	defer f.Close()
}

func TestNewWallet() {

	//creamos e inicializamos objeto WalletECDSA
	var wallet WalletECDSA
	wallet.New()

	//informamos sobre el wallet creado
	fmt.Println("Generate wallet with id = " + wallet.IdAccount)
	fmt.Println("Recuperamos información sobre la wallet...")
	
	//leemos el archivo de la cadena de cuentas
	f, err := os.Open("extra/accounts.txt")

	if err != nil {
		panic(err)
	}

	// necesitamos la privateKey puesta 
	// por el usuario para probar el correcto funcionamiento
	// de las firmas de transacciones


	defer f.Close()

	fmt.Println("\nTestNewWallet no acabado de implementar.")

}
