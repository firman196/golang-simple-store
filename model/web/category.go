package web

type CategoryCreateInput struct {
	Icon         string `json:"icon"`
	CategoryName string `json:"category_name" binding:"required,max=100"`
	Slug         string `json:"slug" binding:"required"`
	Notes        string `json:"notes"`
	IsAktif      string `json:"is_aktif" binding:"required"`
	ParentId     string `json:"parent_id"`
}

type CategoryUpdateInput struct {
	Id           int    `json:"id" binding:"required"`
	Icon         string `json:"icon"`
	CategoryName string `json:"category_name" binding:"required,max=100"`
	Slug         string `json:"slug" binding:"required"`
	Notes        string `json:"notes"`
	IsAktif      string `json:"is_aktif" binding:"required"`
	ParentId     string `json:"parent_id"`
}
