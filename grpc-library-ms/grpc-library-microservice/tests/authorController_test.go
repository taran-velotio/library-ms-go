// Checking testing function - NOT CORRECT
package tests

import (
	"context"
	"fmt"
	"library-comp/controller"
	"library-comp/proto/author/author"
	"log"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (client author.AuthorServiceClient, closer func()) {

	listn := bufconn.Listen(101024 * 1024)

	s := grpc.NewServer()

	authController := controller.AuthorController{}
	author.RegisterAuthorServiceServer(s, &authController)

	go func() {
		if err := s.Serve(listn); err != nil {
			log.Fatalf("s.Serve %v", err)
		}
	}()

	cc, err := grpc.Dial("localhost:8092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer = func() {
		_ = listn.Close()
		s.Stop()
	}

	client = author.NewAuthorServiceClient(cc)

	return
}

func TestGetAuthor(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

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
						Name: "Book3",
					},
				}, err: nil,
			},
		},
	}

	for sc, test := range tests {
		fmt.Println(sc)
		fmt.Println(test.expected.out.Author.Name)

		t.Run(sc, func(t *testing.T) {
			out, err := client.GetAuthor(ctx, test.in)
			if err != nil {
				fmt.Println("Error", err)
			}

			if out == nil || out.Author == nil {
				// Handling unexpected response here
				t.Error("Invalid response")
				return
			}

			fmt.Println(out.Author.Id)

		})
	}

}

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/your/package/author"            // Import your package here
// 	mock_author "github.com/your/package/mocks" // Import the generated mock package here
// )

// func TestGetAuthor(t *testing.T) {
// 	ctx := context.Background()

// 	mockRepo := &mock_author.MockAuthorRepository{} // Create a mock repository instance using Mockery

// 	controller := author.NewAuthorController(mockRepo) // Create an instance of the AuthorController with the mock repository

// 	req := &author.GetAuthorRequest{
// 		Id: 1,
// 	}

// 	expectedResponse := &author.GetAuthorResonse{
// 		Author: &author.Author{
// 			Id:   1,
// 			Name: "John Doe",
// 		},
// 	}

// 	Set up expectations on the mocked repository methods
// 	mockRepo.On("GetAuthor", ctx, req).Return(expectedResponse, nil)

// 	response, err := controller.GetAuthor(ctx, req)

// 	assert.NoError(t, err)                      // Check if there is no error
// 	assert.Equal(t, expectedResponse, response) // Check if the response matches our expectation

// 	mockRepo.AssertExpectations(t) // Verify that all expected calls were made to the mocked repository methods
// }
