package keycloak

import (
	"fmt"
)

const (
	rolesUri          = "%s/auth/admin/realms/%s/users/%s/role-mappings/realm"
	availableRolesUri = "%s/auth/admin/realms/%s/users/%s/role-mappings/realm/available"
	compositeRolesUri = "%s/auth/admin/realms/%s/users/%s/role-mappings/realm/composite"
)

// Attempt to look up available roles for a given user ID
func (c *KeycloakClient) GetAvailableRolesForUser(userId string, realm string) ([]Role, error) {
	url := fmt.Sprintf(availableRolesUri, c.url, realm, userId)

	var roles []Role
	err := c.get(url, &roles)

	return roles, err
}

// Attempt to look up copmosite (effective) roles for a given user ID
func (c *KeycloakClient) GetCompositeRolesForUser(userId string, realm string) ([]Role, error) {
	url := fmt.Sprintf(compositeRolesUri, c.url, realm, userId)

	var roles []Role
	err := c.get(url, &roles)

	return roles, err
}

// This attempts to add a Keycloak role to a user based after looking up the role ID from the available rolesUri.
func (c *KeycloakClient) AddRoleToUser(userId string, roleName string, realm string) (*Role, error) {
	roles, err := c.GetAvailableRolesForUser(userId, realm)
	if err != nil {
		return nil, err
	}

	role, err := FindRoleForUser(roles, roleName)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(rolesUri, c.url, realm, userId)
	body := []Role{*role}

	_, err = c.post(url, &body)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (c *KeycloakClient) RemoveRoleFromUser(userId string, role *Role, realm string) error {
	url := fmt.Sprintf(rolesUri, c.url, realm, userId)
	body := []Role{*role}

	err := c.delete(url, body)
	if err != nil {
		return err
	}

	return nil
}
