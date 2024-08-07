package model

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

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

func SynthesizeNFT(db *gorm.DB, nftIds []int64, newNFTType string, newNFTUserId int64) (*Nft, error) {
	// 查询数据库中已有的NFT
	var existingNfts []Nft
	if err := db.Where("id IN (?)", nftIds).Find(&existingNfts).Error; err != nil {
		return nil, err
	}

	// 提取所有NFT的Name字段到一个新的字符串切片
	var names []string
	for _, nft := range existingNfts {
		names = append(names, nft.Name)
	}

	// 合成新NFT的Name，将已有NFT的Name字段用"、"连接起来
	newNFTName := fmt.Sprintf("%s", strings.Join(names, "、"))

	// 创建新NFT
	newNFT := Nft{
		Name:   newNFTName,
		Type:   newNFTType,
		Nid:    0, // 合成NFT的Nid，这里假设为0，您可以根据业务逻辑来设置
		UserId: newNFTUserId,
	}

	// 保存新NFT到数据库
	if result := db.Create(&newNFT); result.Error != nil {
		return nil, result.Error
	}

	return &newNFT, nil
}
