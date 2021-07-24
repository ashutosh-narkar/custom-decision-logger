module github.com/open-policy-agent/custom-decision-logger

go 1.16

require (
	github.com/open-policy-agent/opa v0.30.2
	github.com/open-policy-agent/opa-envoy-plugin v0.24.0
)

replace github.com/open-policy-agent/opa-envoy-plugin => ../opa-envoy-plugin
