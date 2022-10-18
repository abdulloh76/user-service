package types

type UserBody struct {
	FirstName string       `dynamodbav:"firstName" json:"firstName"`
	LastName  string       `dynamodbav:"lastName" json:"lastName"`
	Email     string       `dynamodbav:"email" json:"email"`
	Address   AddressModel `dynamodbav:"address" json:"address"`
}

type UpdateCredentialsDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UpdateAddressDto struct {
	Address AddressModel `json:"address"`
}

type UpdatePasswordDto struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}
