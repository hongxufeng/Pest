package db

import (
	"model"
	"utils/function"

	"github.com/go-redis/redis"
)

const (
	CACHE_TABLE_SELECTOR_BAR = "W_Redis_Cache_Table_Selector_Bar_Map"
	CACHE_TABLE_SELECTOR     = "W_Redis_Cache_Table_Selector_Existence_bool"
	CACHE_XML_Config         = "W_Redis_Cache_Xml_Config_Byte"
	CACHE_TABLE              = "W_Redis_Cache_Table_String"
	CACHE_TABLE_UPDATE_FALG  = "W_Redis_Cache_Table_Update_Bool"
)

func SetXmlConfigCache(table string, xmlConfig []byte) (err error) {
	err = RedisCache.Set(function.MakeKey(CACHE_XML_Config, table), xmlConfig, model.Xml_Config_Cache_Duration).Err()
	return
}
func GetXmlConfigCache(table string) (xmlConfig []byte, err error) {
	xmlConfig, err = RedisCache.Get(function.MakeKey(CACHE_XML_Config, table)).Bytes()
	return
}

//Persistence用于判断是否更新SelectorBarCache
func SetSelectorCachePersistence(table string, selectorfeild string) (err error) {
	err = RedisCache.Set(function.MakeKey(CACHE_TABLE_SELECTOR, table, selectorfeild), true, model.User_Info_Persistence_Duration).Err()
	return
}

func GetSelectorCachePersistence(table string, selectorfeild string) (being bool, err error) {
	_, err = RedisCache.Get(function.MakeKey(CACHE_TABLE_SELECTOR, table, selectorfeild)).Result()
	if err == redis.Nil {
		being = false
	} else if err != nil {
		being = false
	} else {
		being = true
	}
	return
}

func HSetSelectorBarCache(table string, selectorfeild string, key string, value string) (err error) {
	err = RedisCache.HSet(function.MakeKey(CACHE_TABLE_SELECTOR_BAR, table, selectorfeild), key, value).Err()
	return
}

func HGetSelectorBarCache(table string, selectorfeild string) (selectordata map[string]string, err error) {
	selectordata, _ = RedisCache.HGetAll(function.MakeKey(CACHE_TABLE_SELECTOR_BAR, table, selectorfeild)).Result()
	return
}

func SetTableCache(settings *model.Settings, url string, feild string, value string) (err error) {
	key := function.MakeKey(CACHE_TABLE, settings.ConfigFile, settings.Style, settings.TableID, settings.HasCheckbox, settings.RowList, settings.Condition, settings.Page, settings.Rows, settings.ColPage, settings.Order, url, feild)
	err = RedisCache.Set(key, value, model.Table_Cache_Duration).Err()
	return
}
func GetTableCache(settings *model.Settings, url string, feild string) (value string, err error) {
	key := function.MakeKey(CACHE_TABLE, settings.ConfigFile, settings.Style, settings.TableID, settings.HasCheckbox, settings.RowList, settings.Condition, settings.Page, settings.Rows, settings.ColPage, settings.Order, url, feild)
	value, err = RedisCache.Get(key).Result()
	return
}

func SetTableJsonCache(settings *model.Settings, url string, feild string, value string) (err error) {
	key := function.MakeKey(CACHE_TABLE, settings.ConfigFile, "json", settings.Style, settings.TableID, settings.Page, settings.Rows, settings.Order, url, feild)
	err = RedisCache.Set(key, value, model.Table_Cache_Duration).Err()
	return
}
func GetTableJsonCache(settings *model.Settings, url string, feild string) (value string, err error) {
	key := function.MakeKey(CACHE_TABLE, settings.ConfigFile, "json", settings.Style, settings.TableID, settings.Page, settings.Rows, settings.Order, url, feild)
	value, err = RedisCache.Get(key).Result()
	return
}

func DelTableCache(table string) (err error) {
	cmd := RedisCache.Scan(0, function.MakeKey(CACHE_TABLE, table, "*"), 10000)
	it := cmd.Iterator()
	for it.Next() {
		key := it.Val()
		//fmt.Println(key)
		_, err = RedisCache.Del(key).Result()
		if err != nil {
			return
		}
	}
	return
}
