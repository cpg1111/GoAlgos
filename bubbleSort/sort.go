package bubbleSort

var swapped = false

// Run runs a bubble sort
func Run(arr []int, i int) []int {
	if i == len(arr) && !swapped {
		return arr
	}
	if i == len(arr) && swapped {
		swapped = false
		return Run(arr, 0)
	}
	if i+1 < len(arr) && arr[i] > arr[i+1] {
		temp := arr[i]
		arr[i] = arr[i+1]
		arr[i+1] = temp
		swapped = true
		return Run(arr, i+1)
	}
	return Run(arr, i+1)
}
