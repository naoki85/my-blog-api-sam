package model

type User struct {
	Id                           int
	Email                        string
	Username                     string
	EncryptedPassword            string
	ResetPasswordToken           string
	ResetPasswordSentAt          string
	Provider                     string
	Uid                          string
	ImageUrl                     string
	AuthenticationToken          string
	AuthenticationTokenExpiredAt string
	OnetimeToken                 string
	OnetimeTokenExpiredAt        string
	CreatedAt                    string
	UpdatedAt                    string
}
