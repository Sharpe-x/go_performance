package main

import (
	"encoding/json"
	"fmt"
)

// OrgRegisterToken ...
type OrgRegisterToken struct {
	OrgId     string
	OrgOpenId string
}

// String ...
func (ot *OrgRegisterToken) String() string {
	data, _ := json.Marshal(ot)
	return string(data)
}

func main() {
	token := &OrgRegisterToken{
		OrgId:     "yDRt2UUgygqxuvlqUuO4zjEySqVWqO9J",
		OrgOpenId: "wped0SCQAAHJ0DlsyXkHpPV-tHJaJ_zA",
	}
	fmt.Println(token.String())
}
