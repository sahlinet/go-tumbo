package worker

import "testing"

func TestWorker_Start(t *testing.T) {
	type fields struct {
		Running bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "huhu",
			fields: fields{true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				Running: tt.fields.Running,
			}
			if err := w.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Worker.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
