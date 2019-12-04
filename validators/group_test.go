package validators

import "testing"

func TestGroupName(t *testing.T) {
	tests := []struct {
		name     string
		wantPass bool
		args     string
	}{
		{"太短情况", false, "min"},
		{"正常情况", true, "yinxulai"},
		{"特殊字符", false, "yinxulai123A@"},
		{"特殊字符", false, "yinxulai123A@*&"},
		{"太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := GroupName(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("GroupName() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func TestGroupClass(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"太短情况", false, "min"},
		{"正常情况", true, "yinxulai123A"},
		{"特殊字符", true, "yinxulai123A"},
		{"特殊字符", false, "yinxulai123A@"},
		{"特殊字符", false, "yinxulai123A@*&"},
		{"太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := GroupClass(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("GroupClass() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func TestGroupState(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"太短情况", false, "min"},
		{"正常情况", true, "yinxulai123A"},
		{"特殊字符", true, "yinxulai123A"},
		{"特殊字符", false, "yinxulai123A@"},
		{"特殊字符", false, "yinxulai123A@*&"},
		{"太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := GroupState(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("GroupState() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func TestGroupDescription(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"太短情况", false, "min"},
		{"正常情况", true, "yinxulai"},
		{"太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxuyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulailaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := GroupDescription(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("GroupDescription() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
