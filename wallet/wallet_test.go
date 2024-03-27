package wallet_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/pkg/testkit"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet/mock_handler"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/suite"
)

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}

type WalletTestSuite struct {
	suite.Suite
	mockStore *mock_handler.MockStorer
	handler   *wallet.Handler
}

func (s *WalletTestSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.mockStore = mock_handler.NewMockStorer(ctrl)
	s.handler = wallet.New(s.mockStore)
}

func (s *WalletTestSuite) TestGetAllWallets() {
	wallets := newWallets()

	s.Run("given unable to get wallets should return 500 and error message", func() {
		s.mockStore.EXPECT().
			Wallets(wallet.Filter{}).
			Return(nil, errors.New("error getting wallets"))

		rec := testkit.DoEchoRequest(
			s.handler.GetAllWallets,
			httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil),
		)

		s.Equal(http.StatusInternalServerError, rec.Code)
		s.Contains(rec.Body.String(), "error getting wallets")
	})

	s.Run("given user able to getting wallet should return list of wallets", func() {
		s.mockStore.EXPECT().
			Wallets(wallet.Filter{}).
			Return(wallets, nil)

		rec := testkit.DoEchoRequest(
			s.handler.GetAllWallets,
			httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil),
		)

		s.Equal(http.StatusOK, rec.Code)
		s.Equal(fmt.Sprintln(testkit.JSONStringify(s.T(), wallets)), rec.Body.String())
	})

	s.Run("given wallet type filter should return list of wallets with filter", func() {
		s.mockStore.EXPECT().
			Wallets(wallet.Filter{WalletType: "Savings"}).
			Return(wallets[:1], nil)

		rec := testkit.DoEchoRequest(
			s.handler.GetAllWallets,
			httptest.NewRequest(http.MethodGet, "/api/v1/wallets?wallet_type=Savings", nil),
		)

		s.Equal(http.StatusOK, rec.Code)
		s.Equal(fmt.Sprintln(testkit.JSONStringify(s.T(), wallets[:1])), rec.Body.String())
	})
}

func newWallets() []wallet.Wallet {
	return []wallet.Wallet{
		{
			ID:         1,
			UserID:     1,
			UserName:   "Thirathawat",
			WalletName: "Maz Savings",
			WalletType: "Savings",
			Balance:    100,
			CreatedAt:  time.Date(2024, 3, 25, 14, 19, 0, 729237000, time.UTC),
		},
		{
			ID:         2,
			UserID:     1,
			UserName:   "Thirathawat",
			WalletName: "Maz Credit Card",
			WalletType: "Credit Card",
			Balance:    500,
			CreatedAt:  time.Date(2024, 3, 25, 14, 19, 0, 729237000, time.UTC),
		},
	}
}
