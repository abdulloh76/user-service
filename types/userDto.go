package types

type UserBody struct {
	Name  string `dynamodbav:"name" json:"name"`
	Email string `dynamodbav:"email" json:"email"`
}
