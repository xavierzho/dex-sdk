package dexsdk

import (
	"encoding/json"
	"sort"
	"strings"
)

// Tokens list of unique token
type Tokens struct {
	List []Token
	set  map[string]struct{}
}

// NewTokenList Tokens Constructor
func NewTokenList(list ...Token) Tokens {
	var tokens Tokens
	for _, token := range list {
		tokens.add(token)
	}
	return tokens
}
func (list Tokens) add(token Token) {
	if !list.IsExist(token) {
		var lower, upper = toUpperAndLower(token.Address)
		list.set[lower] = struct{}{}
		list.set[upper] = struct{}{}
		list.List = append(list.List, token)
	}
}

// IsExist check token is exist
func (list Tokens) IsExist(token Token) bool {
	var lower, upper = toUpperAndLower(token.Address)
	if _, ok := list.set[lower]; ok {
		return true
	}
	if _, ok := list.set[upper]; ok {
		return true
	}
	return false
}

// toUpperAndLower address to lower and upper
func toUpperAndLower(s string) (string, string) {
	var lower = strings.ToLower(s)
	var upper = strings.ToUpper(s)
	return lower, upper
}

// AppendAndSort Append new token to list
func (list Tokens) AppendAndSort(tokens ...Token) {
	for _, token := range tokens {
		list.add(token)
	}
	list.Sort()
}

// Sort ...
func (list Tokens) Sort() {
	sort.Slice(list.List, func(i, j int) bool {
		return list.List[i].Address > list.List[j].Address
	})
}

// Unmarshal ...
func (list *Tokens) Unmarshal(bs []byte) error {
	if len(list.List) == 0 {
		return json.Unmarshal(bs, &list.List)
	}
	var newList []Token
	err := json.Unmarshal(bs, &newList)
	if err != nil {
		return err
	}
	list.List = append(list.List, newList...)
	return nil
}

// Marshal ...
func (list *Tokens) Marshal() ([]byte, error) {
	return json.Marshal(list.List)
}
