package objects

type PhysicalInterface int

const (
  I2cInterface      PhysicalInterface = 1 + iota
  OneWireInterface
  UnknownInterface
)


type HasPhysicalInterface interface {
  GetPhysicalInterface() PhysicalInterface
}
