// Copyright 2022 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"os"
)

const (
	DefaultPolicyPath         = ".policy.yml"
	DefaultSharedRepository   = ".github"
	DefaultSharedPolicyPath   = "policy.yml"
	DefaultStatusCheckContext = "policy-bot"
)

type PullEvaluationOptions struct {
	PolicyPath string `yaml:"policy_path"`

	SharedRepository string `yaml:"shared_repository"`
	SharedPolicyPath string `yaml:"shared_policy_path"`

	// StatusCheckContext will be used to create the status context. It will be used in the following
	// pattern: <StatusCheckContext>: <Base Branch Name>
	StatusCheckContext string `yaml:"status_check_context"`

	// HideBranchName will hide the branch name from the status context. This is useful if you want to
	// use a single status context for all branches with predictable branch rules.
	HideBranchName bool `yaml:"hide_branch_name"`

	// PostInsecureStatusChecks enables the sending of a second status using just StatusCheckContext as the context,
	// no templating. This is turned off by default. This is to support legacy workflows that depend on the original
	// context behaviour, and will be removed in 2.0
	PostInsecureStatusChecks bool `yaml:"post_insecure_status_checks"`

	// This field is unused but is left to avoid breaking configuration files.
	// This value is now loaded from the GitHub API.
	//
	// TODO(bkeyes): remove in version 2.0
	Deprecated_AppName string `yaml:"app_name"`

	// This field is unused but is left to avoid breaking configuration files.
	// It enabled a temporary workaround for a GitHub API issue.
	//
	// TODO(bkeyes): remove in version 2.0
	Deprecated_DoNotLoadCommitPushedDate bool `yaml:"do_not_load_commit_pushed_date"`
}

func (p *PullEvaluationOptions) fillDefaults() {
	if p.PolicyPath == "" {
		p.PolicyPath = DefaultPolicyPath
	}
	if p.SharedRepository == "" {
		p.SharedRepository = DefaultSharedRepository
	}
	if p.SharedPolicyPath == "" {
		p.SharedPolicyPath = DefaultSharedPolicyPath
	}
	if p.StatusCheckContext == "" {
		p.StatusCheckContext = DefaultStatusCheckContext
	}
}

func (p *PullEvaluationOptions) SetValuesFromEnv(prefix string) {
	setStringFromEnv("POLICY_PATH", prefix, &p.PolicyPath)
	setStringFromEnv("SHARED_REPOSITORY", prefix, &p.SharedRepository)
	setStringFromEnv("SHARED_POLICY_PATH", prefix, &p.SharedPolicyPath)
	setStringFromEnv("STATUS_CHECK_CONTEXT", prefix, &p.StatusCheckContext)
	p.fillDefaults()
}

func setStringFromEnv(key, prefix string, value *string) bool {
	if v, ok := os.LookupEnv(prefix + key); ok {
		*value = v
		return true
	}
	return false
}
