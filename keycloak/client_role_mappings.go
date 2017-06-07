package keycloak

import "fmt"

const (
	// This URI is used to attach (POST) or remove (DELETE) client role mappings
	clientRolesUri = "%s/auth/admin/realms/%s/users/%s/role-mappings/clients/%s"

	// This URI is used to fetch available client roles per "role-owning" client
	availableClientRolesUri = "%s/auth/admin/realms/%s/users/%s/role-mappings/clients/%s/available"

	// This URI is used to fetch currently active roles per client
	compositeClientRolesUri = "%s/auth/admin/realms/%s/users/%s/role-mappings/clients/%s/composite"
)

func (c *KeycloakClient) GetAvailableRolesForClient(realm string, userId string, clientId string) ([]Role, error) {
	url := fmt.Sprintf(availableClientRolesUri, c.url, realm, userId, clientId)

	var roles []Role
	err := c.get(url, &roles)

	return roles, err
}

func (c *KeycloakClient) GetCompositeRolesForClient(realm string, userId string, clientId string) ([]Role, error) {
	url := fmt.Sprintf(compositeClientRolesUri, c.url, realm, userId, clientId)

	var roles []Role
	err := c.get(url, &roles)

	return roles, err
}

// Get available:
// GET /auth/admin/realms/TomraConnectUsers/users/e0aef983-c14d-4de7-8ef9-c5722fbbbeac/role-mappings/clients/38f5b9a2-7084-469a-ad5a-8a9759e35ae4/available
// Response:
/*
[
  {
    "id": "dba92d03-588a-41ae-9843-a6a7ed08b9fc",
    "name": "view-profile",
    "description": "${role_view-profile}",
    "scopeParamRequired": false,
    "composite": false,
    "clientRole": true,
    // This container ID is the ID of the client that the role "belongs" to, e.g. account
    "containerId": "d063dbba-a4ca-4858-a3f1-d96d0cb72c90"
  }
]
*/

// Get current (per "container" ...):
// GET /auth/admin/realms/TomraConnectUsers/users/e0aef983-c14d-4de7-8ef9-c5722fbbbeac/role-mappings/clients/d063dbba-a4ca-4858-a3f1-d96d0cb72c90/composite
// Response:
/*
[
  {
    "id": "304d9fb1-f4ff-4821-8847-54ea473f45f1",
    "name": "manage-account",
    "description": "${role_manage-account}",
    "scopeParamRequired": false,
    "composite": false,
    "clientRole": true,
    "containerId": "d063dbba-a4ca-4858-a3f1-d96d0cb72c90"
  }
]
*/

// Delete role mapping:
// DELETE /auth/admin/realms/TomraConnectUsers/users/e0aef983-c14d-4de7-8ef9-c5722fbbbeac/role-mappings/clients/d063dbba-a4ca-4858-a3f1-d96d0cb72c90
// Response:
/*
[
  {
    "id": "dba92d03-588a-41ae-9843-a6a7ed08b9fc",
    "name": "view-profile",
    "description": "${role_view-profile}",
    "scopeParamRequired": false,
    "composite": false,
    "clientRole": true,
    "containerId": "d063dbba-a4ca-4858-a3f1-d96d0cb72c90"
  }
]
*/

// Adding client role mapping:
// POST /auth/admin/realms/%s/users/%s/role-mappings/clients/d063dbba-a4ca-4858-a3f1-d96d0cb72c90
// Body:
/*
[
  {
    "id": "dba92d03-588a-41ae-9843-a6a7ed08b9fc",
    "name": "view-profile",
    "description": "${role_view-profile}",
    "scopeParamRequired": false,
    "composite": false,
    "clientRole": true,
    "containerId": "d063dbba-a4ca-4858-a3f1-d96d0cb72c90"
  }
]
*/
// Response: 204
