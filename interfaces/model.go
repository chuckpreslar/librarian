package interfaces

type ModelInterface interface {
  IsNew() bool
  IsModified() bool
  IsValid() bool
  Save() error
  Destroy() error
}
