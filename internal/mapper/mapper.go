package mapper

import (
	"github.com/jinzhu/copier"
)

func MapObject[S, D any](source S) (D, error) {
	var destination D
	if err := copier.Copy(&destination, &source); err != nil {
		return *new(D), err
	}
	return destination, nil
}

func MapList[S, D any](source []S) ([]D, error) {
	destination := []D{}
	if err := copier.Copy(&destination, &source); err != nil {
		return nil, err
	}
	return destination, nil
}
