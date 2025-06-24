package model

import "time"

type Post struct {
	ID        string    `db:"id"`         // UUID tercih edilir
	Username  string    `db:"username"`   // Gönderiyi oluşturan kullanıcı (foreign key gibi)
	Content   string    `db:"content"`    // Gönderi içeriği (text)
	MediaURL  *string   `db:"media_url"`  // Opsiyonel foto/video linki
	CreatedAt time.Time `db:"created_at"` // Oluşturulma tarihi
	UpdatedAt time.Time `db:"updated_at"` // Güncellenme tarihi
	IsDeleted bool      `db:"is_deleted"` // Soft delete için
}
