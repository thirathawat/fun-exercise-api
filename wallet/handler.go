package wallet

//go:generate go run github.com/golang/mock/mockgen -source=./handler.go -destination=./mock_handler/mock_handler.go -package=mock_handler

import (
	"net/http"

	"github.com/KKGo-Software-engineering/fun-exercise-api/pkg/errs"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets(filter Filter) ([]Wallet, error)
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

// GetAllWallets
//
//		@Summary		Get all wallets
//		@Description	Get all wallets
//		@Tags			wallet
//	 	@Param			wallet_type	query	string	false	"wallet type" Enums(Savings, Credit Card, Crypto Wallet)
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	Wallet
//		@Router			/api/v1/wallets [get]
//		@Failure		500	{object}	errs.Err
func (h *Handler) GetAllWallets(c echo.Context) error {
	wallets, err := h.store.Wallets(Filter{
		WalletType: c.QueryParam("wallet_type"),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errs.New(err.Error()))
	}
	return c.JSON(http.StatusOK, wallets)
}

// GetUserWallets
//
//		@Summary		Get user wallets
//		@Description	Get user wallets
//		@Tags			wallet
//	 	@Param			id	path	string	true	"user id"
//	 	@Param			wallet_type	query	string	false	"wallet type" Enums(Savings, Credit Card, Crypto Wallet)
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	Wallet
//		@Router			/api/v1/users/{id}/wallets [get]
//		@Failure		500	{object}	errs.Err
func (h *Handler) GetUserWallets(c echo.Context) error {
	wallets, err := h.store.Wallets(Filter{
		UserID:     c.Param("id"),
		WalletType: c.QueryParam("wallet_type"),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errs.New(err.Error()))
	}
	return c.JSON(http.StatusOK, wallets)
}
