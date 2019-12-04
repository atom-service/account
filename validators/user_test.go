package validators

import "testing"

func TestNickname(t *testing.T) {
	tests := []struct {
		name string

		wantPass bool
		args     string
	}{
		{"太短情况", false, "min"},
		{"正常情况", true, "yinxulai"},
		{"太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
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
		{"太短情况", false, "min"},
		{"正常情况", true, "yinxulai123A"},
		{"特殊字符", true, "yinxulai123A"},
		{"特殊字符", false, "yinxulai123A@"},
		{"特殊字符", false, "yinxulai123A@*&"},
		{"太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
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
		{"太短情况", false, "min"},
		{"正常情况", true, "yinxulai"},
		{"特殊字符", false, "<=>|23333"},
		{"特殊字符", true, "!#$%&()*+,-./:;?@[]"},
		{"太长情况", false, "yinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinxulaiyinlaiyxulaiyinlaiyinxulxulaiyinlaiyinxulinxulaiyinxulai"},
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
