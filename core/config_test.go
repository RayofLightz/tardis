package core
import "testing"

func TestLoadConfig(t *testing.T){
        res, err := LoadConfig()
        if err != nil{
                t.Error("Err not nill")
        }
}
