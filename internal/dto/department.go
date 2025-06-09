package dto

type DepartmentBody struct {
	Name        string `json:"name,omitempty"`
	UserIds     []*int `json:"userIds,omitempty"`
	ParentId    *int   `json:"parentId,omitempty"`
	Description string `json:"description,omitempty"`
}

type UsersIds struct {
	Ids []int `json:"ids"`
}
