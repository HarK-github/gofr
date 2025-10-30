package migrations

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr/migration"
)

// MockPubSub implements the PubSub interface for testing
type MockPubSub struct {
	Calls      []string
	ErrOnTopic map[string]errorpackage migrations

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr/migration"
)

// MockPubSub implements the PubSub interface for testing
type MockPubSub struct {
	Calls      []string
	ErrOnTopic map[string]error
}

func (*MockPubSub) Query(_ context.Context, _ string, _ ...any) ([]byte, error) {
	return []byte{}, nil
}

func (*MockPubSub) DeleteTopic(_ context.Context, _ string) error {
	return nil
}

func (m *MockPubSub) CreateTopic(_ context.Context, topic string) error {
	m.Calls = append(m.Calls, topic)
	if err, ok := m.ErrOnTopic[topic]; ok {
		return err
	}
	return nil
}

func TestCreateTopics(t *testing.T) {
	tests := []struct {
		name          string
		pubsub        migration.PubSub
		errOnProduct  error
		errOnOrder    error
		expectedErr   error
		expectedCalls []string
	}{
		{"success", &MockPubSub{}, nil, nil, nil, []string{"products", "order-logs"}},
		{"error on products", &MockPubSub{
			ErrOnTopic: map[string]error{"products": errors.New("fail products")},
		}, errors.New("fail products"), nil, errors.New("fail products"), []string{"products"}},
		{"error on order-logs", &MockPubSub{
			ErrOnTopic: map[string]error{"order-logs": errors.New("fail order")},
		}, nil, errors.New("fail order"), errors.New("fail order"), []string{"products", "order-logs"}},

		// ❌ New failing test: PubSub is nil
		{
			name:          "nil PubSub should fail",
			pubsub:        nil,
			expectedErr:   errors.New("pubsub not configured"),
			expectedCalls: nil,
		},
	}

	for _, tt := range tests {
		ds := migration.Datasource{PubSub: tt.pubsub}

		err := createTopics().UP(ds)

		if tt.name == "nil PubSub should fail" {
			assert.Error(t, err, tt.name)
			assert.Contains(t, err.Error(), "pubsub", tt.name)
			continue
		}

		assert.Equal(t, tt.expectedErr, err, tt.name)
		if m, ok := tt.pubsub.(*MockPubSub); ok {
			assert.Equal(t, tt.expectedCalls, m.Calls, tt.name)
		}
	}
}
package migrations

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr/migration"
)

// MockPubSub implements the PubSub interface for testing
type MockPubSub struct {
	Calls      []string
	ErrOnTopic map[string]error
}

func (*MockPubSub) Query(_ context.Context, _ string, _ ...any) ([]byte, error) {
	return []byte{}, nil
}

func (*MockPubSub) DeleteTopic(_ context.Context, _ string) error {
	return nil
}

func (m *MockPubSub) CreateTopic(_ context.Context, topic string) error {
	m.Calls = append(m.Calls, topic)
	if err, ok := m.ErrOnTopic[topic]; ok {
		return err
	}
	return nil
}

func TestCreateTopics(t *testing.T) {
	tests := []struct {
		name          string
		pubsub        migration.PubSub
		errOnProduct  error
		errOnOrder    error
		expectedErr   error
		expectedCalls []string
	}{
		{"success", &MockPubSub{}, nil, nil, nil, []string{"products", "order-logs"}},
		{"error on products", &MockPubSub{
			ErrOnTopic: map[string]error{"products": errors.New("fail products")},
		}, errors.New("fail products"), nil, errors.New("fail products"), []string{"products"}},
		{"error on order-logs", &MockPubSub{
			ErrOnTopic: map[string]error{"order-logs": errors.New("fail order")},
		}, nil, errors.New("fail order"), errors.New("fail order"), []string{"products", "order-logs"}},

		// ❌ New failing test: PubSub is nil
		{
			name:          "nil PubSub should fail",
			pubsub:        nil,
			expectedErr:   errors.New("pubsub not configured"),
			expectedCalls: nil,
		},
	}

	for _, tt := range tests {
		ds := migration.Datasource{PubSub: tt.pubsub}

		err := createTopics().UP(ds)

		if tt.name == "nil PubSub should fail" {
			assert.Error(t, err, tt.name)
			assert.Contains(t, err.Error(), "pubsub", tt.name)
			continue
		}

		assert.Equal(t, tt.expectedErr, err, tt.name)
		if m, ok := tt.pubsub.(*MockPubSub); ok {
			assert.Equal(t, tt.expectedCalls, m.Calls, tt.name)
		}
	}
}

}

func (*MockPubSub) Query(_ context.Context, _ string, _ ...any) ([]byte, error) {
	return []byte{}, nil
}

func (*MockPubSub) DeleteTopic(_ context.Context, _ string) error {
	return nil
}

func (m *MockPubSub) CreateTopic(_ context.Context, topic string) error {
	m.Calls = append(m.Calls, topic)
	if err, ok := m.ErrOnTopic[topic]; ok {
		return err
	}
	return nil
}

func TestCreateTopics(t *testing.T) {
	tests := []struct {
		name          string
		errOnProduct  error
		errOnOrder    error
		expectedErr   error
		expectedCalls []string
	}{
		{"success", nil, nil, nil, []string{"products", "order-logs"}},
		{"error on products", errors.New("fail products"), nil,
			errors.New("fail products"), []string{"products"}},
		{"error on order-logs", nil, errors.New("fail order"), errors.New("fail order"), []string{"products", "order-logs"}},
	}

	for _, tt := range tests {
		mockPubSub := &MockPubSub{
			ErrOnTopic: map[string]error{
				"products":   tt.errOnProduct,
				"order-logs": tt.errOnOrder,
			},
		}
		ds := migration.Datasource{PubSub: mockPubSub}

		err := createTopics().UP(ds)

		assert.Equal(t, tt.expectedErr, err, tt.name)
		assert.Equal(t, tt.expectedCalls, mockPubSub.Calls, tt.name)
	}
}
