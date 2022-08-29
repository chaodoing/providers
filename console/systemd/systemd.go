package systemd

type Systemd struct {
	Description string
	Execute     string
	Directory   string
	Username    string
	Group       string
}

func NewSystemd() (data Systemd, err error) {
	var (
		description description
		execute     execute
		directory   directory
		u           *users
	)
	description, err = Description()
	if err != nil {
		return
	}
	execute, err = Execute()
	if err != nil {
		return
	}
	directory, err = Directory()
	if err != nil {
		return
	}
	u, err = Users()
	data = Systemd{
		Description: string(description),
		Execute:     string(execute),
		Directory:   string(directory),
		Username:    u.Username,
		Group:       u.Group,
	}
	return
}
