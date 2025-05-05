package graph

import (
	"runtime"

	"lem-in/utils"
)

func FindPaths(colony *utils.AntFarm) [][]string {
	paths := [][]string{}

	for _, neighbor := range colony.Links[colony.Start.Name] {
		visited := map[string]bool{colony.Start.Name: true, neighbor: true}
		queue := [][]string{{colony.Start.Name, neighbor}}

		for len(queue) > 0 {
			path := queue[0]
			queue = queue[1:]

			lastRoom := path[len(path)-1]
			if lastRoom == colony.End.Name {
				paths = append(paths, path)
				continue
			}

			for _, next := range colony.Links[lastRoom] {
				var memStats runtime.MemStats
				runtime.ReadMemStats(&memStats)

				highMemory := memStats.Alloc > 100*1024*1024

				if highMemory {
					if !contains(path, next) && !visited[next] {
						newPath := append([]string{}, path...)
						newPath = append(newPath, next)
						queue = append(queue, newPath)
						visited[next] = true
					}
				} else {
					if !contains(path, next) {
						newPath := append([]string{}, path...)
						newPath = append(newPath, next)
						queue = append(queue, newPath)
					}
				}
			}
		}
	}

	return paths
}

func contains(path []string, room string) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

func FindDisjointPaths(paths [][]string, colony *utils.AntFarm) [][]string {
	var currentPaths [][]string
	usedNodes := make(map[string]bool)
	var backtrack func(int)
	backtrack = func(start int) {
		if len(currentPaths) > len(utils.Filter) {
			utils.Filter = make([][]string, len(currentPaths))
			copy(utils.Filter, currentPaths)
		}

		for i := start; i < len(paths); i++ {
			path := paths[i]
			keepPath := true

			for _, node := range path[1 : len(path)-1] {
				if usedNodes[node] {
					keepPath = false
					break
				}
			}

			if keepPath {
				currentPaths = append(currentPaths, path)
				for _, node := range path[1 : len(path)-1] {
					usedNodes[node] = true
				}

				backtrack(i + 1)

				currentPaths = currentPaths[:len(currentPaths)-1]
				for _, node := range path[1 : len(path)-1] {
					delete(usedNodes, node)
				}
			}
		}
	}

	backtrack(0)
	return utils.Filter
}

func FindDisjointPaths2(paths [][]string, colony *utils.AntFarm) [][]string {
	n := make(map[string]int)

	smollpath := make(map[int]bool)

	for i := 0; i < len(paths); i++ {
		if len(paths[i]) == 2 {
			utils.Filter = append(utils.Filter, paths[i])
			continue
		}

		for j := 1; j < len(paths[i])-1; j++ {
			room := paths[i][j]
			if k, exists := n[room]; exists {
				if len(paths[k]) > len(paths[i]) {

					smollpath[k] = false
					n[room] = i

					smollpath[i] = true
				} else {
					smollpath[i] = false
					break
				}
			} else {
				n[room] = i

				smollpath[i] = true
			}
		}

	}

	var result [][]string
	for idx, ok := range smollpath {
		if ok {
			result = append(result, paths[idx])
		}
	}
	return result
}
