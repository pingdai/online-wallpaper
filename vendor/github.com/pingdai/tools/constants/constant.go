package constants

// ENV
const (
	// test
	ENV_TEST = "test"
	// pre
	ENV_PRE = "pre"
	// p2p，同bbbank
	ENV_P2P = "p2p"
	// bbbank，同p2p
	ENV_BB_BANK = "bbbank"
	// bank
	ENV_BANK = "bank"
	// x-cros 线上环境
	ENV_PROD = "prod"

	// 单独控制向外输出config信息
	ENV_PRINT_CONFIG = "print_config"

	// 服务名
	EnvVarKeyProjectName = "PROJECT_NAME"
	// 服务环境变量
	EnvVarKeyEnvFlag = "ENV_FLAG"
	// 数据库连接信息
	EnvVarKeyDBMysql = "DB_MYSQL"
	// gin框架环境，线上版本定义为 release
	I_EnvVarKeyGinMode = "GIN_MODE"
	// zk获取配置路径，如 /entry/config/service/athena
	I_EnvVarKeyZKConfigPath = "ZK_CONF_PATH"
	// 配置源标志，1 local，2 remote server
	I_EnvVarKeyConfigSource = "CONF_SOURCE"
	// 配置中心ZK环境变量
	O_EnvVarKeyConfigAddress = "CONF_ADDRESS"
	// 注册中心ZK环境变量
	O_EnvVarKeyRegisterAddress = "REGISTER_ADDRESS"
)

// LOG
const (
	// text
	LOG_FORMAT_TEXT = "text"
	// json
	LOG_FORMAT_JSON = "json"

	LOG_LEVEL_DEBUG = "debug"
	LOG_LEVEL_INFO  = "info"
	LOG_LEVEL_WARN  = "warn"
	LOG_LEVEL_ERROR = "error"
)

// GIN
const (
	// log_id
	HEADER_LOD_ID = "Log-ID"
	// phone
	HEADER_PHONE = "Phone"
	// open_id
	HEADER_OPEN_ID = "Open-ID"
	// user_id
	HEADER_USER_ID = "User-ID"
	// uin
	HEADER_UIN = "Uin"
	// remote_service
	HEADER_REMOTE_SERVICE = "Remote-Service"
	// merchant_code
	HEADER_MERCHANT_CODE = "Merchant-Code"
)

// MQ类型
type MQType int

const (
	// 生产者
	MQ_TYPE_PRODUCER MQType = 1
	// 消费者
	MQ_TYPE_CONSUMER MQType = 2
)
