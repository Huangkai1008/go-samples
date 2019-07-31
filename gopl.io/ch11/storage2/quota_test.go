package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	saved := notifyUser
	defer func() { notifyUser = saved }()

	var notifiedUser, notifiedMsg string
	notifyUser = func(username, msg string) {
		notifiedUser, notifiedMsg = username, msg
	}

	const user = "joe@example.org"
	usage[user] = 980 * 1000 * 1000 // 模拟980MB已使用的情况
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Errorf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}
