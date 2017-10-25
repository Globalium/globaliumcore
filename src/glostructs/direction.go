package glostructs

import (
	"strings"
	"encoding/base64"
	"math/big"
	"golang.org/x/crypto/ripemd160"
	"crypto/sha256"
	"encoding/hex"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

//var aux for save PrivateKey of account Test for passing the test
var auxTestNewDirection string

type Direction struct {
	IdAccount     string
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

func (w *Direction) New() {

	//creamos las claves
	priAndPub, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	//validate generate keys of account
	if err != nil {
		panic(err)
	}
	
	//ese es el id de la wallet, único
	w.IdAccount = CreateIDAccount(priAndPub.PublicKey.X, priAndPub.PublicKey.Y)
	w.Amount = 0

	//imprimimos la clave privada por consola para que se la guarde el cliente
	aux := priAndPub.D.String() + ":" + priAndPub.PublicKey.X.String() + ":" + priAndPub.PublicKey.Y.String()
	
	//for test only!
	auxTestNewDirection = aux

	fmt.Println("Your private Key:")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(aux)))

}

//a partir del string PrivateKey genera un objeto PrivateKey
func GetPrivateKey(pk string) ecdsa.PrivateKey {
	var aux ecdsa.PrivateKey
	return aux
}


/*
	Crea una cuenta y expulsa la privateKey por consola.
	Luego la consola vuelve a pedirte la private.
	A partir de la private es capaz de descifrarla
	Compara la direccion de inicio con la de la direccion privada
	Si son iguales succes!
	else fracaso infinito jeje
*/
func TestNewDirection() {
	
	var testDirection Direction
	testDirection.New()

	fmt.Println("New Account create succes! Direction is " + testDirection.IdAccount)

	var dates []string
	dates = strings.Split(string(auxTestNewDirection), ":")

	n1 := new(big.Int)
	n2 := new(big.Int)

	n1, err1 := n1.SetString(dates[1],10)
	n2, err2 := n2.SetString(dates[2],10)

	if !err1 || !err2 {
		panic(err1 && err2)
	}

	accountComp := CreateIDAccount(n1, n2)

	fmt.Println("Your direction is: " + accountComp)

	if accountComp == testDirection.IdAccount {
		fmt.Println("TestNewDirection Succes!")
	} else {
		fmt.Println("TestNewDirection ERROR!!!!! :(")
	}


}
