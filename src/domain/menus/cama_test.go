package menus

import (
	"testing"
)

func TestGetMenus(t *testing.T) {
	menus := GetMenus()

	if menus.MenuVersion != "2020W" {
		t.Errorf("bad menu_version")
	}
}
