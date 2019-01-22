package arrays

func MaxPairProduct(arr []int) int {
	maxIndex1 := 0
	maxIndex2 := 0
	for i, _ := range arr {
		if arr[i] > arr[maxIndex1] {
			maxIndex1 = i
		} else if arr[i] > arr[maxIndex2] && arr[i] < arr[maxIndex1] {
			maxIndex2 = i
		}
	}
	return arr[maxIndex1] * arr[maxIndex2]
}