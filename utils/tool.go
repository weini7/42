package utils

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"xorm.io/builder"
)

//构建以数组JSON存储的查询cond
func BuildArrayCond(field string, d interface{}) (cond builder.Cond) {
	switch d.(type) {
	case int, int64:
		//return builder.Like{field, fmt.Sprintf("[%d,", d)}.
		//	Or(builder.Like{field, fmt.Sprintf(",%d,", d)}).
		//	Or(builder.Like{field, fmt.Sprintf("[%d]", d)}).
		//	Or(builder.Like{field, fmt.Sprintf(",%d]", d)})
		//return builder.Expr("JSON_CONTAINS(?,'?')", field, d)
		return builder.Expr(fmt.Sprintf("JSON_CONTAINS(%s,'%d')", field, d))
	case string:
		//return builder.Like{field, fmt.Sprintf("[\"%s\",", d)}.
		//	Or(builder.Like{field, fmt.Sprintf(",\"%s\",", d)}).
		//	Or(builder.Like{field, fmt.Sprintf("[\"%s\"]", d)}).
		//	Or(builder.Like{field, fmt.Sprintf(",\"%s\"]", d)})
		//return builder.Expr("JSON_CONTAINS(?,'\\\"?\\\"')", field, d)
		return builder.Expr(fmt.Sprintf("JSON_CONTAINS(%s,'\"%s\"')", field, d))
		//return builder.Expr("JSON_CONTAINS(?,'\"?\"')", field, d)
	default:
		return
	}
}
func BuildArrayCondDead(field string, d interface{}) (cond builder.Cond) {
	switch d.(type) {
	case int, int64:
		return builder.Like{field, fmt.Sprintf("[%d,", d)}.
			Or(builder.Like{field, fmt.Sprintf(",%d,", d)}).
			Or(builder.Like{field, fmt.Sprintf("[%d]", d)}).
			Or(builder.Like{field, fmt.Sprintf(",%d]", d)})
	case string:
		return builder.Like{field, fmt.Sprintf("[\"%s\",", d)}.
			Or(builder.Like{field, fmt.Sprintf(",\"%s\",", d)}).
			Or(builder.Like{field, fmt.Sprintf("[\"%s\"]", d)}).
			Or(builder.Like{field, fmt.Sprintf(",\"%s\"]", d)})
	default:
		return
	}
}

//转GB18030编码
func ConvertByte2Encoding(byte []byte) (encodeBytes []byte) {
	encodeBytes, _ = simplifiedchinese.GB18030.NewEncoder().Bytes(byte)
	return
}
func ConvertStr2Encoding(str string, which string) (encodeString string) {
	switch which {
	case "GB18030":
		var err error
		encodeString, err = simplifiedchinese.GB18030.NewEncoder().String(str)
		if err != nil {
			log.Println(err.Error())
		}
	case "GBK":
		var err error
		encodeString, err = simplifiedchinese.GBK.NewEncoder().String(str)
		if err != nil {
			log.Println(err.Error())
		}
	case "GB2312":
		var err error
		encodeString, err = simplifiedchinese.HZGB2312.NewEncoder().String(str)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return
}

//简单校验-姓名-
func SimpleCheckName(name string) (errMsg string) {
	if len(name) <= 0 {
		return "姓名为空"
	}
	//result, _ := regexp.MatchString(`^(1[3|4|5|7|8][0-9]\d{4,8})$`, mobie)
	result, _ := regexp.MatchString(`\d`, name)
	if result {
		return "姓名不能包含数字"
	}
	return ""
}

//简单校验-身份证
func SimpleCheckIdcard(idcard string) (errMsg string) {
	switch len(idcard) {
	case 15:
		// 15位身份证号码：15位全是数字
		result, _ := regexp.MatchString(`^(\d{15})$`, idcard)
		if !result {
			return "身份证校验未通过"
		}
	case 18:
		// 18位身份证：前17位为数字，第18位为校验位，可能是数字或X
		result, _ := regexp.MatchString(`^(\d{17})([0-9]|X|x)$`, idcard)
		if !result {
			return "身份证校验未通过"
		}
	case 0:
		return "身份证为空"
	default:
		return "身份证位数错误"
	}
	return ""
}

//简单校验-手机号
func SimpleCheckMobile(mobile string) (errMsg string) {
	if len(mobile) <= 0 {
		return "手机号为空"
	}
	//result, _ := regexp.MatchString(`^(1[3|4|5|7|8][0-9]\d{4,8})$`, mobile)
	result, _ := regexp.MatchString(`^(1[0-9][0-9]\d{8})$`, mobile)
	if !result {
		return "手机号校验未通过"
	}
	return ""
}

//简单校验-银行卡号
func SimpleCheckBankAcc(bankAcc string) (errMsg string) {
	if len(bankAcc) <= 0 {
		return "银行卡卡号为空"
	}
	result, _ := regexp.MatchString(`^[0-9]*$`, bankAcc)
	if !result {
		return "银行卡卡号校验未通过"
	}
	return ""
}

//简单校验-金额
//func SimpleCheckMoney(money float64) (errMsg string) {
//	if money == 0 {
//		return "金额为空"
//	}
//	result, _ := regexp.MatchString(`/(^[1-9]{1}[0-9]*$)|(^[0-9]*\.[0-9]{2}$)/`, fmt.Sprintf("%f", money))
//	if !result {
//		return "金额校验未通过"
//	}
//	return ""
//}

//简单校验-大小写字母数字
func SimpleCheckNumberEnglish(str string) (errMsg string) {
	if len(str) <= 0 {
		return "参数为空"
	}
	result, _ := regexp.MatchString(`^[0-9a-zA-Z]*$`, str)
	if !result {
		return "只能包含大小写字母与数字"
	}
	return ""
}

//获取浮点型小数点位数   //位数过大会出现科学计数法
func GetDigitByFloat(f float64) (digit int) { //最多支持算上小数点总18位浮点型   最后一位会四舍五入
	fStr := strconv.FormatFloat(f, 'f', -1, 64)
	fArray := strings.Split(fStr, ".")
	//fArray := strings.Split(fmt.Sprintf("%v", f), ".")
	if len(fArray) == 2 {
		return len(fArray[1])
	} else {
		return 0
	}
}

//现金元单位转无小数分单位|重新过滤一遍浮点型|四舍五入 NWF 待验证
func Float64ConventIntFilter(fee float64) (feeee int, err error) {
	feee, err := strconv.ParseFloat(fmt.Sprintf("%.f", fee*float64(100)), 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return int(feee), nil
}
func Float64ConventInt64(f float64) (fInt64 int64, err error) {
	fff, err := strconv.ParseFloat(fmt.Sprintf("%.f", f), 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return int64(fff), nil

}

//按which获取随机字符串
func GetRandomStringAsWhich(l int, which string) string {
	str := ""
	switch which {
	case "Number":
		str = "0123456789"
	case "Number-UpEnglish":
		str = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "UpEnglish":
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "Number-UpLowEnglish":
		str = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
