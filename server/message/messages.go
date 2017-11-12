package message

const (
	KindConnected = iota + 1
	KindUserJoined
	KindUserLeft
	KindStroke
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type User struct {
	ID    int    `json:"id"`
	Color string `json:"color"`
}

type Connected struct {
	Kind  int    `json:"kind"`
	Color string `json:"color"`
	Users []User `json:"users"`
}

func NewConnected(color string, users []User) *Connected {
	return &Connected{
		Kind:  KindConnected,
		Color: color,
		Users: users,
	}
}

type UserJoined struct {
	Kind int  `json:"kind"`
	User User `json:"user"`
}

func NewUserJoined(userID int, color string) *UserJoined {
	return &UserJoined{
		Kind: KindUserJoined,
		User: User{ID: userID, Color: color},
	}
}

type UserLeft struct {
	Kind   int `json:"kind"`
	UserID int `json:"userId"`
}

func NewUserLeft(userID int) *UserLeft {
	return &UserLeft{
		Kind:   KindUserLeft,
		UserID: userID,
	}
}

type Stroke struct {
	Kind   int     `json:"kind"`
	UserID int     `json:"userId"`
	Points []Point `json:"points"`
	Finish bool    `json:"finish"`
}
