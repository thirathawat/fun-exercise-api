package postgres

import (
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/pkg/sqlkit"
	"github.com/KKGo-Software-engineering/fun-exercise-api/pkg/timekit"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

const table = "user_wallet"

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) Create(wallet *wallet.Wallet) error {
	var id int

	wallet.CreatedAt = timekit.Now()
	query, args := sqlkit.NewQueryBuilder().
		Insert().
		Table(table).
		Set("user_id", wallet.UserID).
		Set("user_name", wallet.UserName).
		Set("wallet_name", wallet.WalletName).
		Set("wallet_type", wallet.WalletType).
		Set("balance", wallet.Balance).
		Set("created_at", wallet.CreatedAt).
		Returning("id").
		Build()

	if err := p.Db.QueryRow(query, args...).Scan(&id); err != nil {
		return err
	}

	wallet.ID = id

	return nil
}

func (p *Postgres) Wallets(filter wallet.Filter) ([]wallet.Wallet, error) {
	query, args := queryWallet(filter)
	rows, err := p.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func queryWallet(filter wallet.Filter) (query string, args []any) {
	b := sqlkit.NewQueryBuilder().Select("*").From(table)
	if filter.WalletType != "" {
		b = b.Where("wallet_type", "=", filter.WalletType)
	}

	return b.Build()
}
