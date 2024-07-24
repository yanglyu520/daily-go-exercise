package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func StructToMap(obj interface{}) map[string]interface{} {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		panic("StructToMap expects a struct")
	}

	result := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		result[field.Name] = value.Interface()
	}
	return result
}

func MapToStruct(m map[string]interface{}, obj interface{}) {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if val, ok := m[field.Name]; ok {
			fieldVal := v.Field(i)
			if fieldVal.CanSet() {
				fieldVal.Set(reflect.ValueOf(val))
			}
		}
	}
}

type Config struct {
	Port  int
	Host  string
	Debug bool
}

func LoadConfigFromEnv(config *Config) {
	envVars := StructToMap(*config)
	for key := range envVars {
		envVal, exists := os.LookupEnv(key)
		if exists {
			switch reflect.TypeOf(envVars[key]).Kind() {
			case reflect.String:
				envVars[key] = envVal
			case reflect.Int:
				intVal, err := strconv.Atoi(envVal)
				if err == nil {
					envVars[key] = intVal
				}
			case reflect.Bool:
				boolVal, err := strconv.ParseBool(envVal)
				if err == nil {
					envVars[key] = boolVal
				}
			}
		}
	}
	MapToStruct(envVars, config)
}

func main() {
	// 设置环境变量
	os.Setenv("Port", "8080")
	os.Setenv("Host", "localhost")
	os.Setenv("Debug", "true")

	// 默认配置
	config := Config{}

	// 从环境变量加载配置
	LoadConfigFromEnv(&config)
	fmt.Println("Config:", config) // Config: {Port:8080 Host:localhost Debug:true}
}
