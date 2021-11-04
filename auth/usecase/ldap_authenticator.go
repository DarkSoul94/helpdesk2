package usecase

import (
	"errors"
	"strings"

	"github.com/go-ldap/ldap"
)

var (
	errLDAPConnection         = errors.New("Ldap connection error")
	errLDAPInvalidCredentials = errors.New("Invalid credentials")
	errLDAPUserNotFound       = errors.New("User not found")
	errLDAPUndefined          = errors.New("Undefined error")
)

type ldapUser struct {
	Name        string
	Department  string
	Email       string
	Description string
}

type ldapAuthenticator struct {
	BaseDN   string
	FilterDN string
	Address  string
}

func NewLdapAuthenticator(addres, baseDN, filterDN string) ldapAuthenticator {
	return ldapAuthenticator{
		BaseDN:   baseDN,
		FilterDN: filterDN,
		Address:  addres,
	}
}

func (l *ldapAuthenticator) Auth(email, password string) (ldapUser, error) {
	conn, err := ldap.Dial("tcp", l.Address)
	defer conn.Close()
	if err != nil {
		return ldapUser{}, errLDAPConnection
	}
	err = conn.Bind(email, password)
	if err != nil {
		return ldapUser{}, errLDAPInvalidCredentials
	}

	res, err := conn.Search(
		ldap.NewSearchRequest(
			l.BaseDN,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			l.filter(email),
			[]string{"sAMAccountName", "cn", "givenName", "mail", "department", "description"},
			nil,
		))
	if err != nil || len(res.Entries) != 1 {
		return ldapUser{}, errLDAPUserNotFound
	}

	for _, entry := range res.Entries {
		return ldapUser{
			Name:        entry.GetAttributeValue("cn"),
			Email:       entry.GetAttributeValue("mail"),
			Department:  entry.GetAttributeValue("department"),
			Description: entry.GetAttributeValue("description"),
		}, nil
	}
	return ldapUser{}, errLDAPUndefined
}

func (l *ldapAuthenticator) filter(needle string) string {
	res := strings.ReplaceAll(
		l.FilterDN,
		"{username}",
		needle,
	)
	return res
}
