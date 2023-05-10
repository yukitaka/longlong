package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewOrganizationManager(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	organizationRep := mockRepository.NewMockOrganizations(ctrl)
	memberRep := mockRepository.NewMockOrganizationMembers(ctrl)
	individualRep := mockRepository.NewMockIndividuals(ctrl)

	organization := entity.NewOrganization(0, 1, "Test")
	memberRep.EXPECT().Members(organization, individualRep).Return(nil, nil)

	itr := NewOrganizationManager(organization, organizationRep, memberRep, individualRep)
	members, err := itr.Members()
	if err != nil {
		t.Errorf("QuitIndividual() = %v", err)
	}
	if members != nil {
		t.Errorf("QuitIndividual() = %v", members)
	}
}
