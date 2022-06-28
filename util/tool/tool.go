package tool

import (
	"bytes"
	"context"
	"encoding/json"
	"fehu/conf"
	"fehu/constant"
	"fehu/model/http_error"
	"fehu/service/base"
	"fehu/util/convert"
	"fehu/util/cryp"
	"fehu/util/math"
	"fehu/util/request"
	"fehu/util/time"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"fehu/util/stringify"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	cos "github.com/tencentyun/cos-go-sdk-v5"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	live "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

const (
	GIF_TYPE = "gif"
	BMP_TYPE = "bmp"
	JPG_TYPE = "jpg"
	PNG_TYPE = "png"
	MP4_TYPE = "mp4"
)

func UpdFile(c *gin.Context, folder string, allowedType []string, limitSize int64) *UploadRes {

	header, _ := c.FormFile("file")

	if header == nil {
		panic(http_error.UploadFileErr)
	}

	contentType := ""
	if contentTypeList, ok := header.Header["Content-Type"]; ok {
		if len(contentTypeList) > 0 {
			contentType = contentTypeList[0]
		}
	}

	if limitSize != 0 && header.Size > limitSize {
		panic(http_error.UploadFileIsTooLarge)
	}

	f, err := header.Open()
	defer f.Close()

	body, err := ioutil.ReadAll(f)

	var itype string
	var remark string

	if convert.InArrayString(PNG_TYPE, allowedType) && contentType == "image/png" {
		itype = PNG_TYPE
		conf, _, _ := image.DecodeConfig(bytes.NewReader(body))
		remark = strconv.FormatInt(int64(conf.Width), 10) + "," + strconv.FormatInt(int64(conf.Height), 10)
	}
	if convert.InArrayString(GIF_TYPE, allowedType) && contentType == "image/gif" {
		itype = GIF_TYPE
		conf, _, _ := image.DecodeConfig(bytes.NewReader(body))
		remark = strconv.FormatInt(int64(conf.Width), 10) + "," + strconv.FormatInt(int64(conf.Height), 10)
	}
	if convert.InArrayString(JPG_TYPE, allowedType) && contentType == "image/jpeg" {
		itype = JPG_TYPE
		conf, _, _ := image.DecodeConfig(bytes.NewReader(body))
		remark = strconv.FormatInt(int64(conf.Width), 10) + "," + strconv.FormatInt(int64(conf.Height), 10)
	}
	if convert.InArrayString(MP4_TYPE, allowedType) && contentType == "video/mp4" {
		itype = MP4_TYPE
	}

	if itype == "" {
		panic("can not identify type")
	}

	baseUrl := conf.Http.UploadDomain

	curlFileData, err := request.CurlFile(baseUrl+conf.Http.UploadExec, folder, header, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	key := curlFileData.Url
	return &UploadRes{key, baseUrl + "/" + key, itype, remark}

}

type UploadRes struct {
	Path   string
	Full   string
	Type   string
	Remark string
}

func UpdTxFile(c *gin.Context, folder string, allowedType []string, limitSize int64) *UploadRes {

	baseUrl := "https://touxiang-1258734358.cos.ap-shanghai.myqcloud.com"
	u, _ := url.Parse(baseUrl)
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDwkJbcOG3wfbIitte8MLNAfuse6NTuZjc",
			SecretKey: "d06xdHekaDar0siKfkbWAh7VdZFq9DyY",
		},
	})

	header, _ := c.FormFile("file")

	if header == nil {
		panic(http_error.UploadFileErr)
	}

	contentType := ""
	if contentTypeList, ok := header.Header["Content-Type"]; ok {
		if len(contentTypeList) > 0 {
			contentType = contentTypeList[0]
		}
	}

	if limitSize != 0 && header.Size > limitSize {
		panic(http_error.UploadFileIsTooLarge)
	}

	f, err := header.Open()
	defer f.Close()

	body, err := ioutil.ReadAll(f)
	md5V := cryp.MD5(string(body))

	var itype string
	var remark string

	if convert.InArrayString(PNG_TYPE, allowedType) && bytes.Equal([]byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}, body[0:8]) {
		itype = PNG_TYPE
		conf, _, _ := image.DecodeConfig(bytes.NewReader(body))
		remark = strconv.FormatInt(int64(conf.Width), 10) + "," + strconv.FormatInt(int64(conf.Height), 10)
	}
	if convert.InArrayString(GIF_TYPE, allowedType) && "GIF" == string(body[0:3]) {
		itype = GIF_TYPE
		conf, _, _ := image.DecodeConfig(bytes.NewReader(body))
		remark = strconv.FormatInt(int64(conf.Width), 10) + "," + strconv.FormatInt(int64(conf.Height), 10)
	}
	if convert.InArrayString(BMP_TYPE, allowedType) && "BM" == string(body[0:2]) {
		itype = BMP_TYPE
	}
	if convert.InArrayString(JPG_TYPE, allowedType) && bytes.Equal([]byte{0xff, 0xd8, 0xff}, body[0:3]) {
		itype = JPG_TYPE
		conf, _, _ := image.DecodeConfig(bytes.NewReader(body))
		remark = strconv.FormatInt(int64(conf.Width), 10) + "," + strconv.FormatInt(int64(conf.Height), 10)
	}
	if convert.InArrayString(MP4_TYPE, allowedType) && contentType == "video/mp4" {
		itype = MP4_TYPE
	}

	if itype == "" {
		panic("can not identify type")
	}

	key := folder + "/" + string(md5V) + "." + itype

	_, err = client.Object.Head(context.Background(), key, nil)
	if err == nil {
		return &UploadRes{key, baseUrl + "/" + key, itype, remark}
	}

	_, err = client.Object.Put(context.Background(), key, bytes.NewReader(body), nil)
	if err != nil {
		panic(err)
	}

	return &UploadRes{key, baseUrl + "/" + key, itype, remark}

}

func PrivateKeyA(host, stream string, typ int) string {

	cdnSwitch, _ := strconv.ParseInt(base.GetConfigPri().CdnSwitch, 10, 64)

	switch cdnSwitch {

	case 2:
		return PrivateKey_tx(host, stream, typ)
	case 4:
		return PrivateKey_ws(host, stream, typ)
	}

	return ""
}

func PrivateKey_tx(host, stream string, typ int) (url string) {
	config := base.GetConfigPri()
	pushKey := config.TxPushkey
	push := config.TxPush
	pull := config.TxPull

	streamKey := strings.Split(stream, ".")[0]
	liveCode := streamKey
	nowTime := time.GetCurrentUnix() + 3*60*60
	txTime := strconv.FormatInt(nowTime, 16)
	txSecret := cryp.MD5(pushKey + liveCode + txTime)
	safeUrl := "txSecret=" + txSecret + "&txTime=" + txTime

	if typ == 1 {
		url = "rtmp://" + push + "/live/" + liveCode + "?" + safeUrl
	} else {
		if host == "flv" {
			url = "https://" + pull + "/live/" + liveCode + ".flv"
		} else {
			url = "https://" + pull + "/live/" + liveCode + ".m3u8"
		}
	}

	return
}

func PrivateKey_ws(host, stream string, typ int) (url string) {
	config := base.GetConfigPri()
	push := config.WsPush
	pull := config.WsPull
	apn := config.WsApn

	streamKey := strings.Split(stream, ".")[0]
	liveCode := streamKey

	if typ == 1 {
		url = "rtmp://" + push + "/" + apn + "/" + liveCode
	} else {
		if host == "flv" {
			url = "https://" + pull + "/" + apn + "/" + liveCode + ".flv"
		} else {
			url = "https://" + pull + "/" + apn + "/" + liveCode + ".m3u8"
		}
	}

	return
}

func SmsSend(mobile, code string) {

	random := strconv.FormatInt(int64(math.Rand(100000, 999999)), 10)
	curTime := time.GetCurrentUnix()

	accesskey := constant.MsmAccessKey
	secretkey := constant.MsmSecretKey
	msg := strings.Replace(constant.MsmMsg, "{code}", code, 1)
	wholeUrl := "https://live.kewail.com/sms/v1/sendsinglesms?accesskey=" + accesskey + "&random=" + random

	data := map[string]interface{}{}

	data["tel"] = map[string]string{"nationcode": "86", "mobile": mobile}
	data["type"] = 0
	data["msg"] = msg
	data["sig"] = cryp.GetSHA256HashCode([]byte("secretkey=" + secretkey + "&random=" + random + "&time=" + strconv.FormatInt(curTime, 10) + "&mobile=" + mobile))
	data["time"] = curTime
	data["extend"] = ""
	data["ext"] = ""

	res, err := request.Curl(wholeUrl, data)

	if err != nil {
		panic(http_error.MsmNetworkErr)
	}

	resData := map[string]interface{}{}
	json.Unmarshal(res, &resData)

	result, _ := resData["result"]
	if stringify.ToString(result) != "0" {
		errmsg, _ := resData["errmsg"]
		panic(http_error.HttpError{
			ErrorCode: http_error.MsmNetworkErr.ErrorCode,
			ErrorMsg:  http_error.MsmNetworkErr.ErrorMsg + "," + stringify.ToString(errmsg),
		})
	}

}

func EmailSend(toEmail, code string) {

	random := strconv.FormatInt(int64(math.Rand(100000, 999999)), 10)
	curTime := time.GetCurrentUnix()

	accesskey := constant.MsmAccessKey
	secretkey := constant.MsmSecretKey
	fromEmail := constant.EmailFrom
	body := strings.Replace(constant.EmailBody, "{code}", code, 1)
	wholeUrl := "https://live.kewail.com/directmail/v1/singleSendMail?accesskey=" + accesskey + "&random=" + random

	data := map[string]interface{}{}
	data["sig"] = cryp.GetSHA256HashCode([]byte("secretkey=" + secretkey + "&random=" + random + "&time=" + strconv.FormatInt(curTime, 10) + "&fromEmail=" + fromEmail))
	data["ext"] = ""
	data["replyEmail"] = ""
	data["fromAlias"] = constant.EmailFromAlias
	data["htmlBody"] = body
	data["needToReply"] = false
	data["subject"] = constant.EmailTitle
	data["clickTrace"] = "0"
	data["time"] = curTime
	data["type"] = 0
	data["toEmail"] = toEmail
	data["fromEmail"] = fromEmail

	res, err := request.Curl(wholeUrl, data)

	if err != nil {
		panic(http_error.EmailNetworkErr)
	}

	resData := map[string]interface{}{}
	json.Unmarshal(res, &resData)

	result, _ := resData["result"]
	if stringify.ToString(result) != "0" {
		errmsg, _ := resData["errmsg"]
		panic(http_error.HttpError{
			ErrorCode: http_error.EmailNetworkErr.ErrorCode,
			ErrorMsg:  http_error.EmailNetworkErr.ErrorMsg + "," + stringify.ToString(errmsg),
		})
	}

}

func QrCode(url string) *[]byte {
	png, _ := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	return &png
}

func CreateCommonMixStream(stream1, stream2 string) {
	credential := common.NewCredential("AKIDwkJbcOG3wfbIitte8MLNAfuse6NTuZjc", "d06xdHekaDar0siKfkbWAh7VdZFq9DyY")
	client, _ := live.NewClient(credential, regions.Shanghai, profile.NewClientProfile())

	request := live.NewCreateCommonMixStreamRequest()

	outStream := stream1 + "_mix"
	request.MixStreamSessionId = &outStream

	var ImageLayer1 int64 = 1
	var ImageLayer2 int64 = 2
	var ImageLayer3 int64 = 3
	var ImageWidth float64 = 800
	var ImageWidth2 float64 = 400
	var ImageHeight float64 = 600
	var LocationX float64 = 400
	var InputType int64 = 3
	var Color string = "0x000000"
	var InputStreamName string = "Layer"

	request.InputStreamList = []*live.CommonMixInputParam{
		{
			InputStreamName: &InputStreamName,
			LayoutParams: &live.CommonMixLayoutParams{
				ImageLayer:  &ImageLayer1,
				ImageWidth:  &ImageWidth,
				ImageHeight: &ImageHeight,
				InputType:   &InputType,
				Color:       &Color,
			},
		},
		{
			InputStreamName: &stream1,
			LayoutParams: &live.CommonMixLayoutParams{
				ImageLayer:  &ImageLayer2,
				ImageWidth:  &ImageWidth2,
				ImageHeight: &ImageHeight,
			},
		},
		{
			InputStreamName: &stream2,
			LayoutParams: &live.CommonMixLayoutParams{
				ImageLayer:  &ImageLayer3,
				ImageWidth:  &ImageWidth2,
				ImageHeight: &ImageHeight,
				LocationX:   &LocationX,
			},
		},
	}

	var OutputStreamType int64 = 1
	request.OutputParams = &live.CommonMixOutputParams{
		OutputStreamName: &outStream,
		OutputStreamType: &OutputStreamType,
	}

	_, err := client.CreateCommonMixStream(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		panic("Mix:An API error has returned")
	}
	if err != nil {
		fmt.Println(err)
		panic("Mix:other errors")
	}
}

func CancelCommonMixStream(stream string) {

	stream += "_mix"
	credential := common.NewCredential("AKIDwkJbcOG3wfbIitte8MLNAfuse6NTuZjc", "d06xdHekaDar0siKfkbWAh7VdZFq9DyY")
	client, _ := live.NewClient(credential, regions.Shanghai, profile.NewClientProfile())

	request := live.NewCancelCommonMixStreamRequest()
	request.MixStreamSessionId = &stream
	_, err := client.CancelCommonMixStream(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		fmt.Printf("An API error has returned: %s", err)
	}
}

func RandName() string {
	tou := []string{"快乐", "冷静", "醉熏", "潇洒", "糊涂", "积极", "冷酷", "深情", "粗暴", "温柔", "可爱", "愉快", "义气", "认真", "威武", "帅气", "传统", "潇洒", "漂亮", "自然", "专一", "听话", "昏睡", "狂野", "等待", "搞怪", "幽默", "魁梧", "活泼", "开心", "高兴", "超帅", "留胡子", "坦率", "直率", "轻松", "痴情", "完美", "精明", "无聊", "有魅力", "丰富", "繁荣", "饱满", "炙热", "暴躁", "碧蓝", "俊逸", "英勇", "健忘", "故意", "无心", "土豪", "朴实", "兴奋", "幸福", "淡定", "不安", "阔达", "孤独", "独特", "疯狂", "时尚", "落后", "风趣", "忧伤", "大胆", "爱笑", "矮小", "健康", "合适", "玩命", "沉默", "斯文", "香蕉", "苹果", "鲤鱼", "鳗鱼", "任性", "细心", "粗心", "大意", "甜甜", "酷酷", "健壮", "英俊", "霸气", "阳光", "默默", "大力", "孝顺", "忧虑", "着急", "紧张", "善良", "凶狠", "害怕", "重要", "危机", "欢喜", "欣慰", "满意", "跳跃", "诚心", "称心", "如意", "怡然", "娇气", "无奈", "无语", "激动", "愤怒", "美好", "感动", "激情", "激昂", "震动", "虚拟", "超级", "寒冷", "精明", "明理", "犹豫", "忧郁", "寂寞", "奋斗", "勤奋", "现代", "过时", "稳重", "热情", "含蓄", "开放", "无辜", "多情", "纯真", "拉长", "热心", "从容", "体贴", "风中", "曾经", "追寻", "儒雅", "优雅", "开朗", "外向", "内向", "清爽", "文艺", "长情", "平常", "单身", "伶俐", "高大", "懦弱", "柔弱", "爱笑", "乐观", "耍酷", "酷炫", "神勇", "年轻", "唠叨", "瘦瘦", "无情", "包容", "顺心", "畅快", "舒适", "靓丽", "负责", "背后", "简单", "谦让", "彩色", "缥缈", "欢呼", "生动", "复杂", "慈祥", "仁爱", "魔幻", "虚幻", "淡然", "受伤", "雪白", "高高", "糟糕", "顺利", "闪闪", "羞涩", "缓慢", "迅速", "优秀", "聪明", "含糊", "俏皮", "淡淡", "坚强", "平淡", "欣喜", "能干", "灵巧", "友好", "机智", "机灵", "正直", "谨慎", "俭朴", "殷勤", "虚心", "辛勤", "自觉", "无私", "无限", "踏实", "老实", "现实", "可靠", "务实", "拼搏", "个性", "粗犷", "活力", "成就", "勤劳", "单纯", "落寞", "朴素", "悲凉", "忧心", "洁净", "清秀", "自由", "小巧", "单薄", "贪玩", "刻苦", "干净", "壮观", "和谐", "文静", "调皮", "害羞", "安详", "自信", "端庄", "坚定", "美满", "舒心", "温暖", "专注", "勤恳", "美丽", "腼腆", "优美", "甜美", "甜蜜", "整齐", "动人", "典雅", "尊敬", "舒服", "妩媚", "秀丽", "喜悦", "甜美", "彪壮", "强健", "大方", "俊秀", "聪慧", "迷人", "陶醉", "悦耳", "动听", "明亮", "结实", "魁梧", "标致", "清脆", "敏感", "光亮", "大气", "老迟到", "知性", "冷傲", "呆萌", "野性", "隐形", "笑点低", "微笑", "笨笨", "难过", "沉静", "火星上", "失眠", "安静", "纯情", "要减肥", "迷路", "烂漫", "哭泣", "贤惠", "苗条", "温婉", "发嗲", "会撒娇", "贪玩", "执着", "眯眯眼", "花痴", "想人陪", "眼睛大", "高贵", "傲娇", "心灵美", "爱撒娇", "细腻", "天真", "怕黑", "感性", "飘逸", "怕孤独", "忐忑", "高挑", "傻傻", "冷艳", "爱听歌", "还单身", "怕孤单", "懵懂"}
	do := []string{"的", "爱", "", "与", "给", "扯", "和", "用", "方", "打", "就", "迎", "向", "踢", "笑", "闻", "有", "等于", "保卫", "演变"}
	wei := []string{"嚓茶", "凉面", "便当", "毛豆", "花生", "可乐", "灯泡", "哈密瓜", "野狼", "背包", "眼神", "缘分", "雪碧", "人生", "牛排", "蚂蚁", "飞鸟", "灰狼", "斑马", "汉堡", "悟空", "巨人", "绿茶", "自行车", "保温杯", "大碗", "墨镜", "魔镜", "煎饼", "月饼", "月亮", "星星", "芝麻", "啤酒", "玫瑰", "大叔", "小伙", "哈密瓜，数据线", "太阳", "树叶", "芹菜", "黄蜂", "蜜粉", "蜜蜂", "信封", "西装", "外套", "裙子", "大象", "猫咪", "母鸡", "路灯", "蓝天", "白云", "星月", "彩虹", "微笑", "摩托", "板栗", "高山", "大地", "大树", "电灯胆", "砖头", "楼房", "水池", "鸡翅", "蜻蜓", "红牛", "咖啡", "机器猫", "枕头", "大船", "诺言", "钢笔", "刺猬", "天空", "飞机", "大炮", "冬天", "洋葱", "春天", "夏天", "秋天", "冬日", "航空", "毛衣", "豌豆", "黑米", "玉米", "眼睛", "老鼠", "白羊", "帅哥", "美女", "季节", "鲜花", "服饰", "裙子", "白开水", "秀发", "大山", "火车", "汽车", "歌曲", "舞蹈", "老师", "导师", "方盒", "大米", "麦片", "水杯", "水壶", "手套", "鞋子", "自行车", "鼠标", "手机", "电脑", "书本", "奇迹", "身影", "香烟", "夕阳", "台灯", "宝贝", "未来", "皮带", "钥匙", "心锁", "故事", "花瓣", "滑板", "画笔", "画板", "学姐", "店员", "电源", "饼干", "宝马", "过客", "大白", "时光", "石头", "钻石", "河马", "犀牛", "西牛", "绿草", "抽屉", "柜子", "往事", "寒风", "路人", "橘子", "耳机", "鸵鸟", "朋友", "苗条", "铅笔", "钢笔", "硬币", "热狗", "大侠", "御姐", "萝莉", "毛巾", "期待", "盼望", "白昼", "黑夜", "大门", "黑裤", "钢铁侠", "哑铃", "板凳", "枫叶", "荷花", "乌龟", "仙人掌", "衬衫", "大神", "草丛", "早晨", "心情", "茉莉", "流沙", "蜗牛", "战斗机", "冥王星", "猎豹", "棒球", "篮球", "乐曲", "电话", "网络", "世界", "中心", "鱼", "鸡", "狗", "老虎", "鸭子", "雨", "羽毛", "翅膀", "外套", "火", "丝袜", "书包", "钢笔", "冷风", "八宝粥", "烤鸡", "大雁", "音响", "招牌", "胡萝卜", "冰棍", "帽子", "菠萝", "蛋挞", "香水", "泥猴桃", "吐司", "溪流", "黄豆", "樱桃", "小鸽子", "小蝴蝶", "爆米花", "花卷", "小鸭子", "小海豚", "日记本", "小熊猫", "小懒猪", "小懒虫", "荔枝", "镜子", "曲奇", "金针菇", "小松鼠", "小虾米", "酒窝", "紫菜", "金鱼", "柚子", "果汁", "百褶裙", "项链", "帆布鞋", "火龙果", "奇异果", "煎蛋", "唇彩", "小土豆", "高跟鞋", "戒指", "雪糕", "睫毛", "铃铛", "手链", "香氛", "红酒", "月光", "酸奶", "银耳汤", "咖啡豆", "小蜜蜂", "小蚂蚁", "蜡烛", "棉花糖", "向日葵", "水蜜桃", "小蝴蝶", "小刺猬", "小丸子", "指甲油", "康乃馨", "糖豆", "薯片", "口红", "超短裙", "乌冬面", "冰淇淋", "棒棒糖", "长颈鹿", "豆芽", "发箍", "发卡", "发夹", "发带", "铃铛", "小馒头", "小笼包", "小甜瓜", "冬瓜", "香菇", "小兔子", "含羞草", "短靴", "睫毛膏", "小蘑菇", "跳跳糖", "小白菜", "草莓", "柠檬", "月饼", "百合", "纸鹤", "小天鹅", "云朵", "芒果", "面包", "海燕", "小猫咪", "龙猫", "唇膏", "鞋垫", "羊", "黑猫", "白猫", "万宝路", "金毛", "山水", "音响", "尊云", "西安"}
	tou_num := math.Rand(0, 331)
	do_num := math.Rand(0, 19)
	wei_num := math.Rand(0, 327)
	typ := math.Rand(0, 1)
	if typ == 0 {
		return tou[tou_num] + do[do_num] + wei[wei_num]
	} else {
		return wei[wei_num] + tou[tou_num]
	}
}

// txSecretId  腾讯云账户密钥对
// txSecretKey 腾讯云账户密钥对
// smsSdkAppId 添加应用后生成的实际 SDKAppID
// templateId  已审核通过的模板ID

func TxSmsSend(phone, code string) error {

	const txSecretId = "AKIDwkJbcOG3wfbIitte8MLNAfuse6NTuZjc"
	const txSecretKey = "d06xdHekaDar0siKfkbWAh7VdZFq9DyY"
	const smsSdkAppId = "1400787878"
	const templateId = "449739"

	/* 必要步骤：
	 * 实例化一个认证对象，入参需要传入腾讯云账户密钥对 secretId 和 secretKey
	 * 本示例采用从环境变量读取的方式，需要预先在环境变量中设置这两个值
	 * 您也可以直接在代码中写入密钥对，但需谨防泄露，不要将代码复制、上传或者分享给他人
	 * CAM 密匙查询: https://console.cloud.tencent.com/cam/capi
	 */
	credential := common.NewCredential(txSecretId, txSecretKey)
	/* 非必要步骤:
	 * 实例化一个客户端配置对象，可以指定超时时间等配置 */
	cpf := profile.NewClientProfile()

	/* SDK 默认使用 POST 方法
	 * 如需使用 GET 方法，可以在此处设置，但 GET 方法无法处理较大的请求 */
	cpf.HttpProfile.ReqMethod = "POST"

	/* SDK 有默认的超时时间，非必要请不要进行调整
	 * 如有需要请在代码中查阅以获取最新的默认值 */
	//cpf.HttpProfile.ReqTimeout = 5

	/* SDK 会自动指定域名，通常无需指定域名，但访问金融区的服务时必须手动指定域名
	 * 例如 SMS 的上海金融区域名为 sms.ap-shanghai-fsi.tencentcloudapi.com */
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"

	/* SDK 默认用 TC3-HMAC-SHA256 进行签名，非必要请不要修改该字段 */
	cpf.SignMethod = "HmacSHA1"

	/* 实例化 SMS 的 client 对象
	 * 第二个参数是地域信息，可以直接填写字符串 ap-guangzhou，或者引用预设的常量 */
	client, _ := sms.NewClient(credential, "ap-shanghai", cpf)

	/* 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	   * 您可以直接查询 SDK 源码确定接口有哪些属性可以设置
	    * 属性可能是基本类型，也可能引用了另一个数据结构
	    * 推荐使用 IDE 进行开发，可以方便地跳转查阅各个接口和数据结构的文档说明 */
	request := sms.NewSendSmsRequest()

	/* 基本类型的设置:
	 * SDK 采用的是指针风格指定参数，即使对于基本类型也需要用指针来对参数赋值。
	 * SDK 提供对基本类型的指针引用封装函数
	 * 帮助链接：
	 * 短信控制台：https://console.cloud.tencent.com/smsv2
	 * sms helper：https://cloud.tencent.com/document/product/382/3773
	 */

	/* 短信应用 ID: 在 [短信控制台] 添加应用后生成的实际 SDKAppID，例如1400006666 */
	request.SmsSdkAppId = common.StringPtr(smsSdkAppId)
	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，可登录 [短信控制台] 查看签名信息 */
	// request.SignName = common.StringPtr("xxx")
	/* 国际/港澳台短信 senderid: 国内短信填空，默认未开通，如需开通请联系 [sms helper] */
	// request.SenderId = common.StringPtr("")
	/* 用户的 session 内容: 可以携带用户侧 ID 等上下文信息，server 会原样返回 */
	// request.SessionContext = common.StringPtr("xxx")
	/* 短信码号扩展号: 默认未开通，如需开通请联系 [sms helper] */
	// request.ExtendCode = common.StringPtr("")
	/* 模板参数: 若无模板参数，则设置为空*/
	request.TemplateParamSet = common.StringPtrs([]string{code})
	/* 模板 ID: 必须填写已审核通过的模板 ID，可登录 [短信控制台] 查看模板 ID */
	request.TemplateId = common.StringPtr(templateId)
	/* 下发手机号码，采用 e.164 标准，+[国家或地区码][手机号]
	 * 例如+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	request.PhoneNumberSet = common.StringPtrs([]string{"+86" + phone})

	// 通过 client 对象调用想要访问的接口，需要传入请求对象
	_, err := client.SendSms(request)
	// 处理异常
	if e, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		panic(http_error.HttpError{
			ErrorCode: http_error.MsmNetworkErr.ErrorCode,
			ErrorMsg:  http_error.MsmNetworkErr.ErrorMsg + "," + stringify.ToString(e.Code) + ":" + e.Message,
		})
	}
	// 非 SDK 异常，直接失败。实际代码中可以加入其他的处理
	if err != nil {
		panic(http_error.MsmNetworkErr)
	}

	return nil
}
