package windicators

// ComponentList ..
type ComponentList []*Component

// FindByID ..
func (complist ComponentList) FindByID(id string) *Component {
	for _, c := range complist {
		if id == c.ID {
			return c
		}
	}
	return nil
}

// IsAllHidden is something changed
func (complist ComponentList) IsAllHidden() bool {
	for _, c := range complist {
		if c.IsVisible {
			return false
		}
	}
	return true
}

// FilterVisible - filter only visibles
func (complist ComponentList) FilterVisible() ComponentList {
	var list ComponentList
	for _, c := range complist {
		if c.IsVisible {
			list = append(list, c)
		}
	}
	return list
}
