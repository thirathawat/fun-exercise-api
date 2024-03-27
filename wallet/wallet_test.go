package wallet_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/pkg/testkit"
	"github.com/KKGo-Software-engineering/fun-exercise-api/pkg/timekit"
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
	timekit.Freeze()
}

func (s *WalletTestSuite) TearDownSuite() {
	timekit.Unfreeze()
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

func (s *WalletTestSuite) TestGetUserWallets() {
	wallets := newWallets()

	s.Run("given unable to get wallets should return 500 and error message", func() {
		s.mockStore.EXPECT().
			Wallets(wallet.Filter{UserID: "1"}).
			Return(nil, errors.New("error getting wallets"))

		rec := testkit.DoEchoRequest(
			s.handler.GetUserWallets,
			httptest.NewRequest(http.MethodGet, "/api/v1/users/1/wallets", nil),
			testkit.WithParams(map[string]string{"id": "1"}),
		)

		s.Equal(http.StatusInternalServerError, rec.Code)
		s.Contains(rec.Body.String(), "error getting wallets")
	})

	s.Run("given user able to getting wallet should return list of wallets", func() {
		s.mockStore.EXPECT().
			Wallets(wallet.Filter{UserID: "1"}).
			Return(wallets, nil)

		rec := testkit.DoEchoRequest(
			s.handler.GetUserWallets,
			httptest.NewRequest(http.MethodGet, "/api/v1/users/1/wallets", nil),
			testkit.WithParams(map[string]string{"id": "1"}),
		)

		s.Equal(http.StatusOK, rec.Code)
		s.Equal(fmt.Sprintln(testkit.JSONStringify(s.T(), wallets)), rec.Body.String())
	})

	s.Run("given wallet type filter should return list of wallets with filter", func() {
		s.mockStore.EXPECT().
			Wallets(wallet.Filter{UserID: "1", WalletType: "Savings"}).
			Return(wallets[:1], nil)

		rec := testkit.DoEchoRequest(
			s.handler.GetUserWallets,
			httptest.NewRequest(http.MethodGet, "/api/v1/users/1/wallets?wallet_type=Savings", nil),
			testkit.WithParams(map[string]string{"id": "1"}),
		)

		s.Equal(http.StatusOK, rec.Code)
		s.Equal(fmt.Sprintln(testkit.JSONStringify(s.T(), wallets[:1])), rec.Body.String())
	})
}

func (s *WalletTestSuite) TestCreateWallet() {
	s.Run("given invalid content type should return 400 and error message", func() {
		rec := testkit.DoEchoRequest(
			s.handler.CreateWallet,
			httptest.NewRequest(http.MethodPost, "/api/v1/wallets", testkit.JSONReader(s.T(), wallet.Request{
				UserID:     1,
				WalletName: "Maz Savings",
				WalletType: "Savings",
				Balance:    100,
			})),
		)

		s.Equal(http.StatusBadRequest, rec.Code)
		s.Equal(`{"message":"code=415, message=Unsupported Media Type"}`, strings.TrimSpace(rec.Body.String()))
	})

	s.Run("given invalid request should return 400 and error message", func() {
		rec := testkit.DoEchoRequest(
			s.handler.CreateWallet,
			httptest.NewRequest(http.MethodPost, "/api/v1/wallets", testkit.JSONReader(s.T(), wallet.Request{
				UserID:     1,
				WalletName: "Maz Savings",
				WalletType: "Savings",
				Balance:    100,
			})),
			testkit.WithJSONContentType(),
		)

		s.Equal(http.StatusBadRequest, rec.Code)
		s.Equal(`{"message":"Key: 'Request.UserName' Error:Field validation for 'UserName' failed on the 'required' tag"}`, strings.TrimSpace(rec.Body.String()))
	})

	s.Run("given valid request should return 201 and created wallet", func() {
		s.mockStore.EXPECT().
			Create(&wallet.Wallet{
				UserID:     1,
				UserName:   "Thirathawat",
				WalletName: "Maz Savings",
				WalletType: "Savings",
				Balance:    100,
			}).
			DoAndReturn(func(w *wallet.Wallet) error {
				w.ID = 1
				w.CreatedAt = timekit.Now()
				return nil
			})

		rec := testkit.DoEchoRequest(
			s.handler.CreateWallet,
			httptest.NewRequest(http.MethodPost, "/api/v1/wallets", testkit.JSONReader(s.T(), wallet.Request{
				UserID:     1,
				UserName:   "Thirathawat",
				WalletName: "Maz Savings",
				WalletType: "Savings",
				Balance:    100,
			})),
			testkit.WithJSONContentType(),
		)

		s.Equal(http.StatusCreated, rec.Code)
		s.Equal(fmt.Sprintln(testkit.JSONStringify(s.T(), &wallet.Wallet{
			ID:         1,
			UserID:     1,
			UserName:   "Thirathawat",
			WalletName: "Maz Savings",
			WalletType: "Savings",
			Balance:    100,
			CreatedAt:  timekit.Now(),
		})), rec.Body.String())
	})

	s.Run("given unable to create wallet should return 500 and error message", func() {
		s.mockStore.EXPECT().
			Create(&wallet.Wallet{
				UserID:     1,
				UserName:   "Thirathawat",
				WalletName: "Maz Savings",
				WalletType: "Savings",
				Balance:    100,
			}).
			Return(errors.New("error creating wallet"))

		rec := testkit.DoEchoRequest(
			s.handler.CreateWallet,
			httptest.NewRequest(http.MethodPost, "/api/v1/wallets", testkit.JSONReader(s.T(), wallet.Request{
				UserID:     1,
				UserName:   "Thirathawat",
				WalletName: "Maz Savings",
				WalletType: "Savings",
				Balance:    100,
			})),
			testkit.WithJSONContentType(),
		)

		s.Equal(http.StatusInternalServerError, rec.Code)
		s.Contains(rec.Body.String(), "error creating wallet")
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
