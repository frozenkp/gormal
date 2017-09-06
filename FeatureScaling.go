package gormal

// FeatureScalingSingle normalizes the first row with "feature scaling".
//
// (Usually only one row in input)
func (d CustomData) FeatureScalingSingle() []float64 {
	// check size
	if d.RawData.RowSize() == 0 || d.RawData.ColSize() == 0 {
		return []float64{}
	}
	// find max and min
	max, min := d.RawData.Value(0, 0), d.RawData.Value(0, 0)
	for i := 0; i < d.RawData.ColSize(); i++ {
		if d.RawData.Value(0, i) > max {
			max = d.RawData.Value(0, i)
		} else if d.RawData.Value(0, i) < min {
			min = d.RawData.Value(0, i)
		}
	}
	// normalization
	out := make([]float64, 0)
	for i := 0; i < d.RawData.ColSize(); i++ {
		out = append(out, (d.RawData.Value(0, i)-min)/(max-min))
	}

	return out
}

// FeatureScaling normalizes every columns with "feature scaling".
func (d CustomData) FeatureScaling() [][]float64 {
	// check size
	if d.RawData.ColSize() == 0 {
		return [][]float64{}
	}

	out := make([][]float64, d.RawData.RowSize())
	for i := 0; i < len(out); i++ {
		out[i] = make([]float64, d.RawData.ColSize())
	}

	for k := 0; k < d.RawData.ColSize(); k++ {
		// find max and min
		max, min := d.RawData.Value(0, k), d.RawData.Value(0, k)
		for i := 0; i < d.RawData.RowSize(); i++ {
			if d.RawData.Value(i, k) > max {
				max = d.RawData.Value(i, k)
			} else if d.RawData.Value(i, k) < min {
				min = d.RawData.Value(i, k)
			}
		}
		// normalization
		for i := 0; i < d.RawData.RowSize(); i++ {
			out[i][k] = (d.RawData.Value(i, k) - min) / (max - min)
		}
	}

	return out
}

// FeatureScalingSingle normalizes the first row with "feature scaling".
//
// (Usually only one row in input)
func FeatureScalingSingle(data []float64) []float64 {
	// check size
	if len(data) == 0 {
		return []float64{}
	}
	// find max and min
	max, min := data[0], data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		} else if data[i] < min {
			min = data[i]
		}
	}
	// normalization
	out := make([]float64, 0)
	for i := 0; i < len(data); i++ {
		out = append(out, (data[i]-min)/(max-min))
	}

	return out
}

// FeatureScaling normalizes every columns with "feature scaling".
func FeatureScaling(data [][]float64) [][]float64 {
	// check size
	if len(data) == 0 {
		return [][]float64{}
	}

	out := make([][]float64, len(data))
	for i := 0; i < len(out); i++ {
		out[i] = make([]float64, len(data[0]))
	}

	for k := 0; k < len(data[0]); k++ {
		// find max and min
		max, min := data[0][k], data[0][k]
		for i := 0; i < len(data); i++ {
			if data[i][k] > max {
				max = data[i][k]
			} else if data[i][k] < min {
				min = data[i][k]
			}
		}
		// normalization
		for i := 0; i < len(data); i++ {
			out[i][k] = (data[i][k] - min) / (max - min)
		}
	}

	return out
}
