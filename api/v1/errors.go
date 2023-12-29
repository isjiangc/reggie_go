package v1

var (
	// common errors
	ErrSuccess             = newError(1, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrEmailAlreadyUse = newError(1001, "The email is already in use.")

	// ErrEmployeeLoginFailed 登录失败
	ErrEmployeeLoginFailed = newError(2001, "Login failed")
	// ErrEmployeeAccountIsDisabled 账号已禁止使用
	ErrEmployeeAccountIsDisabled   = newError(2002, "account is disabled")
	ErrEmployeeAccountAlreadyExist = newError(2003, "The employee is already already exist")
	ErrEmployeeUpdatedFailed       = newError(2004, "update employee failed")
	ErrEmployeeNotExit             = newError(2005, "The employee is not exist")

	// Category
	ErrCreateCategoryFailed       = newError(3001, "save category failed")
	ErrDeleteCategoryFailed       = newError(3002, "delete category failed")
	ErrDeleteCategoryIdIsNotEmpty = newError(3003, "delete category id not empty")
	ErrCategoryHaveSomeDish       = newError(3004, "当前分类下关联了菜品,不能删除")
	ErrCategoryTypeIsIllegal      = newError(3005, "分类类型不合法")
	ErrUpdateCategoryFailed       = newError(3006, "更新分类信息失败")

	ErrDishNotExit = newError(4001, "菜品不存在")
	AddDishIsError = newError(4002, "新增菜品失败")

	ErrDishFlavorNotExit = newError(5001, "菜品口味不存在")
)
