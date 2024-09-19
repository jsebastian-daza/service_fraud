package interfaces

// GetSecret retrieves the secret associated with the given name.
// Returns the secret as a string pointer and any error encountered.
type SecretsVault interface {
	GetSecret(name string) (*string, error)
}
