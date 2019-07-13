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

func TestCashFlow(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.CashFlow("aapl", nil)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestCompany(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Company("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestDelayedQuote(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.DelayedQuote("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestDividends(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Dividends("aapl", Range5Y)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestEarnings(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Earnings("aapl", nil)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestEarningsToday(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.EarningsToday()
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestEstimates(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.Estimates("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestFundOwnership(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.FundOwnership("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestIncomeStmt(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.IncomeStmt("aapl", nil)
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestInsiderRoster(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.InsiderRoster("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestInsiderSum(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.InsiderSum("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestInsiderTXN(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.InsiderTXN("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}

func TestInstlOwnership(t *testing.T) {
	c := NewClient(publishable, secret, nil)
	got, err := c.InstlOwnership("aapl")
	if err != nil {
		t.Errorf("error:\t%v", err.Error())
	}
	t.Logf("got:\t%v", got)
}
