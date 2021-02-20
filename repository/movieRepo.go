package repository

import (
	"fmt"
)

type MovieRepo struct {
}

func (m MovieRepo) CreateTableMovies() {
	fmt.Println("成功")
	/*_ = r.Source.Connect()
	err1 := r.Source.DB().AutoMigrate(&model.Movie{})
	if err1 != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}*/
}
