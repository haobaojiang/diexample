package common

import "github.com/samber/do/v2"

// SafeDoProvide :
// calls Do.Provide safely, it won't panic if same provider is registered
func SafeDoProvide[T any](inj do.Injector, provider do.Provider[T]) {
	name := do.NameOf[T]()
	for _, svc := range inj.ListProvidedServices() {
		if svc.Service == name {
			return
		}
	}
	do.Provide(inj, provider)
}
