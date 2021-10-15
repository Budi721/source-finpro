package model

type Enrollment struct {
	IdUser           uint64 `json:"id_user,omitempty"`
	NamaLengkap      string `json:"nama_lengkap,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	TopikDiminati    string `json:"topik_diminati,omitempty"`
	EnrollmentStatus uint16 `gorm:"default:0" json:"enrollment_status,omitempty"`
	Token            string `gorm:"-" json:"token,omitempty"`
}
