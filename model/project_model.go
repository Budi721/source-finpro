package model

import "time"

type Project struct {
	IdProject          int       `json:"id_project,omitempty" gorm:"primaryKey"`
	KategoriProject    string    `json:"kategori_project,omitempty"`
	NamaProject        string    `json:"nama_project,omitempty"`
	StartDate          time.Time `json:"start_date,omitempty"`
	LinkTrello         string    `json:"link_trello,omitempty"`
	DeskripsiProject   string    `json:"deskripsi_project,omitempty"`
	InvitedUserId      []uint64  `json:"invited_user_id,omitempty" gorm:"-"`
	CollaboratorUserId []uint64  `json:"collaborator_user_id,omitempty" gorm:"-"`
	Creator            uint64    `json:"creator,omitempty"`
	UsersInvited       []User    `json:"-" gorm:"many2many:user_invited;"`
	UsersCollaborator  []User    `json:"-" gorm:"many2many:user_collaborator;"`
}
