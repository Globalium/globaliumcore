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
	id string
	amount uint64 //amount of wallet
}

//constructor de la clase
func (d *Direction) New() {

	//creamos las claves
	priAndPub, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	//validate generate keys of account
	if err != nil {
		panic(err)
	}
	
	//ese es el id de la wallet, único
	d.id = CreateIDAccount(priAndPub.PublicKey.X, priAndPub.PublicKey.Y)
	d.amount = 0

	//imprimimos la clave privada por consola para que se la guarde el cliente
	aux := priAndPub.D.String() + ":" + priAndPub.PublicKey.X.String() + ":" + priAndPub.PublicKey.Y.String()
	
	//for test only!
	auxTestNewDirection = aux

	fmt.Println("Your private Key:")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(aux)))

}

//busca en la blockchain la direccion pasada por parámetro
//busca su amuont total y lo carga
//SIN COMPLETAR
func (d *Direction)Charge(id string) {

}

//devuelve el amount de la direccion
func (d *Direction)GetAmount() uint64 {
	return d.amount
}

//devuelve el amount de la direccion
func (d *Direction)GetDirection() string {
	return d.id
}

//Verificamos si la direccion pasada tiene en nuestra blockchain igual o más coins de los que dice tener el objeto
//SIN ACABAR
func VerifyAmountDirection(direction Direction) bool{
	return true
}

//Dado los dos big int de la clave Publica genera la direccion
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

//String parse to PrivateKey ecdsa
func StringToPrivateKey(pk string) ecdsa.PrivateKey {

	var PK ecdsa.PrivateKey

	var dates []string
	aux, err := base64.StdEncoding.DecodeString(pk)

	if err != nil {
		panic(err)
	}

	dates = strings.Split(string(aux), ":")

	n1 := new(big.Int)
	n2 := new(big.Int)
	nD := new(big.Int)

	n1, err1 := n1.SetString(dates[1],10) //X number publickey
	n2, err2 := n2.SetString(dates[2],10) //Y number publicKey
	nD, err3 := nD.SetString(dates[0],10) //D number PrivateKey

	if !err1 || !err2 || !err3 {
		panic(err1 && err2)
	}

	//mount PrivateKey
	PK.PublicKey.X = n1
	PK.PublicKey.Y = n2
	PK.PublicKey.Curve = elliptic.P256()
	PK.D = nD

	return PK
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

	fmt.Println("New Account create succes! Direction is " + testDirection.id)

	var pK ecdsa.PrivateKey

	pK = StringToPrivateKey(auxTestNewDirection)

	accountComp := CreateIDAccount(pK.PublicKey.X, pK.PublicKey.Y)

	fmt.Println("Your direction is: " + accountComp)

	if accountComp == testDirection.id {
		fmt.Println("TestNewDirection Succes!")
	} else {
		fmt.Println("TestNewDirection ERROR!!!!! :(")
	}


}
