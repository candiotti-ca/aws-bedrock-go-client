package awsbedrockgoclient

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type BedrockMock struct {
	invokeModel                   func(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error)
	invokeModelWithResponseStream func(ctx context.Context, params *bedrockruntime.InvokeModelWithResponseStreamInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error)
}

func (b BedrockMock) InvokeModel(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error) {
	return b.invokeModel(ctx, params)
}
func (b BedrockMock) InvokeModelWithResponseStream(ctx context.Context, params *bedrockruntime.InvokeModelWithResponseStreamInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {
	return b.invokeModelWithResponseStream(ctx, params)
}

func TestClientQuery(t *testing.T) {
	t.Parallel()

	model := cohereCommandV14

	var inputBody []byte
	bedrockMock := BedrockMock{
		invokeModel: func(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error) {
			inputBody = params.Body
			return &bedrockruntime.InvokeModelOutput{
				Body: loadJSONFile(t, "TestClientQuery.output.json"),
			}, nil
		},
	}
	client := New[CohereCommandInput, CohereCommandOutput](bedrockMock, model)

	response, err := client.Query(context.Background(), CohereCommandInput{
		Prompt: "Does this work?",
	})

	require.NoError(t, err)
	expectedInput := string(loadJSONFile(t, "TestClientQuery.input.json"))
	assert.JSONEq(t, expectedInput, string(inputBody))
	assert.Len(t, response.Generations, 1)
	assert.Equal(t, "response", response.Generations[0].Text)
}

func TestClientQuery_invalidInput(t *testing.T) {
	t.Parallel()

	model := cohereCommandV14
	bedrockMock := BedrockMock{}
	client := New[CohereCommandInput, CohereCommandOutput](bedrockMock, model)

	response, err := client.Query(context.Background(), CohereCommandInput{})

	require.ErrorContains(t, err, "invalid input")
	assert.Nil(t, response)
}

func TestClientQuery_invokeModelError(t *testing.T) {
	t.Parallel()

	model := cohereCommandV14
	bedrockMock := BedrockMock{
		invokeModel: func(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error) {
			return nil, errors.New("test fail")
		},
	}
	client := New[CohereCommandInput, CohereCommandOutput](bedrockMock, model)

	response, err := client.Query(context.Background(), CohereCommandInput{Prompt: "test"})

	require.ErrorContains(t, err, "invoke model request has failed")
	assert.Nil(t, response)
}

func TestClientQuery_invalidResponse(t *testing.T) {
	t.Parallel()

	model := cohereCommandV14
	bedrockMock := BedrockMock{
		invokeModel: func(ctx context.Context, params *bedrockruntime.InvokeModelInput, optFns ...func(*bedrockruntime.Options)) (*bedrockruntime.InvokeModelOutput, error) {
			return &bedrockruntime.InvokeModelOutput{Body: []byte("not a json string")}, nil
		},
	}
	client := New[CohereCommandInput, CohereCommandOutput](bedrockMock, model)

	response, err := client.Query(context.Background(), CohereCommandInput{Prompt: "test"})

	require.ErrorContains(t, err, "cannot unmarshal response")
	assert.Nil(t, response)
}

func loadJSONFile(t *testing.T, fileName string) []byte {
	file, err := os.ReadFile("testdata/" + fileName)
	if err != nil {
		t.Fatal(err)
	}
	return file
}
