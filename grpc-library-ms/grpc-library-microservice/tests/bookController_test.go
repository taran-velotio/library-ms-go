package tests

import (
	"context"
	"errors"
	"library-comp/proto/book"
	mock_book "library-comp/proto/tests/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestBookController_GetBook(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *book.GetBookResonse
		err error
	}

	tests := map[string]struct {
		in       *book.GetBookRequest
		expected expectation
	}{
		"Success": {
			in: &book.GetBookRequest{
				Id: 1,
			},
			expected: expectation{
				out: &book.GetBookResonse{
					Book: &book.Book{
						Id:     1,
						Name:   "Taran",
						Author: "Book1",
						Price:  120,
					},
				},
				err: nil,
			},
		},
		"Failure": {
			in: &book.GetBookRequest{
				Id: 4,
			},
			expected: expectation{
				out: &book.GetBookResonse{},
				err: errors.New("record not found"),
			},
		},
	}

	ctrl := gomock.NewController(t)
	client := mock_book.NewMockBookServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			client.EXPECT().GetBook(ctx, tt.in).Return(tt.expected.out, tt.expected.err)
			out, err := client.GetBook(ctx, tt.in)
			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:\n%+v\nGot:\n%+v\n", tt.expected.out, out)
			}

			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected: %q\nGot: %q\n", tt.expected.err, err)
			}

		})
	}
}
func TestBookController_GetListOfBooks(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *book.GetListOfBooksResponse
		err error
	}

	tests := map[string]struct {
		in       *book.GetListOfBooksRequest
		expected expectation
	}{
		"Success": {
			in: &book.GetListOfBooksRequest{},
			expected: expectation{
				out: &book.GetListOfBooksResponse{
					Books: []*book.Book{
						{
							Id:     1,
							Name:   "Taran",
							Author: "F.Scott",
							Price:  110,
						},
						{
							Id:     2,
							Name:   "Karan",
							Author: "Book4",
							Price:  999,
						},
					},
				},
				err: nil,
			},
		},
		"Failure": {
			in: &book.GetListOfBooksRequest{}, // Provide invalid inputs here to simulate failure case.
			expected: expectation{
				out: nil,
				err: errors.New("invalid input"),
			},
		},
	}

	ctrl := gomock.NewController(t)
	client := mock_book.NewMockBookServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			client.EXPECT().GetListOfBooks(ctx, tt.in).Return(tt.expected.out, tt.expected.err)
			out, err := client.GetListOfBooks(ctx, tt.in)

			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:\n%+v\nGot:\n%+v\n", tt.expected.out, out)
			}

			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected: %q\nGot: %q\n", tt.expected.err, err)
			}

		})
	}
}
func TestBookController_CreateBook(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *book.CreateBookResponse
		err error
	}

	tests := map[string]struct {
		in       *book.CreateBookRequest
		expected expectation
	}{
		"Success": {

			in: &book.CreateBookRequest{
				Name:   "Aman",
				Author: "Fitzgerald",
				Price:  100,
			},
			expected: expectation{
				out: &book.CreateBookResponse{
					Book: &book.Book{
						Id:     1,
						Name:   "Aman",
						Author: "Fitzgerald",
						Price:  100,
					},
				},
				err: nil,
			},
		},
		"Failure": {
			in: &book.CreateBookRequest{},
			expected: expectation{
				out: nil,
				err: errors.New("invalid input"),
			},
		},
	}

	ctrl := gomock.NewController(t)
	client := mock_book.NewMockBookServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			client.EXPECT().CreateBook(ctx, tt.in).Return(tt.expected.out, tt.expected.err)
			out, err := client.CreateBook(ctx, tt.in)

			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:\n%+v\nGot:\n%+v\n", tt.expected.out, out)
			}
			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected: %q\nGot: %q\n", tt.expected.err, err)
			}

		})
	}
}
func TestBookController_UpdateBook(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *book.UpdateBookResponse
		err error
	}

	tests := map[string]struct {
		in       *book.UpdateBookRequest
		expected expectation
	}{
		"Success": {
			in: &book.UpdateBookRequest{
				Id:     1,
				Name:   "Book4",
				Author: "Taran",
				Price:  199,
			},
			expected: expectation{
				out: &book.UpdateBookResponse{
					Book: &book.Book{
						Id:     1,
						Name:   "Book4",
						Author: "Taran",
						Price:  199,
					},
				},
				err: nil,
			},
		},
		"Failure": {
			in: &book.UpdateBookRequest{},
			expected: expectation{
				out: nil,
				err: errors.New("invalid input"),
			},
		},
	}

	ctrl := gomock.NewController(t)
	client := mock_book.NewMockBookServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			client.EXPECT().UpdateBook(ctx, tt.in).Return(tt.expected.out, tt.expected.err)
			out, err := client.UpdateBook(ctx, tt.in)

			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:\n%+v\nGot:\n%+v\n", tt.expected.out, out)
			}
			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected: %q\nGot: %q\n", tt.expected.err, err)
			}

		})
	}
}
func TestBookController_DeleteBook(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *book.DeleteBookResponse
		err error
	}

	tests := map[string]struct {
		in       *book.DeleteBookRequest
		expected expectation
	}{
		"Success": {
			in: &book.DeleteBookRequest{
				Id: 1,
			},
			expected: expectation{
				out: &book.DeleteBookResponse{},
				err: nil,
			},
		},
		"Failure": {
			in: &book.DeleteBookRequest{},
			expected: expectation{
				out: nil,
				err: errors.New("invalid input"),
			},
		},
	}

	ctrl := gomock.NewController(t)
	client := mock_book.NewMockBookServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			client.EXPECT().DeleteBook(ctx, tt.in).Return(tt.expected.out, tt.expected.err)
			out, err := client.DeleteBook(ctx, tt.in)

			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:\n%+v\nGot:\n%+v\n", tt.expected.out, out)
			}
			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected: %q\nGot: %q\n", tt.expected.err, err)
			}

		})
	}
}
