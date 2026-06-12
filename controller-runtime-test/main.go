package main

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap/zapcore"
	corev1 "k8s.io/api/core/v1"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PodReconciler struct {
	client.Client
}

func (r *PodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	fmt.Println("Reconcile triggered for:", req.NamespacedName)

	var pod corev1.Pod

	err := r.Get(ctx, req.NamespacedName, &pod)
	if err != nil {
		fmt.Println("Pod not found")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	fmt.Println("Pod phase:", pod.Status.Phase)

	return ctrl.Result{}, nil
}

func main() {
	ctrl.SetLogger(
		zap.New(
			zap.UseDevMode(true),
			zap.WriteTo(os.Stdout),
			zap.Level(zapcore.InfoLevel),
		),
	)

	// Create manager
	mgr, err := ctrl.NewManager(
		ctrl.GetConfigOrDie(),
		ctrl.Options{
			Cache: cache.Options{
				DefaultNamespaces: map[string]cache.Config{
					"default": {},
					"dev":     {},
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	// Register controller
	err = ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Pod{}).
		Complete(&PodReconciler{
			Client: mgr.GetClient(),
		})

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting controller...")

	// Start controller manager
	err = mgr.Start(ctrl.SetupSignalHandler())
	if err != nil {
		panic(err)
	}
}
