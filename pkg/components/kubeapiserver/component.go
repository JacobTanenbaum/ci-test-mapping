package kubeapiserver

import (
	v1 "github.com/openshift-eng/ci-test-mapping/pkg/api/types/v1"
	"github.com/openshift-eng/ci-test-mapping/pkg/config"
	"github.com/openshift-eng/ci-test-mapping/pkg/util"
)

type Component struct {
	*config.Component
}

var KubeApiserverComponent = Component{
	Component: &config.Component{
		Name:                 "kube-apiserver",
		Operators:            []string{"kube-apiserver"},
		DefaultJiraComponent: "kube-apiserver",
		Matchers: []config.ComponentMatcher{
			{
				IncludeAll: []string{"bz-kube-apiserver"},
			},
			{
				IncludeAll: []string{"cache-kube-api-"},
			},
			{
				IncludeAll: []string{"kube-api-", "-connections"},
			},
			{
				IncludeAll: []string{"[sig-api-machinery][Feature:APIServer]"},
			},
			{
				SIG:      "sig-api-machinery",
				Priority: -1,
			},
		},
	},
}

func (c *Component) IdentifyTest(test *v1.TestInfo) (*v1.TestOwnership, error) {
	if matcher := c.FindMatch(test); matcher != nil {
		jira := matcher.JiraComponent
		if jira == "" {
			jira = c.DefaultJiraComponent
		}
		return &v1.TestOwnership{
			Name:          test.Name,
			Component:     c.Name,
			JIRAComponent: jira,
			Priority:      matcher.Priority,
			Capabilities:  append(matcher.Capabilities, identifyCapabilities(test)...),
		}, nil
	}

	return nil, nil
}

func (c *Component) StableID(test *v1.TestInfo) string {
	return util.StableID(test)
}

func (c *Component) JiraComponents() (components []string) {
	components = []string{c.DefaultJiraComponent}
	for _, m := range c.Matchers {
		components = append(components, m.JiraComponent)
	}

	return components
}
