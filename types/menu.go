package types

import "github.com/a20070322/go_fast_admin/ent"

type RepGetUserMenu struct {
	Menu []*MenusTree `json:"menu"`
	Role []string     `json:"role"`
}

type MenusTree struct {
	*ent.AdminMenus
	Children []*MenusTree `json:"children"`
}



type MenuSlice []*ent.AdminMenus

func (a MenuSlice) Len() int {
	return len(a)
}
func (a MenuSlice) Swap(i, j int){
	a[i], a[j] = a[j], a[i]
}
func (a MenuSlice) Less(i, j int) bool {
	return a[j].Sort < a[i].Sort
}