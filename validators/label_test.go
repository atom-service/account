package validators

import "testing"

func TestLabelName(t *testing.T) {
	tests := []struct {
		name     string
		wantPass bool
		args     string
	}{
		{"标签名格式测试：太短情况", false, "min"},
		{"标签名格式测试：正常情况", true, "yinxulai"},
		{"标签名格式测试：特殊字符", false, "yinxulai123A@"},
		{"标签名格式测试：特殊字符", false, "yinxulai123A@*&"},
		{"标签名格式测试：太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := LabelName(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("LabelName() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func TestLabelClass(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"标签类别格式测试：太短情况", false, "min"},
		{"标签类别格式测试：正常情况", true, "yinxulai123A"},
		{"标签类别格式测试：特殊字符", true, "yinxulai123A"},
		{"标签类别格式测试：特殊字符", false, "yinxulai123A@"},
		{"标签类别格式测试：特殊字符", false, "yinxulai123A@*&"},
		{"标签类别格式测试：太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := LabelClass(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("LabelClass() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func TestLabelState(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"标签类别格式测试：太短情况", false, "min"},
		{"标签类别格式测试：正常情况", true, "yinxulai123A"},
		{"标签类别格式测试：特殊字符", true, "yinxulai123A"},
		{"标签类别格式测试：特殊字符", false, "yinxulai123A@"},
		{"标签类别格式测试：特殊字符", false, "yinxulai123A@*&"},
		{"标签状态格式测试：太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := LabelState(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("LabelState() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func TestLabelValue(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"标签值格式测试：太短情况", false, "min"},
		{"标签值格式测试：正常情况", true, "yinxulai"},
		{"标签值格式测试：太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxuyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulailaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := LabelValue(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("LabelValue() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
