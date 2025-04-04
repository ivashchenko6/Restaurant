package table

type Table struct {
	TableID        int
	PeopleQuantity int
	ReceiptAmount  float64
	IsAvailable    bool
	Orders         map[string]int
}

func (table *Table) PrintReceipt() {

}
