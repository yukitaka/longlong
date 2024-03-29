package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/value_object"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/server/core/mock/repository"
)

func TestNewUserAssigned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	organ := entity.NewOrganization(0, 1, "Test")
	user := entity.NewUser(1)
	profile := entity.NewProfile(1, "", "", "")
	individuals := []entity.Individual{*entity.NewIndividual(1, user, profile, "Test")}
	member := entity.NewOrganizationMember(organ, &individuals[0], value_object.MEMBER)
	members := []entity.OrganizationMember{*member}

	individualRep := mockRepository.NewMockIndividuals(ctrl)
	organizationRep := mockRepository.NewMockOrganizations(ctrl)
	memberRep := mockRepository.NewMockOrganizationMembers(ctrl)

	individualRep.EXPECT().FindByUserId(1).Return(&individuals, nil)
	memberRep.EXPECT().IndividualsAssigned(&individuals).Return(&members, nil)
	organizationRep.EXPECT().FindAll(gomock.Any()).Return(&[]entity.Organization{*organ}, nil)

	rep := NewUserAssignedRepository(individualRep, organizationRep, memberRep)
	itr := NewUserAssigned(rep)
	organizations, _ := itr.OrganizationList(member)
	if organizations == nil {
		t.Error("Organization List is nil\n")
	}
}
