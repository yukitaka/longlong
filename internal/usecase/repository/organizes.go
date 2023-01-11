package repository

import "github.com/yukitaka/longlong/internal/entity"

type Organizes interface {
	Find(int64) (*entity.Organize, error)
}
