package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
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
	bins []Bin
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{
		bins,
	}
}
