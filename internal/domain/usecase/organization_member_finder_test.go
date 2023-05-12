package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewOrganizationMemberFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	organizationId := 1
	individualId := 1

	members := []entity.OrganizationMember{
		{Organization: entity.NewOrganization(0, 1, "One"), Individual: entity.NewIndividual(1, 1, 1, "One")},
	}

	memberRep := mockRepository.NewMockOrganizationMembers(ctrl)
	memberRep.EXPECT().Find(organizationId, individualId).Return(&members[0], nil)

	itr := NewOrganizationMemberFinder(memberRep)
	member, err := itr.FindById(1, 1)
	if err != nil {
		t.Errorf("QuitIndividual() = %v\n", err)
	}
	if member.Organization.Id != organizationId || member.Individual.Id != individualId {
		t.Errorf("Found member is invalid (%d, %d) expect (%d, %d)\n", member.Organization.Id, member.Individual.Id, organizationId, individualId)
	}
}
