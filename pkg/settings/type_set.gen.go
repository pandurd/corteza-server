package settings

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// pkg/settings/types.yaml

// 	Hello! This file is auto-generated.

type (

	// ValueSet slice of Value
	//
	// This type is auto-generated.
	ValueSet []*Value
)

// Walk iterates through every slice item and calls w(Value) err
//
// This function is auto-generated.
func (set ValueSet) Walk(w func(*Value) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Value) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set ValueSet) Filter(f func(*Value) (bool, error)) (out ValueSet, err error) {
	var ok bool
	out = ValueSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}
