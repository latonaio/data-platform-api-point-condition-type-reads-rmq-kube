package requests

type Text struct {
	PointConditionType     string  `json:"PointConditionType"`
	Language          		string  `json:"Language"`
	PointConditionTypeName	string  `json:"PointConditionTypeName"`
	CreationDate			string	`json:"CreationDate"`
	LastChangeDate			string	`json:"LastChangeDate"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}
