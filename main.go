package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	sicCodes   []string
	cData      []ContributedData
	dataOwners []DataOwner
	customer
)

type DataOwner struct {
	Id   int
	Name string
}

type Customer struct {
	Id   int
	Sic  string
	Name string
}
type Permission struct {
	Sic     string
	Company int
	Field   string
	Allowed bool
}

type ContributedData struct {
	DataOwner               int
	HouseholdIncomeMonthly  int
	IndividualIncomeMonthly int
	SelfReported            bool
}

func init() {
	for i := 0; i < 1000; i++ {
		sicCodes = append(sicCodes, strconv.Itoa(i))
	}

	for j := 0; j < 100000; j++ {
		cData = append(cData, makecData())
	}
	dataOwners = append(dataOwners, DataOwner{1, "Sprint"})
	dataOwners = append(dataOwners, DataOwner{2, "ATT"})

}

func main() {
	fmt.Println(sicCodes)
	fmt.Println(cData)
}

func makecData() ContributedData {
	cd := ContributedData{
		DataOwner:               1,
		HouseholdIncomeMonthly:  time.Now().Nanosecond(),
		IndividualIncomeMonthly: time.Now().Nanosecond(),
		SelfReported:            true,
	}
	return cd
}
