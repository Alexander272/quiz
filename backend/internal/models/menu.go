package models

type Menu struct {
	Id         string `json:"id" db:"id"`
	RoleId     string `json:"-" db:"role_id"`
	RoleName   string `json:"roleName" db:"name"`
	ItemId     string `json:"menuItemId" db:"menu_item_id"`
	ItemName   string `json:"menuItemName" db:"item_name"`
	ItemMethod string `json:"menuItemMethod" db:"method"`
}

type MenuFull struct {
	Id        string      `json:"id" db:"id"`
	Role      *RoleFull   `json:"role"`
	MenuItems []*MenuItem `json:"menuItems"`
}

type MenuDTO struct {
	Id         string `json:"id"`
	RoleId     string `json:"roleId"`
	MenuItemId string `json:"menuItemId"`
}
