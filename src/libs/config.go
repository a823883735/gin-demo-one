package libs

import (
	"bufio"
	"gin-demo-one/src/utils"
	"io"
	"os"
	"strings"
)

type Configs struct {
	config map[string]string
	node   string
}

var Conf *Configs

func init() {
	Conf = new(Configs)
	Conf.LoadConfig("src/config/config.ini")
}

const MidStr = "?==!"

func (conf *Configs) LoadConfig(path string) {
	conf.config = make(map[string]string)
	file, err := os.Open(path)
	utils.ErrorHandle(err, "this "+path+" file is not exist.")
	defer file.Close()

	buf := bufio.NewReader(file)

	for {
		lines, _, err := buf.ReadLine()
		if err != nil {
			//文件读取结束
			if err == io.EOF {
				//utils.ErrorHandle(err, "file readed.")
				break
			}
			utils.ErrorHandle(err, "unknow error.")
		}
		line := strings.TrimSpace(string(lines))

		//处理注释
		if strings.Index(line, "#") == 0 {
			continue
		}

		//如果是[xxx]
		n := strings.Index(line, "[")
		nl := strings.LastIndex(line, "]")
		if n > -1 && nl > -1 && nl > n+1 {
			conf.node = strings.TrimSpace(line[n+1 : nl])
			continue
		}
		if len(conf.node) == 0 || len(line) == 0 {
			continue
		}
		arr := strings.Split(line, "=")
		key := strings.TrimSpace(arr[0])
		value := strings.TrimSpace(arr[1])
		newKey := conf.node + MidStr + key
		conf.config[newKey] = value
	}
}

func (conf *Configs) Read(node, key string) string {
	key = node + MidStr + key
	if v, ok := conf.config[key]; !ok {
		return ""
	} else {
		return v
	}
}
