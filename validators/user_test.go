package validators

import "testing"

func TestNickname(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"用户昵称格式测试：太短情况", false, "min"},
		{"用户昵称格式测试：正常情况", true, "yinxulai"},
		{"用户昵称格式测试：太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := Nickname(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("Nickname() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func TestUsername(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"用户名格式测试：太短情况", false, "min"},
		{"用户名格式测试：正常情况", true, "yinxulai123A"},
		{"用户名格式测试：特殊字符", true, "yinxulai123A"},
		{"用户名格式测试：特殊字符", false, "yinxulai123A@"},
		{"用户名格式测试：特殊字符", false, "yinxulai123A@*&"},
		{"用户名格式测试：太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := Username(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("Username() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}

func TestPassword(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"用户密码格式测试：太短情况", false, "min"},
		{"用户密码格式测试：正常情况", true, "yinxulai"},
		{"用户密码格式测试：特殊字符", false, "<=>|23333"},
		{"用户密码格式测试：特殊字符", true, "!#$%&()*+,-./:;?@[]"},
		{"用户密码格式测试：太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPass, _ := Password(tt.args)
			if gotPass != tt.wantPass {
				t.Errorf("Password() gotPass = %v, want %v", gotPass, tt.wantPass)
			}
		})
	}
}
