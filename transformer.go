package transformer

func List[E any, T any](rows []E, cb func(val E) T) []T {
	items := make([]T, len(rows))

	for i, w := range rows {
		items[i] = cb(w)
	}

	return items
}

func ListErr[E any, T any](rows []E, cb func(val E) (T, error)) ([]T, error) {
	items := make([]T, len(rows))

	for i, w := range rows {
		val, err := cb(w)
		if err != nil {
			return nil, err
		}
		items[i] = val
	}

	return items, nil
}
