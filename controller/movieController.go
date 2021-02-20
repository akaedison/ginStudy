package controller

import "gomod/service"

type Movie struct {
	Service service.MovieService `inject:""`
}

func (m *Movie) CreateTableMovie() {
	m.Service.CreateTableMoviesS()
}
