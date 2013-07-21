package interfaces

type ModelInterface interface {
  Save() error
  Destroy() error
}
