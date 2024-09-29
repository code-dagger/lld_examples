package docservice

import (
	"errors"
	"fmt"
)

type service struct {
	documents map[string]*document
}

func (s *service) CreateDocument(user, docName, content string) error {
	if _, exists := s.documents[docName]; exists {
		return fmt.Errorf("document already exist with name '%s'", docName)
	}
	doc := newDocument(docName, content, user)
	s.documents[docName] = doc
	fmt.Printf("document '%s' created successfully by '%s'\n", docName, user)
	return nil
}

func (s *service) EditDocument(user, docName, content string) error {
	doc, exists := s.documents[docName]
	if !exists {
		return errors.New("document not found")
	}
	hasPermission := false
	if doc.owner == user || doc.access[user] == AccessWrite {
		hasPermission = true
	}
	if !hasPermission {
		return errors.New("access denied")
	}
	doc.content = content
	fmt.Printf("document '%s' edited successfully by '%s'\n", docName, user)
	return nil
}

func (s *service) ReadDocument(user, docName string) (string, error) {
	doc, exists := s.documents[docName]
	if !exists {
		return "", errors.New("document not found")
	}
	hasPermission := false
	if doc.owner == user || doc.access[user] == AccessRead || doc.access[user] == AccessWrite {
		hasPermission = true
	}
	if !hasPermission {
		return "", errors.New("access denied")
	}
	return doc.content, nil
}

func (s *service) DeleteDocument(user, docName string) error {
	doc, exists := s.documents[docName]
	if !exists {
		return errors.New("document not found")
	}
	if doc.owner != user {
		return errors.New("only the owner can delete the document")
	}
	delete(s.documents, docName)
	fmt.Printf("document '%s' deleted\n", docName)
	return nil
}

func (s *service) GrantAccess(owner, docName, user string, accessType AccessType) error {
	doc, exists := s.documents[docName]
	if !exists {
		return errors.New("document not found")
	}
	if doc.owner != owner {
		return errors.New("only the owner can grant access")
	}
	doc.access[user] = accessType
	fmt.Printf("access '%s' granted to '%s' for document '%s' by the '%s'\n", accessType, user, docName, owner)
	return nil
}

func NewService() *service {
	return &service{
		documents: make(map[string]*document),
	}
}
