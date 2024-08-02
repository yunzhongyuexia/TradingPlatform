package tools

import (
	"fmt"
	"github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/spf13/viper"
	"golang.org/x/exp/rand"
	"strings"
	"time"
)

func SmsVerify(phone string) string {
	accessKeyId := viper.GetString("sms.accessKeyId")
	accessKeySecret := viper.GetString("sms.accessKeySecret")
	endpoint := viper.GetString("sms.endpoint")

	c := &client.Config{AccessKeyId: &accessKeyId, AccessKeySecret: &accessKeySecret, Endpoint: &endpoint}

	newClient, err := dysmsapi20170525.NewClient(c)
	if err != nil {
		panic(err)
	}
	phoneNumber := phone
	templateCode := "SMS_154950909"
	signName := "阿里云短信测试"
	num := GenerateSmsCode(6)
	code := "{\"code\":" + num + "}"
	request := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &phoneNumber,
		TemplateCode:  &templateCode,
		SignName:      &signName,
		TemplateParam: &code,
	}
	sms, err := newClient.SendSms(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(sms)
	return num
}

// GenerateSmsCode 生成验证码;length代表验证码的长度
func GenerateSmsCode(length int) string {
	number := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(uint64(time.Now().Unix()))
	var sb strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&sb, "%d", number[rand.Intn(len(number))])
	}
	return sb.String()
}
