package main

import v1 "test3/v1"

func main() {
	getpi := v1.NewGetPi()
	repo := v1.NewApi(getpi)
	err := repo.Handler()
	if err != nil {
		return
	}
}
