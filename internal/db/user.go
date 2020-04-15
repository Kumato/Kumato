package db

import "github.com/kumato/kumato/internal/types"

func GetUsers(o *types.Option) []User {
	var ts []types.Task

	db.Select("DISTINCT owner_name, owner_qid, owner_email").
		Order("owner_name desc").
		Offset(o.GetOffset()).
		Limit(o.GetLimit()).
		Find(&ts)

	var us []User

	for _, i := range ts {
		us = append(us, User{
			Name:  i.GetOwnerName(),
			Qid:   i.GetOwnerQid(),
			Email: i.GetOwnerEmail(),
		})
	}
	return us
}
