package diff_test

import (
	"fmt"

	"github.com/PieterD/diff"
)

func ExampleStructNew() {
	// Data
	l := []diff.Element{
		diff.Element{
			Id:      "hello",
			Content: []string{"content"},
		},
		diff.Element{
			Id:      "world",
			Content: []string{"content"},
		},
		diff.Element{
			Id:      "content",
			Content: []string{"hello", "world"},
		},
	}
	r := []diff.Element{
		diff.Element{
			Id:      "its",
			Content: []string{"content"},
		},
		diff.Element{
			Id:      "my",
			Content: []string{},
		},
		diff.Element{
			Id:      "world",
			Content: []string{"content"},
		},
		diff.Element{
			Id:      "content",
			Content: []string{"its", "my", "world"},
		},
	}

	// Diff l and r using Strings
	d := diff.NewContent(diff.Elements{
		Left:  l,
		Right: r,
	})
	// Print the diff
	for _, idelta := range d {
		switch idelta.Delta {
		case diff.Both:
			fmt.Printf("  %s\n", l[idelta.Index].Id)
			for _, s := range l[idelta.Index].Content {
				fmt.Printf("    %s\n", s)
			}
		case diff.Left:
			fmt.Printf("- %s\n", l[idelta.Index].Id)
			for _, s := range l[idelta.Index].Content {
				fmt.Printf("-   %s\n", s)
			}
		case diff.Right:
			fmt.Printf("+ %s\n", r[idelta.Index].Id)
			for _, s := range r[idelta.Index].Content {
				fmt.Printf("+   %s\n", s)
			}
		case diff.Content:
			fmt.Printf("  %s\n", l[idelta.Index].Id)

			for _, jdelta := range idelta.ContentDiff {
				switch jdelta.Delta {
				case diff.Both:
					fmt.Printf("    %s\n", l[idelta.Index].Content[jdelta.Index])
				case diff.Left:
					fmt.Printf("  - %s\n", l[idelta.Index].Content[jdelta.Index])
				case diff.Right:
					fmt.Printf("  + %s\n", r[idelta.IndexR].Content[jdelta.Index])
				}
			}
		}
	}
	// Output:
	// - hello
	// -   content
	// + its
	// +   content
	// + my
	//   world
	//     content
	//   content
	//   - hello
	//   + its
	//   + my
	//     world

}
