package helper

func SliceDiff[S ~[]T, T comparable](source S, target S, cmp func(T, T) bool) (add S, remove S) {
	add = make(S, 0, len(source))
	remove = make(S, 0, len(target))

	// Iterate over source and check if each element is in target using cmp
	for _, v := range source {
		found := false
		for _, w := range target {
			if cmp == nil {
				if v == w {
					found = true
					break
				}
			} else {
				if cmp(v, w) {
					found = true
					break
				}
			}
		}
		if !found {
			add = append(add, v)
		}
	}

	// Iterate over target and check if each element is in source using cmp
	for _, v := range target {
		found := false
		for _, w := range source {
			if cmp == nil {
				if v == w {
					found = true
					break
				}
			} else {
				if cmp(v, w) {
					found = true
					break
				}
			}
		}
		if !found {
			remove = append(remove, v)
		}
	}

	return
}
