package tests

import (
	"fmt"
	"testing"

	u "github.com/ivar1309/Api-Go-Boilerplate/internal/utils"
)

func TestHasPermissionDriven(t *testing.T) {
	fullPerm := u.AdminOnly | u.CreatePublic | u.ReadPublic | u.UpdatePublic | u.DeletePublic | u.CreateProtected | u.ReadProtected | u.UpdateProtected | u.DeleteProtected
	publicPerm := u.CreatePublic | u.ReadPublic | u.UpdatePublic | u.DeletePublic
	protectedPerm := u.CreateProtected | u.ReadProtected | u.UpdateProtected | u.DeleteProtected
	viewerPerm := u.ReadPublic | u.ReadProtected
	roleUser := "user"
	roleAdmin := "admin"
	roleViewer := "viewer"

	var tests = []struct {
		role string
		perm u.Permission
		want bool
	}{
		{roleAdmin, fullPerm, true},
		{roleAdmin, publicPerm, true},
		{roleAdmin, protectedPerm, true},
		{roleUser, fullPerm, false},
		{roleUser, publicPerm, true},
		{roleUser, protectedPerm, true},
		{roleViewer, fullPerm, false},
		{roleViewer, publicPerm, false},
		{roleViewer, protectedPerm, false},
		{roleViewer, viewerPerm, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v HasPermission: %v", tt.role, tt.perm)
		t.Run(testname, func(t *testing.T) {
			got := u.HasPermission(tt.role, tt.perm)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
