package mystruct

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agent struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty" jsonschema:"-"`
	CreateAt  *time.Time         `json:"createat,omitempty" bson:"createat,omitempty"`
	UpdateAt  *time.Time         `json:"updateat,omitempty" bson:"updateat,omitempty"`
	DeleteAt  *time.Time         `json:"deleteat,omitempty" bson:"deleteat,omitempty"`
	IsDisable *bool              `json:"isdisable,omitempty" bson:"isdisable,omitempty"`
	IsDelete  *bool              `json:"isdelete,omitempty" bson:"isdelete,omitempty"`

	Pwd        string             `json:"pwd,omitempty" bson:"pwd,omitempty" jsonschema:"minLength=6,maxLength=512,description=pwd"`
	PayPwd     string             `json:"paypwd,omitempty" bson:"paypwd,omitempty" jsonschema:"minLength=0,maxLength=512,description=paypwd"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty" jsonschema:"minLength=1,maxLength=128,description=name"`
	Phone      string             `json:"phone,omitempty" bson:"phone,omitempty" jsonschema:"required,minLength=11,maxLength=11,description=phone"`
	IdCardName string             `json:"idcardname,omitempty" bson:"idcardname,omitempty" jsonschema:"enum=信用代码|身份证|护照|港澳通行证|香港身份证|澳门身份证|国外证照|其他,description=idcardname"`
	IdCardNo   string             `json:"idcardno,omitempty" bson:"idcardno,omitempty" jsonschema:"description=idcardno"`
	WxId       string             `json:"wxid,omitempty" bson:"wxid,omitempty" jsonschema:"minLength=0,maxLength=25,description=wxid"` //微信号
	Remark     string             `json:"remark,omitempty" bson:"remark,omitempty"`
	AuthCode   string             `json:"authcode,omitempty" bson:"authcode,omitempty"`                                                             //授权编码
	Kind       string             `json:"kind,omitempty" bson:"kind,omitempty" jsonschema:"enum=代理商|电商|普通用户,description=kind"`                      //12未注册用户
	IsDocking  *bool              `json:"isdocking,omitempty" bson:"isdocking,omitempty" jsonschema:"description=isdocking"`                        //是否对接代理商
	Level      *float64           `json:"level,omitempty" bson:"level,omitempty" jsonschema:"minimum=0,maximum=9999,description=level"`             //代理商等级AgentLevel
	LevelName  string             `json:"levelname,omitempty" bson:"levelname,omitempty" `                                                          //级别名称
	Leader     primitive.ObjectID `json:"leader,omitempty" bson:"leader,omitempty" jsonschema:"oid,description=leader"`                             //代理商上级oid
	LeaderName string             `json:"leadername,omitempty" bson:"leadername,omitempty"`                                                         //上级姓名
	LvState    string             `json:"lvstate,omitempty" bson:"lvstate,omitempty" jsonschema:"enum=加入首次|升级首次|正常|一次性升级首次,description=lvstate"`    //升级状态
	AcState    string             `json:"acstate,omitempty" bson:"acstate,omitempty" jsonschema:"enum=未确认|未审核|已审核|冻结|退出冻结|已退出,description=acstate"` //账号状态
	WxSvcId    string             `json:"wxsvcid" bson:"wxsvcid"`                                                                                   //服务号id
	OpenId     string             `json:"openid,omitempty" bson:"openid,omitempty"`                                                                 //微信openid
	WxBusId    string             `json:"wxbusid,omitempty" bson:"wxbusid,omitempty"`                                                               //企业号id
	//Property            string        `json:"property,omitempty" bson:"property,omitempty" jsonschema:"enum=企业|个人,description=property"`
	CompanyName         string               `json:"companyname,omitempty" bson:"companyname,omitempty"`
	JoinTime            *time.Time           `json:"jointime,omitempty" bson:"jointime,omitempty"` //加入时间
	Supplier            primitive.ObjectID   `json:"supplier,omitempty" bson:"supplier,omitempty" jsonschema:"oid,description=提货人"`
	DockingSupplier     primitive.ObjectID   `json:"dockingsupplier,omitempty" bson:"dockingsupplier,omitempty" jsonschema:"oid,description=提货对接公司人"`
	DockingSupplierName string               `json:"dockingsuppliername,omitempty" bson:"dockingsuppliername,omitempty" jsonschema:"description=提货对接公司人姓名"`
	SpecialFrom         primitive.ObjectID   `json:"specialfrom,omitempty" bson:"specialfrom,omitempty" jsonschema:"oid,description=在哪个特种兵的团队"`
	SpecialFromName     string               `json:"specialfromname,omitempty" bson:"specialfromname,omitempty" jsonschema:"description=特种兵姓名"`
	IsBigSupply         string               `json:"isbigsupply,omitempty" bson:"isbigsupply,omitempty" jsonschema:"enum=正常|特许|不限量|稽查限量,description=isbigsupply"` //是否开启最大提货量
	Team                string               `json:"team,omitempty" bson:"team,omitempty" jsonschema:"minLength=0,maxLength=512,description=池子"`
	ContractBank        primitive.ObjectID   `json:"contractbank,omitempty" bson:"contractbank,omitempty" jsonschema:"oid"` //公司银行账号
	ContractBankName    string               `json:"contractbankname,omitempty" bson:"contractbankname,omitempty"`          //对接公司名称
	IsSign              *bool                `json:"issign,omitempty" bson:"issign,omitempty" jsonschema:"description=是否已签合同"`
	SignAt              *time.Time           `json:"signat,omitempty" bson:"signat,omitempty"` //签约时间
	OldCount            *float64             `json:"oldcount,omitempty" bson:"oldcount,omitempty" jsonschema:"minimum=0,maximum=999999999,description=历史累积量"`
	OldPayTotal         *float64             `json:"oldpaytotal,omitempty" bson:"oldpaytotal,omitempty" jsonschema:"minimum=0,maximum=999999999,description=历史累积提货金额"`
	Tree                []primitive.ObjectID `json:"tree,omitempty" bson:"tree,omitempty" jsonschema:"oids,description=团队树"`
	Invite              primitive.ObjectID   `json:"invite,omitempty" bson:"invite,omitempty" jsonschema:"oid,description=邀请人"`
	EqualInvite         *bool                `json:"equalinvite,omitempty" bson:"equalinvite,omitempty" jsonschema:"description=是否平推"`
	LowLevelAt          *time.Time           `json:"lowlevelat,omitempty" bson:"lowlevelat,omitempty"` //最后降级时间
	MuState             string               `json:"mustate,omitempty" bson:"mustate,omitempty" jsonschema:"enum=稽查中|稽查已处理-退出|稽查已处理-恢复正常|正常,description=mustate"`
	UpLevelLessAt       *time.Time           `json:"uplevellessat,omitempty" bson:"uplevellessat,omitempty"` //最后优惠升级时间
	IsSpecial           *bool                `json:"isspecial,omitempty" bson:"isspecial,omitempty" jsonschema:"description=是否特种兵"`
	Avatar              string               `json:"avatar,omitempty" bson:"avatar,omitempty" jsonschema:"minLength=0,maxLength=250,description=微信头像地址"`
	FrostState          *bool                `json:"froststate,omitempty" bson:"froststate,omitempty"`                                              //代理的冻结状态
	LastLoginAt         *time.Time           `json:"lastloginat,omitempty" bson:"lastloginat,omitempty"`                                            //最后登陆
	LastReadAt          *time.Time           `json:"lastreadat,omitempty" bson:"lastreadat,omitempty"`                                              //最后读消息的时间
	FranchiseStore      []string             `json:"franchisestore,omitempty" bson:"franchisestore,omitempty"`                                      //专营店授权号
	RecruitCount        int                  `json:"recruitcount,omitempty" bson:"recruitcount,omitempty"  jsonschema:"minimum=0,maximum=99999999"` //活动计数
	August              *bool                `json:"august,omitempty" bson:"august,omitempty"`                                                      //是否新首席
	IsPublish           *bool                `json:"ispublish,omitempty" bson:"ispublish,omitempty"`                                                //是否对公账户
	FromSource          string               `json:"fromsource,omitempty" bson:"fromsource,omitempty"`                                              //来源
	StoreMember         []primitive.ObjectID `json:"storemember,omitempty" bson:"storemember,omitempty" jsonschema:"oids,description=所属专营店店铺"`

	BeforeContractBank primitive.ObjectID `json:"beforecontractbank,omitempty" bson:"beforecontractbank,omitempty" jsonschema:"oid"` //之前的签约银行
	Count136           *int               `json:"count136,omitempty" bson:"count136,omitempty"`                                      //136活动的累计数量
	Time136            *time.Time         `json:"time136,omitempty" bson:"time136,omitempty"`                                        //136活动的上一次达标时间
}

func toJson(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		// fmt.Print(err)
		return nil, err
	}
	// print(string(b))
	return b, nil
}
