package index

import (
	"encoding/base64"
	"fmt"
	"github.com/TeaOSLab/EdgeAdmin/internal/encrypt"
	"github.com/TeaOSLab/EdgeAdmin/internal/ttlcache"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/actions"
	"strconv"
	"strings"
	"time"
)

// 检查是否需要OTP
type RenewalAction struct {
	actionutils.ParentAction
}

func (this *RenewalAction) Init() {
	this.Nav("", "", "")
}
func (this *RenewalAction) RunPost(params struct {
	Secret string

	Must *actions.Must
}) {
	params.Must.Field("secret", params.Secret).Require("请输入续订密钥")

	//判断redis 是否存在改secret
	if ttlcache.DefaultCache.Read(params.Secret) != nil {
		this.ErrorPage(fmt.Errorf("密钥已失效，请不要重复续订"))
		return
	}
	decode,err := base64.StdEncoding.DecodeString(params.Secret)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	secret := encrypt.MagicKeyDecode(decode)
	//code,add_time,time
	secrets := strings.Split(string(secret), ",")
	if len(secrets) != 3 {
		// 日志
		this.CreateLogInfo("系统续订失败 - 无效的密钥:%v", string(secret))
		this.ErrorPage(fmt.Errorf("无效的密钥"))
		return
	}
	config, err := loadServerExpireConfig()
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//未开启 直接通过
	if !config.Expire.On {
		this.Success()
		return
	}
	if config.Expire.Code != secrets[0] {
		// 日志
		this.CreateLogInfo("系统续订失败 - 不支持此密钥:%v != %v", secrets[0], config.Expire.Code)
		this.ErrorPage(fmt.Errorf("不支持此密钥"))
		return
	}
	timeout, _ := strconv.ParseInt(secrets[2], 10, 64)
	if timeout < time.Now().Unix() {
		this.ErrorPage(fmt.Errorf("密钥已失效"))
		return
	}
	renewalT, _ := strconv.ParseInt(secrets[1], 10, 64)
	//续订生效
	if config.Expire.Time == "" {
		expireT := time.Now().Unix() + renewalT
		config.Expire.Time = fmt.Sprintf("%v", expireT)
	} else {
		expireT, err := strconv.ParseInt(config.Expire.Time, 10, 64)
		if err != nil {
			this.ErrorPage(err)
			return
		}
		config.Expire.Time = fmt.Sprintf("%v", expireT+renewalT)
	}
	err = writeServerConfig(config)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	//将secret 写入数据库防止重复续订
	ttlcache.DefaultCache.Write(params.Secret,true,timeout - time.Now().Unix())
	this.Success()
}
