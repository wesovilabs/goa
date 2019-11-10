package testdata

import (
	"github.com/wesovilabs/goa/api"
	"github.com/wesovilabs/goa/api/context"
)

type Person struct {
	firstName string
	age       int
}

func CreatePerson(firstName string, age int) (error, *Person) {
	return nil, &Person{
		firstName: firstName,
		age:       age,
	}
}

type SampleAspect struct {

}

func (a *SampleAspect) Before(ctx *context.GoaContext){

}

func (a *SampleAspect) Returning(ctx *context.GoaContext){

}

func NewSampleAspect() api.Around{
	return &SampleAspect{

	}
}

type SampleBefore struct {

}

func (a *SampleBefore) Before(ctx *context.GoaContext){

}

func NewSampleBefore() api.Before{
	return &SampleBefore{

	}
}

type SampleReturning struct {

}

func (a *SampleReturning) Returning(ctx *context.GoaContext){

}

func NewSampleReturning() api.Returning{
	return &SampleReturning{

	}
}

func Goa()*api.Goa {
	return api.New().
		WithAround("*.*(...)(error,...)",NewSampleAspect).
		WithBefore("*.*(...)...",NewSampleBefore).
		WithReturning("*.*(...)...",NewSampleReturning)
}
