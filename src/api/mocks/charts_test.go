package mocks

import (
	"testing"

	"github.com/arschles/assert"
	"github.com/helm/monocular/src/api/testutil"
)

var chartsImplementation = NewMockCharts()

func TestMockChartsChartFromRepo(t *testing.T) {
	// TODO: validate chart data
	_, err := chartsImplementation.ChartFromRepo(testutil.RepoName, testutil.ChartName)
	assert.NoErr(t, err)
	_, err = chartsImplementation.ChartFromRepo(testutil.BogusRepo, testutil.ChartName)
	assert.ExistsErr(t, err, "sent bogus repo name to Charts.ChartFromRepo()")
	_, err = chartsImplementation.ChartFromRepo(testutil.RepoName, testutil.BogusRepo)
	assert.ExistsErr(t, err, "sent bogus chart name to Charts.ChartFromRepo()")
}

func TestMockChartsChartVersionFromRepo(t *testing.T) {
	chart, err := chartsImplementation.ChartVersionFromRepo(testutil.RepoName, testutil.ChartName, testutil.ChartVersion)
	assert.NoErr(t, err)
	assert.Equal(t, *chart.Name, testutil.ChartName, "chart name")
	assert.Equal(t, *chart.Version, testutil.ChartVersion, "chart version")
	_, err = chartsImplementation.ChartVersionFromRepo(testutil.RepoName, testutil.ChartName, "99.99.99")
	assert.ExistsErr(t, err, "sent bogus chart version to ChartVersionFromRepo")
	_, err = chartsImplementation.ChartVersionFromRepo(testutil.BogusRepo, testutil.ChartName, testutil.ChartVersion)
	assert.ExistsErr(t, err, "sent bogus repo name to Charts.ChartFromRepo()")
	_, err = chartsImplementation.ChartVersionFromRepo(testutil.RepoName, testutil.BogusRepo, testutil.ChartVersion)
	assert.ExistsErr(t, err, "sent bogus chart name to Charts.ChartFromRepo()")
	_, err = chartsImplementation.ChartVersionFromRepo(testutil.UnparseableRepo, testutil.ChartName, testutil.ChartVersion)
	assert.ExistsErr(t, err, "sent unparseable repo name to ChartVersionFromRepo")
}

func TestMockChartsChartVersionsFromRepo(t *testing.T) {
	charts, err := chartsImplementation.ChartVersionsFromRepo(testutil.RepoName, testutil.ChartName)
	assert.NoErr(t, err)
	assert.True(t, len(charts) > 0, "returned charts")
	noCharts, err := chartsImplementation.ChartVersionsFromRepo(testutil.BogusRepo, testutil.ChartName)
	assert.ExistsErr(t, err, "sent bogus repo name to GetChartsInRepo")
	assert.True(t, len(noCharts) == 0, "empty charts slice")
	noCharts, err = chartsImplementation.ChartVersionsFromRepo(testutil.UnparseableRepo, testutil.ChartName)
	assert.ExistsErr(t, err, "sent unparseable repo name to GetChartsInRepo")
	assert.True(t, len(noCharts) == 0, "empty charts slice")
	noCharts, err = chartsImplementation.ChartVersionsFromRepo(testutil.RepoName, testutil.BogusRepo)
	assert.ExistsErr(t, err, "sent bogus chart name to GetChartsInRepo")
	assert.True(t, len(noCharts) == 0, "empty charts slice")
}

func TestMockChartsAll(t *testing.T) {
	_, err := chartsImplementation.All()
	assert.NoErr(t, err)
}

func TestMockChartsAllFromRepo(t *testing.T) {
	charts, err := chartsImplementation.AllFromRepo(testutil.RepoName)
	assert.NoErr(t, err)
	assert.True(t, len(charts) > 0, "returned charts")
	noCharts, err := chartsImplementation.AllFromRepo(testutil.BogusRepo)
	assert.ExistsErr(t, err, "sent bogus repo name to GetChartsInRepo")
	assert.True(t, len(noCharts) == 0, "empty charts slice")
}