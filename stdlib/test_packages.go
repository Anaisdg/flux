// DO NOT EDIT: This file is autogenerated via the builtin command.

package stdlib

import (
	ast "github.com/influxdata/flux/ast"
	naivebayesclassifier "github.com/influxdata/flux/stdlib/contrib/RohanSreerama5/naiveBayesClassifier"
	anomalydetection "github.com/influxdata/flux/stdlib/contrib/anaisdg/anomalydetection"
	statsmodels "github.com/influxdata/flux/stdlib/contrib/anaisdg/statsmodels"
	aggregate "github.com/influxdata/flux/stdlib/contrib/jsternberg/aggregate"
	rows "github.com/influxdata/flux/stdlib/contrib/jsternberg/rows"
	date "github.com/influxdata/flux/stdlib/date"
	experimental "github.com/influxdata/flux/stdlib/experimental"
	aggregate1 "github.com/influxdata/flux/stdlib/experimental/aggregate"
	array "github.com/influxdata/flux/stdlib/experimental/array"
	geo "github.com/influxdata/flux/stdlib/experimental/geo"
	json "github.com/influxdata/flux/stdlib/experimental/json"
	http "github.com/influxdata/flux/stdlib/http"
	monitor "github.com/influxdata/flux/stdlib/influxdata/influxdb/monitor"
	schema "github.com/influxdata/flux/stdlib/influxdata/influxdb/schema"
	secrets "github.com/influxdata/flux/stdlib/influxdata/influxdb/secrets"
	tasks "github.com/influxdata/flux/stdlib/influxdata/influxdb/tasks"
	promql "github.com/influxdata/flux/stdlib/internal/promql"
	interpolate "github.com/influxdata/flux/stdlib/interpolate"
	planner "github.com/influxdata/flux/stdlib/planner"
	regexp "github.com/influxdata/flux/stdlib/regexp"
	strings "github.com/influxdata/flux/stdlib/strings"
	chronograf "github.com/influxdata/flux/stdlib/testing/chronograf"
	influxql "github.com/influxdata/flux/stdlib/testing/influxql"
	kapacitor "github.com/influxdata/flux/stdlib/testing/kapacitor"
	pandas "github.com/influxdata/flux/stdlib/testing/pandas"
	prometheus "github.com/influxdata/flux/stdlib/testing/prometheus"
	promql1 "github.com/influxdata/flux/stdlib/testing/promql"
	usage "github.com/influxdata/flux/stdlib/testing/usage"
	universe "github.com/influxdata/flux/stdlib/universe"
)

var FluxTestPackages = func() []*ast.Package {
	var pkgs []*ast.Package
	pkgs = append(pkgs, naivebayesclassifier.FluxTestPackages...)
	pkgs = append(pkgs, anomalydetection.FluxTestPackages...)
	pkgs = append(pkgs, statsmodels.FluxTestPackages...)
	pkgs = append(pkgs, aggregate.FluxTestPackages...)
	pkgs = append(pkgs, rows.FluxTestPackages...)
	pkgs = append(pkgs, date.FluxTestPackages...)
	pkgs = append(pkgs, experimental.FluxTestPackages...)
	pkgs = append(pkgs, aggregate1.FluxTestPackages...)
	pkgs = append(pkgs, array.FluxTestPackages...)
	pkgs = append(pkgs, geo.FluxTestPackages...)
	pkgs = append(pkgs, json.FluxTestPackages...)
	pkgs = append(pkgs, http.FluxTestPackages...)
	pkgs = append(pkgs, monitor.FluxTestPackages...)
	pkgs = append(pkgs, schema.FluxTestPackages...)
	pkgs = append(pkgs, secrets.FluxTestPackages...)
	pkgs = append(pkgs, tasks.FluxTestPackages...)
	pkgs = append(pkgs, promql.FluxTestPackages...)
	pkgs = append(pkgs, interpolate.FluxTestPackages...)
	pkgs = append(pkgs, planner.FluxTestPackages...)
	pkgs = append(pkgs, regexp.FluxTestPackages...)
	pkgs = append(pkgs, strings.FluxTestPackages...)
	pkgs = append(pkgs, chronograf.FluxTestPackages...)
	pkgs = append(pkgs, influxql.FluxTestPackages...)
	pkgs = append(pkgs, kapacitor.FluxTestPackages...)
	pkgs = append(pkgs, pandas.FluxTestPackages...)
	pkgs = append(pkgs, prometheus.FluxTestPackages...)
	pkgs = append(pkgs, promql1.FluxTestPackages...)
	pkgs = append(pkgs, usage.FluxTestPackages...)
	pkgs = append(pkgs, universe.FluxTestPackages...)
	return pkgs
}()
