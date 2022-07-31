package chaos_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/JYisus/PodChaosMonkey/pkg/chaos"
	"github.com/JYisus/PodChaosMonkey/pkg/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestChaos_Start(t *testing.T) {
	tests := []struct {
		name        string
		namespace   string
		schedule    string
		wantError   bool
		errorString string
	}{
		{
			name:        "Chaos successfully start",
			namespace:   "workloads",
			schedule:    "* * * * * *",
			wantError:   false,
			errorString: "",
		},
		{
			name:        "Chaos start fails on scheduler start",
			namespace:   "workloads",
			schedule:    "* * * * * *",
			wantError:   true,
			errorString: "error adding task to scheduler",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			sch := mocks.NewMockScheduler(ctrl)
			terminatorMock := mocks.NewMockTerminator(ctrl)
			var expectedError error
			if tc.wantError {
				expectedError = fmt.Errorf(tc.errorString)
			}
			ctx := context.Background()

			c := chaos.New(sch, terminatorMock)

			sch.EXPECT().Start(ctx, tc.schedule, gomock.Any()).Return(expectedError)

			err := c.Run(ctx, tc.schedule, tc.namespace)

			if tc.wantError {
				assert.Error(t, err, fmt.Errorf(tc.errorString))
				return
			}
			assert.NoError(t, err)
		})
	}
}
