package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/value_object"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/mock/repository"
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
	authRep.EXPECT().FindToken("TestUser").Return(int64(1), string(token), nil)

	individual := entity.NewIndividual(1, 0, 0, "TestUser")
	memberRep.EXPECT().IndividualsAssigned(&[]entity.Individual{*individual}).Return(&[]entity.OrganizationMember{*entity.NewOrganizationMember(entity.NewOrganization(0, 1, "TestOrganization"), individual, value_object.ADMIN)}, nil)
	organizationRep.EXPECT().Find(int64(1)).Return(entity.NewOrganization(int64(1), int64(1), "TestOrganization"), nil)

	itr := NewAuthentication(authRep, organizationRep, memberRep)
	id, _ := itr.Auth("TestOrganization", "TestUser", "password")
	if id != 1 {
		t.Errorf("Auth = %v", id)
	}
}
