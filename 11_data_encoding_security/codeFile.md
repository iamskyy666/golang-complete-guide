```go 
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 📂 11_data_encoding_security
// 🔵 Marshalling Data

// encoding/decoding ---> marshalling/unmarshalling

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Phone string `json:"ph_no"`
	IsActive bool `json:"is_active"`
}

func main() {
	skyy:=User{
		Name: "Skyy", 
		Age: 30,
		Phone: "123-456-789",
		IsActive: true,
	}

	byteSlice,err:=json.Marshal(skyy)
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	byteSliceIdnt,err:=json.MarshalIndent(skyy,"-"," ")
	// MarshalIndent - a bit heavier in computation, than Marshal
	if err != nil {
		log.Fatal("ERROR:",err.Error())
	}

	fmt.Println(string(byteSlice))

	fmt.Println("########## INDENTED JSON: #########")
	fmt.Println(string(byteSliceIdnt))
}

// O/P:
// $ go run main.go
// {"name":"Skyy","age":30,"ph_no":"123-456-789","is_active":true}
// ########## INDENTED JSON: #########
// {
// - "name": "Skyy",
// - "age": 30,
// - "ph_no": "123-456-789",
// - "is_active": true
// -}
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 📂 11_data_encoding_security
// 🟢 Unarshalling Data

// encoding/decoding ---> marshalling/unmarshalling

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Phone string `json:"ph_no"`
	Password string `json:"-" xml:"-"`
	IsActive bool `json:"is_active"`
	Role string `json:"-" xml:"role"`
	Profile profile `json:"profile"`
}

var payload = `{
		"name": "Skyy", 
		"age": 30,
		"ph_no": "123-456-789",
		"is_active": true,
		"profile": {
		"url":"https://github.com/iamskyy666"
		}
	}`

type profile struct {
	URL string `json:"url"`
}


func main() {
	var u User

	err:=json.Unmarshal([]byte(payload),&u)
	if err != nil {
		log.Fatal("ERROR: ",err)
	}
	
	fmt.Println(u)
	fmt.Printf("%+v\n",u)

	bs,err:=json.MarshalIndent(u,""," ")
	if err != nil {
		log.Fatal("ERROR: ",err)
	}
	fmt.Println("\njson_format:\n",string(bs))
	
}

// $ go run main.go
// {Skyy 30 123-456-789  true  {https://github.com/iamskyy666}}
// {Name:Skyy Age:30 Phone:123-456-789 Password: IsActive:true Role: Profile:{URL:https://github.com/iamskyy666}}

// json_format:
//  {
//  "name": "Skyy",
//  "age": 30,
//  "ph_no": "123-456-789",
//  "is_active": true,
//  "profile": {
//   "url": "https://github.com/iamskyy666"
//  }
// }

```
```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// 📂 11_data_encoding_security
// ENCODER

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Phone string `json:"ph_no"`
	Password string `json:"-" xml:"-"`
	IsActive bool `json:"is_active"`
}



func main() {
	skyy:=User{
		Name: "Skyy", 
		Age: 30,
		Phone: "123-456-789",
		IsActive: true,
	}

	// write on std. output
	enc:=json.NewEncoder(os.Stdout)

	if err:=enc.Encode(&skyy); err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	// wriite on buffer
	buff:=new(bytes.Buffer)
	encBuff:=json.NewEncoder(buff)

	if err:=encBuff.Encode(&skyy); err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	// fmt.Println(string(buff.Bytes()))
	fmt.Println(buff.String())


	
}

// $ go run main.go
// {"name":"Skyy","age":30,"ph_no":"123-456-789","is_active":true}
// {"name":"Skyy","age":30,"ph_no":"123-456-789","is_active":true}
```
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// 📂 11_data_encoding_security
// DECODER

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Phone string `json:"ph_no"`
	Password string `json:"-" xml:"-"`
	IsActive bool `json:"is_active"`
}

var payload=`{"name":"Skyy","age":30,"ph_no":"123-456-789","is_active":true}`

func main() {
	var skyy User

	dec:=json.NewDecoder(strings.NewReader(payload))

	if err:=dec.Decode(&skyy); err!=nil{
		log.Fatal("ERROR: ",err.Error())
	}

	fmt.Println("DECODED:",skyy)
}

// $ go run main.go
// DECODED: {Skyy 30 123-456-789  true}
```