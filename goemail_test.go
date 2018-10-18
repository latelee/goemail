/*
测试：
go test 
*/
package goemail

import (
    "testing"
)

func TestEmail(t *testing.T) {
    serverHost := "smtp.exmail.qq.com"
    serverPort := 465
    fromEmail := "test@latelee.org"
    fromPasswd := "1qaz@WSX"
    
    myToers := "li@latelee.org, latelee@163.com" // 逗号隔开
    myCCers := "" //"readchy@163.com"
    
    subject := "这是主题"
    body := `这是正文<br>
            <h3>这是标题</h3>
             Hello <a href = "http://www.latelee.org">主页</a><br>`
    // 结构体赋值
    myEmail := &EmailParam {
        ServerHost: serverHost,
        ServerPort: serverPort,
        FromEmail:  fromEmail,
        FromPasswd: fromPasswd,
        Toers:      myToers,
        CCers:      myCCers,
    }
    t.Logf("init email.\n")
    InitEmail(myEmail)
    SendEmail(subject, body)
}