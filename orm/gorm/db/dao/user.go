package dao

import (
	"context"
	"golang-awesome/orm/gorm/db"
	"golang-awesome/orm/gorm/db/model"
	"gorm.io/gorm"
	"log"
)

func CreateWallet(ctx context.Context, user *model.User) error {

	db.GetGormDB().Transaction(func(tx *gorm.DB) error {
		userResult := tx.Model(model.User{}).Create(user)
		if userResult.Error != nil {
			log.Fatalf("create user err:%+v", userResult.Error)
			return userResult.Error
		}

		balance := &model.Balance{
			UserId:          user.ID,
			TotalAmount:     0,
			AvailableAmount: 0,
			FrozenAmount:    0,
		}
		balanceResult := tx.Model(model.Balance{}).Create(balance)
		if balanceResult.Error != nil {
			log.Fatalf("create Balance err:%+v", balanceResult.Error)
			return userResult.Error
		}

		balanceChange := &model.BalanceChangeRecord{
			UserId:     user.ID,
			ChangeType: "init_user",
		}
		balanceChangeResult := tx.Model(&model.BalanceChangeRecord{}).Create(balanceChange)
		if balanceChangeResult.Error != nil {
			log.Fatalf("create balanceChange err:%+v", balanceChangeResult.Error)
			return balanceChangeResult.Error
		}
		return nil
	})
	return nil
}

func Recharge(ctx context.Context, userId uint32, available, frozen uint64) error {
	db.GetGormDB().Transaction(func(tx *gorm.DB) error {
		var balance model.Balance
		err := tx.Model(model.Balance{}).Where("user_id=?", userId).First(&balance).Error
		if err != nil {
			log.Fatalf("Balance Recharge err:%+v", err)
			return err
		}
		balanceChange := &model.BalanceChangeRecord{
			UserId:                userId,
			ChangeType:            "recharge",
			BeforeAvailableAmount: balance.AvailableAmount,
			BeforeFrozenAmount:    balance.FrozenAmount,
			BeforeTotalAmount:     balance.TotalAmount,
			AfterAvailableAmount:  balance.AvailableAmount + available,
			AfterFrozenAmount:     balance.FrozenAmount + frozen,
			AfterTotalAmount:      balance.TotalAmount + available + frozen,
		}
		err = tx.Model(model.BalanceChangeRecord{}).Create(balanceChange).Error
		if err != nil {
			log.Fatalf("Recharge BalanceChangeRecord err:%+v", err)
			return err
		}
		err = tx.Model(model.Balance{}).Where("user_id=?", userId).Updates(&model.Balance{
			AvailableAmount: balance.AvailableAmount + available,
			FrozenAmount:    balance.FrozenAmount + frozen,
			TotalAmount:     balance.TotalAmount + available + frozen,
		}).Error

		if err != nil {
			log.Fatalf("Recharge updateBalance err:%+v", err)
			return err
		}
		return nil
	})
	return nil
}

func GetUserById(ctx context.Context, userId uint32) (*model.User, error) {
	var userInfo model.User
	err := db.GetGormDB().Model(model.User{}).Where("id=?", userId).First(&userInfo).Error
	if err != nil {
		return &userInfo, err
	}
	return &userInfo, nil
}
