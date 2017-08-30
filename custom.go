package gormal

import(

)

// Data is a interface for custom data.
//
// You should implement Value(), RowSize(), and ColSize() for your own Data type.
type Data interface{
  Value(i, j int)float64
  RowSize()int
  ColSize()int
}

// DefaultData is used for normal [][]float64 type,
// and all member functions have been implemented.
type DefaultData struct{
  RawData   [][]float64
}

// CustomData is used to call all normalization functions,
// and you should add your own Data to RawData.
type CustomData struct{
  RawData   Data
}

// Value is the default Value() for DefaultData.
func (d DefaultData)Value(i, j int)float64{
  return d.RawData[i][j]
}

// RowSize is the default RowSize() for DefaultData.
func (d DefaultData)RowSize()int{
  return len(d.RawData)
}

// ColSize is the default ColSize() for DefaultData.
func (d DefaultData)ColSize()int{
  if len(d.RawData) == 0 {
    return 0
  }
  return len(RawData[0])
}

// New return a CustomData.
func New(rawData Data)CustomData{
  return CustomData{rawData}
}

// NewFloat64Data return a CustomData with DefaultData.
func NewFloat64Data(rawData [][]float64)CustomData{
  return CustomData{DefaultData{rawData}}
}

func (d CustomData)StandardScoreSingle()[]float64{
}

func (d CustomData)StandardScore()[][]float64{
}
