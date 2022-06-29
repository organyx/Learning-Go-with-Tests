package maps

type Dictionary map[string]string

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }

func (d Dictionary) Search(word string) string {
	return d[word]
}
