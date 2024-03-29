//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/server/core/pkg/domain/entity"
	"github.com/yukitaka/longlong/server/core/pkg/domain/value_object"
)

type OrganizationMembers interface {
	Find(organizationId, individualId int) (*entity.OrganizationMember, error)
	Entry(organizationId, individualId int, role value_object.Role) error
	Leave(organizationId, individualId int, reason string) error
	Members(organization *entity.Organization, individualRepository Individuals) (*[]entity.OrganizationMember, error)
	IndividualsAssigned(individual *[]entity.Individual) (*[]entity.OrganizationMember, error)
}
