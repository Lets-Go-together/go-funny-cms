package dispatcher

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "-"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d = New()
			d.RegisterTaskType(TaskTypeCron, func() *TaskResult {
				return nil
			})
			d.Start()
		})
	}
}
