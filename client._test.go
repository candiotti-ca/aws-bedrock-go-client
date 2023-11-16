package awsbedrockgoclient

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/stretchr/testify/assert"
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
				Body: testDataJSON(t, "TestClientQuery.output.json"),
			}, nil
		},
	}
	client := New[CohereCommandInput, CohereCommandOutput](bedrockMock, model)

	response, err := client.Query(context.Background(), CohereCommandInput{
		Prompt:         "Does this work?",
		NumGenerations: 1,
	})

	assert.NoError(t, err)
	expectedInput := string(testDataJSON(t, "TestClientQuery.input.json"))
	assert.JSONEq(t, expectedInput, string(inputBody))
	assert.Len(t, response.Generations, 1)
	assert.Equal(t, "response", response.Generations[0].Text)
}

func testDataJSON(t *testing.T, fileName string) []byte {
	file, err := os.ReadFile("testdata/" + fileName)
	if err != nil {
		t.Fatal(err)
	}
	return file
}
