package core

import(
        "encoding/json"
        "io/ioutil"
)

type Config struct {
        FakeSuccessPage string
        EnableFakeSuccessPage bool
        LogPath string
        Port int
}

func LoadConfig() (Config, error){
        rawConf, err := ioutil.ReadFile("config/config.json")
        var conf Config
        if err != nil{
            return conf, err
        }
        err = json.Unmarshal(rawConf, &conf)
        if err != nil{
                return conf, err
        }
        return conf, nil
}
