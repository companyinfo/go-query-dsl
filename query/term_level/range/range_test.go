package ranging

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// ptr returns a pointer to v. Bound fields (GT/GTE/LT/LTE/Boost) are now typed as pointers so a
// legitimate zero value can be distinguished from "unset" — this helper just makes building the
// expected *Param values in these tests convenient.
func ptr[T any](v T) *T {
	return &v
}

// RangeTestSuite exercises the string-bound case (e.g. date strings, keyword terms), mirroring
// the original suite. See TestNew_Numeric, TestNew_Time, TestValidate, and
// TestZeroValueBoundIsEmitted below for coverage of the other Value types and the behavior that
// changed when GT/GTE/LT/LTE moved from bare string to *T.
type RangeTestSuite struct {
	suite.Suite
}

func (r *RangeTestSuite) TestNewRange_Success() {
	require := r.Require()
	rang := &Range[string]{Range: map[string]*Param[string]{"date": {}}}

	require.Equal(rang, New("date", &Param[string]{}))
}

func (r *RangeTestSuite) TestNewParam_Success() {
	require := r.Require()
	param := &Param[string]{}

	require.Equal(param, NewParam[string]())
}

func (r *RangeTestSuite) TestSetGT_Success() {
	require := r.Require()
	param := &Param[string]{GT: ptr("2022-04-17T06:00:00")}

	require.Equal(param, NewParam[string]().SetGT("2022-04-17T06:00:00"))
}

func (r *RangeTestSuite) TestSetGTE_Success() {
	require := r.Require()
	param := &Param[string]{GTE: ptr("2022-04-17T06:00:00")}

	require.Equal(param, NewParam[string]().SetGTE("2022-04-17T06:00:00"))
}

func (r *RangeTestSuite) TestSetLT_Success() {
	require := r.Require()
	param := &Param[string]{LT: ptr("2022-04-17T06:00:00")}

	require.Equal(param, NewParam[string]().SetLT("2022-04-17T06:00:00"))
}

func (r *RangeTestSuite) TestSetLTE_Success() {
	require := r.Require()
	param := &Param[string]{LTE: ptr("2022-04-17T06:00:00")}

	require.Equal(param, NewParam[string]().SetLTE("2022-04-17T06:00:00"))
}

func (r *RangeTestSuite) TestSetFormat_Success() {
	require := r.Require()
	param := &Param[string]{Format: "epoch_millis"}

	require.Equal(param, NewParam[string]().SetFormat("epoch_millis"))
}

func (r *RangeTestSuite) TestSetRelation_Success() {
	require := r.Require()
	param := &Param[string]{Relation: IntersectsRelationType}

	require.Equal(param, NewParam[string]().SetRelation(IntersectsRelationType))
}

func (r *RangeTestSuite) TestSetBoost_Success() {
	require := r.Require()
	param := &Param[string]{Boost: ptr(1.0)}

	require.Equal(param, NewParam[string]().SetBoost(1.0))
}

func (r *RangeTestSuite) TestSetTimezone_Success() {
	require := r.Require()
	param := &Param[string]{Timezone: "-04:00"}

	require.Equal(param, NewParam[string]().SetTimezone("-04:00"))
}

// TestNew_Numeric covers int64/float64 bounds, the main case the generic rewrite was for.
func (r *RangeTestSuite) TestNew_Numeric() {
	require := r.Require()
	got := New("turnover", NewParam[int64]().SetGTE(100000).SetLTE(500000))
	want := &Range[int64]{Range: map[string]*Param[int64]{
		"turnover": {GTE: ptr(int64(100000)), LTE: ptr(int64(500000))},
	}}

	require.Equal(want, got)

	b, err := json.Marshal(got)

	require.NoError(err)
	require.JSONEq(`{"range":{"turnover":{"gte":100000,"lte":500000}}}`, string(b))
}

func (r *RangeTestSuite) TestNew_Float() {
	require := r.Require()
	got := New("profit", NewParam[float64]().SetGT(0))
	want := &Range[float64]{Range: map[string]*Param[float64]{
		"profit": {GT: ptr(0.0)},
	}}

	require.Equal(want, got)

	b, err := json.Marshal(got)

	require.NoError(err)
	require.JSONEq(`{"range":{"profit":{"gt":0}}}`, string(b))
}

// TestNew_Time covers a time.Time bound, marshaled as RFC3339 by default.
func (r *RangeTestSuite) TestNew_Time() {
	require := r.Require()
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	got := New("createdAt", NewParam[time.Time]().SetGTE(start).SetLT(end))
	b, err := json.Marshal(got)

	require.NoError(err)
	require.JSONEq(`{"range":{"createdAt":{"gte":"2026-01-01T00:00:00Z","lt":"2026-02-01T00:00:00Z"}}}`, string(b))
}

// TestZeroValueBoundIsEmitted guards against the bug the old string+omitempty fields had:
// an explicit zero bound must still appear in the JSON body, not be dropped as "unset".
func (r *RangeTestSuite) TestZeroValueBoundIsEmitted() {
	require := r.Require()
	got := New("profit", NewParam[int64]().SetGTE(0))
	b, err := json.Marshal(got)

	require.NoError(err)
	require.JSONEq(`{"range":{"profit":{"gte":0}}}`, string(b))
	require.NotNil(got.Range["profit"].GTE)
	require.Equal(int64(0), *got.Range["profit"].GTE)
}

// TestZeroBoost_Emitted guards the same behavior for Boost, which is now *float64.
func (r *RangeTestSuite) TestZeroBoost_Emitted() {
	require := r.Require()
	p := NewParam[string]().SetGTE("a").SetBoost(0)
	b, err := json.Marshal(p)

	require.NoError(err)
	require.JSONEq(`{"gte":"a","boost":0}`, string(b))
}

func (r *RangeTestSuite) TestValidate() {
	require := r.Require()
	tests := []struct {
		name    string
		param   *Param[int64]
		wantErr bool
	}{
		{
			name:  "valid single lower bound",
			param: NewParam[int64]().SetGTE(1),
		},
		{
			name:  "valid lower and upper bound",
			param: NewParam[int64]().SetGTE(1).SetLTE(10),
		},
		{
			name:    "GT and GTE both set",
			param:   NewParam[int64]().SetGT(1).SetGTE(2),
			wantErr: true,
		},
		{
			name:    "LT and LTE both set",
			param:   NewParam[int64]().SetLT(1).SetLTE(2),
			wantErr: true,
		},
		{
			name:    "no bounds set",
			param:   NewParam[int64](),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		r.Run(tt.name, func() {
			err := tt.param.Validate()
			if tt.wantErr {
				require.Error(err)
				return
			}
			require.NoError(err)
		})
	}
}

func TestRange(t *testing.T) {
	suite.Run(t, new(RangeTestSuite))
}
