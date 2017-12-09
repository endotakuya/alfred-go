package alfred

// New item
func NewItem() *Item {
	return &Item{
		Valid: true,
		Type:  TYPE_DEFAULT,
	}
}

// Get Items
func (a *Alfred) GetItems() []Item {
	return a.Items
}
