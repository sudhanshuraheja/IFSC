package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBankIndexKeywords(t *testing.T) {
	frequency := addBankIndexKeywords("Abhyudaya Bank Bldg., B.No.71, Nehru Nagar, Kurla (E), Mumbai-400024", 2)
	assert.Equal(t, frequency, map[string]int{"abhyudaya": 2, "bldg": 2, "b": 2, "nehru": 2, "kurla": 2, "e": 2, "mumbai": 2, "400024": 2, "bank": 2, "no": 2, "71": 2, "nagar": 2})

	frequency = addBankIndexKeywords("Abhyudaya Cooperative Bank Limited", 5)
	assert.Equal(t, frequency, map[string]int{"abhyudaya": 5, "cooper": 5, "bank": 5, "limit": 5})
}

// func Test_buildIndex(t *testing.T) {

// 	inputSubitems := []SubItem{
// 		SubItem{Key: "bank", Value: , Weight: 5},
// 		SubItem{Key: "ifsc", Value: "Abhy0065001", Weight: 5},
// 		SubItem{Key: "micr", Value: "400065001", Weight: 2},
// 		SubItem{Key: "branch", Value: "Rtgs-Ho", Weight: 2},
// 		SubItem{Key: "address", Value: , Weight: 1},
// 		SubItem{Key: "city", Value: "Mumbai", Weight: 2},
// 		SubItem{Key: "district", Value: "Greater Mumbai", Weight: 2},
// 		SubItem{Key: "state", Value: "Maharashtra", Weight: 2},
// 		SubItem{Key: "contact", Value: "25260173", Weight: 2},
// 	}
// 	input := Item{ID: 1, SubItems: inputSubitems}
// 	input.AddIndex()

// 	assert.Equal(t, input.Index, map[string]int{"mumbai": 5, "abhyudaya": 6, "nagar": 1, "kurla": 1, "cooper": 5, "rtg": 2, "abhy0065001": 5, "400024": 1, "b": 1, "bldg": 1, "25260173": 2, "greater": 2, "e": 1, "limit": 5, "400065001": 2, "bank": 6, "maharashtra": 2, "71": 1, "no": 1, "nehru": 1, "ho": 2})

// }

// func Test_globalIndex(t *testing.T) {
// 	config.Load()
// 	logger.Setup()
// 	// db.Init()
// 	// defer db.Close()

// 	// var inx WordIndex

// 	input1 := Item{ID: 5, SubItems: []SubItem{
// 		SubItem{Key: "test1", Value: "one two three", Weight: 1},
// 	}}
// 	input1.AddIndex()

// 	input2 := Item{ID: 6, SubItems: []SubItem{
// 		SubItem{Key: "test2", Value: "two three four", Weight: 2},
// 	}}
// 	input2.AddIndex()

// 	input3 := Item{ID: 7, SubItems: []SubItem{
// 		SubItem{Key: "test3", Value: "three four five", Weight: 3},
// 	}}
// 	input3.AddIndex()

// 	// inx.AddLookup(input1)
// 	// inx.AddLookup(input2)
// 	// inx.AddLookup(input3)

// 	// ids, err := inx.Find("three")
// 	// assert.Equal(t, nil, err)
// 	// assert.Equal(t, []model.Branch{model.Branch{DBId: 5, Bank: "Abhyudaya Cooperative Bank Limited", Ifsc: "Abhy0065005", Micr: "400065005", Branch: "Darukhana", Address: "Potia Ind.Estate, Reay Road (E), Darukhana, Mumbai-400010", City: "Mumbai", District: "Greater Mumbai", State: "Maharashtra", Contact: "23778164", CreatedAt: "2018-01-01T16:51:08.449753Z", UpdatedAt: "2018-01-01T16:51:08.449753Z"}, model.Branch{DBId: 6, Bank: "Abhyudaya Cooperative Bank Limited", Ifsc: "Abhy0065006", Micr: "400065006", Branch: "Fort", Address: "Abhyudaya Bank Bldg., 251, Perin Nariman Street, Fort, Mumbai-400001", City: "Mumbai", District: "Greater Mumbai", State: "Maharashtra", Contact: "22614468", CreatedAt: "2018-01-01T16:51:08.451523Z", UpdatedAt: "2018-01-01T16:51:08.451523Z"}, model.Branch{DBId: 7, Bank: "Abhyudaya Cooperative Bank Limited", Ifsc: "Abhy0065007", Micr: "400065007", Branch: "Ghatkopar", Address: "Unit No 2 & 3, Silver Harmony Bldg,New Maniklal Estate, Ghatkopar (West), Mumbai-400086", City: "Mumbai", District: "Greater Mumbai", State: "Maharashtra", Contact: "25116673", CreatedAt: "2018-01-01T16:51:08.453326Z", UpdatedAt: "2018-01-01T16:51:08.453326Z"}}, ids)

// 	// ids, err = inx.Find("two")
// 	// assert.Equal(t, nil, err)
// 	// assert.Equal(t, []model.Branch{model.Branch{DBId: 5, Bank: "Abhyudaya Cooperative Bank Limited", Ifsc: "Abhy0065005", Micr: "400065005", Branch: "Darukhana", Address: "Potia Ind.Estate, Reay Road (E), Darukhana, Mumbai-400010", City: "Mumbai", District: "Greater Mumbai", State: "Maharashtra", Contact: "23778164", CreatedAt: "2018-01-01T16:51:08.449753Z", UpdatedAt: "2018-01-01T16:51:08.449753Z"}, model.Branch{DBId: 6, Bank: "Abhyudaya Cooperative Bank Limited", Ifsc: "Abhy0065006", Micr: "400065006", Branch: "Fort", Address: "Abhyudaya Bank Bldg., 251, Perin Nariman Street, Fort, Mumbai-400001", City: "Mumbai", District: "Greater Mumbai", State: "Maharashtra", Contact: "22614468", CreatedAt: "2018-01-01T16:51:08.451523Z", UpdatedAt: "2018-01-01T16:51:08.451523Z"}}, ids)

// 	// ids, err = inx.Find("six")
// 	// assert.Equal(t, []model.Branch{}, ids)
// }
