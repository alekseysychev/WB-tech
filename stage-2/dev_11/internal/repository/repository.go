package repository

type Repository interface{}

type repository struct{}

func New() (Repository, error) {
	return &repository{}, nil
}
