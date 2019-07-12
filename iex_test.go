package iex

import "testing"

const (
	publishable = "REPLACE_WITH_PUBLISHABLE_TOKEN"
	secret      = "REPLACE_WITH_SECRET_TOKEN"
)

func TestNewClient(t *testing.T) {
	got := NewClient(publishable, secret, nil)
	if got != nil {
		t.Logf("got:\t%v", got)
	}
}

func TestAdvancedStats(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.AdvancedStats("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestBalanceSheet(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.BalanceSheet("aapl", nil)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}
