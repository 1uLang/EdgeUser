package remotelogs

import (
	_ "github.com/iwind/TeaGo/bootstrap"
	"testing"
)

func TestPrintln(t *testing.T) {
	Println("test", "123")
	err := uploadLogs()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
