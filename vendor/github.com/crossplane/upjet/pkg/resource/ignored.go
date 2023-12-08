// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"fmt"
	"sort"
)

// GetTerraformIgnoreChanges returns a sorted Terraform `ignore_changes`
// lifecycle meta-argument expression by looking for differences between
// the `initProvider` and `forProvider` maps. The ignored fields are the ones
// that are present in initProvider, but not in forProvider.
func GetTerraformIgnoreChanges(forProvider, initProvider map[string]any) []string {
	ignored := getIgnoredFieldsMap("%s", forProvider, initProvider)
	// sort the ignored fields so that we can compare them easily
	sort.Strings(ignored)
	return ignored
}

func getIgnoredFieldsMap(format string, forProvider, initProvider map[string]any) []string {
	ignored := []string{}

	for k := range initProvider {
		if _, ok := forProvider[k]; !ok {
			ignored = append(ignored, fmt.Sprintf(format, k))
		} else {
			// both are the same type so we dont need to check for forProvider type
			if _, ok = initProvider[k].(map[string]any); ok {
				ignored = append(ignored, getIgnoredFieldsMap(fmt.Sprintf(format, k)+"[%q]", forProvider[k].(map[string]any), initProvider[k].(map[string]any))...)
			}
			// if its an array, we need to check if its an array of maps or not
			if _, ok = initProvider[k].([]any); ok {
				ignored = append(ignored, getIgnoredFieldsArray(fmt.Sprintf(format, k), forProvider[k].([]any), initProvider[k].([]any))...)
			}

		}
	}
	return ignored
}

func getIgnoredFieldsArray(format string, forProvider, initProvider []any) []string {
	ignored := []string{}
	for i := range initProvider {
		// Construct the full field path with array index and prefix.
		fieldPath := fmt.Sprintf("%s[%d]", format, i)
		if i < len(forProvider) {
			if _, ok := initProvider[i].(map[string]any); ok {
				ignored = append(ignored, getIgnoredFieldsMap(fieldPath+".%s", forProvider[i].(map[string]any), initProvider[i].(map[string]any))...)
			}
			if _, ok := initProvider[i].([]any); ok {
				ignored = append(ignored, getIgnoredFieldsArray(fieldPath+"%s", forProvider[i].([]any), initProvider[i].([]any))...)
			}
		} else {
			ignored = append(ignored, fieldPath)
		}

	}
	return ignored
}
