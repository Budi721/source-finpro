package model

import (
	"encoding/json"
)

type Project struct {
	IdProject          int    `json:"id_project" gorm:"primaryKey"`
	KategoriProject    string `json:"kategori_project"`
	NamaProject        string `json:"nama_project"`
	TanggalMulai       int64  `json:"tanggal_mulai"`
	LinkTrello         string `json:"link_trello"`
	DeskripsiProject   string `json:"deskripsi_project"`
	InvitedUserId      []User `json:"invited_user_id" gorm:"many2many:user_invited;"`
	CollaboratorUserId []User `json:"collaborator_user_id" gorm:"many2many:user_collaborator;"`
	Admin              int    `json:"admin"`
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
