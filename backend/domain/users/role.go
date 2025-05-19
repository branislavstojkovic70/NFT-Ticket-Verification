package domain

type Role string

const (
	RoleUser      Role = "user"
	RoleOrganizer Role = "organizer"
	RoleAdmin     Role = "admin"
)
