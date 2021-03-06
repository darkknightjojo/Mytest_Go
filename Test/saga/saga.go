package saga

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type TxOptions struct {
}

type SubTx struct {
	Name           string
	Func           interface{}
	CompensateFunc interface{}
	Options        *TxOptions
}

type Saga struct {
	Name   string
	subTxs []*SubTx
}

type Result struct {
	ExecutionError   error
	CompensateErrors []error
}

func NewSaga(name string) *Saga {
	return &Saga{
		Name: name,
	}
}

func (saga *Saga) AddSubTx(tx *SubTx) error {
	if err := checkSubTx(tx); err != nil {
		return err
	}
	saga.subTxs = append(saga.subTxs, tx)
	return nil
}

func checkSubTx(subTx *SubTx) error {
	funcType := reflect.TypeOf(subTx.Func)
	if funcType.Kind() != reflect.Func {
		return fmt.Errorf("func field is not a func, but %s", funcType.Kind())
	}

	compensateType := reflect.TypeOf(subTx.CompensateFunc)
	if compensateType.Kind() != reflect.Func {
		return fmt.Errorf("func field is not a func, but %s", compensateType.Kind())
	}

	if funcType.NumIn() != 1 || funcType.In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
		return errors.New("func must have strictly one parameter context.Context")
	}
	if funcType.NumOut() == 0 {
		return errors.New("func must have at least one out value of type error")
	}
	if !funcType.Out(funcType.NumOut() - 1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		return errors.New("last out parameter of func must be of type error")
	}

	if compensateType.NumIn() == 0 {
		return errors.New("compensate must have at least one parameter context.Context")
	}
	if compensateType.In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
		return errors.New("first parameter of a compensate must be of type context.Context")
	}
	if compensateType.NumOut() != 1 {
		return errors.New("compensate must must return single value of type error")
	}

	if compensateType.NumIn() != funcType.NumOut() {
		return errors.New("compensate in params not matched to func return values")
	}

	for i := 0; i < compensateType.NumIn()-1; i++ {
		if compensateType.In(i+1) != funcType.Out(i) {
			return fmt.Errorf("param %d not matched in func and compensate", i)
		}
	}

	return nil
}
