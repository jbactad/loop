package ports

import "time"

//go:generate mockery --name=UUIDGenerator --output=./mocks --outpkg=mocks --with-expecter
type UUIDGenerator interface {
	Generate() string
}

//go:generate mockery --name=TimeProvider --output=./mocks --outpkg=mocks --with-expecter
type TimeProvider interface {
	Now() time.Time
}
