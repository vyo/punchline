package model

//config defaults
const defaultPunchDir string = "/Users/manu/.punch"
const defaultConfigDir string = defaultPunchDir
const defaultBinDir string = defaultPunchDir
const defaultDataDir string = defaultPunchDir + "/data"
const defaultExportDir string = defaultPunchDir + "/export"
const defaultTemplateDir string = defaultPunchDir + "/template"

const defaultChunkSize int = 30 //measure times in 30 minute chunks


type Config struct {
	Dir       map[string]string `json:"dir"`
	ChunkSize int `json:"chunk_size"`
}

func NewConfig(dir map[string]string, chunkSize int) *Config {
	if (dir == nil) {
		dir = make(map[string]string)
	}

	if (dir["punch"] == "") {
		dir["punch"] = defaultPunchDir
	}

	if (dir["config"] == "") {
		dir["config"] = defaultConfigDir
	}

	if (dir["bin"] == "") {
		dir["bin"] = defaultBinDir
	}

	if (dir["data"] == "") {
		dir["data"] = defaultDataDir
	}

	if (dir["export"] == "") {
		dir["export"] = defaultExportDir
	}

	if (dir["template"] == "") {
		dir["template"] = defaultTemplateDir
	}

	if (chunkSize == 0) {
		chunkSize = defaultChunkSize
	}

	return &Config{dir, chunkSize}
}

func (c *Config) GetDir() *map[string]string {
	return &c.Dir
}

func (c *Config) GetChunkSize() *int {
	return &c.ChunkSize
}
