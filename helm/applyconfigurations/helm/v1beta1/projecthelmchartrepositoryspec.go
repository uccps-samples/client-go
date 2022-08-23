// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// ProjectHelmChartRepositorySpecApplyConfiguration represents an declarative configuration of the ProjectHelmChartRepositorySpec type for use
// with apply.
type ProjectHelmChartRepositorySpecApplyConfiguration struct {
	Disabled                *bool                                              `json:"disabled,omitempty"`
	DisplayName             *string                                            `json:"name,omitempty"`
	Description             *string                                            `json:"description,omitempty"`
	ProjectConnectionConfig *ConnectionConfigNamespaceScopedApplyConfiguration `json:"connectionConfig,omitempty"`
}

// ProjectHelmChartRepositorySpecApplyConfiguration constructs an declarative configuration of the ProjectHelmChartRepositorySpec type for use with
// apply.
func ProjectHelmChartRepositorySpec() *ProjectHelmChartRepositorySpecApplyConfiguration {
	return &ProjectHelmChartRepositorySpecApplyConfiguration{}
}

// WithDisabled sets the Disabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Disabled field is set to the value of the last call.
func (b *ProjectHelmChartRepositorySpecApplyConfiguration) WithDisabled(value bool) *ProjectHelmChartRepositorySpecApplyConfiguration {
	b.Disabled = &value
	return b
}

// WithDisplayName sets the DisplayName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisplayName field is set to the value of the last call.
func (b *ProjectHelmChartRepositorySpecApplyConfiguration) WithDisplayName(value string) *ProjectHelmChartRepositorySpecApplyConfiguration {
	b.DisplayName = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *ProjectHelmChartRepositorySpecApplyConfiguration) WithDescription(value string) *ProjectHelmChartRepositorySpecApplyConfiguration {
	b.Description = &value
	return b
}

// WithProjectConnectionConfig sets the ProjectConnectionConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProjectConnectionConfig field is set to the value of the last call.
func (b *ProjectHelmChartRepositorySpecApplyConfiguration) WithProjectConnectionConfig(value *ConnectionConfigNamespaceScopedApplyConfiguration) *ProjectHelmChartRepositorySpecApplyConfiguration {
	b.ProjectConnectionConfig = value
	return b
}
