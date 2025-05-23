// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivers

import (
	"github.com/go-logr/logr"
	rbacv1 "k8s.io/api/rbac/v1"
)

type k8seventsConfig struct{}

func generatek8seventsRbacRules(_ logr.Logger, _ k8seventsConfig) ([]rbacv1.PolicyRule, error) {
	// The k8s Events Receiver needs the get permissions on the following resources always.
	return []rbacv1.PolicyRule{
		{
			APIGroups: []string{""},
			Resources: []string{
				"events",
				"namespaces",
				"namespaces/status",
				"nodes",
				"nodes/spec",
				"pods",
				"pods/status",
				"replicationcontrollers",
				"replicationcontrollers/status",
				"resourcequotas",
				"services",
			},
			Verbs: []string{"get", "list", "watch"},
		},
		{
			APIGroups: []string{"apps"},
			Resources: []string{
				"daemonsets",
				"deployments",
				"replicasets",
				"statefulsets",
			},
			Verbs: []string{"get", "list", "watch"},
		},
		{
			APIGroups: []string{"extensions"},
			Resources: []string{
				"daemonsets",
				"deployments",
				"replicasets",
			},
			Verbs: []string{"get", "list", "watch"},
		},
		{
			APIGroups: []string{"batch"},
			Resources: []string{
				"jobs",
				"cronjobs",
			},
			Verbs: []string{"get", "list", "watch"},
		},
		{
			APIGroups: []string{"autoscaling"},
			Resources: []string{
				"horizontalpodautoscalers",
			},
			Verbs: []string{"get", "list", "watch"},
		},
	}, nil
}
