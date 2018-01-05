package models

func remove(elements []ConsumableCard, elem ConsumableCard) []ConsumableCard {
	index := 0
	for key, element := range elements {
		if element == elem {
			index = key
			break
		}
	}
	elements[len(elements)-1], elements[index] = elements[index], elements[len(elements)-1]
	elements = elements[:len(elements)-1]
	return elements
}
