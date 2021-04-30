package schemas

/**
 * @Description: Post请求过来的添加新成员参数
 */
type NewCustomer struct {
	Uid       string `json:"uid" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	QQ        string `json:"qq" binding:"required"`
	ClassName string `json:"className" binding:"required"`
	Profile   string `json:"profile" binding:"required"`
}
