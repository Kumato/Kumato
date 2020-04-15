package auth

import (
	"errors"
	ldap "github.com/korylprince/go-ad-auth/v3"
	"github.com/kumato/kumato/internal/db"
)

func authWithLDAP(username, password string) (db.User, error) {
	config := &ldap.Config{
		Server: "131.181.196.156",
		Port:   389,
		BaseDN: "DC=qut,DC=edu,DC=au",
	}

	status, entry, _, err := ldap.AuthenticateExtended(config, username, password, []string{}, []string{})

	if err != nil {
		return db.User{}, err
	}

	if !status {
		return db.User{}, errors.New("ldap auth status is not true for user: " + username)
	}

	user := db.User{
		Qid:   entry.GetAttributeValue("cn"),
		Name:  entry.GetAttributeValue("displayName"),
		Email: entry.GetAttributeValue("qutPrimaryEmail"),
	}

	if user.Qid == "" || user.Email == "" || user.Name == "" {
		return db.User{}, errors.New("ldap auth entry cannot fill all fields for user: " + username)
	}

	return user, nil
}
