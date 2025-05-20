package http

// listResourceRequest represents the request body for listing resources
type listResourceRequest struct {
	Page int64 `query:"page" form:"page" binding:"min=1" example:"0"`
	Size int64 `query:"size" form:"size" binding:"min=1,max=1000" example:"20"`
}
