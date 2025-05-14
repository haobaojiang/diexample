package common

// the package lists the common interfaces and types used in the project.

//-------------------------------------

type Orders interface {
	GetById() (Order, error)
}

type Order interface {
	Id() int64
	Pay() error
}

//---------------------------------------

type Payment interface {
	Pay(id int64) error
}

//----------------------------------------
