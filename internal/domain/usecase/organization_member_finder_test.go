package usecase

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	mockRepository "github.com/yukitaka/longlong/mock/repository"
)

func TestNewOrganizationMemberFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var members []entity.OrganizationMember
	memberRep := mockRepository.NewMockOrganizationMembers(ctrl)
	for oid := 1; oid <= 5; oid++ {
		for iid := 1; iid < 5; iid++ {
			member := entity.OrganizationMember{Organization: entity.NewOrganization(0, oid, "Organization "+strconv.Itoa(oid)), Individual: entity.NewIndividual(iid, iid, iid, "Individual "+strconv.Itoa(iid))}
			members = append(members, member)
			memberRep.EXPECT().Find(gomock.Any(), gomock.Any()).DoAndReturn(
				func(oid int, iid int) (*entity.OrganizationMember, error) {
					member := entity.OrganizationMember{Organization: entity.NewOrganization(0, oid, "Organization "+strconv.Itoa(oid)), Individual: entity.NewIndividual(iid, iid, iid, "Individual "+strconv.Itoa(iid))}
					return &member, nil
				}).AnyTimes()
		}
	}

	itr := NewOrganizationMemberFinder(memberRep)
	member, err := itr.FindById(1, 2)
	if err != nil {
		t.Errorf("QuitIndividual() = %v\n", err)
	}
	if member.Organization.Id != 1 || member.Individual.Id != 2 {
		t.Errorf("Found member is invalid (%d, %d) expect (1, 2)\n", member.Organization.Id, member.Individual.Id)
	}
}
