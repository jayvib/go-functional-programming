package name

type NameMonoid interface {
	Append(s string) NameMonoid
	Zero() string
}

type NameContainer struct {
	name string
}

func (s NameContainer) Append(name string) NameMonoid {
	s.name = s.name + name
	return s
}

func (s NameContainer) String() string {
	return s.name
}

func (NameContainer) Zero() string {
	return ""
}

func WrapName(s string) NameMonoid {
	return NameContainer{name: s}
}
