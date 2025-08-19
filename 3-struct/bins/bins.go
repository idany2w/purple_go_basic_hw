package bins

import (
	"fmt"
	"strconv"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func (bin *Bin) OutputBin() {
	fmt.Println("Id: " + bin.Id)
	fmt.Println("Private: " + strconv.FormatBool(bin.Private))
	fmt.Println("CreatedAt: " + bin.CreatedAt.String())
	fmt.Println("Name: " + bin.Name)
}

func NewList() *BinList {
	return &BinList{
		Bins: []Bin{},
	}
}

func (binList *BinList) AddToList(bin *Bin) {
	binList.Bins = append(binList.Bins, *bin)
}

func (binList *BinList) OutputList() {
	for _, bin := range binList.Bins {
		bin.OutputBin()
	}
}
