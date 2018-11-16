//工具包,config类与方法
package Tools

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

var path string

type cfg struct {
	section     string
	Discription string
	Data        map[string]string
}
type Cfgs map[string]*cfg

func (this Cfgs) GetBool(section, key string) (bool, error) {
	_, ok := this[section]
	if ok {
		v, ok := this[section].Data[key]
		if ok {
			Value, err := strconv.ParseBool(v)
			return Value, err
		} else {
			return false, errors.New("can not find key: " + key + " in  section: " + section)
		}
	} else {
		return false, errors.New("can not find section " + section)
	}

}
func (this Cfgs) GetInt(section, key string) (int, error) {
	_, ok := this[section]
	if ok {
		v, ok := this[section].Data[key]
		if ok {
			Value, err := strconv.Atoi(v)
			return int(Value), err
		} else {
			return 0, errors.New("can not find key: " + key + " in  section: " + section)
		}
	} else {
		return 0, errors.New("can not find section " + section)
	}

}
func (this Cfgs) GetFloat(section, key string) (float64, error) {
	_, ok := this[section]
	if ok {
		v, ok := this[section].Data[key]
		if ok {
			Value, err := strconv.ParseFloat(v, 64)
			return Value, err
		} else {
			return 0, errors.New("can not find key: " + key + " in  section: " + section)
		}
	} else {
		return 0, errors.New("can not find section " + section)
	}

}
func (this Cfgs) GetValue(section, key string) (string, error) {
	v, ok := this[section]
	if ok {
		if value, ok := v.Data[key]; ok {
			return value, nil
		}
		return "", errors.New("can not find key: " + key + " in  section: " + section)
	} else {
		return "", errors.New("can not find section " + section)
	}

}

func (this Cfgs) SetBool(section, key string, value bool) error {
	_, ok := this[section]
	if ok {
		_, ok := this[section].Data[key]
		if ok {
			this[section].Data[key] = strconv.FormatBool(value)
			return nil
		} else {
			return errors.New("can not find key: " + key + " in  section: " + section)
		}
	} else {
		return errors.New("can not find section " + section)
	}

}
func (this Cfgs) SetInt(section, key string, value int) error {
	_, ok := this[section]
	if ok {
		_, ok := this[section].Data[key]
		if ok {
			this[section].Data[key] = strconv.Itoa(value)
			return nil
		} else {
			return errors.New("can not find key: " + key + " in  section: " + section)
		}
	} else {
		return errors.New("can not find section " + section)
	}

}
func (this Cfgs) SetFloat(section, key string, value float64) error {
	_, ok := this[section]
	if ok {
		_, ok := this[section].Data[key]
		if ok {
			this[section].Data[key] = strconv.FormatFloat(value, 'E', 5, 32)
			return nil
		} else {
			return errors.New("can not find key: " + key + " in  section: " + section)
		}
	} else {
		return errors.New("can not find section " + section)
	}

}
func (this Cfgs) SetValue(section, key string, value string) error {
	_, ok := this[section]
	if ok {
		_, ok := this[section].Data[key]
		if ok {
			this[section].Data[key] = value
			return nil
		} else {
			return errors.New("can not find key: " + key + " in  section: " + section)
		}
	} else {
		return errors.New("can not find section " + section)
	}

}

func (this Cfgs) GetDiscription(section string) (string, error) {
	_, ok := this[section]
	if ok {
		return this[section].Discription, nil
	} else {
		return "", errors.New("can not find section " + section)
	}
}
func (this Cfgs) ToString() string {
	str := ""
	for section := range this {
		str += "[" + section + "]\r\n"
		if this[section].Discription != "" {
			str += this[section].Discription + "\r\n"
		}

		for key := range this[section].Data {
			str += key + "=" + this[section].Data[key] + "\r\n"
		}
	}
	return str
}
func (this Cfgs) SaveConfigAs(path string) error {
	str := this.ToString()
	err := WritAllText(path, str)
	return err
	return nil
}

func (this Cfgs) SaveConfig() error {
	str := this.ToString()
	err := WritAllText(path, str)
	return err
}
func LoadConfig(cfgpath string) (Cfgs, error) {
	var cfgMP Cfgs
	path = cfgpath
	cfgMP = make(map[string]*cfg)

	data := make(map[string]string)
	cfgMP["default"] = &cfg{section: "default", Data: data}

	cfg_str, err := ReadAllText(path)

	if err != nil {
		return cfgMP, err
	}
	cfg_strs := strings.Split(cfg_str, "\r\n")
	for i := 0; i < len(cfg_strs)-1; i++ {
		bb := ([]byte(cfg_strs[i]))
		if len(bb) < 1 {
			continue
		} else if len(bb) > 0 && bb[0] == '#' {
			log.Println(cfgMP["default"].Discription)
			cfgMP["default"].Discription += cfg_strs[i]
		} else if len(bb) > 2 && bb[0] == '[' && bb[len(bb)-1] == ']' {
			name := string(bb[1 : len(bb)-1])
			data := make(map[string]string)
			cfgMP[name] = &cfg{section: name, Data: data}
			for i < len(cfg_strs)-1 {
				i++
				bb := ([]byte(cfg_strs[i]))
				if len(bb) < 1 {
					continue
				} else if len(bb) > 0 && bb[0] == '#' {
					cfgMP[name].Discription += cfg_strs[i]
				} else if len(bb) > 2 && bb[0] == '[' && bb[len(bb)-1] == ']' {
					i--
					break
				} else {
					data := strings.Split(cfg_strs[i], "=")
					if len(data) == 2 {
						cfgMP[name].Data[data[0]] = data[1]

					}
				}
			}
		} else {
			data := strings.Split(cfg_strs[i], "=")
			if len(data) == 2 {
				cfgMP["default"].Data[data[0]] = data[1]
			}
		}
	}
	return cfgMP, nil
}

func WritAllText(path, Text string) error {
	if fin, err := os.Open(path); err != nil {
		err = fin.Close()
		return err
	} else {
		_, err = fin.WriteString(Text)
		err = fin.Close()
		return err
	}
}
func ReadAllText(path string) (string, error) {
	if fin, err := os.Open(path); err != nil {

		return "", err
	} else {
		str := ""
		for {
			buf := make([]byte, 1024, 1024)
			n, err := fin.Read(buf)
			if err != nil {
				fin.Close()
				return str, err
			}
			if n < 1024 {
				str += string(buf[:n])
				break
			} else {
				str += string(buf)
			}

		}
		fin.Close()
		return str, nil
	}
}

//func main() {
//	cfg, err := LoadConfig("cfg.ini")
//	if err == nil {
//		value, err := cfg.GetValue("MainForm", "BackColor")
//		cfg.SetBool("MainForm", "BackColor", true)
//		log.Println(err, value)
//		value, err = cfg.GetValue("MainForm", "BackColor")
//		log.Println(err, value)
//		log.Println(cfg.ToString())
//	}
//}
