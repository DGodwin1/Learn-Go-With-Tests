package pointers_and_errors

import (
	"testing"
)

func TestBalance(t *testing.T) {
	wallet := Wallet{10}

	want := Bitcoin(10)

	AssertBalance(t, wallet, want)

}

func TestWithdraw(t *testing.T) {
	t.Run("Withdraw from an account with enough cash", func(t *testing.T) {
		wallet := Wallet{100}

		err := wallet.Withdraw(Bitcoin(10))

		BalanceWanted := Bitcoin(90)

		AssertBalance(t, wallet, BalanceWanted)

		DidntGetErrorCheck(t, err)

	})
	t.Run("Attempt to withdraw more than you have", func(t *testing.T) {
		wallet := Wallet{20}
		err := wallet.Withdraw(Bitcoin(30))

		// The transaction should not take place. This means the wallet's
		// balance should be the same as it was when it was first setup.

		AssertBalance(t, wallet, 20)

		GotErrorChecker(t, err, "You've tried to take out more money than you have")
	})
}
func AssertBalance(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func GotErrorChecker(t *testing.T, got error, message string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error when I wanted one")
	}
	if got.Error() != message {
		// Using %q because then you get quotes which makes it easier
		// to distinguish between the print out and the strings you're
		// supposed to be looking at.
		t.Errorf("Got %q but expected %q", got, message)
	}
}

func DidntGetErrorCheck(t *testing.T, got error) {
	t.Helper()
	// Error should be nil. Ie, everything should be fine.
	if got != nil {
		t.Fatalf("Got the error %q when I shouldn't have done", got.Error())
	}
}
