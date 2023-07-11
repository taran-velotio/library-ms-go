// Checking testing function - NOT CORRECT
package tests

import (
	"context"
	"errors"
	"fmt"
	"library-comp/proto/author/author"
	"library-comp/repository"
	"reflect"
	"testing"
)

func TestAuthorRepository_GetAuthor(t *testing.T) {
	ctx := context.Background()

	repo := repository.NewAuthorRepository()

	type expectation struct {
		out *author.GetAuthorResonse
		err error
	}

	tests := map[string]struct {
		in       *author.GetAuthorRequest
		expected expectation
	}{
		"Must_Success": {
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
		"Not_Found_Id": {
			in: &author.GetAuthorRequest{
				Id: 4,
			},
			expected: expectation{
				out: &author.GetAuthorResonse{},
				err: errors.New("record not found"),
			},
		},
	}

	for scenario, tt := range tests {

		t.Run(scenario, func(t *testing.T) {
			fmt.Println("coming here")
			out, err := repo.GetAuthor(ctx, tt.in)
			fmt.Println("Going after = ", out)

			if !reflect.DeepEqual(out, tt.expected.out) {
				t.Errorf("Error:\nExpected: %+vead\nGot: %+v\n", tt.expected.out, out)
			}
			if !reflect.DeepEqual(err, tt.expected.err) {
				t.Errorf("Error:\nExpected: %q\nGot: %q\n", tt.expected.err, err)
			}

		})
	}
}
