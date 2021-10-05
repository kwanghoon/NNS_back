package train

import "nns_back/query"

//go:generate mockery --name TrainRepository --inpackage
type TrainRepository interface {
	FindNextTrainNo(userId int64) (int64, error)
	CountCurrentTraining(userId int64) (int, error)

	Insert(train Train) (int64, error)
	Delete(opts ...Option) error
	Find(opts ...Option) (Train, error)
	FindAll(opts ...query.Option) ([]Train, error)
	Update(train Train, opts ...Option) error
}