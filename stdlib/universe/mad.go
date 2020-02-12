package universe

import (
	"fmt"

	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/math"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/arrow"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/plan"
)

const MadKind = "mad"

type MadOpSpec struct {
	execute.AggregateConfig
}

func init() {
	madSignature := execute.AggregateSignature(nil, nil)

	flux.RegisterPackageValue("universe", MadKind, flux.FunctionValue(MadKind, createMadOpSpec, madSignature))
	flux.RegisterOpSpec(MadKind, newMadOp)
	plan.RegisterProcedureSpec(MadKind, newMadProcedure, MadKind)
	execute.RegisterTransformation(MadKind, createMadTransformation)
}

func createMadOpSpec(args flux.Arguments, a *flux.Administration) (flux.OperationSpec, error) {
	if err := a.AddParentFromArgs(args); err != nil {
		return nil, err
	}
	s := new(MadOpSpec)
	if err := s.AggregateConfig.ReadArgs(args); err != nil {
		return s, err
	}
	return s, nil
}

func newMadOp() flux.OperationSpec {
	return new(MadOpSpec)
}

func (s *MadOpSpec) Kind() flux.OperationKind {
	return MadKind
}

type MadProcedureSpec struct {
	execute.AggregateConfig
}

func newMadProcedure(qs flux.OperationSpec, a plan.Administration) (plan.ProcedureSpec, error) {
	spec, ok := qs.(*MadOpSpec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type %T", qs)
	}
	return &MadProcedureSpec{
		AggregateConfig: spec.AggregateConfig,
	}, nil
}

func (s *MadProcedureSpec) Kind() plan.ProcedureKind {
	return MadKind
}

func (s *MadProcedureSpec) Copy() plan.ProcedureSpec {
	return &MadProcedureSpec{
		AggregateConfig: s.AggregateConfig,
	}
}

// TriggerSpec implements plan.TriggerAwareProcedureSpec
func (s *MadProcedureSpec) TriggerSpec() plan.TriggerSpec {
	return plan.NarrowTransformationTriggerSpec{}
}

func (s *MadProcedureSpec) AggregateMethod() string {
	return MadKind
}
func (s *MadProcedureSpec) ReAggregateSpec() plan.ProcedureSpec {
	return new(MadProcedureSpec)
}

type MadAgg struct{}

func createMadTransformation(id execute.DatasetID, mode execute.AccumulationMode, spec plan.ProcedureSpec, a execute.Administration) (execute.Transformation, execute.Dataset, error) {
	s, ok := spec.(*MadProcedureSpec)
	if !ok {
		return nil, nil, fmt.Errorf("invalid spec type %T", spec)
	}

	t, d := execute.NewAggregateTransformationAndDataset(id, mode, new(MadAgg), s.AggregateConfig, a.Allocator())
	return t, d, nil
}

func (a *MadAgg) NewBoolAgg() execute.DoBoolAgg {
	return nil
}
func (a *MadAgg) NewIntAgg() execute.DoIntAgg {
	return new(MadIntAgg)
}
func (a *MadAgg) NewUIntAgg() execute.DoUIntAgg {
	return new(MadUIntAgg)
}
func (a *MadAgg) NewFloatAgg() execute.DoFloatAgg {
	return new(MadFloatAgg)
}
func (a *MadAgg) NewStringAgg() execute.DoStringAgg {
	return nil
}

type MadIntAgg struct {
	mad int64
	ok  bool
}

func (a *MadIntAgg) DoInt(vs *array.Int64) {
	if l := vs.Len() - vs.NullN(); l > 0 {
		if vs.NullN() == 0 {
			a.mad += math.Int64.Sum(vs)
			a.ok = true
		} else {
			for i := 0; i < vs.Len(); i++ {
				if vs.IsValid(i) {
					a.mad += vs.Value(i)
					a.ok = true
				}
			}
		}
	}
}
func (a *MadIntAgg) Type() flux.ColType {
	return flux.TInt
}
func (a *MadIntAgg) ValueInt() int64 {
	return a.mad
}
func (a *MadIntAgg) IsNull() bool {
	return !a.ok
}

type MadUIntAgg struct {
	mad uint64
	ok  bool
}

func (a *MadUIntAgg) DoUInt(vs *array.Uint64) {
	if l := vs.Len() - vs.NullN(); l > 0 {
		if vs.NullN() == 0 {
			a.mad += math.Uint64.Sum(vs)
			a.ok = true
		} else {
			for i := 0; i < vs.Len(); i++ {
				if vs.IsValid(i) {
					a.mad += vs.Value(i)
					a.ok = true
				}
			}
		}
	}
}
func (a *MadUIntAgg) Type() flux.ColType {
	return flux.TUInt
}
func (a *MadUIntAgg) ValueUInt() uint64 {
	return a.mad
}
func (a *MadUIntAgg) IsNull() bool {
	return !a.ok
}

type MadFloatAgg struct {
	mad []*array.Float64
	ok  bool
	q   *QuantileAgg
}

func (a *MadFloatAgg) DoFloat(vs *array.Float64) {
	vs.Retain()
	a.mad = append(a.mad, vs)
	a.q.DoFloat(vs)
}

// test if we can order the dataset CHECK
// test if we can find the median CHECK
// test if we can find the differences between median and each value
// test if we can find the sort that difference
// test if we can find the median value of that

//vs.Retain tells arrow. says I still need this memory.
//DoFloat gets called values get passed in, calling retain says don't release and keep a reference to it. It doesn't reuse it.
//Implementation detail: Arrow can allocate memory and free memory (in the future sync.pool to not have to call make everytime because make is expensive)

func (a *MadFloatAgg) Type() flux.ColType {
	return flux.TFloat
}
func (a *MadFloatAgg) ValueFloat() float64 {
	//Compute median
	median := a.q.ValueFloat()
	//Initialize quanitlie agg
	q := &QuantileAgg{}
	//Iterate over each of the columns
	for _, vs := range a.mad {
		b := arrow.NewFloatBuilder(nil)
		b.Resize(vs.Len())
		for i, l := 0, vs.Len(); i < l; i++ {
			if vs.IsNull(i) {
				b.AppendNull()
				continue
			}
			v := vs.Value(i)
			diff := v - median
			b.Append(diff)
		}
		vs.Release()
		//Construct the Array
		diffs := b.NewFloat64Array()
		q.DoFloat(diffs)
		diffs.Release()
	}
	a.mad = nil
	return q.ValueFloat()
}

//For each col, compute diff, send them to quantile agg
//Return quanitile agg

//do calculations here. This is for retaining the data
//use dofloat (interface that gets called) to store and keep the arrays using retain on vs
//value float you actually perform the calculation

func (a *MadFloatAgg) IsNull() bool {
	return !a.ok
}
