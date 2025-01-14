// Copyright (c) 2020 Red Hat, Inc.

package metrics

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"github.com/open-cluster-management/cluster-lifecycle-e2e/pkg/clients"
	libgocmd "github.com/open-cluster-management/library-e2e-go/pkg/cmd"
	"k8s.io/klog"
)

func init() {
	klog.SetOutput(GinkgoWriter)
	klog.InitFlags(nil)

	libgocmd.InitFlags(nil)
}

var _ = BeforeSuite(func() {
	hubClients = clients.GetHubClients()
})

func TestMetrics(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter(fmt.Sprintf("%s-%d.xml", "/results/result-metrics", config.GinkgoConfig.ParallelNode))
	RunSpecsWithDefaultAndCustomReporters(t, "Metrics Suite", []Reporter{junitReporter})
}
