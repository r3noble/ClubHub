package main

type Club struct {
	clubName    string
	memberCount int
	execMembers map[string]string
}

type Clubber interface {
	IncrimentFollowerCount()
	GetFollowerCount() int
	UpdateName(name string)
	GetName() string
	SetExec(map[string]string)
	GetExec() map[string]string
}

func (c *Club) IncrimentFollowerCount() {
	c.memberCount++
}

func (c *Club) GetFollowerCount() int {
	return c.memberCount
}

func (c *Club) UpdateName(name string) {
	c.clubName = name
}

func (c *Club) GetName() string {
	return c.clubName
}

func (c *Club) SetExec(exec map[string]string) {
	c.execMembers = exec
}

func (c *Club) GetExec() map[string]string {
	return c.execMembers
}
