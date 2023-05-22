package entity

import (
	"github.com/yukitaka/longlong/internal/domain/value_object"
	"testing"
)

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
