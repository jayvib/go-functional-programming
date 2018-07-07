package workflow

type Data interface {}

type Monad func(e error) (Data, error)