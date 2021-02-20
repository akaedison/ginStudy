package service

import "gomod/repository"

type MovieService struct {
	Repo repository.MovieRepoI
}

func (s MovieService) CreateTableMoviesS() {
	s.Repo.CreateTableMovies()
}
