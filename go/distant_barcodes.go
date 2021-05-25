package main

// question: https://leetcode.com/problems/distant-barcodes/
func rearrangeBarcodes(barcodes []int) []int {
	countMap := make(map[int]int)
	result := make([]int, len(barcodes))
	maxCount := 0
	maxBarcode := 0

	for _, barcode := range barcodes {
		if _, contained := countMap[barcode]; contained {
			countMap[barcode]++
		} else {
			countMap[barcode] = 1
		}

		if maxCount < countMap[barcode] {
			maxCount = countMap[barcode]
			maxBarcode = barcode
		}
	}

	position := 0

	for countMap[maxBarcode] > 0 {
		result[position] = maxBarcode
		countMap[maxBarcode]--
		position += 2
	}

	for k, v := range countMap {
		for v > 0 {
			if position >= len(result) {
				position = 1
			}

			result[position] = k
			position += 2
			v--
		}
	}

	return result
}
