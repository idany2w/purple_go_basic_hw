package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		id,
		private,
		time.Now(),
		name,
	}
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{
		bins,
	}
}
