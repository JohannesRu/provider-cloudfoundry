/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/upbound/upjet-provider-template/internal/controller/providerconfig"
	instance "github.com/upbound/upjet-provider-template/internal/controller/serviceinstance/instance"
	space "github.com/upbound/upjet-provider-template/internal/controller/space/space"
	users "github.com/upbound/upjet-provider-template/internal/controller/spaceusers/users"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		instance.Setup,
		space.Setup,
		users.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
