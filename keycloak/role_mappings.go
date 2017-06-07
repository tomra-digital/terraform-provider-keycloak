// This file contains shared code for the different role mapping types

package keycloak

import "fmt"

type Role struct {
	Id                 string `json:"id"`
	Name               string `json:"name,omitempty"`
	ScopeParamRequired bool   `json:"scopeParamRequired"`
	ClientRole         *bool  `json:"clientRole,omitempty"`
	ContainerId        string `json:"containerId,omitempty"`
}

// Find a role for a given user based on the role ID.
// The idea is to hide the complexity of the randomly generated role IDs from the user.
// TODO: Evaluate whether this is the most sensible approach vs. some sort of data provider
func FindRoleForUser(roles []Role, roleIdentifier string) (*Role, error) {
	var role Role
	for _, value := range roles {
		if value.Name == roleIdentifier || value.Id == roleIdentifier {
			role = value
		}
	}

	if role.Id == "" {
		return nil, fmt.Errorf("Role %s not found", roleIdentifier)
	}

	return &role, nil
}
