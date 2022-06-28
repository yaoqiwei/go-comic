package constant

const WsSignKey = "76576076c1f5f657b634e966c8836a06"

const MsmAccessKey = ""
const MsmSecretKey = ""
const MsmMsg = "【好看视频】您的验证码是：{code}。"
const EmailFrom = "admin@test.pgc.api.yimisaas.com"
const EmailFromAlias = "好看视频邮件管理员"
const EmailBody = "【好看视频】您的验证码是：{code}。"
const EmailTitle = "【好看视频】验证码"

const ApiAesKey = "46cc793c53dc451b"

var SDK = map[string]map[string]string{
	"ios": {
		"codingmode":         "2",     //编码 0自动，1软编，2硬编
		"resolution":         "5",     //分辨率
		"isauto":             "1",     //是否自适应 0否1是
		"fps":                "20",    //帧数
		"fps_min":            "20",    //最低帧数
		"fps_max":            "30",    //最高帧数
		"gop":                "3",     //关键帧间隔
		"bitrate":            "800",   //初始码率  kbps
		"bitrate_min":        "800",   //最低码率
		"bitrate_max":        "1200",  //最高码率
		"audiorate":          "44100", //音频采样率  Hz
		"audiobitrate":       "48",    //音频码率 kbps
		"preview_fps":        "15",    //预览帧数
		"preview_resolution": "1",     //预览分辨率
	},
	"android": {
		"codingmode":         "3",     //编码 1自动，3软编，2硬编
		"resolution":         "1",     //分辨率
		"isauto":             "1",     //是否自适应 0否1是
		"fps":                "20",    //帧数
		"fps_min":            "20",    //最低帧数
		"fps_max":            "30",    //最高帧数
		"gop":                "3",     //关键帧间隔
		"bitrate":            "500",   //初始码率  kbps
		"bitrate_min":        "500",   //最低码率
		"bitrate_max":        "800",   //最高码率
		"audiorate":          "44100", //音频采样率  Hz
		"audiobitrate":       "48",    //音频码率 kbps
		"preview_fps":        "15",    //预览帧数
		"preview_resolution": "1",     //预览分辨率
	},
}
