package entity

type Book interface {
	GetID() uint
	GetTitle() string
	IsAvailable() bool
	// GetEntity() BookEntity
}

type BookEntity struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	SoldOut bool   `json:"soldOut"`
}

func (b *BookEntity) GetID() uint {
	return b.ID
}

func (b *BookEntity) GetTitle() string {
	return b.Title
}

func (b *BookEntity) IsAvailable() bool {
	return !(b.SoldOut)

}

// func (b *BookEntity) GetEntity() BookEntity {
// 	return *b
// }
