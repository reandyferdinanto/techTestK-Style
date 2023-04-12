package request

type MemberCreateRequest struct {
	ID        int8   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string `gorm:"type:varchar(50)" json:"username"`
	Gender    string `gorm:"type:enum('Male', 'Female')" json:"gender"`
	SkinType  string `gorm:"type:enum('Oily', 'Dry', 'Combination', 'Sensitive', 'Normal')" json:"skintype"`
	SkinColor string `gorm:"type:enum('Fair', 'Medium', 'Olive', 'Brown', 'Dark')" json:"skincolor"`
}

type MemberUpdateRequest struct {
	Username  string `gorm:"type:varchar(50)" json:"username"`
	Gender    string `gorm:"type:enum('Male', 'Female')" json:"gender"`
	SkinType  string `gorm:"type:enum('Oily', 'Dry', 'Combination', 'Sensitive', 'Normal')" json:"skintype"`
	SkinColor string `gorm:"type:enum('Fair', 'Medium', 'Olive', 'Brown', 'Dark')" json:"skincolor"`
}