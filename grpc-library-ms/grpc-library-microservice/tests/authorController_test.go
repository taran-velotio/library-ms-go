package tests

import (
	"context"
	"errors"
	"reflect"
	"testing"

	author "library-comp/proto/author/author"
	mock_author "library-comp/proto/author/tests/mocks"

	"github.com/golang/mock/gomock"
)

func TestAuthorController_GetAuthor(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *author.GetAuthorResonse
		err error
	}

	tests := map[string]struct {
		in       *author.GetAuthorRequest
		expected expectation
	}{
		"Success": {
			in: &author.GetAuthorRequest{
				Id: 1,
			},
			expected: expectation{
				out: &author.GetAuthorResonse{
					Author: &author.Author{
						Id:   1,
						Name: "Taran",
					},
				},
				err: nil,
			},
		},
		"Failure": {
			in: &author.GetAuthorRequest{
				Id: 4,
			},
			expected: expectation{
				out: &author.GetAuthorResonse{},
				err: errors.New("record not found"),
			},
		},
	}

	ctrl := gomock.NewController(t)
	client := mock_author.NewMockAuthorServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			client.EXPECT().GetAuthor(ctx, tt.in).Return(tt.expected.out, tt.expected.err)
			out, err := client.GetAuthor(ctx, tt.in)
			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:\n%+v\nGot:\n%+v\n", tt.expected.out, out)
			}

			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected: %q\nGot: %q\n", tt.expected.err, err)
			}
		})
	}
}
func TestAuthorController_GetListOfAuthors(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *author.GetListOfAuthorsResponse
		err error
	}

	tests := map[string]struct {
		in       *author.GetListOfAuthorsRequest
		expected expectation
	}{
		"Success": {
			in: &author.GetListOfAuthorsRequest{},
			expected: expectation{
				out: &author.GetListOfAuthorsResponse{
					Authors: []*author.Author{
						{Id: 1, Name: "Taran"},
						{Id: 2, Name: "Karan"},
					},
				},
				err: nil,
			},
		},
		"Failure": {
			in: &author.GetListOfAuthorsRequest{},
			expected: expectation{
				out: &author.GetListOfAuthorsResponse{},
				err: errors.New("record not found"),
			},
		},
	}

	ctrl := gomock.NewController(t)
	client := mock_author.NewMockAuthorServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			// instance of gomock controller
			client.EXPECT().GetListofAuthors(ctx, tt.in).Return(tt.expected.out, tt.expected.err)

			//actual AuthorController GetListOfAuthors call
			out, err := client.GetListofAuthors(ctx, tt.in)
			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:\n%+v\nGot:\n%+v\n", tt.expected.out, out)
			}

			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected:%q\nGot:%q\n", tt.expected.err, err)
			}
		})
	}
}
func TestAuthorController_CreateAuthor(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *author.CreateAuthorResponse
		err error
	}

	tests := map[string]struct {
		in       *author.CreateAuthorRequest
		expected expectation
	}{
		"Success": {
			in: &author.CreateAuthorRequest{
				Name: "Taran",
			},
			expected: expectation{
				out: &author.CreateAuthorResponse{
					Author: &author.Author{
						Id:   1,
						Name: "Taran",
					},
				},
				err: nil,
			},
		}, "Failure": {
			in: &author.CreateAuthorRequest{},
			expected: expectation{
				out: &author.CreateAuthorResponse{},
				err: errors.New("record not found"),
			},
		},
	}

	ctrl := gomock.NewController(t)
	client := mock_author.NewMockAuthorServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			client.EXPECT().CreateAuthor(ctx, tt.in).Return(tt.expected.out, tt.expected.err)

			out, err := client.CreateAuthor(ctx, tt.in)
			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:\n%+v\nGot:\n%+v\n", tt.expected.out, out)
			}

			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected:%q\nGot:%q\n", tt.expected.err, err)
			}
		})
	}
}

func TestAuthorController_UpdateAuthor(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *author.UpdateAuthorResponse
		err error
	}

	tests := map[string]struct {
		in       *author.UpdateAuthorRequest
		expected expectation
	}{
		"Success": {
			in: &author.UpdateAuthorRequest{
				Id:   1,
				Name: "UpdatedTaran",
			},
			expected: expectation{
				out: &author.UpdateAuthorResponse{
					Author: &author.Author{
						Id:   1,
						Name: "UpdatedTaran",
					},
				},
				err: nil,
			},
		}, "Failure": {
			in: &author.UpdateAuthorRequest{
				Id:   2,
				Name: "Invalid Author",
			},
			expected: expectation{
				out: nil,
				err: errors.New("Internal Server Error"),
			},
		},
	}
	ctrl := gomock.NewController(t)
	client := mock_author.NewMockAuthorServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			client.EXPECT().UpdateAuthor(ctx, tt.in).Return(tt.expected.out, tt.expected.err)

			out, err := client.UpdateAuthor(ctx, tt.in)
			// fmt.Println(out, tt.expected.out, err, tt.expected.err)
			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:%+v\nGot:%+v\n", tt.expected.out, out)
			}

			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected: %q\nGot: %q\n", tt.expected.err, err)
			}
		})
	}
}
func TestAuthorController_DeleteAuthor(t *testing.T) {
	ctx := context.Background()

	type expectation struct {
		out *author.DeleteAuthorResponse
		err error
	}

	tests := map[string]struct {
		in       *author.DeleteAuthorRequest
		expected expectation
	}{
		"Success": {
			in: &author.DeleteAuthorRequest{
				Id: 1,
			},
			expected: expectation{
				out: &author.DeleteAuthorResponse{},
				err: nil,
			},
		},
	}
	ctrl := gomock.NewController(t)
	client := mock_author.NewMockAuthorServiceClient(ctrl)

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {

			client.EXPECT().DeleteAuthor(ctx, tt.in).Return(tt.expected.out, tt.expected.err)

			out, err := client.DeleteAuthor(ctx, tt.in)

			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected:%+vead\nGot:%+v\n", tt.expected.out, out)
			}

			if !errors.Is(err, tt.expected.err) {
				t.Errorf("Error:\nExpected: %q\nGot: %q\n", tt.expected.err.Error(), err.Error())
			}

		})

	}
}
