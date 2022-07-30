/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"transwarp.io/tos/demo/pkg/monitor"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	demov1alpha1 "transwarp.io/tos/demo/api/v1alpha1"
)

// RollingMonitor12Reconciler reconciles a RollingMonitor12 object
type RollingMonitor12Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

const MonitorFinalizer = "transwarp.io.tos.monitored/finalizer"

//+kubebuilder:rbac:groups=demo.transwarp.io,resources=rollingmonitor12s,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.transwarp.io,resources=rollingmonitor12s/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.transwarp.io,resources=rollingmonitor12s/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RollingMonitor12 object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *RollingMonitor12Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("start reconcile")
	rollingMonitor := &demov1alpha1.RollingMonitor12{}
	key := client.ObjectKey{Namespace: req.Namespace, Name: req.Name}
	if err := r.Get(ctx, key, rollingMonitor); err != nil {
		logger.Error(err, "get target obj failed")
		return ctrl.Result{}, err
	}

	if rollingMonitor.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(rollingMonitor, MonitorFinalizer) {
			controllerutil.AddFinalizer(rollingMonitor, MonitorFinalizer)
			if err := r.Update(ctx, rollingMonitor); err != nil {
				logger.Error(err, "add MonitorFinalizer failed")
				return ctrl.Result{}, err
			}
		}
	} else {
		if controllerutil.ContainsFinalizer(rollingMonitor, MonitorFinalizer) {
			monitor.RemoveMonitoredDeploy(rollingMonitor.Namespace, rollingMonitor.Spec.DeploymentName)
			controllerutil.RemoveFinalizer(rollingMonitor, MonitorFinalizer)
			if err := r.Update(ctx, rollingMonitor); err != nil {
				logger.Error(err, "remove MonitorFinalizer failed")
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	monitor.AddMonitorDeploy(rollingMonitor.Namespace, rollingMonitor.Spec.DeploymentName)

	// TODO update cr status

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RollingMonitor12Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1alpha1.RollingMonitor12{}).
		Complete(r)
}
