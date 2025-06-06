// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// KubernetesProxyConfigApplyConfiguration represents a declarative configuration of the KubernetesProxyConfig type for use
// with apply.
type KubernetesProxyConfigApplyConfiguration struct {
	Deployment     *ProxyDeploymentApplyConfiguration  `json:"deployment,omitempty"`
	EnvoyContainer *EnvoyContainerApplyConfiguration   `json:"envoyContainer,omitempty"`
	SdsContainer   *SdsContainerApplyConfiguration     `json:"sdsContainer,omitempty"`
	PodTemplate    *PodApplyConfiguration              `json:"podTemplate,omitempty"`
	Service        *ServiceApplyConfiguration          `json:"service,omitempty"`
	ServiceAccount *ServiceAccountApplyConfiguration   `json:"serviceAccount,omitempty"`
	Istio          *IstioIntegrationApplyConfiguration `json:"istio,omitempty"`
	Stats          *StatsConfigApplyConfiguration      `json:"stats,omitempty"`
	AiExtension    *AiExtensionApplyConfiguration      `json:"aiExtension,omitempty"`
	AgentGateway   *AgentGatewayApplyConfiguration     `json:"agentGateway,omitempty"`
	FloatingUserId *bool                               `json:"floatingUserId,omitempty"`
}

// KubernetesProxyConfigApplyConfiguration constructs a declarative configuration of the KubernetesProxyConfig type for use with
// apply.
func KubernetesProxyConfig() *KubernetesProxyConfigApplyConfiguration {
	return &KubernetesProxyConfigApplyConfiguration{}
}

// WithDeployment sets the Deployment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Deployment field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithDeployment(value *ProxyDeploymentApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.Deployment = value
	return b
}

// WithEnvoyContainer sets the EnvoyContainer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnvoyContainer field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithEnvoyContainer(value *EnvoyContainerApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.EnvoyContainer = value
	return b
}

// WithSdsContainer sets the SdsContainer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SdsContainer field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithSdsContainer(value *SdsContainerApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.SdsContainer = value
	return b
}

// WithPodTemplate sets the PodTemplate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodTemplate field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithPodTemplate(value *PodApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.PodTemplate = value
	return b
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithService(value *ServiceApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.Service = value
	return b
}

// WithServiceAccount sets the ServiceAccount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccount field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithServiceAccount(value *ServiceAccountApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.ServiceAccount = value
	return b
}

// WithIstio sets the Istio field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Istio field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithIstio(value *IstioIntegrationApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.Istio = value
	return b
}

// WithStats sets the Stats field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Stats field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithStats(value *StatsConfigApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.Stats = value
	return b
}

// WithAiExtension sets the AiExtension field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AiExtension field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithAiExtension(value *AiExtensionApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.AiExtension = value
	return b
}

// WithAgentGateway sets the AgentGateway field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AgentGateway field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithAgentGateway(value *AgentGatewayApplyConfiguration) *KubernetesProxyConfigApplyConfiguration {
	b.AgentGateway = value
	return b
}

// WithFloatingUserId sets the FloatingUserId field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FloatingUserId field is set to the value of the last call.
func (b *KubernetesProxyConfigApplyConfiguration) WithFloatingUserId(value bool) *KubernetesProxyConfigApplyConfiguration {
	b.FloatingUserId = &value
	return b
}
