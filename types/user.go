package types

type User struct {
	ID        string       `dynamodbav:"id" json:"id"`
	FirstName string       `dynamodbav:"firstName" json:"firstName"`
	LastName  string       `dynamodbav:"lastName" json:"lastName"`
	Email     string       `dynamodbav:"email" json:"email"`
	Address   AddressModel `dynamodbav:"address" json:"address"`
}

type AddressModel struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"postCode"`
	Country  string `json:"country"`
}
