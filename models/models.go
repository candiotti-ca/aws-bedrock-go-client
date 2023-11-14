package models

type Model string

const (
	Cohere_CommandV14     Model = "cohere.command-text-v14"
	Cohere_EmbedEnglishV3 Model = "cohere.embed-english-v3"
	Cohere_EmbedMultiV3   Model = "cohere.embed-multilingual-v3"

	Amazon_TitanTextLarge        Model = "amazon.titan-tg1-large"
	Amazon_TitanTextEmbeddingsV1 Model = "amazon.titan-embed-text-v1"
	Amazon_TitanTextEmbeddingsV2 Model = "amazon.titan-embed-g1-text-02"
	Amazon_TitanTextLiteV1       Model = "amazon.titan-text-lite-v1"
	Amazon_TitanTextExpressV1    Model = "amazon.titan-text-express-v1"

	StabilityAI_StableDiffusionXLV0 Model = "stability.stable-diffusion-xl-v0"
	StabilityAI_StableDiffusionXLV1 Model = "stability.stable-diffusion-xl-v1"

	AI21_J2GrandeInstruct Model = "ai21.j2-grande-instruct"
	AI21_J2JumboInstruct  Model = "ai21.j2-jumbo-instruct"
	AI21_Jurassic2Mid     Model = "ai21.j2-mid"
	AI21_Jurassic2Midv1   Model = "ai21.j2-mid-v1"
	AI21_Jurassic2Ultra   Model = "ai21.j2-ultra"
	AI21_Jurassic2UltraV1 Model = "ai21.j2-ultra-v1"

	Anthropic_ClaudeInstantV1 Model = "anthropic.claude-instant-v1"
	Anthropic_ClaudeV1        Model = "anthropic.claude-v1"
	Anthropic_ClaudeV2        Model = "anthropic.claude-v2"

	Meta_Llama2ChatV1 Model = "meta.llama2-13b-chat-v1"
)
