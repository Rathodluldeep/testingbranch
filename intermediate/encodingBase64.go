package intermediate

import (
	"encoding/base64"
	"fmt"
	
)


func main(){
	data :=[]byte("Hello, Base64 Encoding")

	//Encode Base 64
	encoded :=base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	//Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil{
		fmt.Println("Error in decoding ",err)
		return
	}
	fmt.Println(string(decoded))

	// url safe,avoid "/" and '+'

	urlSafeEncoded := base64.URLEncoding.EncodeToString((data))
	fmt.Println(urlSafeEncoded)

	//decode
	urlDecode,err:= base64.URLEncoding.DecodeString(urlSafeEncoded)
	if err != nil{
		fmt.Println("Error in decoding ",err)
		return
	}
	fmt.Println(string(urlDecode))


}
