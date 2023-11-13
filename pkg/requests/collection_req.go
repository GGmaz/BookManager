package requests

type CreateBookCollectionRequest struct {
	Name string `json:"name"`
}

type CollectionByIdRequest struct {
	Id int64 `uri:"collectionId" binding:"required"`
}

type BookCollectionRequest struct {
	CollectionId int64 `uri:"collectionId" binding:"required"`
	BookId       int64 `uri:"bookId" binding:"required"`
}
