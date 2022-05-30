package domain

type CustomerRepoStub struct {
	customers []Customer
}

func (stub CustomerRepoStub) FindAll() ([]Customer, error) {
	return stub.customers, nil
}

func NewCustomerRepoStub() CustomerRepoStub {
	customers := []Customer{
		{"1", "Izaak", "Vancouver", "V6B5X6", "1993-08-28", "1"},
		{"2", "Mahtab", "Vancouver", "V6B5X6", "1996-07-07", "1"},
	}
	return CustomerRepoStub{customers}
}
