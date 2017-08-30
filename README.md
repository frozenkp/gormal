# gormal
a go package for data normalization

## TODO list
- [ ] member function (interface and struct)
  - [ ] Standard Score for []float64
  - [ ] Standard Score for [][]float64 (data with multiple attributes)
- [ ] function
  - [ ] Standard Score for []float64
  - [ ] Standard Score for [][]float64 (data with multiple attributes)

### structure
```go
type CustomData struct{
  RawData   Data
}

type Data interface{
  Value(i, j int)float64
  RowSize()int
  ColSize()int
}

type DefaultData struct{
  RawData   [][]float64
}
```
