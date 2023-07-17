package helper

func DuplicateInArray(arr []uint) bool {
	visited := make(map[uint]string, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == "visited" {
			return true
		}
		visited[arr[i]] = "visited"
	}
	return false
}
