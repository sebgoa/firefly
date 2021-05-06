// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package persistencemocks

import (
	context "context"

	fftypes "github.com/kaleido-io/firefly/internal/fftypes"
	mock "github.com/stretchr/testify/mock"

	persistence "github.com/kaleido-io/firefly/internal/persistence"

	uuid "github.com/google/uuid"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// Capabilities provides a mock function with given fields:
func (_m *Plugin) Capabilities() *persistence.Capabilities {
	ret := _m.Called()

	var r0 *persistence.Capabilities
	if rf, ok := ret.Get(0).(func() *persistence.Capabilities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.Capabilities)
		}
	}

	return r0
}

// ConfigInterface provides a mock function with given fields:
func (_m *Plugin) ConfigInterface() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetBatchById provides a mock function with given fields: ctx, ns, id
func (_m *Plugin) GetBatchById(ctx context.Context, ns string, id *uuid.UUID) (*fftypes.Batch, error) {
	ret := _m.Called(ctx, ns, id)

	var r0 *fftypes.Batch
	if rf, ok := ret.Get(0).(func(context.Context, string, *uuid.UUID) *fftypes.Batch); ok {
		r0 = rf(ctx, ns, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *uuid.UUID) error); ok {
		r1 = rf(ctx, ns, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatches provides a mock function with given fields: ctx, skip, limit, filter
func (_m *Plugin) GetBatches(ctx context.Context, skip uint64, limit uint64, filter persistence.Filter) ([]*fftypes.Batch, error) {
	ret := _m.Called(ctx, skip, limit, filter)

	var r0 []*fftypes.Batch
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, persistence.Filter) []*fftypes.Batch); ok {
		r0 = rf(ctx, skip, limit, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, persistence.Filter) error); ok {
		r1 = rf(ctx, skip, limit, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetData provides a mock function with given fields: ctx, skip, limit, filter
func (_m *Plugin) GetData(ctx context.Context, skip uint64, limit uint64, filter persistence.Filter) ([]*fftypes.Data, error) {
	ret := _m.Called(ctx, skip, limit, filter)

	var r0 []*fftypes.Data
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, persistence.Filter) []*fftypes.Data); ok {
		r0 = rf(ctx, skip, limit, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Data)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, persistence.Filter) error); ok {
		r1 = rf(ctx, skip, limit, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataById provides a mock function with given fields: ctx, ns, id
func (_m *Plugin) GetDataById(ctx context.Context, ns string, id *uuid.UUID) (*fftypes.Data, error) {
	ret := _m.Called(ctx, ns, id)

	var r0 *fftypes.Data
	if rf, ok := ret.Get(0).(func(context.Context, string, *uuid.UUID) *fftypes.Data); ok {
		r0 = rf(ctx, ns, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Data)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *uuid.UUID) error); ok {
		r1 = rf(ctx, ns, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessageById provides a mock function with given fields: ctx, ns, id
func (_m *Plugin) GetMessageById(ctx context.Context, ns string, id *uuid.UUID) (*fftypes.Message, error) {
	ret := _m.Called(ctx, ns, id)

	var r0 *fftypes.Message
	if rf, ok := ret.Get(0).(func(context.Context, string, *uuid.UUID) *fftypes.Message); ok {
		r0 = rf(ctx, ns, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *uuid.UUID) error); ok {
		r1 = rf(ctx, ns, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessages provides a mock function with given fields: ctx, skip, limit, filter
func (_m *Plugin) GetMessages(ctx context.Context, skip uint64, limit uint64, filter persistence.Filter) ([]*fftypes.Message, error) {
	ret := _m.Called(ctx, skip, limit, filter)

	var r0 []*fftypes.Message
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, persistence.Filter) []*fftypes.Message); ok {
		r0 = rf(ctx, skip, limit, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, persistence.Filter) error); ok {
		r1 = rf(ctx, skip, limit, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchemaById provides a mock function with given fields: ctx, ns, id
func (_m *Plugin) GetSchemaById(ctx context.Context, ns string, id *uuid.UUID) (*fftypes.Schema, error) {
	ret := _m.Called(ctx, ns, id)

	var r0 *fftypes.Schema
	if rf, ok := ret.Get(0).(func(context.Context, string, *uuid.UUID) *fftypes.Schema); ok {
		r0 = rf(ctx, ns, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Schema)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *uuid.UUID) error); ok {
		r1 = rf(ctx, ns, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSchemas provides a mock function with given fields: ctx, skip, limit, filter
func (_m *Plugin) GetSchemas(ctx context.Context, skip uint64, limit uint64, filter persistence.Filter) ([]*fftypes.Schema, error) {
	ret := _m.Called(ctx, skip, limit, filter)

	var r0 []*fftypes.Schema
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, persistence.Filter) []*fftypes.Schema); ok {
		r0 = rf(ctx, skip, limit, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Schema)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, persistence.Filter) error); ok {
		r1 = rf(ctx, skip, limit, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionById provides a mock function with given fields: ctx, ns, id
func (_m *Plugin) GetTransactionById(ctx context.Context, ns string, id *uuid.UUID) (*fftypes.Transaction, error) {
	ret := _m.Called(ctx, ns, id)

	var r0 *fftypes.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, string, *uuid.UUID) *fftypes.Transaction); ok {
		r0 = rf(ctx, ns, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *uuid.UUID) error); ok {
		r1 = rf(ctx, ns, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactions provides a mock function with given fields: ctx, skip, limit, filter
func (_m *Plugin) GetTransactions(ctx context.Context, skip uint64, limit uint64, filter persistence.Filter) ([]*fftypes.Transaction, error) {
	ret := _m.Called(ctx, skip, limit, filter)

	var r0 []*fftypes.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, persistence.Filter) []*fftypes.Transaction); ok {
		r0 = rf(ctx, skip, limit, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, persistence.Filter) error); ok {
		r1 = rf(ctx, skip, limit, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: ctx, config, events
func (_m *Plugin) Init(ctx context.Context, config interface{}, events persistence.Events) error {
	ret := _m.Called(ctx, config, events)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, persistence.Events) error); ok {
		r0 = rf(ctx, config, events)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertBatch provides a mock function with given fields: ctx, data
func (_m *Plugin) UpsertBatch(ctx context.Context, data *fftypes.Batch) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Batch) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertData provides a mock function with given fields: ctx, data
func (_m *Plugin) UpsertData(ctx context.Context, data *fftypes.Data) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Data) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertMessage provides a mock function with given fields: ctx, message
func (_m *Plugin) UpsertMessage(ctx context.Context, message *fftypes.Message) error {
	ret := _m.Called(ctx, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Message) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertSchema provides a mock function with given fields: ctx, data
func (_m *Plugin) UpsertSchema(ctx context.Context, data *fftypes.Schema) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Schema) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertTransaction provides a mock function with given fields: ctx, data
func (_m *Plugin) UpsertTransaction(ctx context.Context, data *fftypes.Transaction) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Transaction) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
