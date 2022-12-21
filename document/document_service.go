package document

import (
	"document-service-design/user"
	"errors"
	"fmt"
)

var (
	ErrReadAccessNotFound   = errors.New("read access not found")
	ErrDeleteAccessNotFound = errors.New("delete access not found")
	ErrGrantAccessNotFound  = errors.New("grant access not found")
)

type DocumentService struct {
	documentsMap map[int64]Document
}

func NewDocumentService() *DocumentService {
	return &DocumentService{
		documentsMap: make(map[int64]Document),
	}
}

func (d *DocumentService) GetDocumentsList() []string {
	docList := make([]string, len(d.documentsMap))

	for _, v := range d.documentsMap {
		docList = append(docList, v.getName())
	}

	return docList
}

func (d *DocumentService) Create(documentName string, owner user.User) Document {
	newDoc := newDocument(documentName, owner)
	d.documentsMap[newDoc.GetId()] = newDoc

	return newDoc
}

func (d *DocumentService) AddContent(doc *Document, content string) {
	doc.AddContent(content)
	d.documentsMap[doc.GetId()] = *doc
}

func (d *DocumentService) Read(doc Document, reader user.User) (string, error) {
	docOwner := doc.GetOwner()
	if reader == docOwner {
		return doc.GetContent(), nil
	}

	accessType, err := doc.GetAccessTypeByUser(reader)
	if err != nil {
		return "", err
	}

	if accessType == NoAccess {
		return "", ErrReadAccessNotFound
	}

	return doc.GetContent(), nil
}

func (d *DocumentService) GrantAccess(grantor, user user.User, doc Document, accessType AccessType) error {
	docOwner := doc.GetOwner()

	if grantor == docOwner {
		doc.AddAccess(user, accessType)
		return nil
	}

	return ErrGrantAccessNotFound
}

func (d *DocumentService) Delete(doc Document, user user.User) error {
	docOwner := doc.GetOwner()
	if user == docOwner {
		d.delete(doc.GetId())
		fmt.Printf("document: %s deleted successfully", doc.name)
		return nil
	}

	return ErrDeleteAccessNotFound
}

func (d *DocumentService) delete(docId int64) {
	delete(d.documentsMap, docId)
}
