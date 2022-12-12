package applications

import "github.com/upbound/upjet/pkg/config"

const group = "applications"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_application", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_application_certificate", func(r *config.Resource) {
		r.References["application_object_id"] = config.Reference{
			Type: "Application",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
	p.AddResourceConfigurator("azuread_application_password", func(r *config.Resource) {
		r.References["application_object_id"] = config.Reference{
			Type: "Application",
		}
		// We need to override the default group that upjet generated for
		// this resource, which would be "azuread"
		r.ShortGroup = group
	})
}
