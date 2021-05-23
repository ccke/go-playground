package services

import (
	"encoding/base64"
	"errors"
	"github.com/ccke/go-playground/study_02/homework_01/models"
	"net/url"
	"reflect"
)

/**
 * @Description: kms服务
 */
type KmsServer struct{}

/**
 * @Description: 构造函数
 * @return *KmsServer
 */
func NewKmsServer() *KmsServer {
	return &KmsServer{}
}

/**
 * @Description: 加密实现
 * @receiver kms
 * @param data
 * @return error
 */
func (kms *KmsServer) Encrypt(data interface{}) error {
	//要求
	//1.通过反射获取每个字段的编码方式
	//2.对加密数据分片，开多个协程进行加密
	//3.将多个协程加密的结果汇总
	//4.将汇总的数据通过反射设置回data

	newData, ok := data.([]models.Employee)
	if !ok {
		return errors.New("参数类型错误")
	}

	for i, item := range newData {
		t := reflect.TypeOf(item)
		v := reflect.ValueOf(&item).Elem()
		for j := 0; j < t.NumField(); j++ {
			field := t.Field(j)
			tag := field.Tag.Get("kms")
			switch tag {
			case "type=urlEncode":
				str := v.Field(j).String()
				v.Field(j).SetString(url.QueryEscape(str))
			case "type=base64":
				str := v.Field(j).String()
				v.Field(j).SetString(base64.StdEncoding.EncodeToString([]byte(str)))
			}
		}
		newData[i] = item
	}
	return nil
}

/**
 * @Description: 解密实现
 * @receiver kms
 * @param data
 * @return error
 */
func (kms *KmsServer) Decrypt(data interface{}) error {
	//要求
	//1.通过反射获取每个字段的编码方式
	//2.对解密数据分片，开多个协程进行解密
	//3.将多个协程解密的结果汇总
	//4.将汇总的数据通过反射设置回data

	newData, ok := data.([]models.Employee)
	if !ok {
		return errors.New("参数类型错误")
	}

	for i, item := range newData {
		t := reflect.TypeOf(item)
		v := reflect.ValueOf(&item).Elem()
		for j := 0; j < t.NumField(); j++ {
			field := t.Field(j)
			tag := field.Tag.Get("kms")
			switch tag {
			case "type=urlEncode":
				str, error := url.QueryUnescape(v.Field(j).String())
				if error != nil {
					errors.New("urlEncode 解码失败")
				}
				v.Field(j).SetString(str)
			case "type=base64":
				str, error := base64.StdEncoding.DecodeString(v.Field(j).String())
				if error != nil {
					errors.New("base64 解码失败")
				}
				v.Field(j).SetString(string(str))
			}
		}
		newData[i] = item
	}
	return nil
}
