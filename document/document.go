package document

import (
	"document-service-design/user"
	"errors"
)

var ErrUserAccessNotFound error = errors.New("user access not found")

var idCurrentValue int64 = 0

type Document struct {
	id int64

	name string

	content string

	owner user.User

	accessList map[user.User]AccessType
}

func newDocument(name string, owner user.User) Document {
	idCurrentValue++
	return Document{
		id:         idCurrentValue,
		name:       name,
		content:    "",
		owner:      owner,
		accessList: make(map[user.User]AccessType),
	}
}

func (d *Document) getId() int64 {
	return d.id
}

func (d *Document) getName() string {
	return d.name
}

func (d *Document) addContent(content string) {
	d.content = content
}

func (d *Document) getOwner() user.User {
	return d.owner
}

func (d *Document) getContent() string {
	return d.content
}

func (d *Document) addAccess(user user.User, accessType AccessType) {
	d.accessList[user] = accessType
}

func (d *Document) getAccessTypeByUser(user user.User) (AccessType, error) {
	accessType, exists := d.accessList[user]
	if exists {
		return accessType, nil
	}
	return NoAccess, ErrUserAccessNotFound
}
