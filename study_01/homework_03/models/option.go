package models

// Option 期权授予信息
type Option struct {
	Id            int
	CompanyId     int
	EmployeeId    int
	Code          string // 授予编号
	Num           int    // 已归属数量
	UnExerciseNum int    // 待行权数量
}
