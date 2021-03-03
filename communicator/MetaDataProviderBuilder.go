package communicator

import (
	"github.com/processout/connect-sdk-go/communicator/communication"
	"github.com/processout/connect-sdk-go/domain/metadata"
)

// MetaDataProviderBuilder represents a builder for a MetaDataProvider object.
type MetaDataProviderBuilder struct {
	integrator               string
	ShoppingCartExtension    *metadata.ShoppingCartExtension
	AdditionalRequestHeaders []communication.Header
}

// NewMetaDataProviderBuilder creates a MetaDataProviderBuilder with the given Integrator
func NewMetaDataProviderBuilder(integrator string) *MetaDataProviderBuilder {
	return &MetaDataProviderBuilder{integrator, nil, nil}
}

// Build creates a fully initialized MetaDataProvider
func (m *MetaDataProviderBuilder) Build() (*MetaDataProvider, error) {
	return NewMetaDataProviderWithBuilder(*m)
}
