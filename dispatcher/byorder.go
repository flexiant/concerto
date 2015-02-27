package dispatcher

// ByOrder implements sort.Interface for []ScriptCharacterization based on the Order field
type ByOrder []ScriptCharacterization

func (a ByOrder) Len() int {
	return len(a)
}
func (a ByOrder) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByOrder) Less(i, j int) bool {
	return a[i].Order < a[j].Order
}
