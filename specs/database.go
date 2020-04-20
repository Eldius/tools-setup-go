package specs

import (
	"encoding/json"
	"fmt"

	"github.com/Eldius/tools-setup-go/config"
	"github.com/Eldius/tools-setup-go/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func configDb() *gorm.DB {
	if db, err := gorm.Open("sqlite3", config.GetDBFile()); err != nil {
		fmt.Println("failed to connect database", err.Error())
		panic(err.Error())
	} else {
		db.AutoMigrate(&ToolSpec{})
		return db
	}
}

/*
PersistSpec persists a new spec
*/
func PersistSpec(spec ToolSpec) {
	db := configDb()
	defer db.Close()
	db.Create(spec)
}

/*
PersistSpecs persists a list of specs
*/
func PersistSpecs(specs []ToolSpec) []ToolSpec {
	for _, s := range specs {
		PersistSpec(s)
	}

	return specs
}

/*
FindSpec fins a spec by name and version
*/
func FindSpec(name, version string) ToolSpec {
	db := configDb()
	defer db.Close()
	var result ToolSpec

	db.Where("name = ?, version = ?", name, version).First(&result)

	return result
}

/*
ListSpecs lists all specs
*/
func ListSpecs() []ToolSpec {
	db := configDb()
	defer db.Close()
	var specs []ToolSpec
	db.Find(&specs)

	return specs
}

/*
ListSpecsByName lists all specs by app name
*/
func ListSpecsByName(name string) []ToolSpec {
	db := configDb()
	defer db.Close()
	specs := make([]ToolSpec, 1)
	db.Where("name = ?", name).Find(specs)

	return specs
}

func getKey(spec ToolSpec) []byte {
	return []byte(fmt.Sprintf("%s-%s", spec.Name, spec.Version))
}

func getKeyAlt(name, version string) []byte {
	return []byte(fmt.Sprintf("%s-%s", name, version))
}

func marshal(obj interface{}) []byte {
	jsonb, err := json.Marshal(obj)
	if err != nil {
		panic(err.Error())
	}
	logger.Debug(fmt.Sprintf("marshaled: %s", string(jsonb)))
	return jsonb
}
