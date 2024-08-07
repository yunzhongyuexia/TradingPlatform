package model

type Nft struct {
	Id     int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	Name   string `gorm:"column:name;type:varchar(255);comment:nft_name" json:"name"`
	UserId int64  `gorm:"column:user_id;type:bigint(20)" json:"user_id"`
	Type   string `gorm:"column:type;type:varchar(255);comment:类型" json:"type"`
	Nid    int64  `gorm:"column:nid;type:bigint(20);comment:nftid" json:"nid"`
}

func (m *Nft) TableName() string {
	return "nft"
}
