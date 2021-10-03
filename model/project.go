package model

import (
	"encoding/json"
)

type Project struct {
	IdProject          int    `json:"id_project,omitempty" gorm:"primaryKey"`
	KategoriProject    string `json:"kategori_project,omitempty"`
	NamaProject        string `json:"nama_project,omitempty"`
	TanggalMulai       int64  `json:"tanggal_mulai,omitempty"`
	LinkTrello         string `json:"link_trello,omitempty"`
	DeskripsiProject   string `json:"deskripsi_project,omitempty"`
	InvitedUserId      []int  `json:"invited_user_id,omitempty" gorm:"-"`
	CollaboratorUserId []int  `json:"collaborator_user_id,omitempty" gorm:"-"`
	Admin              int    `json:"admin,omitempty"`
	UsersInvited       []User `json:"-" gorm:"many2many:user_invited;"`
	UsersCollaborator  []User `json:"-" gorm:"many2many:user_collaborator;"`
}

func (p Project) GetInvitedUserId(dataJson string) []int {
	var result []int

	json.Unmarshal([]byte(dataJson), &result)
	return result
}

func (p Project) SetInvitedUserId(invitedId []int) string {
	var result []int
	for _, id := range invitedId {
		result = append(result, id)
	}
	data, _ := json.Marshal(result)
	return string(data)
}

func (p Project) GetCollaboratorUserId(dataJson string) []int {
	var result []int

	json.Unmarshal([]byte(dataJson), &result)
	return result
}

func (p Project) SetCollaboratorUserId(collaboratorId []int) string {
	var result []int
	for _, id := range collaboratorId {
		result = append(result, id)
	}
	data, _ := json.Marshal(result)
	return string(data)
}
