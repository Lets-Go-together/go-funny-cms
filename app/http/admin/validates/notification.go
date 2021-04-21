package validates

type NotificationSaveValidate struct {
	Title        string `validate:"required" json:"title"`
	Submitter_id int    `json:"submitter_id"`
	Description  string `validate:"required" json:"description"`
	FollowIds    []int  `validate:"required" json:"fllow_ds"`
}
