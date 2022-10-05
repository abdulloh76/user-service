package types

type User struct {
	ID        string       `dynamodbav:"id" json:"-"`
	Email     string       `dynamodbav:"email" json:"email" binding:"required"`
	Password  string       `dynamodbav:"password" json:"password" binding:"required"`
	FirstName string       `dynamodbav:"firstName" json:"firstName"`
	LastName  string       `dynamodbav:"lastName" json:"lastName"`
	Address   AddressModel `dynamodbav:"address" json:"address"`
}

type AddressModel struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"postCode"`
	Country  string `json:"country"`
}

type UserCredentials struct {
	Email    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
