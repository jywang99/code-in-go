// +build pro
// +build enterprise

package buildtags

func init() {
	features = append(features,
		"Enterprise Feature #1",
		"Enterprise Feature #2",
	)
}

// if build tags are put on the same line, they are interpreted in OR logic
// e.g. 
// +build pro enterprise
// in the current example, it's in AND logic
