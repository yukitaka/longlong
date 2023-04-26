//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/yukitaka/longlong/internal/domain/entity"
)

type OrganizationBelongings interface {
	Entry(individualId int64) error
	Leave(individualId int64, reason string) error
	Members(organization *entity.Organization, individualRepository Individuals) (*[]entity.OrganizationBelonging, error)
	IndividualsAssigned(individual *[]entity.Individual) (*[]entity.OrganizationBelonging, error)
	Close()
}
