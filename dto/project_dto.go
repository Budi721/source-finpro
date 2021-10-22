package dto

import "time"

type Project struct {
	KategoriProject  string    `json:"kategori_project" binding:"required"`
	NamaProject      string    `json:"nama_project" binding:"required"`
	StartDate        string    `json:"start_date" binding:"required"`
	Date             time.Time `json:"-"`
	LinkTrello       string    `json:"link_trello" binding:"required"`
	DeskripsiProject string    `json:"deskripsi_project" binding:"required"`
	InvitedUserId    []uint64  `json:"invited_user_id" binding:"required"`
	CollaboratorUserId []uint16 `json:"collaborator_user_id"`
	Creator          uint64    `json:"creator"`
}
