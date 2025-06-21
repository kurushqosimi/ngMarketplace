package category

type getCategoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
