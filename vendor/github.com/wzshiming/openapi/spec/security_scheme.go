package spec

// securityScheme an API key (either as a header or as a query parameter), OAuth2's common flows (implicit, password, application and access code) as defined in RFC6749, and OpenID Connect Discovery.
type securityScheme struct {

	// Any	REQUIRED.
	// The type of the security scheme.
	// Valid values are "apiKey", "http", "oauth2", "openIdConnect".
	Type string `json:"type"`

	// Any	A short description for security scheme.
	// CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// apiKey	REQUIRED.
	// The name of the header, query or cookie parameter to be used.
	Name string `json:"name,omitempty"`

	// apiKey	REQUIRED.
	// The location of the API key.
	// Valid values are "query", "header" or "cookie".
	In string `json:"in,omitempty"`

	// http	REQUIRED.
	// The name of the HTTP Authorization scheme to be used in the Authorization header as defined in RFC7235.
	Scheme string `json:"scheme,omitempty"`

	// http ("bearer")	A hint to the client to identify how the bearer token is formatted.
	// Bearer tokens are usually generated by an authorization server, so this information is primarily for documentation purposes.
	BearerFormat string `json:"bearerFormat,omitempty"`

	// oauth2	REQUIRED.
	// An object containing configuration information for the flow types supported.
	Flows *OAuthFlows `json:"flows,omitempty"`

	// openIdConnect	REQUIRED.
	// OpenId Connect URL to discover OAuth2 configuration values.
	// This MUST be in the form of a URL.
	OpenIDConnectURL string `json:"openIdConnectUrl,omitempty"`
}