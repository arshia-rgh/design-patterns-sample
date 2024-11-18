package solid

type Document struct{}

type Operations interface {
	Print()
	Scan()
	Fax()
}

// this is ok because MultiFunctionMachine can handle all of these operations

type MultiFunctionMachine struct{}

func (mf *MultiFunctionMachine) Print() {
	//ok
}
func (mf *MultiFunctionMachine) Scan() {
	//ok
}
func (mf *MultiFunctionMachine) Fax() {
	//ok
}

type OldMachine struct{}

func (om *OldMachine) Print() {
	//ok
}
func (om *OldMachine) Scan() {
	// OldMachine doesn't support Scan, but it is forced to implement this, and it will break Interface Segregation
}
func (om *OldMachine) Fax() {
	// OldMachine doesn't support Fax, but it is forced to implement this, and it will break Interface Segregation
}

// we should separate the interfaces

type Printer interface {
	Print()
}
type Scanner interface {
	Scan()
}

// and ...

// here each struct can implement interfaces it need and does not force to implement unnecessary
