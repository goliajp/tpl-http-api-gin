package env

import (
	"github.com/goliajp/envx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// app config
var (
	HttpListenAddr        = envx.Get("http_listen_addr", "0.0.0.0")
	HttpListenPort        = envx.Get("http_listen_port", 60001)
	HttpReadTimeout       = envx.Get("http_read_timeout", time.Second*120)
	HttpReadHeaderTimeout = envx.Get("http_read_header_timeout", time.Second*120)
	HttpWriteTimeout      = envx.Get("http_write_timeout", time.Second*120)
	HttpIdleTimeout       = envx.Get("http_idle_timeout", time.Second*120)
)

// business config
var (
	CryptCost   = envx.Get("crypt_cost", bcrypt.DefaultCost)
	ExpireHours = envx.Get("expire_hours", 24*30) // for user session
)

// relational database config
var (
	RdHost     = envx.Get("rd_host", "localhost")
	RdPort     = envx.Get("rd_port", 5432)
	RdTz       = envx.Get("rd_tz", "Asia/Shanghai")
	RdUser     = envx.Get("rd_user", "postgres")
	RdPassword = envx.Get("rd_password", "postgres")
	RdName     = envx.Get("rd_name", "demo")
	RdRebuild  = envx.Get("rd_rebuild", true)
	RdMockData = envx.Get("rd_mock_data", true)
)

// key-value database config
var (
	KvHost     = envx.Get("kv_host", "localhost")
	KvPort     = envx.Get("kv_port", 6379)
	KvPassword = envx.Get("kv_password", "")
	KvDb       = envx.Get("kv_db", 0)
	KvRebuild  = envx.Get("kv_rebuild", false)
)
