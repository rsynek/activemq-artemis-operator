module github.com/artemiscloud/activemq-artemis-operator

go 1.16

require (
	github.com/Azure/go-amqp v0.17.4
	//this module is problematic
	github.com/RHsyseng/operator-utils v1.4.7
	github.com/artemiscloud/activemq-artemis-management v0.0.0-20220211120143-717f9a910005
	github.com/go-logr/logr v0.4.0
	github.com/onsi/ginkgo/v2 v2.1.4
	github.com/onsi/gomega v1.19.0
	github.com/openshift/api v3.9.0+incompatible
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/client-go v0.22.2
	sigs.k8s.io/controller-runtime v0.10.0
)
