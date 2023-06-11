package main

import (
	"fmt"
)

func main() {
	fmt.Println("let's start")
}

// func createTestingFiles() {
// 	fileChanges := []string{
// 		".github/workflows/go-merge-req-dev.yaml",
// 		".github/workflows/go-pull-req-feat-dev.yaml",
// 		"Dockerfile",
// 		"book/.gitkeep",
// 		"book/entity/book.go",
// 		"book/entity/book_test.go",
// 	}

// 	pattern := `.*_test\.go`
// 	r := regexp.MustCompile(pattern)

// 	files := []string{}
// 	for _, v := range fileChanges {
// 		if r.MatchString(v) {
// 			dir := filepath.Dir(v)
// 			dir = strings.TrimPrefix(dir, "/")
// 			dir = "./" + dir
// 			// fmt.Println(dir)
// 			files = append(files, dir)
// 		}

// 	}
// 	fmt.Println(files)
// 	testingFile := strings.Join(files, " ")

// 	fmt.Println(testingFile)

// }
