package models

type User struct {
	Name               string `db:"name"`
	Email              string `db:"email"`
	PhoneNumber        string `db:"phone_number"`
	Password           string `db:"password"`
	Avatar             string `db:"avatar"`
	Birthday           string `db:"birthday"`
	RecognitionReceive string `db:"recognition_receive`
	RecognitionSend    string `db:"recognition_send`
}
