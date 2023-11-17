package mockete_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jcleira/mockete/mocks"
)

type testInterface interface {
	DoSomething(context.Context, string) (string, error)
	FinishSomething(context.Context, string) error
}

type Client struct {
	testInterface
}

func (c *Client) Perform(ctx context.Context, arg0 string) error {
	value, err := c.testInterface.DoSomething(ctx, arg0)
	if err != nil {
		return err
	}

	if err = c.testInterface.FinishSomething(ctx, value); err != nil {
		return err
	}

	return nil
}

func TestSomeInterface(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name            string
		arg0            string
		DoSomething     *mocks.DoSomethingMock
		FinishSomething *mocks.FinishSomethingMock
		wantError       error
	}{
		{
			name:            "test1",
			arg0:            "test1",
			DoSomething:     mocks.DoSomething(t, ctx, "test1").Return("test1", nil),
			FinishSomething: mocks.FinishSomething(t, ctx, "test1").Return(nil),
			wantError:       nil,
		},
		{
			name:            "test2",
			arg0:            "test2",
			DoSomething:     mocks.DoSomething(t, ctx, "test2").Return("", errors.New("error")),
			FinishSomething: mocks.FinishSomethingNotCalled(t),
			wantError:       errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				testInterface: mocks.NewMockTestInterface(
					tt.DoSomething, tt.FinishSomething,
				),
			}

			err := client.Perform(ctx, tt.name)

			if err == nil && tt.wantError != nil {
				t.Errorf("Perform() got = %v, want %v", err, tt.wantError)
			}

			if err != nil && err.Error() != tt.wantError.Error() {
				t.Errorf("Perform() got = %v, want %v", err, tt.wantError)
			}

		})
	}
}
