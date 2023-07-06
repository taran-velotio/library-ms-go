// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "library-comp/models"

	mock "github.com/stretchr/testify/mock"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: bookDetails
func (_m *BookRepository) CreateBook(bookDetails models.Book) (models.Book, error) {
	ret := _m.Called(bookDetails)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(models.Book) models.Book); ok {
		r0 = rf(bookDetails)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Book) error); ok {
		r1 = rf(bookDetails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBook provides a mock function with given fields: id
func (_m *BookRepository) DeleteBook(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBook provides a mock function with given fields: id
func (_m *BookRepository) GetBook(id int) (models.Book, error) {
	ret := _m.Called(id)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(int) models.Book); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListOfBooks provides a mock function with given fields:
func (_m *BookRepository) GetListOfBooks() ([]models.Book, error) {
	ret := _m.Called()

	var r0 []models.Book
	if rf, ok := ret.Get(0).(func() []models.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBook provides a mock function with given fields: bookDetails
func (_m *BookRepository) UpdateBook(bookDetails models.Book) (models.Book, error) {
	ret := _m.Called(bookDetails)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(models.Book) models.Book); ok {
		r0 = rf(bookDetails)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Book) error); ok {
		r1 = rf(bookDetails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
