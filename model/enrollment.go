package model

type Enrollment struct {
    IdUser           int    `json:"id_user"`
    NamaLengkap      string `json:"nama_lengkap"`
    Username         string `json:"username"`
    Password         string `json:"password"`
    TopikDiminati    string `json:"topik_diminati"`
    EnrollmentStatus int    `json:"enrollment_status"`
}
