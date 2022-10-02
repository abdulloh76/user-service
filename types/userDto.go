package types

type UserBody struct {
	FirstName string       `dynamodbav:"firstName" json:"firstName"`
	LastName  string       `dynamodbav:"lastName" json:"lastName"`
	Email     string       `dynamodbav:"email" json:"email"`
	Address   AddressModel `dynamodbav:"address" json:"address"`
}
