package znet

import (
	"crypto/tls"
	"fmt"

	log "github.com/sirupsen/logrus"
	ldap "gopkg.in/ldap.v2"
)

func NewLDAPClient(config LdapConfig) (*ldap.Conn, error) {

	log.Warnf("%+v", config)

	l, err := ldap.DialTLS(
		"tcp",
		fmt.Sprintf("%s:%d", config.Host, 636),
		&tls.Config{InsecureSkipVerify: true},
	)
	if err != nil {
		return &ldap.Conn{}, err
	}
	// defer l.Close()

	// First bind with a read only user
	err = l.Bind(config.BindDN, config.BindPW)
	if err != nil {
		return &ldap.Conn{}, err
	}

	return l, nil

}

func stringValues(a *ldap.EntryAttribute) []string {
	var values []string

	for _, b := range a.ByteValues {
		values = append(values, string(b))
	}

	return values
}
