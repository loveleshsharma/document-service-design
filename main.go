package main

import (
	"document-service-design/document"
	"document-service-design/user"
	"fmt"
)

func main() {

	ds := document.NewDocumentService()
	userRamesh := user.NewUser("Ramesh")
	userSuresh := user.NewUser("Suresh")
	userDinesh := user.NewUser("Dinesh")

	doc1 := ds.Create("doc1", userRamesh)

	ds.AddContent(&doc1, "the quick brown fox jumps over the lazy dog")

	if err := ds.GrantAccess(userRamesh, userSuresh, doc1, document.Read); err != nil {
		fmt.Println("cannot grant access to suresh: ", err)
	}

	if err := ds.GrantAccess(userRamesh, userDinesh, doc1, document.Edit); err != nil {
		fmt.Println("cannot grant access to dinesh:", err)
	}

	content, err := ds.Read(doc1, userDinesh)
	if err != nil {
		fmt.Println("cannot read: ", err)
		return
	}

	fmt.Println("content: ", content)

	if err := ds.Delete(doc1, userRamesh); err != nil {
		fmt.Println("cannot delete doc: ", err)
	}

}
