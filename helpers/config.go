package helpers

import (
	"github.com/goravel-community/goravel-admin/config"
	"github.com/goravel/framework/facades"
)

func GetConfig(key string) any {
	return facades.Config().Get("goravel_admin."+key, config.DefaultGoravelConfig[key])
}

func GetConfigString(key string) string {
	return GetConfig(key).(string)
}

func GetConfigInt(key string) int {
	return GetConfig(key).(int)
}

func GetConfigBool(key string) bool {
	return GetConfig(key).(bool)
}
