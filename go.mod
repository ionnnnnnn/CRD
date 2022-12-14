module transwarp.io/tos/demo

go 1.15

require (
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/klog/v2 v2.4.0
	sigs.k8s.io/controller-runtime v0.8.3
)

replace (
	k8s.io/api v0.20.2 => k8s.io/api v0.19.6
	k8s.io/apimachinery v0.20.2 => k8s.io/apimachinery v0.19.6
	k8s.io/client-go v0.20.2 => k8s.io/client-go v0.19.6
)
