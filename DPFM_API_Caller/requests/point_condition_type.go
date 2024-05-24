package requests

type PointConditionType struct {
	PointConditionType	string	`json:"PointConditionType"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
