package main

import (
	"fmt"
)

type UserManager interface {
	//insert functions admins should be able to perform here
	getType() int
	setType(rank int) error
	IncType()
	DecType()
}

type User struct {
	//insert data admin should store
	/*user type will be set with integers and a known code of level. The lower the value the lower the privilege
	0 - member
	1 - exec
	2 - admin
	*/
	Name string
	Pass string
	Type int
	Corp bool //default is false
}

func (u *User) getType() int {
	return u.Type
}

func (u *User) setType(rank int) error {
	if rank < 0 || rank > 2 {
		return fmt.Errorf("rank must be between 0 and 2, got %d", rank)
	}
	u.Type = rank
	return nil
}

func (u *User) IncRank() {
	if u.Type < 2 {
		u.Type++
	}
}

func (u *User) DecRank() {
	if u.Type > 0 {
		u.Type--
	}
}
