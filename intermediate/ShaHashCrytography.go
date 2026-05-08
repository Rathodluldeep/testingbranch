package intermediate

import (
	// "crypto/sha256"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)


func main (){
	password :="password123"
	// hash256 := sha512.Sum512([]byte(password))

	// fmt.Println(password)
	// fmt.Println(hash256)
	// //converting hex value
	// fmt.Printf("SHA-256 hash hex value: %x\n",hash256)

	//saltstring

	salt,err := generateSalt()
	if err != nil{
		fmt.Println("Error generating salt: ",err)
		return
	}
	//hash the password with salt
	signUpHash := hashPassword(password,salt)

	//store the salt and password in datatbase,just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt: ",saltStr) //simulate as storing in database
	fmt.Println("Hash: ",signUpHash) //simulate as storing in database

	//verify password
	//reteive the saltStr and decode it
	decodedSalt,err :=base64.StdEncoding.DecodeString(saltStr)
	if err != nil{
		fmt.Println("unable to decode salt: ",err)
		return
	}
	loginHash :=hashPassword("password123",decodedSalt)

	//compare the stored signuphash with loginhash
	if signUpHash == loginHash{
		fmt.Println("Password is correct.You are logged in ")
	}else{
		fmt.Println("Password login fail.Please check user credentials.")
	}


}

func generateSalt() ([]byte, error){
	salt :=make([]byte,16)
	_, err := io.ReadFull(rand.Reader,salt)
	if err != nil{
		return nil, err
	}
	return salt, nil
}

//function to hash password
func hashPassword (password string, salt []byte)string{
	saltedPassword := append(salt,[]byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}