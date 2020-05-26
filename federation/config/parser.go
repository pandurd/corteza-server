package config

import (
	"log"
	"time"

	"gopkg.in/yaml.v3"
)

type (
	ConfigSettings struct {
		Settings struct {
			Sync time.Duration `yaml:"sync"`
		}
	}

	ConfigModules struct {
		Modules []ConfigModule `yaml:"modules"`
	}

	ConfigModule struct {
		Handle       string                `yaml:"handle"`
		SourceServer string                `yaml:"source_server"`
		Namespace    ConfigModuleNamespace `yaml:"namespace"`
		Fields       []ConfigModuleField   `yaml:"fields"`
	}

	ConfigModuleNamespace struct {
		ID     uint64 `yaml: "id"`
		Handle string `yaml: "handle"`
	}

	ConfigModuleField string

	ConfigStructure struct {
		Mapped []ConfigStructureMapped `yaml:"mapped"`
	}

	ConfigStructureMapped struct {
		Source       ConfigStructureModule  `yaml:"source"`
		SourceServer ConfigSourceServer     `yaml:"source_server"`
		Destination  ConfigStructureModule  `yaml:"destination"`
		Fields       []ConfigStructureField `yaml:"fields"`
	}

	ConfigStructureModule struct {
		ID        uint64 `yaml:"id"`
		Namespace uint64 `yaml:"namespace"`
		Handle    string `yaml:"handle"`
		Name      string `yaml:"name"`
	}

	ConfigStructureServer struct {
		Doit string `yaml:"doit"`
	}

	ConfigStructureField struct {
		Source      string `yaml:"source"`
		Destination string `yaml:"destination"`
	}

	ConfigSource struct {
		Servers []ConfigSourceServer `yaml:"servers"`
	}

	ConfigSourceServer struct {
		ID   string `yaml:"id"`
		URI  string `yaml:"uri"`
		Name string `yaml:"name"`
	}

	Parser struct {
		Config []byte
	}

	structureParser interface {
		Data() (ConfigModules, error)
		Structure() (ConfigStructure, error)
		Source() (ConfigSource, error)
	}
)

func (cs ConfigSource) FindSource(id string) *ConfigSourceServer {
	for _, s := range cs.Servers {
		if s.ID == id {
			return &s
		}
	}
	return nil
}

func (p Parser) Modules() (ConfigModules, error) {
	config := ConfigModules{}
	err := yaml.Unmarshal(p.Config, &config)

	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return config, err
}

// rename to Mapping
func (p Parser) Structure() (ConfigStructure, error) {
	config := ConfigStructure{}
	err := yaml.Unmarshal(p.Config, &config)

	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return config, err
}

func (p Parser) Source() (ConfigSource, error) {
	config := ConfigSource{}
	err := yaml.Unmarshal(p.Config, &config)

	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return config, err
}

func (p Parser) Settings() (ConfigSettings, error) {
	config := ConfigSettings{}
	err := yaml.Unmarshal(p.Config, &config)

	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return config, err
}

func (cs ConfigModules) FindAllHandles() []string {
	list := []string{}

	for _, i := range cs.Modules {
		list = append(list, i.Handle)
	}

	return list
}

func (cs ConfigModules) FindFields(handle string) *[]ConfigModuleField {
	for _, m := range cs.Modules {
		if m.Handle == handle {
			return &m.Fields
		}
	}

	return nil
}

func (cm ConfigModules) FindModules(server string) []*ConfigModule {
	list := []*ConfigModule{}

	for _, module := range cm.Modules {
		if module.SourceServer == server {
			list = append(list, &module)
		}
	}

	return list
}

func (cs ConfigStructure) FindModules(serverId string) []*ConfigStructureMapped {
	list := []*ConfigStructureMapped{}

	for _, module := range cs.Mapped {
		if module.SourceServer.ID == serverId {
			// if module.SourceServer.ID == serverId {
			list = append(list, &module)
		}
	}

	return list
}
