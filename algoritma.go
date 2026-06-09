package main

func SequentialSearch(data Data, nData int, namaCari string) int {
	var index, i int

	index = -1

	i = 0
	for i < nData && index == -1 {
		if data[i].nama == namaCari {
			index = i
		}

		i++
	}

	return index
}

func BinarySearch(data Data, nData int, namaCari string) int {
	var index, left, right, mid int

	index = -1
	left = 0
	right = nData - 1

	for left <= right && index == -1 {
		mid = (left + right) / 2

		if data[mid].nama == namaCari {
			index = mid
		} else if data[mid].nama < namaCari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return index
}

func InsertionSortAsc(data *Data, nData int) {
	var pass, i int
	var temp Negara

	pass = 1
	for pass < nData {
		temp = data[pass]

		i = pass
		for i > 0 && temp.id < data[i-1].id {
			data[i] = data[i-1]
			i--
		}

		data[i] = temp

		pass++
	}
}

func SelectionSortDesc(data *Data, nData int) {
	var pass, index, i int
	var temp Negara

	pass = 1
	for pass < nData-1 {

		index = pass - 1

		i = pass
		for i < nData {
			if data[i].medali.emas > data[index].medali.emas {
				index = i
			} else if data[i].medali.emas == data[index].medali.emas {
				if data[i].medali.perak > data[index].medali.perak {
					index = i
				} else if data[i].medali.perak == data[index].medali.perak {
					if data[i].medali.perunggu > data[index].medali.perunggu {
						index = i
					}
				}
			}

			i++
		}

		temp = data[pass-1]
		data[pass-1] = data[index]
		data[index] = temp

		pass++
	}
}
