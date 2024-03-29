package usecase

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/value_object"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/server/core/mock/repository"
)

func TestAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctrl2 := gomock.NewController(t)
	defer ctrl2.Finish()

	token, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	authRep := mockRepository.NewMockAuthentications(ctrl)
	organizationRep := mockRepository.NewMockOrganizations(ctrl)
	memberRep := mockRepository.NewMockOrganizationMembers(ctrl2)
	authRep.EXPECT().FindToken("TestUser").Return(1, string(token), nil)

	individual := entity.NewIndividual(1, entity.NewUser(0), entity.NewProfile(0, "", "", ""), "TestUser")
	memberRep.EXPECT().IndividualsAssigned(&[]entity.Individual{*individual}).Return(&[]entity.OrganizationMember{*entity.NewOrganizationMember(entity.NewOrganization(0, 1, "TestOrganization"), individual, value_object.ADMIN)}, nil)
	organizationRep.EXPECT().Find(1).Return(entity.NewOrganization(1, 1, "TestOrganization"), nil)

	rep := NewAuthenticationRepository(authRep, organizationRep, memberRep)
	itr := NewAuthentication(rep)
	id, _ := itr.Auth("TestOrganization", "TestUser", "password")
	if id != 1 {
		t.Errorf("Auth = %v", id)
	}
}
