package points

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)

		assert.Equal(t, expected, actual)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.WithDraw(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(0)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}

		err := wallet.WithDraw(Bitcoin(100))

		assert.Error(t, err)
		assert.Equal(t, startingBalance, wallet.Balance())
	})
}
