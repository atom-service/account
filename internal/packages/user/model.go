package user

type User struct {
	ID       string   `bson:"_"`
	Admin    bool     `bson:"admin"`
	Email    string   `bson:"email"`
	Username string   `bson:"username"`
	Password string   `bson:"password"`
}
