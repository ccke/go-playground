package models

type Employee struct {
	// 员工id
	Id int
	// 公司id
	CompanyId int
	// 员工名字
	Name string `kms:"type=urlEncode"` //该字段采用urlEncode进行编码
	// uid
	Uid string `kms:"type=base64"` //该字段采用base64进行编码
	// id_code
	IdCode string `kms:"type=base64"` //该字段采用base64进行编码
}
