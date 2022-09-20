package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"
)

type DescribeOfficialWebFormAnswersRequestOperation struct {
	Caller    *Caller                     `protobuf:"bytes,1,opt,name=Caller,proto3" json:"Caller,omitempty"`
	Filters   []*FilterOperator           `protobuf:"bytes,4,rep,name=Filters,proto3" json:"Filters,omitempty"`
	StartTime int64                       `protobuf:"varint,5,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime   int64                       `protobuf:"varint,6,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Offset    int64                       `protobuf:"varint,7,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     int64                       `protobuf:"varint,8,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Answers   []*FormAnswerRecordOperator `protobuf:"bytes,3,rep,name=Answers,proto3" json:"Answers,omitempty"`
}

type FormAnswerRecordOperator struct {
	AnswerId string            `protobuf:"bytes,1,opt,name=answer_id,json=answerId,proto3" json:"answer_id,omitempty"`
	Answers  map[string]string `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

type Caller struct {
	OrganizationId    string `protobuf:"bytes,1,opt,name=OrganizationId,proto3" json:"OrganizationId,omitempty"`
	ApplicationId     string `protobuf:"bytes,2,opt,name=ApplicationId,proto3" json:"ApplicationId,omitempty"`
	SubOrganizationId string `protobuf:"bytes,3,opt,name=SubOrganizationId,proto3" json:"SubOrganizationId,omitempty"`
	OperatorId        string `protobuf:"bytes,4,opt,name=OperatorId,proto3" json:"OperatorId,omitempty"`
	SourceIp          string `protobuf:"bytes,5,opt,name=SourceIp,proto3" json:"SourceIp,omitempty"`
}

type FilterOperator struct {
	Key    string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=Values,proto3" json:"Values,omitempty"`
}

type DescribeOfficialWebFormAnswersRequest struct {
	Organization *OrganizationInfo   `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Operator     *UserInfo           `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Agent        *Agent              `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	Filters      []*Filter           `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
	StartTime    int64               `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime      int64               `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Offset       int64               `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit        int64               `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`
	Answers      []*FormAnswerRecord //`protobuf:"bytes,3,rep,name=Answers,proto3" json:"Answers,omitempty"`
}

type FormAnswerRecord struct {
	AnswerId string            `protobuf:"bytes,1,opt,name=AnswerId,proto3" json:"AnswerId,omitempty"`
	Answers  map[string]string `protobuf:"bytes,2,rep,name=Answers,proto3" json:"Answers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

type OrganizationInfo struct {
	OrganizationId     string `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	ClientIp           string `protobuf:"bytes,2,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	ChannelType        string `protobuf:"bytes,6,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	Channel            string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	OrganizationOpenId string `protobuf:"bytes,4,opt,name=organization_open_id,json=organizationOpenId,proto3" json:"organization_open_id,omitempty"`
	ProxyIp            string `protobuf:"bytes,5,opt,name=proxy_ip,json=proxyIp,proto3" json:"proxy_ip,omitempty"`
}

type Filter struct {
	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

type UserInfo struct {
	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientIp string `protobuf:"bytes,2,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	Channel  string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	OpenId   string `protobuf:"bytes,4,opt,name=open_id,json=openId,proto3" json:"open_id,omitempty"`
	ProxyIp  string `protobuf:"bytes,5,opt,name=proxy_ip,json=proxyIp,proto3" json:"proxy_ip,omitempty"`
}

type Agent struct {
	Channel             string    `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	AppId               string    `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ProxyAppId          string    `protobuf:"bytes,3,opt,name=proxy_app_id,json=proxyAppId,proto3" json:"proxy_app_id,omitempty"`
	ProxyOrganizationId string    `protobuf:"bytes,4,opt,name=proxy_organization_id,json=proxyOrganizationId,proto3" json:"proxy_organization_id,omitempty"`
	ProxyOperator       *UserInfo `protobuf:"bytes,5,opt,name=proxy_operator,json=proxyOperator,proto3" json:"proxy_operator,omitempty"`
}

type UrchinTrackingModuleInfo struct {
	Source   string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Medium   string `protobuf:"bytes,2,opt,name=medium,proto3" json:"medium,omitempty"`
	Campaign string `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	Unit     string `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	Term     string `protobuf:"bytes,5,opt,name=term,proto3" json:"term,omitempty"`
}

func main() {
	utmInfo := &UrchinTrackingModuleInfo{
		Source: "source",
		Medium: "medium",
	}

	t := reflect.TypeOf(utmInfo).Elem()
	v := reflect.ValueOf(utmInfo).Elem()
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		value := v.FieldByName(t.Field(i).Name).String()
		fmt.Println(name, " = ", value)
	}

	fmt.Println(strings.Split("sharpezhang", ";"))

	fmt.Println("===============")

	request := &DescribeOfficialWebFormAnswersRequestOperation{
		Caller: &Caller{
			OperatorId: "sharpezhang",
		},
		Filters: []*FilterOperator{
			{
				Key:    "key1",
				Values: []string{"a", "b", "c"},
			},
			{
				Key:    "key2",
				Values: []string{"a", "b", "c"},
			},
			{
				Key:    "key3",
				Values: []string{"a", "b", "c"},
			},
		},
		StartTime: 100,
		EndTime:   200,
		Limit:     300,
		Offset:    400,
	}

	describeRequest := new(DescribeOfficialWebFormAnswersRequest)
	if err := copier.Copy(describeRequest, request); err != nil {
		fmt.Println("err = ", err)
	}

	fmt.Printf("%v\n", describeRequest)
	fmt.Printf("%+v\n", describeRequest)
	fmt.Printf("%#v\n", describeRequest)
	fmt.Printf("%#v\n", describeRequest.Filters[0])

	fmt.Println("===============")
	describeRequest.Answers = []*FormAnswerRecord{
		{
			AnswerId: "1",
			Answers: map[string]string{
				"1":  "1",
				"11": "1",
			},
		},
		{
			AnswerId: "2",
			Answers: map[string]string{
				"2":  "2",
				"22": "2",
			},
		},
	}

	newRequestOperation := new(DescribeOfficialWebFormAnswersRequestOperation)
	if err := copier.Copy(newRequestOperation, describeRequest); err != nil {
		fmt.Println("err = ", err)
	}

	for _, answer := range newRequestOperation.Answers {
		fmt.Printf("%#v\n", *answer)
	}

}
