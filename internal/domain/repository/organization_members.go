//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
	"github.com/yukitaka/longlong/internal/domain/value_object"
)

type OrganizationMembers interface {
	Find(organizationId, individualId int64) (*entity.OrganizationMember, error)
	Entry(organizationId, individualId int64, role value_object.Role) error
	Leave(individualId int64, reason string) error
	Members(organization *entity.Organization, individualRepository Individuals) (*[]entity.OrganizationMember, error)
	IndividualsAssigned(individual *[]entity.Individual) (*[]entity.OrganizationMember, error)
	Close()
}
