package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/value_object"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewOrganizationManagerMembers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	organizationRep := mockRepository.NewMockOrganizations(ctrl)
	memberRep := mockRepository.NewMockOrganizationMembers(ctrl)
	individualRep := mockRepository.NewMockIndividuals(ctrl)

	organization := entity.NewOrganization(0, 1, "Test")
	members := &[]entity.OrganizationMember{
		{organization, entity.NewIndividual(
			1,
			*entity.NewUser(1),
			*entity.NewProfile(1, "", "", ""),
			"Test1"), value_object.MEMBER},
		{organization, entity.NewIndividual(
			2,
			*entity.NewUser(2),
			*entity.NewProfile(2, "", "", ""),
			"Test2"), value_object.MEMBER},
	}
	memberRep.EXPECT().Members(organization, individualRep).Return(members, nil)

	itr := NewOrganizationManager(organization, organizationRep, memberRep, individualRep)
	members, err := itr.Members()
	if err != nil {
		t.Errorf("QuitIndividual() = %v\n", err)
	}
	if members == nil {
		t.Errorf("Member is nil\n")
	}
	for i, m := range *members {
		if m.Individual.Id != i+1 || m.Individual.Name != "Test"+strconv.Itoa(i+1) {
			t.Errorf("Member is %v expect number %d\n", m.Individual, i+1)
		}
	}
}
