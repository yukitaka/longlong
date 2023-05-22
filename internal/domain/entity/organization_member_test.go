package entity

import (
	"github.com/yukitaka/longlong/internal/domain/value_object"
	"testing"
)

func TestOrganizationMember_AccessLevel(t *testing.T) {
	owner := NewOrganizationMember(NewOrganization(0, 0, ""), NewIndividual(0, NewUser(0), NewProfile(0, "", "", ""), ""), value_object.OWNER)
	admin := NewOrganizationMember(NewOrganization(0, 0, ""), NewIndividual(1, NewUser(1), NewProfile(1, "", "", ""), ""), value_object.ADMIN)
	member := NewOrganizationMember(NewOrganization(0, 0, ""), NewIndividual(2, NewUser(2), NewProfile(2, "", "", ""), ""), value_object.MEMBER)

	if !owner.IsOwner() {
		t.Errorf("IsOwner = %v, the user is owner.", owner.IsOwner())
	}

	if owner.IsAdmin() || owner.IsMember() {
		t.Errorf("IsAdmin = %v, IsMember = %v, the user is owner.", owner.IsAdmin(), owner.IsMember())
	}

	if !admin.IsAdmin() {
		t.Errorf("IsAdmin = %v, the user is admin.", admin.IsAdmin())
	}

	if admin.IsOwner() || admin.IsMember() {
		t.Errorf("IsOwner = %v, IsMember = %v, the user is admin.", admin.IsOwner(), admin.IsMember())
	}

	if !member.IsMember() {
		t.Errorf("IsMember = %v, the user is member.", member.IsMember())
	}

	if member.IsOwner() || member.IsAdmin() {
		t.Errorf("IsOwner = %v, IsAdmin = %v, the user is member.", member.IsOwner(), member.IsAdmin())
	}
}

func TestOrganizationMember_CanManage(t *testing.T) {
	operator := NewOrganizationMember(NewOrganization(0, 0, ""), NewIndividual(0, NewUser(0), NewProfile(0, "", "", ""), ""), value_object.ADMIN)
	owner := NewOrganizationMember(NewOrganization(0, 0, ""), NewIndividual(1, NewUser(1), NewProfile(1, "", "", ""), ""), value_object.OWNER)
	admin := NewOrganizationMember(NewOrganization(0, 0, ""), NewIndividual(2, NewUser(2), NewProfile(2, "", "", ""), ""), value_object.ADMIN)

	if operator.CanManage(operator) {
		t.Errorf("CanManage = %v, the user can't manage own.", operator.CanManage(operator))
	}

	if operator.CanManage(owner) {
		t.Errorf("CanManage = %v, the user can't manage above access level.", operator.CanManage(owner))
	}

	if !operator.CanManage(admin) {
		t.Errorf("CanManage = %v, the user can manage same access level.", operator.CanManage(admin))
	}
}
