package constants

type Gender int

const (
	Male    Gender = 1
	Female  Gender = 2
	Unknown Gender = 0
)

// 状态
const (
	StatusOk        = 0  //正常
	StatusDeleted   = 1  //删除
	StatusForbidden = -1 //禁用
)

const (
	DefaultJwtExp      = 7     //默认jwt失效时间
	DefaultCacheTime   = 86400 //默认缓存时间
	DefaultSuccessCode = 10000
)
