package systemd

import (
	`os/user`
)

type users struct {
	Username string
	Group    string
}

// Users 获取当前用户
func Users() (*users, error) {
	u, err := user.Lookup("daemon")
	if err != nil {
		u, err = user.Current()
		if err != nil {
			return nil, err
		}
	}
	group, err := user.LookupGroupId(u.Gid)
	if err != nil {
		return nil, err
	}
	
	return &users{
		Username: u.Username,
		Group:    group.Name,
	}, nil
}
