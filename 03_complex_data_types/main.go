package main

import "fmt"

// 📂 03_complex_data_types
// Project: CONTACT MANAGEMENT SYSTEM


type Contact struct {
	ID int
	Name string
	Email string
	Phone string
}

var contactList []Contact
var contactIdxByName map[string]int
var nextID int = 1

func init(){
	contactList = make([]Contact, 0)
	contactIdxByName = make(map[string]int)
}

func addContact(name,email,phone string){
	if _,exists:=contactIdxByName[name];exists{
		fmt.Printf("Contact already exists: %v\n",name)
		return
	}

	newContact:=Contact{
		ID: nextID,
		Name:name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIdxByName[name] = len(contactList)-1
	fmt.Printf("Contact Added: %v\n",name)
}

func findContactByName(name string)*Contact{
	idx,exists:=contactIdxByName[name]
	if exists{
		return &contactList[idx]
	}
	return nil
}

func listContacts(){
	fmt.Println("------ Listing Contacts ------")
	if len(contactList) == 0{
		fmt.Println("No Contacts Found!")
		return
	}

	for i,contact:=range contactList{
		fmt.Printf("%d ID: %d, Name: %s, Email: %s, Phone: %s\n",i+1,contact.ID, contact.Name, contact.Email, contact.Phone)
	}

	fmt.Println("----------------------------")

	skyy:=findContactByName("skyy")
	if skyy==nil{
		fmt.Println("skyy not found!")
	}else{
		fmt.Println("skyy found")
	}
}

func main() {
	addContact("skyy","skyy@test.com","111-333")
	addContact("alice","alice@test.com","222-333")
	addContact("jim","jim@test.com","333-333")
	listContacts()
}

// $ go run main.go




