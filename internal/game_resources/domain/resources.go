package domain

type GameItem[T any] struct {
	Id          int
	Title       string
	Description string
	Metadata    []GameItemMetadata[T]
}

type GameItemMetadata[T any] struct {
	Type       string
	GameItemID int
	Data       T
}
