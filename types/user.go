package types

type User struct {
	ID    string `dynamodbav:"id" json:"id"`
	Name  string `dynamodbav:"name" json:"name"`
	Email string `dynamodbav:"email" json:"email"`
}
