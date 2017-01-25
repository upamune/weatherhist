package weatherhist

// auto generated code

type ObservationID string

func GetObservatoriesKey(id string, groupNumber string) ObservationID {
	return ObservationID(id + "_" + groupNumber)
}

func GetAllObservatories() Observatries {
	return allObservatories
}

type Observation struct {
	ID          ObservationID
	GroupNumber string
	Name        string
}

type Observatries map[ObservationID]Observation

var allObservatories = Observatries{
	"0002_11": Observation{
		ID:          ObservationID("0002"),
		GroupNumber: "11",
		Name:        "沓形",
	},
	"0003_11": Observation{
		ID:          ObservationID("0003"),
		GroupNumber: "11",
		Name:        "浜頓別",
	},
	"0005_11": Observation{
		ID:          ObservationID("0005"),
		GroupNumber: "11",
		Name:        "歌登",
	},
	"1051_11": Observation{
		ID:          ObservationID("1051"),
		GroupNumber: "11",
		Name:        "中頓別",
	},
	"1054_11": Observation{
		ID:          ObservationID("1054"),
		GroupNumber: "11",
		Name:        "豊富",
	},
	"1203_11": Observation{
		ID:          ObservationID("1203"),
		GroupNumber: "11",
		Name:        "沼川",
	},
	"1207_11": Observation{
		ID:          ObservationID("1207"),
		GroupNumber: "11",
		Name:        "船泊",
	},
	"1284_11": Observation{
		ID:          ObservationID("1284"),
		GroupNumber: "11",
		Name:        "宗谷岬",
	},
	"1285_11": Observation{
		ID:          ObservationID("1285"),
		GroupNumber: "11",
		Name:        "浜鬼志別",
	},
	"1512_11": Observation{
		ID:          ObservationID("1512"),
		GroupNumber: "11",
		Name:        "本泊",
	},
	"1528_11": Observation{
		ID:          ObservationID("1528"),
		GroupNumber: "11",
		Name:        "声問",
	},
	"1546_11": Observation{
		ID:          ObservationID("1546"),
		GroupNumber: "11",
		Name:        "礼文",
	},
	"1573_11": Observation{
		ID:          ObservationID("1573"),
		GroupNumber: "11",
		Name:        "幌延",
	},
	"1593_11": Observation{
		ID:          ObservationID("1593"),
		GroupNumber: "11",
		Name:        "幌泊",
	},
	"47401_11": Observation{
		ID:          ObservationID("47401"),
		GroupNumber: "11",
		Name:        "稚内",
	},
	"47402_11": Observation{
		ID:          ObservationID("47402"),
		GroupNumber: "11",
		Name:        "北見枝幸",
	},
	"0006_12": Observation{
		ID:          ObservationID("0006"),
		GroupNumber: "12",
		Name:        "美深",
	},
	"0007_12": Observation{
		ID:          ObservationID("0007"),
		GroupNumber: "12",
		Name:        "ピヤシリ山",
	},
	"0008_12": Observation{
		ID:          ObservationID("0008"),
		GroupNumber: "12",
		Name:        "名寄",
	},
	"0009_12": Observation{
		ID:          ObservationID("0009"),
		GroupNumber: "12",
		Name:        "西風連",
	},
	"0010_12": Observation{
		ID:          ObservationID("0010"),
		GroupNumber: "12",
		Name:        "下川",
	},
	"0011_12": Observation{
		ID:          ObservationID("0011"),
		GroupNumber: "12",
		Name:        "士別",
	},
	"0012_12": Observation{
		ID:          ObservationID("0012"),
		GroupNumber: "12",
		Name:        "犬牛別峠",
	},
	"0013_12": Observation{
		ID:          ObservationID("0013"),
		GroupNumber: "12",
		Name:        "雄鷹峰",
	},
	"0014_12": Observation{
		ID:          ObservationID("0014"),
		GroupNumber: "12",
		Name:        "蓬莱山",
	},
	"0015_12": Observation{
		ID:          ObservationID("0015"),
		GroupNumber: "12",
		Name:        "上川",
	},
	"0017_12": Observation{
		ID:          ObservationID("0017"),
		GroupNumber: "12",
		Name:        "瑞穂",
	},
	"0018_12": Observation{
		ID:          ObservationID("0018"),
		GroupNumber: "12",
		Name:        "終沢",
	},
	"0019_12": Observation{
		ID:          ObservationID("0019"),
		GroupNumber: "12",
		Name:        "旭岳",
	},
	"0020_12": Observation{
		ID:          ObservationID("0020"),
		GroupNumber: "12",
		Name:        "白金",
	},
	"0021_12": Observation{
		ID:          ObservationID("0021"),
		GroupNumber: "12",
		Name:        "富良野",
	},
	"0022_12": Observation{
		ID:          ObservationID("0022"),
		GroupNumber: "12",
		Name:        "前富良野岳",
	},
	"0023_12": Observation{
		ID:          ObservationID("0023"),
		GroupNumber: "12",
		Name:        "前時雨山",
	},
	"1047_12": Observation{
		ID:          ObservationID("1047"),
		GroupNumber: "12",
		Name:        "朝日",
	},
	"1048_12": Observation{
		ID:          ObservationID("1048"),
		GroupNumber: "12",
		Name:        "比布",
	},
	"1049_12": Observation{
		ID:          ObservationID("1049"),
		GroupNumber: "12",
		Name:        "層雲峡",
	},
	"1052_12": Observation{
		ID:          ObservationID("1052"),
		GroupNumber: "12",
		Name:        "美瑛",
	},
	"1053_12": Observation{
		ID:          ObservationID("1053"),
		GroupNumber: "12",
		Name:        "幌加内",
	},
	"1055_12": Observation{
		ID:          ObservationID("1055"),
		GroupNumber: "12",
		Name:        "和寒",
	},
	"1184_12": Observation{
		ID:          ObservationID("1184"),
		GroupNumber: "12",
		Name:        "東川",
	},
	"1185_12": Observation{
		ID:          ObservationID("1185"),
		GroupNumber: "12",
		Name:        "志比内",
	},
	"1188_12": Observation{
		ID:          ObservationID("1188"),
		GroupNumber: "12",
		Name:        "幾寅",
	},
	"1189_12": Observation{
		ID:          ObservationID("1189"),
		GroupNumber: "12",
		Name:        "占冠",
	},
	"1190_12": Observation{
		ID:          ObservationID("1190"),
		GroupNumber: "12",
		Name:        "上富良野",
	},
	"1191_12": Observation{
		ID:          ObservationID("1191"),
		GroupNumber: "12",
		Name:        "江丹別",
	},
	"1197_12": Observation{
		ID:          ObservationID("1197"),
		GroupNumber: "12",
		Name:        "音威子府",
	},
	"1201_12": Observation{
		ID:          ObservationID("1201"),
		GroupNumber: "12",
		Name:        "中川",
	},
	"1279_12": Observation{
		ID:          ObservationID("1279"),
		GroupNumber: "12",
		Name:        "朱鞠内",
	},
	"1280_12": Observation{
		ID:          ObservationID("1280"),
		GroupNumber: "12",
		Name:        "麓郷",
	},
	"1458_12": Observation{
		ID:          ObservationID("1458"),
		GroupNumber: "12",
		Name:        "東神楽",
	},
	"1570_12": Observation{
		ID:          ObservationID("1570"),
		GroupNumber: "12",
		Name:        "小車",
	},
	"1572_12": Observation{
		ID:          ObservationID("1572"),
		GroupNumber: "12",
		Name:        "剣淵",
	},
	"1574_12": Observation{
		ID:          ObservationID("1574"),
		GroupNumber: "12",
		Name:        "金山",
	},
	"47407_12": Observation{
		ID:          ObservationID("47407"),
		GroupNumber: "12",
		Name:        "旭川",
	},
	"0024_13": Observation{
		ID:          ObservationID("0024"),
		GroupNumber: "13",
		Name:        "天塩",
	},
	"0025_13": Observation{
		ID:          ObservationID("0025"),
		GroupNumber: "13",
		Name:        "遠別",
	},
	"0028_13": Observation{
		ID:          ObservationID("0028"),
		GroupNumber: "13",
		Name:        "増毛",
	},
	"1112_13": Observation{
		ID:          ObservationID("1112"),
		GroupNumber: "13",
		Name:        "古丹別",
	},
	"1213_13": Observation{
		ID:          ObservationID("1213"),
		GroupNumber: "13",
		Name:        "焼尻",
	},
	"1216_13": Observation{
		ID:          ObservationID("1216"),
		GroupNumber: "13",
		Name:        "初山別",
	},
	"1217_13": Observation{
		ID:          ObservationID("1217"),
		GroupNumber: "13",
		Name:        "達布",
	},
	"1218_13": Observation{
		ID:          ObservationID("1218"),
		GroupNumber: "13",
		Name:        "幌糠",
	},
	"47404_13": Observation{
		ID:          ObservationID("47404"),
		GroupNumber: "13",
		Name:        "羽幌",
	},
	"47406_13": Observation{
		ID:          ObservationID("47406"),
		GroupNumber: "13",
		Name:        "留萌",
	},
	"0029_14": Observation{
		ID:          ObservationID("0029"),
		GroupNumber: "14",
		Name:        "阿曽岩",
	},
	"0030_14": Observation{
		ID:          ObservationID("0030"),
		GroupNumber: "14",
		Name:        "新篠津",
	},
	"0031_14": Observation{
		ID:          ObservationID("0031"),
		GroupNumber: "14",
		Name:        "山口",
	},
	"0032_14": Observation{
		ID:          ObservationID("0032"),
		GroupNumber: "14",
		Name:        "手稲山",
	},
	"0034_14": Observation{
		ID:          ObservationID("0034"),
		GroupNumber: "14",
		Name:        "小金湯",
	},
	"0035_14": Observation{
		ID:          ObservationID("0035"),
		GroupNumber: "14",
		Name:        "恵庭島松",
	},
	"0036_14": Observation{
		ID:          ObservationID("0036"),
		GroupNumber: "14",
		Name:        "島松山",
	},
	"0037_14": Observation{
		ID:          ObservationID("0037"),
		GroupNumber: "14",
		Name:        "支笏湖畔",
	},
	"1076_14": Observation{
		ID:          ObservationID("1076"),
		GroupNumber: "14",
		Name:        "西野幌",
	},
	"1085_14": Observation{
		ID:          ObservationID("1085"),
		GroupNumber: "14",
		Name:        "石狩",
	},
	"1193_14": Observation{
		ID:          ObservationID("1193"),
		GroupNumber: "14",
		Name:        "浜益",
	},
	"1194_14": Observation{
		ID:          ObservationID("1194"),
		GroupNumber: "14",
		Name:        "厚田",
	},
	"1459_14": Observation{
		ID:          ObservationID("1459"),
		GroupNumber: "14",
		Name:        "千歳",
	},
	"1507_14": Observation{
		ID:          ObservationID("1507"),
		GroupNumber: "14",
		Name:        "江別",
	},
	"47412_14": Observation{
		ID:          ObservationID("47412"),
		GroupNumber: "14",
		Name:        "札幌",
	},
	"0038_15": Observation{
		ID:          ObservationID("0038"),
		GroupNumber: "15",
		Name:        "石狩沼田",
	},
	"0039_15": Observation{
		ID:          ObservationID("0039"),
		GroupNumber: "15",
		Name:        "深川",
	},
	"0040_15": Observation{
		ID:          ObservationID("0040"),
		GroupNumber: "15",
		Name:        "常盤山",
	},
	"0041_15": Observation{
		ID:          ObservationID("0041"),
		GroupNumber: "15",
		Name:        "滝川",
	},
	"0042_15": Observation{
		ID:          ObservationID("0042"),
		GroupNumber: "15",
		Name:        "神威岳",
	},
	"0043_15": Observation{
		ID:          ObservationID("0043"),
		GroupNumber: "15",
		Name:        "芦別",
	},
	"0044_15": Observation{
		ID:          ObservationID("0044"),
		GroupNumber: "15",
		Name:        "美唄",
	},
	"0046_15": Observation{
		ID:          ObservationID("0046"),
		GroupNumber: "15",
		Name:        "下桂沢",
	},
	"0047_15": Observation{
		ID:          ObservationID("0047"),
		GroupNumber: "15",
		Name:        "長沼",
	},
	"0048_15": Observation{
		ID:          ObservationID("0048"),
		GroupNumber: "15",
		Name:        "丁未山",
	},
	"0049_15": Observation{
		ID:          ObservationID("0049"),
		GroupNumber: "15",
		Name:        "夕張",
	},
	"0050_15": Observation{
		ID:          ObservationID("0050"),
		GroupNumber: "15",
		Name:        "登川山",
	},
	"0999_15": Observation{
		ID:          ObservationID("0999"),
		GroupNumber: "15",
		Name:        "月形",
	},
	"1068_15": Observation{
		ID:          ObservationID("1068"),
		GroupNumber: "15",
		Name:        "栗沢",
	},
	"1270_15": Observation{
		ID:          ObservationID("1270"),
		GroupNumber: "15",
		Name:        "晩生内",
	},
	"1271_15": Observation{
		ID:          ObservationID("1271"),
		GroupNumber: "15",
		Name:        "桜山",
	},
	"1287_15": Observation{
		ID:          ObservationID("1287"),
		GroupNumber: "15",
		Name:        "空知吉野",
	},
	"1400_15": Observation{
		ID:          ObservationID("1400"),
		GroupNumber: "15",
		Name:        "浦臼",
	},
	"1401_15": Observation{
		ID:          ObservationID("1401"),
		GroupNumber: "15",
		Name:        "鹿島",
	},
	"1402_15": Observation{
		ID:          ObservationID("1402"),
		GroupNumber: "15",
		Name:        "沼の沢",
	},
	"1419_15": Observation{
		ID:          ObservationID("1419"),
		GroupNumber: "15",
		Name:        "新城",
	},
	"1566_15": Observation{
		ID:          ObservationID("1566"),
		GroupNumber: "15",
		Name:        "赤平",
	},
	"1567_15": Observation{
		ID:          ObservationID("1567"),
		GroupNumber: "15",
		Name:        "雨竜",
	},
	"47413_15": Observation{
		ID:          ObservationID("47413"),
		GroupNumber: "15",
		Name:        "岩見沢",
	},
	"0051_16": Observation{
		ID:          ObservationID("0051"),
		GroupNumber: "16",
		Name:        "美国",
	},
	"0052_16": Observation{
		ID:          ObservationID("0052"),
		GroupNumber: "16",
		Name:        "余市",
	},
	"0054_16": Observation{
		ID:          ObservationID("0054"),
		GroupNumber: "16",
		Name:        "共和",
	},
	"0055_16": Observation{
		ID:          ObservationID("0055"),
		GroupNumber: "16",
		Name:        "蘭越",
	},
	"0057_16": Observation{
		ID:          ObservationID("0057"),
		GroupNumber: "16",
		Name:        "脇方台地",
	},
	"0059_16": Observation{
		ID:          ObservationID("0059"),
		GroupNumber: "16",
		Name:        "桂台",
	},
	"0060_16": Observation{
		ID:          ObservationID("0060"),
		GroupNumber: "16",
		Name:        "喜茂別",
	},
	"0061_16": Observation{
		ID:          ObservationID("0061"),
		GroupNumber: "16",
		Name:        "黒松内",
	},
	"1072_16": Observation{
		ID:          ObservationID("1072"),
		GroupNumber: "16",
		Name:        "真狩",
	},
	"1091_16": Observation{
		ID:          ObservationID("1091"),
		GroupNumber: "16",
		Name:        "赤井川",
	},
	"1208_16": Observation{
		ID:          ObservationID("1208"),
		GroupNumber: "16",
		Name:        "神恵内",
	},
	"1615_16": Observation{
		ID:          ObservationID("1615"),
		GroupNumber: "16",
		Name:        "ニセコ",
	},
	"47411_16": Observation{
		ID:          ObservationID("47411"),
		GroupNumber: "16",
		Name:        "小樽",
	},
	"47421_16": Observation{
		ID:          ObservationID("47421"),
		GroupNumber: "16",
		Name:        "寿都",
	},
	"47433_16": Observation{
		ID:          ObservationID("47433"),
		GroupNumber: "16",
		Name:        "倶知安",
	},
	"0063_17": Observation{
		ID:          ObservationID("0063"),
		GroupNumber: "17",
		Name:        "興部",
	},
	"0065_17": Observation{
		ID:          ObservationID("0065"),
		GroupNumber: "17",
		Name:        "湧別",
	},
	"0066_17": Observation{
		ID:          ObservationID("0066"),
		GroupNumber: "17",
		Name:        "滝上",
	},
	"0067_17": Observation{
		ID:          ObservationID("0067"),
		GroupNumber: "17",
		Name:        "上藻別",
	},
	"0068_17": Observation{
		ID:          ObservationID("0068"),
		GroupNumber: "17",
		Name:        "常呂",
	},
	"0069_17": Observation{
		ID:          ObservationID("0069"),
		GroupNumber: "17",
		Name:        "遠軽",
	},
	"0070_17": Observation{
		ID:          ObservationID("0070"),
		GroupNumber: "17",
		Name:        "佐呂間",
	},
	"0072_17": Observation{
		ID:          ObservationID("0072"),
		GroupNumber: "17",
		Name:        "生田原",
	},
	"0073_17": Observation{
		ID:          ObservationID("0073"),
		GroupNumber: "17",
		Name:        "仁頃山",
	},
	"0074_17": Observation{
		ID:          ObservationID("0074"),
		GroupNumber: "17",
		Name:        "北見",
	},
	"0075_17": Observation{
		ID:          ObservationID("0075"),
		GroupNumber: "17",
		Name:        "小清水",
	},
	"0076_17": Observation{
		ID:          ObservationID("0076"),
		GroupNumber: "17",
		Name:        "斜里",
	},
	"0077_17": Observation{
		ID:          ObservationID("0077"),
		GroupNumber: "17",
		Name:        "留辺蘂",
	},
	"0078_17": Observation{
		ID:          ObservationID("0078"),
		GroupNumber: "17",
		Name:        "留辺蘂山",
	},
	"0079_17": Observation{
		ID:          ObservationID("0079"),
		GroupNumber: "17",
		Name:        "美幌",
	},
	"0080_17": Observation{
		ID:          ObservationID("0080"),
		GroupNumber: "17",
		Name:        "藻琴山",
	},
	"0081_17": Observation{
		ID:          ObservationID("0081"),
		GroupNumber: "17",
		Name:        "チミケップ山",
	},
	"0082_17": Observation{
		ID:          ObservationID("0082"),
		GroupNumber: "17",
		Name:        "津別",
	},
	"0083_17": Observation{
		ID:          ObservationID("0083"),
		GroupNumber: "17",
		Name:        "北見中山",
	},
	"0981_17": Observation{
		ID:          ObservationID("0981"),
		GroupNumber: "17",
		Name:        "宇登呂",
	},
	"1073_17": Observation{
		ID:          ObservationID("1073"),
		GroupNumber: "17",
		Name:        "西興部",
	},
	"1092_17": Observation{
		ID:          ObservationID("1092"),
		GroupNumber: "17",
		Name:        "白滝",
	},
	"1209_17": Observation{
		ID:          ObservationID("1209"),
		GroupNumber: "17",
		Name:        "境野",
	},
	"1263_17": Observation{
		ID:          ObservationID("1263"),
		GroupNumber: "17",
		Name:        "東藻琴",
	},
	"1420_17": Observation{
		ID:          ObservationID("1420"),
		GroupNumber: "17",
		Name:        "丸瀬布",
	},
	"1460_17": Observation{
		ID:          ObservationID("1460"),
		GroupNumber: "17",
		Name:        "女満別",
	},
	"1487_17": Observation{
		ID:          ObservationID("1487"),
		GroupNumber: "17",
		Name:        "紋別小向",
	},
	"1568_17": Observation{
		ID:          ObservationID("1568"),
		GroupNumber: "17",
		Name:        "津別二又",
	},
	"1569_17": Observation{
		ID:          ObservationID("1569"),
		GroupNumber: "17",
		Name:        "置戸常元",
	},
	"1617_17": Observation{
		ID:          ObservationID("1617"),
		GroupNumber: "17",
		Name:        "山園",
	},
	"47405_17": Observation{
		ID:          ObservationID("47405"),
		GroupNumber: "17",
		Name:        "雄武",
	},
	"47409_17": Observation{
		ID:          ObservationID("47409"),
		GroupNumber: "17",
		Name:        "網走",
	},
	"47435_17": Observation{
		ID:          ObservationID("47435"),
		GroupNumber: "17",
		Name:        "紋別",
	},
	"0084_18": Observation{
		ID:          ObservationID("0084"),
		GroupNumber: "18",
		Name:        "羅臼",
	},
	"0085_18": Observation{
		ID:          ObservationID("0085"),
		GroupNumber: "18",
		Name:        "標津",
	},
	"0086_18": Observation{
		ID:          ObservationID("0086"),
		GroupNumber: "18",
		Name:        "中標津",
	},
	"0087_18": Observation{
		ID:          ObservationID("0087"),
		GroupNumber: "18",
		Name:        "計根別",
	},
	"0088_18": Observation{
		ID:          ObservationID("0088"),
		GroupNumber: "18",
		Name:        "別海",
	},
	"1056_18": Observation{
		ID:          ObservationID("1056"),
		GroupNumber: "18",
		Name:        "納沙布",
	},
	"1196_18": Observation{
		ID:          ObservationID("1196"),
		GroupNumber: "18",
		Name:        "厚床",
	},
	"1266_18": Observation{
		ID:          ObservationID("1266"),
		GroupNumber: "18",
		Name:        "糸櫛別",
	},
	"1489_18": Observation{
		ID:          ObservationID("1489"),
		GroupNumber: "18",
		Name:        "根室中標津",
	},
	"1547_18": Observation{
		ID:          ObservationID("1547"),
		GroupNumber: "18",
		Name:        "上標津",
	},
	"47420_18": Observation{
		ID:          ObservationID("47420"),
		GroupNumber: "18",
		Name:        "根室",
	},
	"0090_19": Observation{
		ID:          ObservationID("0090"),
		GroupNumber: "19",
		Name:        "川湯",
	},
	"0091_19": Observation{
		ID:          ObservationID("0091"),
		GroupNumber: "19",
		Name:        "阿寒湖畔",
	},
	"0092_19": Observation{
		ID:          ObservationID("0092"),
		GroupNumber: "19",
		Name:        "弟子屈",
	},
	"0093_19": Observation{
		ID:          ObservationID("0093"),
		GroupNumber: "19",
		Name:        "標茶",
	},
	"0094_19": Observation{
		ID:          ObservationID("0094"),
		GroupNumber: "19",
		Name:        "幌呂台地",
	},
	"0095_19": Observation{
		ID:          ObservationID("0095"),
		GroupNumber: "19",
		Name:        "中徹別",
	},
	"0096_19": Observation{
		ID:          ObservationID("0096"),
		GroupNumber: "19",
		Name:        "阿寒",
	},
	"0097_19": Observation{
		ID:          ObservationID("0097"),
		GroupNumber: "19",
		Name:        "霧多布",
	},
	"0098_19": Observation{
		ID:          ObservationID("0098"),
		GroupNumber: "19",
		Name:        "白糠",
	},
	"1107_19": Observation{
		ID:          ObservationID("1107"),
		GroupNumber: "19",
		Name:        "二俣",
	},
	"1186_19": Observation{
		ID:          ObservationID("1186"),
		GroupNumber: "19",
		Name:        "鶴居",
	},
	"1187_19": Observation{
		ID:          ObservationID("1187"),
		GroupNumber: "19",
		Name:        "弟子屈",
	},
	"1192_19": Observation{
		ID:          ObservationID("1192"),
		GroupNumber: "19",
		Name:        "太田",
	},
	"1195_19": Observation{
		ID:          ObservationID("1195"),
		GroupNumber: "19",
		Name:        "榊町",
	},
	"1283_19": Observation{
		ID:          ObservationID("1283"),
		GroupNumber: "19",
		Name:        "知方学",
	},
	"1403_19": Observation{
		ID:          ObservationID("1403"),
		GroupNumber: "19",
		Name:        "塘路",
	},
	"1417_19": Observation{
		ID:          ObservationID("1417"),
		GroupNumber: "19",
		Name:        "茶内原野",
	},
	"1461_19": Observation{
		ID:          ObservationID("1461"),
		GroupNumber: "19",
		Name:        "鶴丘",
	},
	"47418_19": Observation{
		ID:          ObservationID("47418"),
		GroupNumber: "19",
		Name:        "釧路",
	},
	"0100_20": Observation{
		ID:          ObservationID("0100"),
		GroupNumber: "20",
		Name:        "軍艦山",
	},
	"0101_20": Observation{
		ID:          ObservationID("0101"),
		GroupNumber: "20",
		Name:        "陸別",
	},
	"0102_20": Observation{
		ID:          ObservationID("0102"),
		GroupNumber: "20",
		Name:        "殖産高地",
	},
	"0103_20": Observation{
		ID:          ObservationID("0103"),
		GroupNumber: "20",
		Name:        "糠平",
	},
	"0104_20": Observation{
		ID:          ObservationID("0104"),
		GroupNumber: "20",
		Name:        "小坂山",
	},
	"0105_20": Observation{
		ID:          ObservationID("0105"),
		GroupNumber: "20",
		Name:        "上足寄",
	},
	"0106_20": Observation{
		ID:          ObservationID("0106"),
		GroupNumber: "20",
		Name:        "西ヌプカウシ山",
	},
	"0107_20": Observation{
		ID:          ObservationID("0107"),
		GroupNumber: "20",
		Name:        "上士幌",
	},
	"0108_20": Observation{
		ID:          ObservationID("0108"),
		GroupNumber: "20",
		Name:        "幌安山",
	},
	"0109_20": Observation{
		ID:          ObservationID("0109"),
		GroupNumber: "20",
		Name:        "足寄",
	},
	"0110_20": Observation{
		ID:          ObservationID("0110"),
		GroupNumber: "20",
		Name:        "佐幌岳",
	},
	"0111_20": Observation{
		ID:          ObservationID("0111"),
		GroupNumber: "20",
		Name:        "三角山",
	},
	"0112_20": Observation{
		ID:          ObservationID("0112"),
		GroupNumber: "20",
		Name:        "本別",
	},
	"0113_20": Observation{
		ID:          ObservationID("0113"),
		GroupNumber: "20",
		Name:        "新得",
	},
	"0114_20": Observation{
		ID:          ObservationID("0114"),
		GroupNumber: "20",
		Name:        "鹿追",
	},
	"0115_20": Observation{
		ID:          ObservationID("0115"),
		GroupNumber: "20",
		Name:        "芽室",
	},
	"0117_20": Observation{
		ID:          ObservationID("0117"),
		GroupNumber: "20",
		Name:        "池田",
	},
	"0118_20": Observation{
		ID:          ObservationID("0118"),
		GroupNumber: "20",
		Name:        "浦幌",
	},
	"0119_20": Observation{
		ID:          ObservationID("0119"),
		GroupNumber: "20",
		Name:        "帯広岳",
	},
	"0120_20": Observation{
		ID:          ObservationID("0120"),
		GroupNumber: "20",
		Name:        "ひょうたん沢",
	},
	"0121_20": Observation{
		ID:          ObservationID("0121"),
		GroupNumber: "20",
		Name:        "大樹",
	},
	"1099_20": Observation{
		ID:          ObservationID("1099"),
		GroupNumber: "20",
		Name:        "三国山",
	},
	"1204_20": Observation{
		ID:          ObservationID("1204"),
		GroupNumber: "20",
		Name:        "駒場",
	},
	"1210_20": Observation{
		ID:          ObservationID("1210"),
		GroupNumber: "20",
		Name:        "上札内",
	},
	"1211_20": Observation{
		ID:          ObservationID("1211"),
		GroupNumber: "20",
		Name:        "更別",
	},
	"1214_20": Observation{
		ID:          ObservationID("1214"),
		GroupNumber: "20",
		Name:        "大津",
	},
	"1286_20": Observation{
		ID:          ObservationID("1286"),
		GroupNumber: "20",
		Name:        "糠内",
	},
	"1404_20": Observation{
		ID:          ObservationID("1404"),
		GroupNumber: "20",
		Name:        "小利別",
	},
	"1405_20": Observation{
		ID:          ObservationID("1405"),
		GroupNumber: "20",
		Name:        "押帯",
	},
	"1409_20": Observation{
		ID:          ObservationID("1409"),
		GroupNumber: "20",
		Name:        "柏倉",
	},
	"1410_20": Observation{
		ID:          ObservationID("1410"),
		GroupNumber: "20",
		Name:        "上螺湾",
	},
	"1418_20": Observation{
		ID:          ObservationID("1418"),
		GroupNumber: "20",
		Name:        "留真",
	},
	"1510_20": Observation{
		ID:          ObservationID("1510"),
		GroupNumber: "20",
		Name:        "帯広泉",
	},
	"1571_20": Observation{
		ID:          ObservationID("1571"),
		GroupNumber: "20",
		Name:        "三股",
	},
	"47417_20": Observation{
		ID:          ObservationID("47417"),
		GroupNumber: "20",
		Name:        "帯広",
	},
	"47440_20": Observation{
		ID:          ObservationID("47440"),
		GroupNumber: "20",
		Name:        "広尾",
	},
	"0123_21": Observation{
		ID:          ObservationID("0123"),
		GroupNumber: "21",
		Name:        "安平",
	},
	"0124_21": Observation{
		ID:          ObservationID("0124"),
		GroupNumber: "21",
		Name:        "厚真",
	},
	"0125_21": Observation{
		ID:          ObservationID("0125"),
		GroupNumber: "21",
		Name:        "穂別",
	},
	"0126_21": Observation{
		ID:          ObservationID("0126"),
		GroupNumber: "21",
		Name:        "森野",
	},
	"0128_21": Observation{
		ID:          ObservationID("0128"),
		GroupNumber: "21",
		Name:        "大岸",
	},
	"0129_21": Observation{
		ID:          ObservationID("0129"),
		GroupNumber: "21",
		Name:        "有珠山",
	},
	"0130_21": Observation{
		ID:          ObservationID("0130"),
		GroupNumber: "21",
		Name:        "白老",
	},
	"0131_21": Observation{
		ID:          ObservationID("0131"),
		GroupNumber: "21",
		Name:        "鵡川",
	},
	"0132_21": Observation{
		ID:          ObservationID("0132"),
		GroupNumber: "21",
		Name:        "伊達",
	},
	"0133_21": Observation{
		ID:          ObservationID("0133"),
		GroupNumber: "21",
		Name:        "登別山",
	},
	"0134_21": Observation{
		ID:          ObservationID("0134"),
		GroupNumber: "21",
		Name:        "登別",
	},
	"1067_21": Observation{
		ID:          ObservationID("1067"),
		GroupNumber: "21",
		Name:        "大滝",
	},
	"1329_21": Observation{
		ID:          ObservationID("1329"),
		GroupNumber: "21",
		Name:        "洞爺湖温泉",
	},
	"1429_21": Observation{
		ID:          ObservationID("1429"),
		GroupNumber: "21",
		Name:        "カルルス",
	},
	"1496_21": Observation{
		ID:          ObservationID("1496"),
		GroupNumber: "21",
		Name:        "豊浦",
	},
	"1498_21": Observation{
		ID:          ObservationID("1498"),
		GroupNumber: "21",
		Name:        "花和",
	},
	"1499_21": Observation{
		ID:          ObservationID("1499"),
		GroupNumber: "21",
		Name:        "虻田",
	},
	"1500_21": Observation{
		ID:          ObservationID("1500"),
		GroupNumber: "21",
		Name:        "壮瞥",
	},
	"1501_21": Observation{
		ID:          ObservationID("1501"),
		GroupNumber: "21",
		Name:        "入江",
	},
	"1502_21": Observation{
		ID:          ObservationID("1502"),
		GroupNumber: "21",
		Name:        "月浦",
	},
	"1503_21": Observation{
		ID:          ObservationID("1503"),
		GroupNumber: "21",
		Name:        "北有珠",
	},
	"1504_21": Observation{
		ID:          ObservationID("1504"),
		GroupNumber: "21",
		Name:        "洞爺湖畔",
	},
	"47423_21": Observation{
		ID:          ObservationID("47423"),
		GroupNumber: "21",
		Name:        "室蘭",
	},
	"47424_21": Observation{
		ID:          ObservationID("47424"),
		GroupNumber: "21",
		Name:        "苫小牧",
	},
	"0136_22": Observation{
		ID:          ObservationID("0136"),
		GroupNumber: "22",
		Name:        "日高門別",
	},
	"0137_22": Observation{
		ID:          ObservationID("0137"),
		GroupNumber: "22",
		Name:        "新和",
	},
	"0138_22": Observation{
		ID:          ObservationID("0138"),
		GroupNumber: "22",
		Name:        "ヌモトル山",
	},
	"0139_22": Observation{
		ID:          ObservationID("0139"),
		GroupNumber: "22",
		Name:        "笹山",
	},
	"0140_22": Observation{
		ID:          ObservationID("0140"),
		GroupNumber: "22",
		Name:        "静内",
	},
	"0141_22": Observation{
		ID:          ObservationID("0141"),
		GroupNumber: "22",
		Name:        "ピセナイ山",
	},
	"1097_22": Observation{
		ID:          ObservationID("1097"),
		GroupNumber: "22",
		Name:        "仁世宇",
	},
	"1198_22": Observation{
		ID:          ObservationID("1198"),
		GroupNumber: "22",
		Name:        "三石",
	},
	"1202_22": Observation{
		ID:          ObservationID("1202"),
		GroupNumber: "22",
		Name:        "日高",
	},
	"1264_22": Observation{
		ID:          ObservationID("1264"),
		GroupNumber: "22",
		Name:        "幌満",
	},
	"1276_22": Observation{
		ID:          ObservationID("1276"),
		GroupNumber: "22",
		Name:        "目黒",
	},
	"1288_22": Observation{
		ID:          ObservationID("1288"),
		GroupNumber: "22",
		Name:        "中杵臼",
	},
	"1340_22": Observation{
		ID:          ObservationID("1340"),
		GroupNumber: "22",
		Name:        "えりも岬",
	},
	"1430_22": Observation{
		ID:          ObservationID("1430"),
		GroupNumber: "22",
		Name:        "旭",
	},
	"1594_22": Observation{
		ID:          ObservationID("1594"),
		GroupNumber: "22",
		Name:        "平取",
	},
	"47426_22": Observation{
		ID:          ObservationID("47426"),
		GroupNumber: "22",
		Name:        "浦河",
	},
	"0143_23": Observation{
		ID:          ObservationID("0143"),
		GroupNumber: "23",
		Name:        "長万部",
	},
	"0144_23": Observation{
		ID:          ObservationID("0144"),
		GroupNumber: "23",
		Name:        "八雲",
	},
	"0145_23": Observation{
		ID:          ObservationID("0145"),
		GroupNumber: "23",
		Name:        "森",
	},
	"0146_23": Observation{
		ID:          ObservationID("0146"),
		GroupNumber: "23",
		Name:        "大沼",
	},
	"0147_23": Observation{
		ID:          ObservationID("0147"),
		GroupNumber: "23",
		Name:        "川汲",
	},
	"0148_23": Observation{
		ID:          ObservationID("0148"),
		GroupNumber: "23",
		Name:        "北斗",
	},
	"0150_23": Observation{
		ID:          ObservationID("0150"),
		GroupNumber: "23",
		Name:        "汐首",
	},
	"0151_23": Observation{
		ID:          ObservationID("0151"),
		GroupNumber: "23",
		Name:        "小谷石",
	},
	"0152_23": Observation{
		ID:          ObservationID("0152"),
		GroupNumber: "23",
		Name:        "松前",
	},
	"0155_23": Observation{
		ID:          ObservationID("0155"),
		GroupNumber: "23",
		Name:        "熊石",
	},
	"0917_23": Observation{
		ID:          ObservationID("0917"),
		GroupNumber: "23",
		Name:        "木古内",
	},
	"1265_23": Observation{
		ID:          ObservationID("1265"),
		GroupNumber: "23",
		Name:        "千軒",
	},
	"1416_23": Observation{
		ID:          ObservationID("1416"),
		GroupNumber: "23",
		Name:        "蛾眉野",
	},
	"1462_23": Observation{
		ID:          ObservationID("1462"),
		GroupNumber: "23",
		Name:        "高松",
	},
	"1543_23": Observation{
		ID:          ObservationID("1543"),
		GroupNumber: "23",
		Name:        "戸井泊",
	},
	"1552_23": Observation{
		ID:          ObservationID("1552"),
		GroupNumber: "23",
		Name:        "知内",
	},
	"47430_23": Observation{
		ID:          ObservationID("47430"),
		GroupNumber: "23",
		Name:        "函館",
	},
	"0153_24": Observation{
		ID:          ObservationID("0153"),
		GroupNumber: "24",
		Name:        "せたな",
	},
	"0154_24": Observation{
		ID:          ObservationID("0154"),
		GroupNumber: "24",
		Name:        "今金",
	},
	"0156_24": Observation{
		ID:          ObservationID("0156"),
		GroupNumber: "24",
		Name:        "潮見",
	},
	"1199_24": Observation{
		ID:          ObservationID("1199"),
		GroupNumber: "24",
		Name:        "奥尻",
	},
	"1200_24": Observation{
		ID:          ObservationID("1200"),
		GroupNumber: "24",
		Name:        "鶉",
	},
	"1205_24": Observation{
		ID:          ObservationID("1205"),
		GroupNumber: "24",
		Name:        "石崎",
	},
	"1513_24": Observation{
		ID:          ObservationID("1513"),
		GroupNumber: "24",
		Name:        "米岡",
	},
	"47428_24": Observation{
		ID:          ObservationID("47428"),
		GroupNumber: "24",
		Name:        "江差",
	},
	"0158_31": Observation{
		ID:          ObservationID("0158"),
		GroupNumber: "31",
		Name:        "障子山",
	},
	"0160_31": Observation{
		ID:          ObservationID("0160"),
		GroupNumber: "31",
		Name:        "蟹田",
	},
	"0162_31": Observation{
		ID:          ObservationID("0162"),
		GroupNumber: "31",
		Name:        "野辺地",
	},
	"0163_31": Observation{
		ID:          ObservationID("0163"),
		GroupNumber: "31",
		Name:        "鰺ケ沢",
	},
	"0164_31": Observation{
		ID:          ObservationID("0164"),
		GroupNumber: "31",
		Name:        "五所川原",
	},
	"0166_31": Observation{
		ID:          ObservationID("0166"),
		GroupNumber: "31",
		Name:        "弘前",
	},
	"0167_31": Observation{
		ID:          ObservationID("0167"),
		GroupNumber: "31",
		Name:        "黒石",
	},
	"0168_31": Observation{
		ID:          ObservationID("0168"),
		GroupNumber: "31",
		Name:        "八甲田山",
	},
	"0169_31": Observation{
		ID:          ObservationID("0169"),
		GroupNumber: "31",
		Name:        "三沢",
	},
	"0170_31": Observation{
		ID:          ObservationID("0170"),
		GroupNumber: "31",
		Name:        "四兵衛森",
	},
	"0171_31": Observation{
		ID:          ObservationID("0171"),
		GroupNumber: "31",
		Name:        "毛無山",
	},
	"0172_31": Observation{
		ID:          ObservationID("0172"),
		GroupNumber: "31",
		Name:        "十和田",
	},
	"0174_31": Observation{
		ID:          ObservationID("0174"),
		GroupNumber: "31",
		Name:        "大鰐",
	},
	"0175_31": Observation{
		ID:          ObservationID("0175"),
		GroupNumber: "31",
		Name:        "碇ケ関",
	},
	"0176_31": Observation{
		ID:          ObservationID("0176"),
		GroupNumber: "31",
		Name:        "空岱山",
	},
	"0177_31": Observation{
		ID:          ObservationID("0177"),
		GroupNumber: "31",
		Name:        "休屋",
	},
	"0178_31": Observation{
		ID:          ObservationID("0178"),
		GroupNumber: "31",
		Name:        "朝日奈岳",
	},
	"0179_31": Observation{
		ID:          ObservationID("0179"),
		GroupNumber: "31",
		Name:        "三戸",
	},
	"1026_31": Observation{
		ID:          ObservationID("1026"),
		GroupNumber: "31",
		Name:        "今別",
	},
	"1027_31": Observation{
		ID:          ObservationID("1027"),
		GroupNumber: "31",
		Name:        "六ケ所",
	},
	"1041_31": Observation{
		ID:          ObservationID("1041"),
		GroupNumber: "31",
		Name:        "大間",
	},
	"1042_31": Observation{
		ID:          ObservationID("1042"),
		GroupNumber: "31",
		Name:        "大和山",
	},
	"1079_31": Observation{
		ID:          ObservationID("1079"),
		GroupNumber: "31",
		Name:        "酸ケ湯",
	},
	"1108_31": Observation{
		ID:          ObservationID("1108"),
		GroupNumber: "31",
		Name:        "岳",
	},
	"1110_31": Observation{
		ID:          ObservationID("1110"),
		GroupNumber: "31",
		Name:        "長慶平",
	},
	"1119_31": Observation{
		ID:          ObservationID("1119"),
		GroupNumber: "31",
		Name:        "市浦",
	},
	"1122_31": Observation{
		ID:          ObservationID("1122"),
		GroupNumber: "31",
		Name:        "小田野沢",
	},
	"1124_31": Observation{
		ID:          ObservationID("1124"),
		GroupNumber: "31",
		Name:        "脇野沢",
	},
	"1183_31": Observation{
		ID:          ObservationID("1183"),
		GroupNumber: "31",
		Name:        "七戸",
	},
	"1272_31": Observation{
		ID:          ObservationID("1272"),
		GroupNumber: "31",
		Name:        "戸来",
	},
	"1406_31": Observation{
		ID:          ObservationID("1406"),
		GroupNumber: "31",
		Name:        "温川",
	},
	"1495_31": Observation{
		ID:          ObservationID("1495"),
		GroupNumber: "31",
		Name:        "青森大谷",
	},
	"1559_31": Observation{
		ID:          ObservationID("1559"),
		GroupNumber: "31",
		Name:        "湯野川",
	},
	"47574_31": Observation{
		ID:          ObservationID("47574"),
		GroupNumber: "31",
		Name:        "深浦",
	},
	"47575_31": Observation{
		ID:          ObservationID("47575"),
		GroupNumber: "31",
		Name:        "青森",
	},
	"47576_31": Observation{
		ID:          ObservationID("47576"),
		GroupNumber: "31",
		Name:        "むつ",
	},
	"47581_31": Observation{
		ID:          ObservationID("47581"),
		GroupNumber: "31",
		Name:        "八戸",
	},
	"0180_32": Observation{
		ID:          ObservationID("0180"),
		GroupNumber: "32",
		Name:        "田代岳",
	},
	"0181_32": Observation{
		ID:          ObservationID("0181"),
		GroupNumber: "32",
		Name:        "陣場",
	},
	"0182_32": Observation{
		ID:          ObservationID("0182"),
		GroupNumber: "32",
		Name:        "杉沢山",
	},
	"0183_32": Observation{
		ID:          ObservationID("0183"),
		GroupNumber: "32",
		Name:        "能代",
	},
	"0184_32": Observation{
		ID:          ObservationID("0184"),
		GroupNumber: "32",
		Name:        "鷹巣",
	},
	"0185_32": Observation{
		ID:          ObservationID("0185"),
		GroupNumber: "32",
		Name:        "鹿角",
	},
	"0186_32": Observation{
		ID:          ObservationID("0186"),
		GroupNumber: "32",
		Name:        "湯瀬",
	},
	"0187_32": Observation{
		ID:          ObservationID("0187"),
		GroupNumber: "32",
		Name:        "森吉山",
	},
	"0188_32": Observation{
		ID:          ObservationID("0188"),
		GroupNumber: "32",
		Name:        "五城目",
	},
	"0189_32": Observation{
		ID:          ObservationID("0189"),
		GroupNumber: "32",
		Name:        "男鹿本山",
	},
	"0191_32": Observation{
		ID:          ObservationID("0191"),
		GroupNumber: "32",
		Name:        "太平山",
	},
	"0192_32": Observation{
		ID:          ObservationID("0192"),
		GroupNumber: "32",
		Name:        "岩見三内",
	},
	"0193_32": Observation{
		ID:          ObservationID("0193"),
		GroupNumber: "32",
		Name:        "角館",
	},
	"0194_32": Observation{
		ID:          ObservationID("0194"),
		GroupNumber: "32",
		Name:        "田沢湖",
	},
	"0195_32": Observation{
		ID:          ObservationID("0195"),
		GroupNumber: "32",
		Name:        "大曲",
	},
	"0196_32": Observation{
		ID:          ObservationID("0196"),
		GroupNumber: "32",
		Name:        "本荘",
	},
	"0197_32": Observation{
		ID:          ObservationID("0197"),
		GroupNumber: "32",
		Name:        "保呂羽山",
	},
	"0198_32": Observation{
		ID:          ObservationID("0198"),
		GroupNumber: "32",
		Name:        "横手",
	},
	"0199_32": Observation{
		ID:          ObservationID("0199"),
		GroupNumber: "32",
		Name:        "象潟",
	},
	"0200_32": Observation{
		ID:          ObservationID("0200"),
		GroupNumber: "32",
		Name:        "矢島",
	},
	"0201_32": Observation{
		ID:          ObservationID("0201"),
		GroupNumber: "32",
		Name:        "三森山",
	},
	"0202_32": Observation{
		ID:          ObservationID("0202"),
		GroupNumber: "32",
		Name:        "湯沢",
	},
	"0203_32": Observation{
		ID:          ObservationID("0203"),
		GroupNumber: "32",
		Name:        "姥井戸山",
	},
	"0204_32": Observation{
		ID:          ObservationID("0204"),
		GroupNumber: "32",
		Name:        "湯の岱",
	},
	"0912_32": Observation{
		ID:          ObservationID("0912"),
		GroupNumber: "32",
		Name:        "大館",
	},
	"1035_32": Observation{
		ID:          ObservationID("1035"),
		GroupNumber: "32",
		Name:        "大潟",
	},
	"1036_32": Observation{
		ID:          ObservationID("1036"),
		GroupNumber: "32",
		Name:        "男鹿",
	},
	"1043_32": Observation{
		ID:          ObservationID("1043"),
		GroupNumber: "32",
		Name:        "八森",
	},
	"1050_32": Observation{
		ID:          ObservationID("1050"),
		GroupNumber: "32",
		Name:        "大正寺",
	},
	"1109_32": Observation{
		ID:          ObservationID("1109"),
		GroupNumber: "32",
		Name:        "東成瀬",
	},
	"1115_32": Observation{
		ID:          ObservationID("1115"),
		GroupNumber: "32",
		Name:        "比立内",
	},
	"1117_32": Observation{
		ID:          ObservationID("1117"),
		GroupNumber: "32",
		Name:        "東由利",
	},
	"1131_32": Observation{
		ID:          ObservationID("1131"),
		GroupNumber: "32",
		Name:        "阿仁合",
	},
	"1167_32": Observation{
		ID:          ObservationID("1167"),
		GroupNumber: "32",
		Name:        "鎧畑",
	},
	"1273_32": Observation{
		ID:          ObservationID("1273"),
		GroupNumber: "32",
		Name:        "藤里",
	},
	"1278_32": Observation{
		ID:          ObservationID("1278"),
		GroupNumber: "32",
		Name:        "桧木内",
	},
	"1289_32": Observation{
		ID:          ObservationID("1289"),
		GroupNumber: "32",
		Name:        "八幡平",
	},
	"1407_32": Observation{
		ID:          ObservationID("1407"),
		GroupNumber: "32",
		Name:        "笹子",
	},
	"1408_32": Observation{
		ID:          ObservationID("1408"),
		GroupNumber: "32",
		Name:        "藤原",
	},
	"1422_32": Observation{
		ID:          ObservationID("1422"),
		GroupNumber: "32",
		Name:        "男鹿真山",
	},
	"1425_32": Observation{
		ID:          ObservationID("1425"),
		GroupNumber: "32",
		Name:        "仁別",
	},
	"1463_32": Observation{
		ID:          ObservationID("1463"),
		GroupNumber: "32",
		Name:        "脇神",
	},
	"1511_32": Observation{
		ID:          ObservationID("1511"),
		GroupNumber: "32",
		Name:        "雄和",
	},
	"47582_32": Observation{
		ID:          ObservationID("47582"),
		GroupNumber: "32",
		Name:        "秋田",
	},
	"0205_33": Observation{
		ID:          ObservationID("0205"),
		GroupNumber: "33",
		Name:        "種市",
	},
	"0206_33": Observation{
		ID:          ObservationID("0206"),
		GroupNumber: "33",
		Name:        "軽米",
	},
	"0207_33": Observation{
		ID:          ObservationID("0207"),
		GroupNumber: "33",
		Name:        "二戸",
	},
	"0208_33": Observation{
		ID:          ObservationID("0208"),
		GroupNumber: "33",
		Name:        "折爪岳",
	},
	"0209_33": Observation{
		ID:          ObservationID("0209"),
		GroupNumber: "33",
		Name:        "久慈",
	},
	"0210_33": Observation{
		ID:          ObservationID("0210"),
		GroupNumber: "33",
		Name:        "西岳",
	},
	"0211_33": Observation{
		ID:          ObservationID("0211"),
		GroupNumber: "33",
		Name:        "葛巻",
	},
	"0212_33": Observation{
		ID:          ObservationID("0212"),
		GroupNumber: "33",
		Name:        "袖山",
	},
	"0213_33": Observation{
		ID:          ObservationID("0213"),
		GroupNumber: "33",
		Name:        "茶臼岳",
	},
	"0214_33": Observation{
		ID:          ObservationID("0214"),
		GroupNumber: "33",
		Name:        "岩手松尾",
	},
	"0215_33": Observation{
		ID:          ObservationID("0215"),
		GroupNumber: "33",
		Name:        "岩手山",
	},
	"0216_33": Observation{
		ID:          ObservationID("0216"),
		GroupNumber: "33",
		Name:        "岩洞",
	},
	"0217_33": Observation{
		ID:          ObservationID("0217"),
		GroupNumber: "33",
		Name:        "大森山",
	},
	"0218_33": Observation{
		ID:          ObservationID("0218"),
		GroupNumber: "33",
		Name:        "岩泉",
	},
	"0219_33": Observation{
		ID:          ObservationID("0219"),
		GroupNumber: "33",
		Name:        "駒ケ岳",
	},
	"0220_33": Observation{
		ID:          ObservationID("0220"),
		GroupNumber: "33",
		Name:        "害鷹森",
	},
	"0221_33": Observation{
		ID:          ObservationID("0221"),
		GroupNumber: "33",
		Name:        "雫石",
	},
	"0224_33": Observation{
		ID:          ObservationID("0224"),
		GroupNumber: "33",
		Name:        "薬師岳",
	},
	"0225_33": Observation{
		ID:          ObservationID("0225"),
		GroupNumber: "33",
		Name:        "黒森山",
	},
	"0226_33": Observation{
		ID:          ObservationID("0226"),
		GroupNumber: "33",
		Name:        "豊沢",
	},
	"0227_33": Observation{
		ID:          ObservationID("0227"),
		GroupNumber: "33",
		Name:        "花巻",
	},
	"0228_33": Observation{
		ID:          ObservationID("0228"),
		GroupNumber: "33",
		Name:        "大迫",
	},
	"0229_33": Observation{
		ID:          ObservationID("0229"),
		GroupNumber: "33",
		Name:        "湯田",
	},
	"0230_33": Observation{
		ID:          ObservationID("0230"),
		GroupNumber: "33",
		Name:        "北上",
	},
	"0231_33": Observation{
		ID:          ObservationID("0231"),
		GroupNumber: "33",
		Name:        "遠野",
	},
	"0232_33": Observation{
		ID:          ObservationID("0232"),
		GroupNumber: "33",
		Name:        "五葉山",
	},
	"0233_33": Observation{
		ID:          ObservationID("0233"),
		GroupNumber: "33",
		Name:        "釜石",
	},
	"0234_33": Observation{
		ID:          ObservationID("0234"),
		GroupNumber: "33",
		Name:        "岳山",
	},
	"0235_33": Observation{
		ID:          ObservationID("0235"),
		GroupNumber: "33",
		Name:        "若柳",
	},
	"0236_33": Observation{
		ID:          ObservationID("0236"),
		GroupNumber: "33",
		Name:        "江刺",
	},
	"0238_33": Observation{
		ID:          ObservationID("0238"),
		GroupNumber: "33",
		Name:        "一関",
	},
	"0239_33": Observation{
		ID:          ObservationID("0239"),
		GroupNumber: "33",
		Name:        "千厩",
	},
	"0240_33": Observation{
		ID:          ObservationID("0240"),
		GroupNumber: "33",
		Name:        "室根山",
	},
	"1028_33": Observation{
		ID:          ObservationID("1028"),
		GroupNumber: "33",
		Name:        "祭畤",
	},
	"1032_33": Observation{
		ID:          ObservationID("1032"),
		GroupNumber: "33",
		Name:        "好摩",
	},
	"1033_33": Observation{
		ID:          ObservationID("1033"),
		GroupNumber: "33",
		Name:        "山田",
	},
	"1037_33": Observation{
		ID:          ObservationID("1037"),
		GroupNumber: "33",
		Name:        "大野",
	},
	"1045_33": Observation{
		ID:          ObservationID("1045"),
		GroupNumber: "33",
		Name:        "衣川",
	},
	"1078_33": Observation{
		ID:          ObservationID("1078"),
		GroupNumber: "33",
		Name:        "藪川",
	},
	"1113_33": Observation{
		ID:          ObservationID("1113"),
		GroupNumber: "33",
		Name:        "米里",
	},
	"1118_33": Observation{
		ID:          ObservationID("1118"),
		GroupNumber: "33",
		Name:        "沢内",
	},
	"1120_33": Observation{
		ID:          ObservationID("1120"),
		GroupNumber: "33",
		Name:        "荒屋",
	},
	"1123_33": Observation{
		ID:          ObservationID("1123"),
		GroupNumber: "33",
		Name:        "普代",
	},
	"1128_33": Observation{
		ID:          ObservationID("1128"),
		GroupNumber: "33",
		Name:        "紫波",
	},
	"1168_33": Observation{
		ID:          ObservationID("1168"),
		GroupNumber: "33",
		Name:        "葛根田",
	},
	"1206_33": Observation{
		ID:          ObservationID("1206"),
		GroupNumber: "33",
		Name:        "住田",
	},
	"1212_33": Observation{
		ID:          ObservationID("1212"),
		GroupNumber: "33",
		Name:        "小本",
	},
	"1215_33": Observation{
		ID:          ObservationID("1215"),
		GroupNumber: "33",
		Name:        "奥中山",
	},
	"1219_33": Observation{
		ID:          ObservationID("1219"),
		GroupNumber: "33",
		Name:        "山形",
	},
	"1224_33": Observation{
		ID:          ObservationID("1224"),
		GroupNumber: "33",
		Name:        "川井",
	},
	"1225_33": Observation{
		ID:          ObservationID("1225"),
		GroupNumber: "33",
		Name:        "門馬",
	},
	"1332_33": Observation{
		ID:          ObservationID("1332"),
		GroupNumber: "33",
		Name:        "下戸鎖",
	},
	"1447_33": Observation{
		ID:          ObservationID("1447"),
		GroupNumber: "33",
		Name:        "区界",
	},
	"1508_33": Observation{
		ID:          ObservationID("1508"),
		GroupNumber: "33",
		Name:        "大槌",
	},
	"1557_33": Observation{
		ID:          ObservationID("1557"),
		GroupNumber: "33",
		Name:        "大東",
	},
	"1558_33": Observation{
		ID:          ObservationID("1558"),
		GroupNumber: "33",
		Name:        "金ヶ崎",
	},
	"1560_33": Observation{
		ID:          ObservationID("1560"),
		GroupNumber: "33",
		Name:        "滝沢",
	},
	"1563_33": Observation{
		ID:          ObservationID("1563"),
		GroupNumber: "33",
		Name:        "附馬牛",
	},
	"1598_33": Observation{
		ID:          ObservationID("1598"),
		GroupNumber: "33",
		Name:        "厳美",
	},
	"47512_33": Observation{
		ID:          ObservationID("47512"),
		GroupNumber: "33",
		Name:        "大船渡",
	},
	"47584_33": Observation{
		ID:          ObservationID("47584"),
		GroupNumber: "33",
		Name:        "盛岡",
	},
	"47585_33": Observation{
		ID:          ObservationID("47585"),
		GroupNumber: "33",
		Name:        "宮古",
	},
	"0241_34": Observation{
		ID:          ObservationID("0241"),
		GroupNumber: "34",
		Name:        "栗駒山",
	},
	"0242_34": Observation{
		ID:          ObservationID("0242"),
		GroupNumber: "34",
		Name:        "気仙沼",
	},
	"0243_34": Observation{
		ID:          ObservationID("0243"),
		GroupNumber: "34",
		Name:        "川渡",
	},
	"0244_34": Observation{
		ID:          ObservationID("0244"),
		GroupNumber: "34",
		Name:        "築館",
	},
	"0245_34": Observation{
		ID:          ObservationID("0245"),
		GroupNumber: "34",
		Name:        "箕輪山",
	},
	"0246_34": Observation{
		ID:          ObservationID("0246"),
		GroupNumber: "34",
		Name:        "志津川",
	},
	"0247_34": Observation{
		ID:          ObservationID("0247"),
		GroupNumber: "34",
		Name:        "古川",
	},
	"0248_34": Observation{
		ID:          ObservationID("0248"),
		GroupNumber: "34",
		Name:        "大衡",
	},
	"0249_34": Observation{
		ID:          ObservationID("0249"),
		GroupNumber: "34",
		Name:        "鹿島台",
	},
	"0251_34": Observation{
		ID:          ObservationID("0251"),
		GroupNumber: "34",
		Name:        "新川",
	},
	"0252_34": Observation{
		ID:          ObservationID("0252"),
		GroupNumber: "34",
		Name:        "鷹巣山",
	},
	"0254_34": Observation{
		ID:          ObservationID("0254"),
		GroupNumber: "34",
		Name:        "川崎",
	},
	"0255_34": Observation{
		ID:          ObservationID("0255"),
		GroupNumber: "34",
		Name:        "不忘山",
	},
	"0256_34": Observation{
		ID:          ObservationID("0256"),
		GroupNumber: "34",
		Name:        "白石",
	},
	"0257_34": Observation{
		ID:          ObservationID("0257"),
		GroupNumber: "34",
		Name:        "亘理",
	},
	"0985_34": Observation{
		ID:          ObservationID("0985"),
		GroupNumber: "34",
		Name:        "雄勝",
	},
	"1029_34": Observation{
		ID:          ObservationID("1029"),
		GroupNumber: "34",
		Name:        "米山",
	},
	"1030_34": Observation{
		ID:          ObservationID("1030"),
		GroupNumber: "34",
		Name:        "塩釜",
	},
	"1126_34": Observation{
		ID:          ObservationID("1126"),
		GroupNumber: "34",
		Name:        "駒ノ湯",
	},
	"1220_34": Observation{
		ID:          ObservationID("1220"),
		GroupNumber: "34",
		Name:        "丸森",
	},
	"1267_34": Observation{
		ID:          ObservationID("1267"),
		GroupNumber: "34",
		Name:        "花山",
	},
	"1290_34": Observation{
		ID:          ObservationID("1290"),
		GroupNumber: "34",
		Name:        "江ノ島",
	},
	"1298_34": Observation{
		ID:          ObservationID("1298"),
		GroupNumber: "34",
		Name:        "筆甫",
	},
	"1365_34": Observation{
		ID:          ObservationID("1365"),
		GroupNumber: "34",
		Name:        "泉ケ岳",
	},
	"1464_34": Observation{
		ID:          ObservationID("1464"),
		GroupNumber: "34",
		Name:        "名取",
	},
	"1526_34": Observation{
		ID:          ObservationID("1526"),
		GroupNumber: "34",
		Name:        "鴬沢",
	},
	"1561_34": Observation{
		ID:          ObservationID("1561"),
		GroupNumber: "34",
		Name:        "加美",
	},
	"1564_34": Observation{
		ID:          ObservationID("1564"),
		GroupNumber: "34",
		Name:        "蔵王",
	},
	"1597_34": Observation{
		ID:          ObservationID("1597"),
		GroupNumber: "34",
		Name:        "栗駒深山",
	},
	"1600_34": Observation{
		ID:          ObservationID("1600"),
		GroupNumber: "34",
		Name:        "栗駒",
	},
	"47590_34": Observation{
		ID:          ObservationID("47590"),
		GroupNumber: "34",
		Name:        "仙台",
	},
	"47592_34": Observation{
		ID:          ObservationID("47592"),
		GroupNumber: "34",
		Name:        "石巻",
	},
	"0258_35": Observation{
		ID:          ObservationID("0258"),
		GroupNumber: "35",
		Name:        "鳥海山",
	},
	"0260_35": Observation{
		ID:          ObservationID("0260"),
		GroupNumber: "35",
		Name:        "差首鍋",
	},
	"0261_35": Observation{
		ID:          ObservationID("0261"),
		GroupNumber: "35",
		Name:        "太平山",
	},
	"0262_35": Observation{
		ID:          ObservationID("0262"),
		GroupNumber: "35",
		Name:        "金山",
	},
	"0263_35": Observation{
		ID:          ObservationID("0263"),
		GroupNumber: "35",
		Name:        "鶴岡",
	},
	"0264_35": Observation{
		ID:          ObservationID("0264"),
		GroupNumber: "35",
		Name:        "狩川",
	},
	"0266_35": Observation{
		ID:          ObservationID("0266"),
		GroupNumber: "35",
		Name:        "瀬見",
	},
	"0267_35": Observation{
		ID:          ObservationID("0267"),
		GroupNumber: "35",
		Name:        "温海岳",
	},
	"0268_35": Observation{
		ID:          ObservationID("0268"),
		GroupNumber: "35",
		Name:        "櫛引",
	},
	"0269_35": Observation{
		ID:          ObservationID("0269"),
		GroupNumber: "35",
		Name:        "月山",
	},
	"0270_35": Observation{
		ID:          ObservationID("0270"),
		GroupNumber: "35",
		Name:        "銀山",
	},
	"0271_35": Observation{
		ID:          ObservationID("0271"),
		GroupNumber: "35",
		Name:        "楯岡",
	},
	"0272_35": Observation{
		ID:          ObservationID("0272"),
		GroupNumber: "35",
		Name:        "左沢",
	},
	"0274_35": Observation{
		ID:          ObservationID("0274"),
		GroupNumber: "35",
		Name:        "葉山",
	},
	"0275_35": Observation{
		ID:          ObservationID("0275"),
		GroupNumber: "35",
		Name:        "長井",
	},
	"0276_35": Observation{
		ID:          ObservationID("0276"),
		GroupNumber: "35",
		Name:        "蔵王山",
	},
	"0277_35": Observation{
		ID:          ObservationID("0277"),
		GroupNumber: "35",
		Name:        "小国",
	},
	"0278_35": Observation{
		ID:          ObservationID("0278"),
		GroupNumber: "35",
		Name:        "高峰",
	},
	"0279_35": Observation{
		ID:          ObservationID("0279"),
		GroupNumber: "35",
		Name:        "米沢",
	},
	"0280_35": Observation{
		ID:          ObservationID("0280"),
		GroupNumber: "35",
		Name:        "西吾妻山",
	},
	"0910_35": Observation{
		ID:          ObservationID("0910"),
		GroupNumber: "35",
		Name:        "向町",
	},
	"1038_35": Observation{
		ID:          ObservationID("1038"),
		GroupNumber: "35",
		Name:        "上草津",
	},
	"1039_35": Observation{
		ID:          ObservationID("1039"),
		GroupNumber: "35",
		Name:        "尾花沢",
	},
	"1040_35": Observation{
		ID:          ObservationID("1040"),
		GroupNumber: "35",
		Name:        "鼠ケ関",
	},
	"1103_35": Observation{
		ID:          ObservationID("1103"),
		GroupNumber: "35",
		Name:        "山毛欅潰山",
	},
	"1125_35": Observation{
		ID:          ObservationID("1125"),
		GroupNumber: "35",
		Name:        "肘折",
	},
	"1132_35": Observation{
		ID:          ObservationID("1132"),
		GroupNumber: "35",
		Name:        "高畠",
	},
	"1169_35": Observation{
		ID:          ObservationID("1169"),
		GroupNumber: "35",
		Name:        "落合",
	},
	"1291_35": Observation{
		ID:          ObservationID("1291"),
		GroupNumber: "35",
		Name:        "飛島",
	},
	"1292_35": Observation{
		ID:          ObservationID("1292"),
		GroupNumber: "35",
		Name:        "大井沢",
	},
	"1357_35": Observation{
		ID:          ObservationID("1357"),
		GroupNumber: "35",
		Name:        "荒沢",
	},
	"1372_35": Observation{
		ID:          ObservationID("1372"),
		GroupNumber: "35",
		Name:        "中津川",
	},
	"1428_35": Observation{
		ID:          ObservationID("1428"),
		GroupNumber: "35",
		Name:        "白鷹山",
	},
	"1465_35": Observation{
		ID:          ObservationID("1465"),
		GroupNumber: "35",
		Name:        "浜中",
	},
	"1488_35": Observation{
		ID:          ObservationID("1488"),
		GroupNumber: "35",
		Name:        "東根",
	},
	"1539_35": Observation{
		ID:          ObservationID("1539"),
		GroupNumber: "35",
		Name:        "村山",
	},
	"1562_35": Observation{
		ID:          ObservationID("1562"),
		GroupNumber: "35",
		Name:        "上山中山",
	},
	"47520_35": Observation{
		ID:          ObservationID("47520"),
		GroupNumber: "35",
		Name:        "新庄",
	},
	"47587_35": Observation{
		ID:          ObservationID("47587"),
		GroupNumber: "35",
		Name:        "酒田",
	},
	"47588_35": Observation{
		ID:          ObservationID("47588"),
		GroupNumber: "35",
		Name:        "山形",
	},
	"0281_36": Observation{
		ID:          ObservationID("0281"),
		GroupNumber: "36",
		Name:        "梁川",
	},
	"0282_36": Observation{
		ID:          ObservationID("0282"),
		GroupNumber: "36",
		Name:        "稲荷峠",
	},
	"0283_36": Observation{
		ID:          ObservationID("0283"),
		GroupNumber: "36",
		Name:        "吾妻山",
	},
	"0285_36": Observation{
		ID:          ObservationID("0285"),
		GroupNumber: "36",
		Name:        "相馬",
	},
	"0286_36": Observation{
		ID:          ObservationID("0286"),
		GroupNumber: "36",
		Name:        "喜多方",
	},
	"0287_36": Observation{
		ID:          ObservationID("0287"),
		GroupNumber: "36",
		Name:        "八木沢",
	},
	"0288_36": Observation{
		ID:          ObservationID("0288"),
		GroupNumber: "36",
		Name:        "原町",
	},
	"0289_36": Observation{
		ID:          ObservationID("0289"),
		GroupNumber: "36",
		Name:        "飯谷山",
	},
	"0290_36": Observation{
		ID:          ObservationID("0290"),
		GroupNumber: "36",
		Name:        "猪苗代",
	},
	"0291_36": Observation{
		ID:          ObservationID("0291"),
		GroupNumber: "36",
		Name:        "二本松",
	},
	"0292_36": Observation{
		ID:          ObservationID("0292"),
		GroupNumber: "36",
		Name:        "白馬石山",
	},
	"0294_36": Observation{
		ID:          ObservationID("0294"),
		GroupNumber: "36",
		Name:        "船引",
	},
	"0295_36": Observation{
		ID:          ObservationID("0295"),
		GroupNumber: "36",
		Name:        "浪江",
	},
	"0296_36": Observation{
		ID:          ObservationID("0296"),
		GroupNumber: "36",
		Name:        "浅草岳",
	},
	"0297_36": Observation{
		ID:          ObservationID("0297"),
		GroupNumber: "36",
		Name:        "只見",
	},
	"0298_36": Observation{
		ID:          ObservationID("0298"),
		GroupNumber: "36",
		Name:        "博士峠",
	},
	"0299_36": Observation{
		ID:          ObservationID("0299"),
		GroupNumber: "36",
		Name:        "郡山",
	},
	"0300_36": Observation{
		ID:          ObservationID("0300"),
		GroupNumber: "36",
		Name:        "富岡",
	},
	"0301_36": Observation{
		ID:          ObservationID("0301"),
		GroupNumber: "36",
		Name:        "南郷",
	},
	"0302_36": Observation{
		ID:          ObservationID("0302"),
		GroupNumber: "36",
		Name:        "駒止峠",
	},
	"0303_36": Observation{
		ID:          ObservationID("0303"),
		GroupNumber: "36",
		Name:        "田島",
	},
	"0304_36": Observation{
		ID:          ObservationID("0304"),
		GroupNumber: "36",
		Name:        "小野新町",
	},
	"0305_36": Observation{
		ID:          ObservationID("0305"),
		GroupNumber: "36",
		Name:        "観音山",
	},
	"0307_36": Observation{
		ID:          ObservationID("0307"),
		GroupNumber: "36",
		Name:        "石川",
	},
	"0308_36": Observation{
		ID:          ObservationID("0308"),
		GroupNumber: "36",
		Name:        "山王峠",
	},
	"0309_36": Observation{
		ID:          ObservationID("0309"),
		GroupNumber: "36",
		Name:        "大辷山",
	},
	"0310_36": Observation{
		ID:          ObservationID("0310"),
		GroupNumber: "36",
		Name:        "平",
	},
	"0311_36": Observation{
		ID:          ObservationID("0311"),
		GroupNumber: "36",
		Name:        "田代山",
	},
	"0312_36": Observation{
		ID:          ObservationID("0312"),
		GroupNumber: "36",
		Name:        "東白川",
	},
	"1031_36": Observation{
		ID:          ObservationID("1031"),
		GroupNumber: "36",
		Name:        "西会津",
	},
	"1034_36": Observation{
		ID:          ObservationID("1034"),
		GroupNumber: "36",
		Name:        "広野",
	},
	"1044_36": Observation{
		ID:          ObservationID("1044"),
		GroupNumber: "36",
		Name:        "金山",
	},
	"1116_36": Observation{
		ID:          ObservationID("1116"),
		GroupNumber: "36",
		Name:        "鷲倉",
	},
	"1127_36": Observation{
		ID:          ObservationID("1127"),
		GroupNumber: "36",
		Name:        "上遠野",
	},
	"1129_36": Observation{
		ID:          ObservationID("1129"),
		GroupNumber: "36",
		Name:        "川内",
	},
	"1130_36": Observation{
		ID:          ObservationID("1130"),
		GroupNumber: "36",
		Name:        "飯舘",
	},
	"1150_36": Observation{
		ID:          ObservationID("1150"),
		GroupNumber: "36",
		Name:        "津島",
	},
	"1268_36": Observation{
		ID:          ObservationID("1268"),
		GroupNumber: "36",
		Name:        "長沼",
	},
	"1269_36": Observation{
		ID:          ObservationID("1269"),
		GroupNumber: "36",
		Name:        "舘岩",
	},
	"1274_36": Observation{
		ID:          ObservationID("1274"),
		GroupNumber: "36",
		Name:        "川前",
	},
	"1282_36": Observation{
		ID:          ObservationID("1282"),
		GroupNumber: "36",
		Name:        "桧原",
	},
	"1293_36": Observation{
		ID:          ObservationID("1293"),
		GroupNumber: "36",
		Name:        "茂庭",
	},
	"1294_36": Observation{
		ID:          ObservationID("1294"),
		GroupNumber: "36",
		Name:        "湯本",
	},
	"1295_36": Observation{
		ID:          ObservationID("1295"),
		GroupNumber: "36",
		Name:        "桧枝岐",
	},
	"1359_36": Observation{
		ID:          ObservationID("1359"),
		GroupNumber: "36",
		Name:        "吉尾峠",
	},
	"1364_36": Observation{
		ID:          ObservationID("1364"),
		GroupNumber: "36",
		Name:        "六十里越",
	},
	"1386_36": Observation{
		ID:          ObservationID("1386"),
		GroupNumber: "36",
		Name:        "湖南",
	},
	"1466_36": Observation{
		ID:          ObservationID("1466"),
		GroupNumber: "36",
		Name:        "玉川",
	},
	"1607_36": Observation{
		ID:          ObservationID("1607"),
		GroupNumber: "36",
		Name:        "山田",
	},
	"47570_36": Observation{
		ID:          ObservationID("47570"),
		GroupNumber: "36",
		Name:        "若松",
	},
	"47595_36": Observation{
		ID:          ObservationID("47595"),
		GroupNumber: "36",
		Name:        "福島",
	},
	"47597_36": Observation{
		ID:          ObservationID("47597"),
		GroupNumber: "36",
		Name:        "白河",
	},
	"47598_36": Observation{
		ID:          ObservationID("47598"),
		GroupNumber: "36",
		Name:        "小名浜",
	},
	"0314_40": Observation{
		ID:          ObservationID("0314"),
		GroupNumber: "40",
		Name:        "花園",
	},
	"0315_40": Observation{
		ID:          ObservationID("0315"),
		GroupNumber: "40",
		Name:        "北茨城",
	},
	"0316_40": Observation{
		ID:          ObservationID("0316"),
		GroupNumber: "40",
		Name:        "大子",
	},
	"0317_40": Observation{
		ID:          ObservationID("0317"),
		GroupNumber: "40",
		Name:        "神峰山",
	},
	"0318_40": Observation{
		ID:          ObservationID("0318"),
		GroupNumber: "40",
		Name:        "笠間",
	},
	"0320_40": Observation{
		ID:          ObservationID("0320"),
		GroupNumber: "40",
		Name:        "古河",
	},
	"0321_40": Observation{
		ID:          ObservationID("0321"),
		GroupNumber: "40",
		Name:        "愛宕山",
	},
	"0322_40": Observation{
		ID:          ObservationID("0322"),
		GroupNumber: "40",
		Name:        "下妻",
	},
	"0323_40": Observation{
		ID:          ObservationID("0323"),
		GroupNumber: "40",
		Name:        "坂東",
	},
	"0324_40": Observation{
		ID:          ObservationID("0324"),
		GroupNumber: "40",
		Name:        "土浦",
	},
	"0325_40": Observation{
		ID:          ObservationID("0325"),
		GroupNumber: "40",
		Name:        "鹿嶋",
	},
	"0911_40": Observation{
		ID:          ObservationID("0911"),
		GroupNumber: "40",
		Name:        "江戸崎",
	},
	"1011_40": Observation{
		ID:          ObservationID("1011"),
		GroupNumber: "40",
		Name:        "日立",
	},
	"1012_40": Observation{
		ID:          ObservationID("1012"),
		GroupNumber: "40",
		Name:        "柿岡",
	},
	"1013_40": Observation{
		ID:          ObservationID("1013"),
		GroupNumber: "40",
		Name:        "美野里",
	},
	"1014_40": Observation{
		ID:          ObservationID("1014"),
		GroupNumber: "40",
		Name:        "龍ケ崎",
	},
	"1082_40": Observation{
		ID:          ObservationID("1082"),
		GroupNumber: "40",
		Name:        "筑波山",
	},
	"1170_40": Observation{
		ID:          ObservationID("1170"),
		GroupNumber: "40",
		Name:        "中野",
	},
	"1173_40": Observation{
		ID:          ObservationID("1173"),
		GroupNumber: "40",
		Name:        "徳田",
	},
	"1245_40": Observation{
		ID:          ObservationID("1245"),
		GroupNumber: "40",
		Name:        "鉾田",
	},
	"1331_40": Observation{
		ID:          ObservationID("1331"),
		GroupNumber: "40",
		Name:        "小瀬",
	},
	"1388_40": Observation{
		ID:          ObservationID("1388"),
		GroupNumber: "40",
		Name:        "門井",
	},
	"1421_40": Observation{
		ID:          ObservationID("1421"),
		GroupNumber: "40",
		Name:        "大能",
	},
	"1530_40": Observation{
		ID:          ObservationID("1530"),
		GroupNumber: "40",
		Name:        "下館",
	},
	"47629_40": Observation{
		ID:          ObservationID("47629"),
		GroupNumber: "40",
		Name:        "水戸",
	},
	"47646_40": Observation{
		ID:          ObservationID("47646"),
		GroupNumber: "40",
		Name:        "つくば(館野)",
	},
	"0326_41": Observation{
		ID:          ObservationID("0326"),
		GroupNumber: "41",
		Name:        "那須",
	},
	"0327_41": Observation{
		ID:          ObservationID("0327"),
		GroupNumber: "41",
		Name:        "黒田原",
	},
	"0328_41": Observation{
		ID:          ObservationID("0328"),
		GroupNumber: "41",
		Name:        "八方が原",
	},
	"0329_41": Observation{
		ID:          ObservationID("0329"),
		GroupNumber: "41",
		Name:        "黒磯",
	},
	"0330_41": Observation{
		ID:          ObservationID("0330"),
		GroupNumber: "41",
		Name:        "八溝山",
	},
	"0331_41": Observation{
		ID:          ObservationID("0331"),
		GroupNumber: "41",
		Name:        "大田原",
	},
	"0333_41": Observation{
		ID:          ObservationID("0333"),
		GroupNumber: "41",
		Name:        "方塞山",
	},
	"0334_41": Observation{
		ID:          ObservationID("0334"),
		GroupNumber: "41",
		Name:        "高根沢",
	},
	"0335_41": Observation{
		ID:          ObservationID("0335"),
		GroupNumber: "41",
		Name:        "鹿沼",
	},
	"0337_41": Observation{
		ID:          ObservationID("0337"),
		GroupNumber: "41",
		Name:        "葛生",
	},
	"0338_41": Observation{
		ID:          ObservationID("0338"),
		GroupNumber: "41",
		Name:        "真岡",
	},
	"0339_41": Observation{
		ID:          ObservationID("0339"),
		GroupNumber: "41",
		Name:        "足利",
	},
	"0340_41": Observation{
		ID:          ObservationID("0340"),
		GroupNumber: "41",
		Name:        "栃木",
	},
	"0341_41": Observation{
		ID:          ObservationID("0341"),
		GroupNumber: "41",
		Name:        "小山",
	},
	"1015_41": Observation{
		ID:          ObservationID("1015"),
		GroupNumber: "41",
		Name:        "五十里",
	},
	"1016_41": Observation{
		ID:          ObservationID("1016"),
		GroupNumber: "41",
		Name:        "足尾",
	},
	"1017_41": Observation{
		ID:          ObservationID("1017"),
		GroupNumber: "41",
		Name:        "烏山",
	},
	"1018_41": Observation{
		ID:          ObservationID("1018"),
		GroupNumber: "41",
		Name:        "佐野",
	},
	"1221_41": Observation{
		ID:          ObservationID("1221"),
		GroupNumber: "41",
		Name:        "土呂部",
	},
	"1333_41": Observation{
		ID:          ObservationID("1333"),
		GroupNumber: "41",
		Name:        "今市",
	},
	"1334_41": Observation{
		ID:          ObservationID("1334"),
		GroupNumber: "41",
		Name:        "塩谷",
	},
	"1605_41": Observation{
		ID:          ObservationID("1605"),
		GroupNumber: "41",
		Name:        "那須烏山",
	},
	"47615_41": Observation{
		ID:          ObservationID("47615"),
		GroupNumber: "41",
		Name:        "宇都宮",
	},
	"47690_41": Observation{
		ID:          ObservationID("47690"),
		GroupNumber: "41",
		Name:        "奥日光(日光)",
	},
	"0342_42": Observation{
		ID:          ObservationID("0342"),
		GroupNumber: "42",
		Name:        "前武尊",
	},
	"0343_42": Observation{
		ID:          ObservationID("0343"),
		GroupNumber: "42",
		Name:        "草津",
	},
	"0344_42": Observation{
		ID:          ObservationID("0344"),
		GroupNumber: "42",
		Name:        "野反湖",
	},
	"0345_42": Observation{
		ID:          ObservationID("0345"),
		GroupNumber: "42",
		Name:        "雨見山",
	},
	"0346_42": Observation{
		ID:          ObservationID("0346"),
		GroupNumber: "42",
		Name:        "沼田",
	},
	"0347_42": Observation{
		ID:          ObservationID("0347"),
		GroupNumber: "42",
		Name:        "赤城山",
	},
	"0348_42": Observation{
		ID:          ObservationID("0348"),
		GroupNumber: "42",
		Name:        "田代",
	},
	"0349_42": Observation{
		ID:          ObservationID("0349"),
		GroupNumber: "42",
		Name:        "榛名山",
	},
	"0351_42": Observation{
		ID:          ObservationID("0351"),
		GroupNumber: "42",
		Name:        "桐生",
	},
	"0352_42": Observation{
		ID:          ObservationID("0352"),
		GroupNumber: "42",
		Name:        "一の字山",
	},
	"0353_42": Observation{
		ID:          ObservationID("0353"),
		GroupNumber: "42",
		Name:        "上里見",
	},
	"0354_42": Observation{
		ID:          ObservationID("0354"),
		GroupNumber: "42",
		Name:        "西野牧",
	},
	"0355_42": Observation{
		ID:          ObservationID("0355"),
		GroupNumber: "42",
		Name:        "藤岡",
	},
	"0356_42": Observation{
		ID:          ObservationID("0356"),
		GroupNumber: "42",
		Name:        "館林",
	},
	"0357_42": Observation{
		ID:          ObservationID("0357"),
		GroupNumber: "42",
		Name:        "稲含山",
	},
	"1019_42": Observation{
		ID:          ObservationID("1019"),
		GroupNumber: "42",
		Name:        "みなかみ",
	},
	"1020_42": Observation{
		ID:          ObservationID("1020"),
		GroupNumber: "42",
		Name:        "中之条",
	},
	"1021_42": Observation{
		ID:          ObservationID("1021"),
		GroupNumber: "42",
		Name:        "伊勢崎",
	},
	"1174_42": Observation{
		ID:          ObservationID("1174"),
		GroupNumber: "42",
		Name:        "片品",
	},
	"1222_42": Observation{
		ID:          ObservationID("1222"),
		GroupNumber: "42",
		Name:        "藤原",
	},
	"1223_42": Observation{
		ID:          ObservationID("1223"),
		GroupNumber: "42",
		Name:        "神流",
	},
	"1259_42": Observation{
		ID:          ObservationID("1259"),
		GroupNumber: "42",
		Name:        "黒保根",
	},
	"1375_42": Observation{
		ID:          ObservationID("1375"),
		GroupNumber: "42",
		Name:        "高手山",
	},
	"47624_42": Observation{
		ID:          ObservationID("47624"),
		GroupNumber: "42",
		Name:        "前橋",
	},
	"0359_43": Observation{
		ID:          ObservationID("0359"),
		GroupNumber: "43",
		Name:        "久喜",
	},
	"0361_43": Observation{
		ID:          ObservationID("0361"),
		GroupNumber: "43",
		Name:        "堂平",
	},
	"0362_43": Observation{
		ID:          ObservationID("0362"),
		GroupNumber: "43",
		Name:        "飯能",
	},
	"0363_43": Observation{
		ID:          ObservationID("0363"),
		GroupNumber: "43",
		Name:        "さいたま",
	},
	"0364_43": Observation{
		ID:          ObservationID("0364"),
		GroupNumber: "43",
		Name:        "越谷",
	},
	"1000_43": Observation{
		ID:          ObservationID("1000"),
		GroupNumber: "43",
		Name:        "三峰",
	},
	"1009_43": Observation{
		ID:          ObservationID("1009"),
		GroupNumber: "43",
		Name:        "寄居",
	},
	"1069_43": Observation{
		ID:          ObservationID("1069"),
		GroupNumber: "43",
		Name:        "鴻巣",
	},
	"1070_43": Observation{
		ID:          ObservationID("1070"),
		GroupNumber: "43",
		Name:        "所沢",
	},
	"1159_43": Observation{
		ID:          ObservationID("1159"),
		GroupNumber: "43",
		Name:        "浦山",
	},
	"1182_43": Observation{
		ID:          ObservationID("1182"),
		GroupNumber: "43",
		Name:        "上吉田",
	},
	"1232_43": Observation{
		ID:          ObservationID("1232"),
		GroupNumber: "43",
		Name:        "鳩山",
	},
	"1497_43": Observation{
		ID:          ObservationID("1497"),
		GroupNumber: "43",
		Name:        "ときがわ",
	},
	"47626_43": Observation{
		ID:          ObservationID("47626"),
		GroupNumber: "43",
		Name:        "熊谷",
	},
	"47641_43": Observation{
		ID:          ObservationID("47641"),
		GroupNumber: "43",
		Name:        "秩父",
	},
	"0365_44": Observation{
		ID:          ObservationID("0365"),
		GroupNumber: "44",
		Name:        "小河内",
	},
	"0366_44": Observation{
		ID:          ObservationID("0366"),
		GroupNumber: "44",
		Name:        "八王子",
	},
	"0367_44": Observation{
		ID:          ObservationID("0367"),
		GroupNumber: "44",
		Name:        "調布",
	},
	"0368_44": Observation{
		ID:          ObservationID("0368"),
		GroupNumber: "44",
		Name:        "世田谷",
	},
	"0370_44": Observation{
		ID:          ObservationID("0370"),
		GroupNumber: "44",
		Name:        "江戸川臨海",
	},
	"0371_44": Observation{
		ID:          ObservationID("0371"),
		GroupNumber: "44",
		Name:        "羽田",
	},
	"1001_44": Observation{
		ID:          ObservationID("1001"),
		GroupNumber: "44",
		Name:        "青梅",
	},
	"1002_44": Observation{
		ID:          ObservationID("1002"),
		GroupNumber: "44",
		Name:        "練馬",
	},
	"1133_44": Observation{
		ID:          ObservationID("1133"),
		GroupNumber: "44",
		Name:        "府中",
	},
	"1134_44": Observation{
		ID:          ObservationID("1134"),
		GroupNumber: "44",
		Name:        "新島",
	},
	"1180_44": Observation{
		ID:          ObservationID("1180"),
		GroupNumber: "44",
		Name:        "小沢",
	},
	"1467_44": Observation{
		ID:          ObservationID("1467"),
		GroupNumber: "44",
		Name:        "大島北ノ山",
	},
	"1468_44": Observation{
		ID:          ObservationID("1468"),
		GroupNumber: "44",
		Name:        "八重見ヶ原",
	},
	"1505_44": Observation{
		ID:          ObservationID("1505"),
		GroupNumber: "44",
		Name:        "神津島",
	},
	"1506_44": Observation{
		ID:          ObservationID("1506"),
		GroupNumber: "44",
		Name:        "三宅坪田",
	},
	"1509_44": Observation{
		ID:          ObservationID("1509"),
		GroupNumber: "44",
		Name:        "三宅島阿古",
	},
	"1522_44": Observation{
		ID:          ObservationID("1522"),
		GroupNumber: "44",
		Name:        "新島",
	},
	"1523_44": Observation{
		ID:          ObservationID("1523"),
		GroupNumber: "44",
		Name:        "神津島",
	},
	"1525_44": Observation{
		ID:          ObservationID("1525"),
		GroupNumber: "44",
		Name:        "三宅伊豆",
	},
	"1595_44": Observation{
		ID:          ObservationID("1595"),
		GroupNumber: "44",
		Name:        "母島",
	},
	"47662_44": Observation{
		ID:          ObservationID("47662"),
		GroupNumber: "44",
		Name:        "東京",
	},
	"47675_44": Observation{
		ID:          ObservationID("47675"),
		GroupNumber: "44",
		Name:        "大島",
	},
	"47677_44": Observation{
		ID:          ObservationID("47677"),
		GroupNumber: "44",
		Name:        "三宅島",
	},
	"47678_44": Observation{
		ID:          ObservationID("47678"),
		GroupNumber: "44",
		Name:        "八丈島",
	},
	"47971_44": Observation{
		ID:          ObservationID("47971"),
		GroupNumber: "44",
		Name:        "父島",
	},
	"47991_44": Observation{
		ID:          ObservationID("47991"),
		GroupNumber: "44",
		Name:        "南鳥島",
	},
	"0375_45": Observation{
		ID:          ObservationID("0375"),
		GroupNumber: "45",
		Name:        "香取",
	},
	"0376_45": Observation{
		ID:          ObservationID("0376"),
		GroupNumber: "45",
		Name:        "我孫子",
	},
	"0377_45": Observation{
		ID:          ObservationID("0377"),
		GroupNumber: "45",
		Name:        "東庄",
	},
	"0378_45": Observation{
		ID:          ObservationID("0378"),
		GroupNumber: "45",
		Name:        "成田",
	},
	"0381_45": Observation{
		ID:          ObservationID("0381"),
		GroupNumber: "45",
		Name:        "茂原",
	},
	"0382_45": Observation{
		ID:          ObservationID("0382"),
		GroupNumber: "45",
		Name:        "木更津",
	},
	"0383_45": Observation{
		ID:          ObservationID("0383"),
		GroupNumber: "45",
		Name:        "鋸南",
	},
	"0384_45": Observation{
		ID:          ObservationID("0384"),
		GroupNumber: "45",
		Name:        "鴨川",
	},
	"0916_45": Observation{
		ID:          ObservationID("0916"),
		GroupNumber: "45",
		Name:        "佐倉",
	},
	"1003_45": Observation{
		ID:          ObservationID("1003"),
		GroupNumber: "45",
		Name:        "横芝光",
	},
	"1004_45": Observation{
		ID:          ObservationID("1004"),
		GroupNumber: "45",
		Name:        "大多喜",
	},
	"1236_45": Observation{
		ID:          ObservationID("1236"),
		GroupNumber: "45",
		Name:        "船橋",
	},
	"1238_45": Observation{
		ID:          ObservationID("1238"),
		GroupNumber: "45",
		Name:        "牛久",
	},
	"1241_45": Observation{
		ID:          ObservationID("1241"),
		GroupNumber: "45",
		Name:        "坂畑",
	},
	"47648_45": Observation{
		ID:          ObservationID("47648"),
		GroupNumber: "45",
		Name:        "銚子",
	},
	"47672_45": Observation{
		ID:          ObservationID("47672"),
		GroupNumber: "45",
		Name:        "館山",
	},
	"47674_45": Observation{
		ID:          ObservationID("47674"),
		GroupNumber: "45",
		Name:        "勝浦",
	},
	"47682_45": Observation{
		ID:          ObservationID("47682"),
		GroupNumber: "45",
		Name:        "千葉",
	},
	"0387_46": Observation{
		ID:          ObservationID("0387"),
		GroupNumber: "46",
		Name:        "相模湖",
	},
	"0388_46": Observation{
		ID:          ObservationID("0388"),
		GroupNumber: "46",
		Name:        "海老名",
	},
	"0390_46": Observation{
		ID:          ObservationID("0390"),
		GroupNumber: "46",
		Name:        "箱根",
	},
	"0391_46": Observation{
		ID:          ObservationID("0391"),
		GroupNumber: "46",
		Name:        "江ノ島",
	},
	"0392_46": Observation{
		ID:          ObservationID("0392"),
		GroupNumber: "46",
		Name:        "三浦",
	},
	"1005_46": Observation{
		ID:          ObservationID("1005"),
		GroupNumber: "46",
		Name:        "相模原中央",
	},
	"1006_46": Observation{
		ID:          ObservationID("1006"),
		GroupNumber: "46",
		Name:        "日吉",
	},
	"1007_46": Observation{
		ID:          ObservationID("1007"),
		GroupNumber: "46",
		Name:        "平塚",
	},
	"1008_46": Observation{
		ID:          ObservationID("1008"),
		GroupNumber: "46",
		Name:        "小田原",
	},
	"1096_46": Observation{
		ID:          ObservationID("1096"),
		GroupNumber: "46",
		Name:        "丹沢湖",
	},
	"1443_46": Observation{
		ID:          ObservationID("1443"),
		GroupNumber: "46",
		Name:        "辻堂",
	},
	"47670_46": Observation{
		ID:          ObservationID("47670"),
		GroupNumber: "46",
		Name:        "横浜",
	},
	"0393_48": Observation{
		ID:          ObservationID("0393"),
		GroupNumber: "48",
		Name:        "野沢温泉",
	},
	"0394_48": Observation{
		ID:          ObservationID("0394"),
		GroupNumber: "48",
		Name:        "信濃町",
	},
	"0395_48": Observation{
		ID:          ObservationID("0395"),
		GroupNumber: "48",
		Name:        "飯山",
	},
	"0396_48": Observation{
		ID:          ObservationID("0396"),
		GroupNumber: "48",
		Name:        "白馬",
	},
	"0398_48": Observation{
		ID:          ObservationID("0398"),
		GroupNumber: "48",
		Name:        "笠岳",
	},
	"0399_48": Observation{
		ID:          ObservationID("0399"),
		GroupNumber: "48",
		Name:        "大町",
	},
	"0400_48": Observation{
		ID:          ObservationID("0400"),
		GroupNumber: "48",
		Name:        "信州新町",
	},
	"0401_48": Observation{
		ID:          ObservationID("0401"),
		GroupNumber: "48",
		Name:        "四阿屋山",
	},
	"0402_48": Observation{
		ID:          ObservationID("0402"),
		GroupNumber: "48",
		Name:        "上田",
	},
	"0403_48": Observation{
		ID:          ObservationID("0403"),
		GroupNumber: "48",
		Name:        "燕岳",
	},
	"0404_48": Observation{
		ID:          ObservationID("0404"),
		GroupNumber: "48",
		Name:        "穂高",
	},
	"0407_48": Observation{
		ID:          ObservationID("0407"),
		GroupNumber: "48",
		Name:        "鹿教湯",
	},
	"0408_48": Observation{
		ID:          ObservationID("0408"),
		GroupNumber: "48",
		Name:        "佐久",
	},
	"0409_48": Observation{
		ID:          ObservationID("0409"),
		GroupNumber: "48",
		Name:        "高ボッチ",
	},
	"0410_48": Observation{
		ID:          ObservationID("0410"),
		GroupNumber: "48",
		Name:        "茂来山",
	},
	"0412_48": Observation{
		ID:          ObservationID("0412"),
		GroupNumber: "48",
		Name:        "八ケ岳",
	},
	"0413_48": Observation{
		ID:          ObservationID("0413"),
		GroupNumber: "48",
		Name:        "木曽平沢",
	},
	"0414_48": Observation{
		ID:          ObservationID("0414"),
		GroupNumber: "48",
		Name:        "原村",
	},
	"0415_48": Observation{
		ID:          ObservationID("0415"),
		GroupNumber: "48",
		Name:        "野辺山",
	},
	"0416_48": Observation{
		ID:          ObservationID("0416"),
		GroupNumber: "48",
		Name:        "御嶽山",
	},
	"0417_48": Observation{
		ID:          ObservationID("0417"),
		GroupNumber: "48",
		Name:        "木曽福島",
	},
	"0418_48": Observation{
		ID:          ObservationID("0418"),
		GroupNumber: "48",
		Name:        "高遠",
	},
	"0419_48": Observation{
		ID:          ObservationID("0419"),
		GroupNumber: "48",
		Name:        "入笠山",
	},
	"0420_48": Observation{
		ID:          ObservationID("0420"),
		GroupNumber: "48",
		Name:        "宮田高原",
	},
	"0421_48": Observation{
		ID:          ObservationID("0421"),
		GroupNumber: "48",
		Name:        "松峯",
	},
	"0422_48": Observation{
		ID:          ObservationID("0422"),
		GroupNumber: "48",
		Name:        "須原",
	},
	"0423_48": Observation{
		ID:          ObservationID("0423"),
		GroupNumber: "48",
		Name:        "摺古木山",
	},
	"0424_48": Observation{
		ID:          ObservationID("0424"),
		GroupNumber: "48",
		Name:        "飯島",
	},
	"0425_48": Observation{
		ID:          ObservationID("0425"),
		GroupNumber: "48",
		Name:        "笹山",
	},
	"0427_48": Observation{
		ID:          ObservationID("0427"),
		GroupNumber: "48",
		Name:        "恵那山",
	},
	"0428_48": Observation{
		ID:          ObservationID("0428"),
		GroupNumber: "48",
		Name:        "鶯巣",
	},
	"0992_48": Observation{
		ID:          ObservationID("0992"),
		GroupNumber: "48",
		Name:        "菅平",
	},
	"0993_48": Observation{
		ID:          ObservationID("0993"),
		GroupNumber: "48",
		Name:        "東御",
	},
	"0994_48": Observation{
		ID:          ObservationID("0994"),
		GroupNumber: "48",
		Name:        "上高地",
	},
	"1022_48": Observation{
		ID:          ObservationID("1022"),
		GroupNumber: "48",
		Name:        "南木曽",
	},
	"1090_48": Observation{
		ID:          ObservationID("1090"),
		GroupNumber: "48",
		Name:        "南小谷",
	},
	"1162_48": Observation{
		ID:          ObservationID("1162"),
		GroupNumber: "48",
		Name:        "鬼無里",
	},
	"1300_48": Observation{
		ID:          ObservationID("1300"),
		GroupNumber: "48",
		Name:        "立科",
	},
	"1312_48": Observation{
		ID:          ObservationID("1312"),
		GroupNumber: "48",
		Name:        "奈川",
	},
	"1314_48": Observation{
		ID:          ObservationID("1314"),
		GroupNumber: "48",
		Name:        "開田高原",
	},
	"1317_48": Observation{
		ID:          ObservationID("1317"),
		GroupNumber: "48",
		Name:        "浪合",
	},
	"1319_48": Observation{
		ID:          ObservationID("1319"),
		GroupNumber: "48",
		Name:        "辰野",
	},
	"1323_48": Observation{
		ID:          ObservationID("1323"),
		GroupNumber: "48",
		Name:        "南信濃",
	},
	"1339_48": Observation{
		ID:          ObservationID("1339"),
		GroupNumber: "48",
		Name:        "阿南",
	},
	"1360_48": Observation{
		ID:          ObservationID("1360"),
		GroupNumber: "48",
		Name:        "聖高原",
	},
	"1392_48": Observation{
		ID:          ObservationID("1392"),
		GroupNumber: "48",
		Name:        "網掛山",
	},
	"1394_48": Observation{
		ID:          ObservationID("1394"),
		GroupNumber: "48",
		Name:        "大鹿",
	},
	"1397_48": Observation{
		ID:          ObservationID("1397"),
		GroupNumber: "48",
		Name:        "杉島",
	},
	"1399_48": Observation{
		ID:          ObservationID("1399"),
		GroupNumber: "48",
		Name:        "十石峠",
	},
	"1411_48": Observation{
		ID:          ObservationID("1411"),
		GroupNumber: "48",
		Name:        "小谷",
	},
	"1445_48": Observation{
		ID:          ObservationID("1445"),
		GroupNumber: "48",
		Name:        "伊那",
	},
	"1529_48": Observation{
		ID:          ObservationID("1529"),
		GroupNumber: "48",
		Name:        "松本今井",
	},
	"1549_48": Observation{
		ID:          ObservationID("1549"),
		GroupNumber: "48",
		Name:        "北相木",
	},
	"1550_48": Observation{
		ID:          ObservationID("1550"),
		GroupNumber: "48",
		Name:        "白樺湖",
	},
	"1551_48": Observation{
		ID:          ObservationID("1551"),
		GroupNumber: "48",
		Name:        "高遠",
	},
	"1591_48": Observation{
		ID:          ObservationID("1591"),
		GroupNumber: "48",
		Name:        "宮田",
	},
	"47610_48": Observation{
		ID:          ObservationID("47610"),
		GroupNumber: "48",
		Name:        "長野",
	},
	"47618_48": Observation{
		ID:          ObservationID("47618"),
		GroupNumber: "48",
		Name:        "松本",
	},
	"47620_48": Observation{
		ID:          ObservationID("47620"),
		GroupNumber: "48",
		Name:        "諏訪",
	},
	"47622_48": Observation{
		ID:          ObservationID("47622"),
		GroupNumber: "48",
		Name:        "軽井沢",
	},
	"47637_48": Observation{
		ID:          ObservationID("47637"),
		GroupNumber: "48",
		Name:        "飯田",
	},
	"0429_49": Observation{
		ID:          ObservationID("0429"),
		GroupNumber: "49",
		Name:        "剣ノ峰",
	},
	"0430_49": Observation{
		ID:          ObservationID("0430"),
		GroupNumber: "49",
		Name:        "日向山",
	},
	"0431_49": Observation{
		ID:          ObservationID("0431"),
		GroupNumber: "49",
		Name:        "大菩薩",
	},
	"0433_49": Observation{
		ID:          ObservationID("0433"),
		GroupNumber: "49",
		Name:        "勝沼",
	},
	"0434_49": Observation{
		ID:          ObservationID("0434"),
		GroupNumber: "49",
		Name:        "大月",
	},
	"0435_49": Observation{
		ID:          ObservationID("0435"),
		GroupNumber: "49",
		Name:        "八町山",
	},
	"0436_49": Observation{
		ID:          ObservationID("0436"),
		GroupNumber: "49",
		Name:        "古関",
	},
	"0438_49": Observation{
		ID:          ObservationID("0438"),
		GroupNumber: "49",
		Name:        "山中",
	},
	"0439_49": Observation{
		ID:          ObservationID("0439"),
		GroupNumber: "49",
		Name:        "南部",
	},
	"1023_49": Observation{
		ID:          ObservationID("1023"),
		GroupNumber: "49",
		Name:        "大泉",
	},
	"1024_49": Observation{
		ID:          ObservationID("1024"),
		GroupNumber: "49",
		Name:        "韮崎",
	},
	"1165_49": Observation{
		ID:          ObservationID("1165"),
		GroupNumber: "49",
		Name:        "上野原",
	},
	"1234_49": Observation{
		ID:          ObservationID("1234"),
		GroupNumber: "49",
		Name:        "切石",
	},
	"1599_49": Observation{
		ID:          ObservationID("1599"),
		GroupNumber: "49",
		Name:        "乙女湖",
	},
	"47638_49": Observation{
		ID:          ObservationID("47638"),
		GroupNumber: "49",
		Name:        "甲府",
	},
	"47639_49": Observation{
		ID:          ObservationID("47639"),
		GroupNumber: "49",
		Name:        "富士山",
	},
	"47640_49": Observation{
		ID:          ObservationID("47640"),
		GroupNumber: "49",
		Name:        "河口湖",
	},
	"0440_50": Observation{
		ID:          ObservationID("0440"),
		GroupNumber: "50",
		Name:        "白糸",
	},
	"0441_50": Observation{
		ID:          ObservationID("0441"),
		GroupNumber: "50",
		Name:        "御殿場",
	},
	"0442_50": Observation{
		ID:          ObservationID("0442"),
		GroupNumber: "50",
		Name:        "富士",
	},
	"0444_50": Observation{
		ID:          ObservationID("0444"),
		GroupNumber: "50",
		Name:        "越木平",
	},
	"0445_50": Observation{
		ID:          ObservationID("0445"),
		GroupNumber: "50",
		Name:        "川根本町",
	},
	"0446_50": Observation{
		ID:          ObservationID("0446"),
		GroupNumber: "50",
		Name:        "大山",
	},
	"0448_50": Observation{
		ID:          ObservationID("0448"),
		GroupNumber: "50",
		Name:        "高根山",
	},
	"0450_50": Observation{
		ID:          ObservationID("0450"),
		GroupNumber: "50",
		Name:        "霧山",
	},
	"0451_50": Observation{
		ID:          ObservationID("0451"),
		GroupNumber: "50",
		Name:        "天竜",
	},
	"0452_50": Observation{
		ID:          ObservationID("0452"),
		GroupNumber: "50",
		Name:        "島田",
	},
	"0453_50": Observation{
		ID:          ObservationID("0453"),
		GroupNumber: "50",
		Name:        "湯ケ島",
	},
	"0454_50": Observation{
		ID:          ObservationID("0454"),
		GroupNumber: "50",
		Name:        "天城山",
	},
	"0456_50": Observation{
		ID:          ObservationID("0456"),
		GroupNumber: "50",
		Name:        "松崎",
	},
	"0457_50": Observation{
		ID:          ObservationID("0457"),
		GroupNumber: "50",
		Name:        "稲取",
	},
	"0986_50": Observation{
		ID:          ObservationID("0986"),
		GroupNumber: "50",
		Name:        "佐久間",
	},
	"0987_50": Observation{
		ID:          ObservationID("0987"),
		GroupNumber: "50",
		Name:        "土肥",
	},
	"0988_50": Observation{
		ID:          ObservationID("0988"),
		GroupNumber: "50",
		Name:        "三ヶ日",
	},
	"0989_50": Observation{
		ID:          ObservationID("0989"),
		GroupNumber: "50",
		Name:        "掛川",
	},
	"1105_50": Observation{
		ID:          ObservationID("1105"),
		GroupNumber: "50",
		Name:        "熊",
	},
	"1106_50": Observation{
		ID:          ObservationID("1106"),
		GroupNumber: "50",
		Name:        "三倉",
	},
	"1114_50": Observation{
		ID:          ObservationID("1114"),
		GroupNumber: "50",
		Name:        "梅ケ島",
	},
	"1243_50": Observation{
		ID:          ObservationID("1243"),
		GroupNumber: "50",
		Name:        "清水",
	},
	"1244_50": Observation{
		ID:          ObservationID("1244"),
		GroupNumber: "50",
		Name:        "磐田",
	},
	"1335_50": Observation{
		ID:          ObservationID("1335"),
		GroupNumber: "50",
		Name:        "菊川牧之原",
	},
	"1338_50": Observation{
		ID:          ObservationID("1338"),
		GroupNumber: "50",
		Name:        "井川",
	},
	"1393_50": Observation{
		ID:          ObservationID("1393"),
		GroupNumber: "50",
		Name:        "三倉",
	},
	"1440_50": Observation{
		ID:          ObservationID("1440"),
		GroupNumber: "50",
		Name:        "鍵穴",
	},
	"1601_50": Observation{
		ID:          ObservationID("1601"),
		GroupNumber: "50",
		Name:        "静岡空港",
	},
	"47639_50": Observation{
		ID:          ObservationID("47639"),
		GroupNumber: "50",
		Name:        "富士山",
	},
	"47654_50": Observation{
		ID:          ObservationID("47654"),
		GroupNumber: "50",
		Name:        "浜松",
	},
	"47655_50": Observation{
		ID:          ObservationID("47655"),
		GroupNumber: "50",
		Name:        "御前崎",
	},
	"47656_50": Observation{
		ID:          ObservationID("47656"),
		GroupNumber: "50",
		Name:        "静岡",
	},
	"47657_50": Observation{
		ID:          ObservationID("47657"),
		GroupNumber: "50",
		Name:        "三島",
	},
	"47666_50": Observation{
		ID:          ObservationID("47666"),
		GroupNumber: "50",
		Name:        "石廊崎",
	},
	"47668_50": Observation{
		ID:          ObservationID("47668"),
		GroupNumber: "50",
		Name:        "網代",
	},
	"0460_51": Observation{
		ID:          ObservationID("0460"),
		GroupNumber: "51",
		Name:        "一宮",
	},
	"0461_51": Observation{
		ID:          ObservationID("0461"),
		GroupNumber: "51",
		Name:        "豊山",
	},
	"0462_51": Observation{
		ID:          ObservationID("0462"),
		GroupNumber: "51",
		Name:        "茶臼山",
	},
	"0464_51": Observation{
		ID:          ObservationID("0464"),
		GroupNumber: "51",
		Name:        "豊田",
	},
	"0465_51": Observation{
		ID:          ObservationID("0465"),
		GroupNumber: "51",
		Name:        "東海",
	},
	"0466_51": Observation{
		ID:          ObservationID("0466"),
		GroupNumber: "51",
		Name:        "出来山",
	},
	"0467_51": Observation{
		ID:          ObservationID("0467"),
		GroupNumber: "51",
		Name:        "岡崎",
	},
	"0468_51": Observation{
		ID:          ObservationID("0468"),
		GroupNumber: "51",
		Name:        "作手",
	},
	"0469_51": Observation{
		ID:          ObservationID("0469"),
		GroupNumber: "51",
		Name:        "一色",
	},
	"0470_51": Observation{
		ID:          ObservationID("0470"),
		GroupNumber: "51",
		Name:        "豊橋",
	},
	"0472_51": Observation{
		ID:          ObservationID("0472"),
		GroupNumber: "51",
		Name:        "田原",
	},
	"0982_51": Observation{
		ID:          ObservationID("0982"),
		GroupNumber: "51",
		Name:        "稲武",
	},
	"0983_51": Observation{
		ID:          ObservationID("0983"),
		GroupNumber: "51",
		Name:        "蟹江",
	},
	"0984_51": Observation{
		ID:          ObservationID("0984"),
		GroupNumber: "51",
		Name:        "南知多",
	},
	"1166_51": Observation{
		ID:          ObservationID("1166"),
		GroupNumber: "51",
		Name:        "小原",
	},
	"1341_51": Observation{
		ID:          ObservationID("1341"),
		GroupNumber: "51",
		Name:        "鳳来",
	},
	"1344_51": Observation{
		ID:          ObservationID("1344"),
		GroupNumber: "51",
		Name:        "蒲郡",
	},
	"1346_51": Observation{
		ID:          ObservationID("1346"),
		GroupNumber: "51",
		Name:        "愛西",
	},
	"1541_51": Observation{
		ID:          ObservationID("1541"),
		GroupNumber: "51",
		Name:        "新城",
	},
	"1555_51": Observation{
		ID:          ObservationID("1555"),
		GroupNumber: "51",
		Name:        "セントレア",
	},
	"1576_51": Observation{
		ID:          ObservationID("1576"),
		GroupNumber: "51",
		Name:        "阿蔵",
	},
	"47636_51": Observation{
		ID:          ObservationID("47636"),
		GroupNumber: "51",
		Name:        "名古屋",
	},
	"47653_51": Observation{
		ID:          ObservationID("47653"),
		GroupNumber: "51",
		Name:        "伊良湖",
	},
	"0473_52": Observation{
		ID:          ObservationID("0473"),
		GroupNumber: "52",
		Name:        "神岡",
	},
	"0474_52": Observation{
		ID:          ObservationID("0474"),
		GroupNumber: "52",
		Name:        "流葉山",
	},
	"0475_52": Observation{
		ID:          ObservationID("0475"),
		GroupNumber: "52",
		Name:        "森茂",
	},
	"0477_52": Observation{
		ID:          ObservationID("0477"),
		GroupNumber: "52",
		Name:        "乗鞍岳",
	},
	"0478_52": Observation{
		ID:          ObservationID("0478"),
		GroupNumber: "52",
		Name:        "大日岳",
	},
	"0479_52": Observation{
		ID:          ObservationID("0479"),
		GroupNumber: "52",
		Name:        "新渕山",
	},
	"0480_52": Observation{
		ID:          ObservationID("0480"),
		GroupNumber: "52",
		Name:        "船山",
	},
	"0481_52": Observation{
		ID:          ObservationID("0481"),
		GroupNumber: "52",
		Name:        "白尾山",
	},
	"0482_52": Observation{
		ID:          ObservationID("0482"),
		GroupNumber: "52",
		Name:        "白鳥",
	},
	"0483_52": Observation{
		ID:          ObservationID("0483"),
		GroupNumber: "52",
		Name:        "萩原",
	},
	"0484_52": Observation{
		ID:          ObservationID("0484"),
		GroupNumber: "52",
		Name:        "蕪山",
	},
	"0485_52": Observation{
		ID:          ObservationID("0485"),
		GroupNumber: "52",
		Name:        "八幡",
	},
	"0486_52": Observation{
		ID:          ObservationID("0486"),
		GroupNumber: "52",
		Name:        "樽見",
	},
	"0487_52": Observation{
		ID:          ObservationID("0487"),
		GroupNumber: "52",
		Name:        "金山",
	},
	"0488_52": Observation{
		ID:          ObservationID("0488"),
		GroupNumber: "52",
		Name:        "三界山",
	},
	"0489_52": Observation{
		ID:          ObservationID("0489"),
		GroupNumber: "52",
		Name:        "権現山",
	},
	"0490_52": Observation{
		ID:          ObservationID("0490"),
		GroupNumber: "52",
		Name:        "伽藍",
	},
	"0491_52": Observation{
		ID:          ObservationID("0491"),
		GroupNumber: "52",
		Name:        "美濃加茂",
	},
	"0492_52": Observation{
		ID:          ObservationID("0492"),
		GroupNumber: "52",
		Name:        "柄石峠",
	},
	"0493_52": Observation{
		ID:          ObservationID("0493"),
		GroupNumber: "52",
		Name:        "恵那",
	},
	"0494_52": Observation{
		ID:          ObservationID("0494"),
		GroupNumber: "52",
		Name:        "中津川",
	},
	"0495_52": Observation{
		ID:          ObservationID("0495"),
		GroupNumber: "52",
		Name:        "関ケ原",
	},
	"0496_52": Observation{
		ID:          ObservationID("0496"),
		GroupNumber: "52",
		Name:        "大垣",
	},
	"0498_52": Observation{
		ID:          ObservationID("0498"),
		GroupNumber: "52",
		Name:        "三森山",
	},
	"1057_52": Observation{
		ID:          ObservationID("1057"),
		GroupNumber: "52",
		Name:        "美濃",
	},
	"1058_52": Observation{
		ID:          ObservationID("1058"),
		GroupNumber: "52",
		Name:        "多治見",
	},
	"1059_52": Observation{
		ID:          ObservationID("1059"),
		GroupNumber: "52",
		Name:        "上石津",
	},
	"1061_52": Observation{
		ID:          ObservationID("1061"),
		GroupNumber: "52",
		Name:        "栃尾",
	},
	"1066_52": Observation{
		ID:          ObservationID("1066"),
		GroupNumber: "52",
		Name:        "宮地",
	},
	"1153_52": Observation{
		ID:          ObservationID("1153"),
		GroupNumber: "52",
		Name:        "平井",
	},
	"1299_52": Observation{
		ID:          ObservationID("1299"),
		GroupNumber: "52",
		Name:        "長滝",
	},
	"1301_52": Observation{
		ID:          ObservationID("1301"),
		GroupNumber: "52",
		Name:        "揖斐川",
	},
	"1305_52": Observation{
		ID:          ObservationID("1305"),
		GroupNumber: "52",
		Name:        "六厩",
	},
	"1306_52": Observation{
		ID:          ObservationID("1306"),
		GroupNumber: "52",
		Name:        "白川",
	},
	"1308_52": Observation{
		ID:          ObservationID("1308"),
		GroupNumber: "52",
		Name:        "河合",
	},
	"1311_52": Observation{
		ID:          ObservationID("1311"),
		GroupNumber: "52",
		Name:        "宮之前",
	},
	"1313_52": Observation{
		ID:          ObservationID("1313"),
		GroupNumber: "52",
		Name:        "黒川",
	},
	"1363_52": Observation{
		ID:          ObservationID("1363"),
		GroupNumber: "52",
		Name:        "大日",
	},
	"1384_52": Observation{
		ID:          ObservationID("1384"),
		GroupNumber: "52",
		Name:        "付知",
	},
	"1387_52": Observation{
		ID:          ObservationID("1387"),
		GroupNumber: "52",
		Name:        "御母衣",
	},
	"1390_52": Observation{
		ID:          ObservationID("1390"),
		GroupNumber: "52",
		Name:        "小津",
	},
	"1396_52": Observation{
		ID:          ObservationID("1396"),
		GroupNumber: "52",
		Name:        "十二岳",
	},
	"1436_52": Observation{
		ID:          ObservationID("1436"),
		GroupNumber: "52",
		Name:        "ひるがの",
	},
	"1437_52": Observation{
		ID:          ObservationID("1437"),
		GroupNumber: "52",
		Name:        "清見",
	},
	"1438_52": Observation{
		ID:          ObservationID("1438"),
		GroupNumber: "52",
		Name:        "丹生川",
	},
	"1577_52": Observation{
		ID:          ObservationID("1577"),
		GroupNumber: "52",
		Name:        "関市板取",
	},
	"47617_52": Observation{
		ID:          ObservationID("47617"),
		GroupNumber: "52",
		Name:        "高山",
	},
	"47632_52": Observation{
		ID:          ObservationID("47632"),
		GroupNumber: "52",
		Name:        "岐阜",
	},
	"0499_53": Observation{
		ID:          ObservationID("0499"),
		GroupNumber: "53",
		Name:        "北勢",
	},
	"0500_53": Observation{
		ID:          ObservationID("0500"),
		GroupNumber: "53",
		Name:        "桑名",
	},
	"0501_53": Observation{
		ID:          ObservationID("0501"),
		GroupNumber: "53",
		Name:        "雲母峰",
	},
	"0503_53": Observation{
		ID:          ObservationID("0503"),
		GroupNumber: "53",
		Name:        "亀山",
	},
	"0505_53": Observation{
		ID:          ObservationID("0505"),
		GroupNumber: "53",
		Name:        "笠取山",
	},
	"0507_53": Observation{
		ID:          ObservationID("0507"),
		GroupNumber: "53",
		Name:        "名張",
	},
	"0508_53": Observation{
		ID:          ObservationID("0508"),
		GroupNumber: "53",
		Name:        "国見山",
	},
	"0509_53": Observation{
		ID:          ObservationID("0509"),
		GroupNumber: "53",
		Name:        "小俣",
	},
	"0510_53": Observation{
		ID:          ObservationID("0510"),
		GroupNumber: "53",
		Name:        "粥見",
	},
	"0511_53": Observation{
		ID:          ObservationID("0511"),
		GroupNumber: "53",
		Name:        "藤坂峠",
	},
	"0513_53": Observation{
		ID:          ObservationID("0513"),
		GroupNumber: "53",
		Name:        "八幡峠",
	},
	"0514_53": Observation{
		ID:          ObservationID("0514"),
		GroupNumber: "53",
		Name:        "熊野",
	},
	"0918_53": Observation{
		ID:          ObservationID("0918"),
		GroupNumber: "53",
		Name:        "大台",
	},
	"0990_53": Observation{
		ID:          ObservationID("0990"),
		GroupNumber: "53",
		Name:        "松阪",
	},
	"0991_53": Observation{
		ID:          ObservationID("0991"),
		GroupNumber: "53",
		Name:        "磯部",
	},
	"1230_53": Observation{
		ID:          ObservationID("1230"),
		GroupNumber: "53",
		Name:        "鳥羽",
	},
	"1255_53": Observation{
		ID:          ObservationID("1255"),
		GroupNumber: "53",
		Name:        "宮川",
	},
	"1258_53": Observation{
		ID:          ObservationID("1258"),
		GroupNumber: "53",
		Name:        "南伊勢",
	},
	"1349_53": Observation{
		ID:          ObservationID("1349"),
		GroupNumber: "53",
		Name:        "紀伊長島",
	},
	"1356_53": Observation{
		ID:          ObservationID("1356"),
		GroupNumber: "53",
		Name:        "白山",
	},
	"1381_53": Observation{
		ID:          ObservationID("1381"),
		GroupNumber: "53",
		Name:        "阿児",
	},
	"1423_53": Observation{
		ID:          ObservationID("1423"),
		GroupNumber: "53",
		Name:        "御浜",
	},
	"1532_53": Observation{
		ID:          ObservationID("1532"),
		GroupNumber: "53",
		Name:        "熊野新鹿",
	},
	"47649_53": Observation{
		ID:          ObservationID("47649"),
		GroupNumber: "53",
		Name:        "上野",
	},
	"47651_53": Observation{
		ID:          ObservationID("47651"),
		GroupNumber: "53",
		Name:        "津",
	},
	"47663_53": Observation{
		ID:          ObservationID("47663"),
		GroupNumber: "53",
		Name:        "尾鷲",
	},
	"47684_53": Observation{
		ID:          ObservationID("47684"),
		GroupNumber: "53",
		Name:        "四日市",
	},
	"0515_54": Observation{
		ID:          ObservationID("0515"),
		GroupNumber: "54",
		Name:        "村上",
	},
	"0516_54": Observation{
		ID:          ObservationID("0516"),
		GroupNumber: "54",
		Name:        "鷲ケ巣山",
	},
	"0518_54": Observation{
		ID:          ObservationID("0518"),
		GroupNumber: "54",
		Name:        "両津",
	},
	"0520_54": Observation{
		ID:          ObservationID("0520"),
		GroupNumber: "54",
		Name:        "二王子岳",
	},
	"0521_54": Observation{
		ID:          ObservationID("0521"),
		GroupNumber: "54",
		Name:        "新津",
	},
	"0522_54": Observation{
		ID:          ObservationID("0522"),
		GroupNumber: "54",
		Name:        "巻",
	},
	"0523_54": Observation{
		ID:          ObservationID("0523"),
		GroupNumber: "54",
		Name:        "宝珠山",
	},
	"0524_54": Observation{
		ID:          ObservationID("0524"),
		GroupNumber: "54",
		Name:        "寺泊",
	},
	"0525_54": Observation{
		ID:          ObservationID("0525"),
		GroupNumber: "54",
		Name:        "三条",
	},
	"0526_54": Observation{
		ID:          ObservationID("0526"),
		GroupNumber: "54",
		Name:        "村松",
	},
	"0527_54": Observation{
		ID:          ObservationID("0527"),
		GroupNumber: "54",
		Name:        "津川",
	},
	"0528_54": Observation{
		ID:          ObservationID("0528"),
		GroupNumber: "54",
		Name:        "粟ヶ岳",
	},
	"0529_54": Observation{
		ID:          ObservationID("0529"),
		GroupNumber: "54",
		Name:        "長岡",
	},
	"0530_54": Observation{
		ID:          ObservationID("0530"),
		GroupNumber: "54",
		Name:        "栃尾",
	},
	"0531_54": Observation{
		ID:          ObservationID("0531"),
		GroupNumber: "54",
		Name:        "守門岳",
	},
	"0532_54": Observation{
		ID:          ObservationID("0532"),
		GroupNumber: "54",
		Name:        "柏崎",
	},
	"0533_54": Observation{
		ID:          ObservationID("0533"),
		GroupNumber: "54",
		Name:        "桜坂峠",
	},
	"0534_54": Observation{
		ID:          ObservationID("0534"),
		GroupNumber: "54",
		Name:        "小出",
	},
	"0536_54": Observation{
		ID:          ObservationID("0536"),
		GroupNumber: "54",
		Name:        "安塚",
	},
	"0537_54": Observation{
		ID:          ObservationID("0537"),
		GroupNumber: "54",
		Name:        "十日町",
	},
	"0538_54": Observation{
		ID:          ObservationID("0538"),
		GroupNumber: "54",
		Name:        "枝折峠",
	},
	"0539_54": Observation{
		ID:          ObservationID("0539"),
		GroupNumber: "54",
		Name:        "糸魚川",
	},
	"0540_54": Observation{
		ID:          ObservationID("0540"),
		GroupNumber: "54",
		Name:        "能生",
	},
	"0541_54": Observation{
		ID:          ObservationID("0541"),
		GroupNumber: "54",
		Name:        "光ケ原",
	},
	"0542_54": Observation{
		ID:          ObservationID("0542"),
		GroupNumber: "54",
		Name:        "大沢峠",
	},
	"0543_54": Observation{
		ID:          ObservationID("0543"),
		GroupNumber: "54",
		Name:        "津南",
	},
	"0544_54": Observation{
		ID:          ObservationID("0544"),
		GroupNumber: "54",
		Name:        "湯沢",
	},
	"0545_54": Observation{
		ID:          ObservationID("0545"),
		GroupNumber: "54",
		Name:        "袴岳",
	},
	"0995_54": Observation{
		ID:          ObservationID("0995"),
		GroupNumber: "54",
		Name:        "粟島",
	},
	"0996_54": Observation{
		ID:          ObservationID("0996"),
		GroupNumber: "54",
		Name:        "羽茂",
	},
	"0997_54": Observation{
		ID:          ObservationID("0997"),
		GroupNumber: "54",
		Name:        "入広瀬",
	},
	"0998_54": Observation{
		ID:          ObservationID("0998"),
		GroupNumber: "54",
		Name:        "関山",
	},
	"1060_54": Observation{
		ID:          ObservationID("1060"),
		GroupNumber: "54",
		Name:        "下関",
	},
	"1102_54": Observation{
		ID:          ObservationID("1102"),
		GroupNumber: "54",
		Name:        "鍵取",
	},
	"1111_54": Observation{
		ID:          ObservationID("1111"),
		GroupNumber: "54",
		Name:        "赤谷",
	},
	"1154_54": Observation{
		ID:          ObservationID("1154"),
		GroupNumber: "54",
		Name:        "小国",
	},
	"1155_54": Observation{
		ID:          ObservationID("1155"),
		GroupNumber: "54",
		Name:        "松代",
	},
	"1256_54": Observation{
		ID:          ObservationID("1256"),
		GroupNumber: "54",
		Name:        "高根",
	},
	"1297_54": Observation{
		ID:          ObservationID("1297"),
		GroupNumber: "54",
		Name:        "平岩",
	},
	"1302_54": Observation{
		ID:          ObservationID("1302"),
		GroupNumber: "54",
		Name:        "大潟",
	},
	"1315_54": Observation{
		ID:          ObservationID("1315"),
		GroupNumber: "54",
		Name:        "弾崎",
	},
	"1318_54": Observation{
		ID:          ObservationID("1318"),
		GroupNumber: "54",
		Name:        "中条",
	},
	"1369_54": Observation{
		ID:          ObservationID("1369"),
		GroupNumber: "54",
		Name:        "三面",
	},
	"1370_54": Observation{
		ID:          ObservationID("1370"),
		GroupNumber: "54",
		Name:        "大湯",
	},
	"1395_54": Observation{
		ID:          ObservationID("1395"),
		GroupNumber: "54",
		Name:        "塩沢",
	},
	"1398_54": Observation{
		ID:          ObservationID("1398"),
		GroupNumber: "54",
		Name:        "宮寄上",
	},
	"1424_54": Observation{
		ID:          ObservationID("1424"),
		GroupNumber: "54",
		Name:        "樽本",
	},
	"1426_54": Observation{
		ID:          ObservationID("1426"),
		GroupNumber: "54",
		Name:        "川谷",
	},
	"1427_54": Observation{
		ID:          ObservationID("1427"),
		GroupNumber: "54",
		Name:        "筒方",
	},
	"1444_54": Observation{
		ID:          ObservationID("1444"),
		GroupNumber: "54",
		Name:        "室谷",
	},
	"1469_54": Observation{
		ID:          ObservationID("1469"),
		GroupNumber: "54",
		Name:        "松浜",
	},
	"1514_54": Observation{
		ID:          ObservationID("1514"),
		GroupNumber: "54",
		Name:        "秋津",
	},
	"47602_54": Observation{
		ID:          ObservationID("47602"),
		GroupNumber: "54",
		Name:        "相川",
	},
	"47604_54": Observation{
		ID:          ObservationID("47604"),
		GroupNumber: "54",
		Name:        "新潟",
	},
	"47612_54": Observation{
		ID:          ObservationID("47612"),
		GroupNumber: "54",
		Name:        "高田",
	},
	"0546_55": Observation{
		ID:          ObservationID("0546"),
		GroupNumber: "55",
		Name:        "泊",
	},
	"0547_55": Observation{
		ID:          ObservationID("0547"),
		GroupNumber: "55",
		Name:        "魚津",
	},
	"0548_55": Observation{
		ID:          ObservationID("0548"),
		GroupNumber: "55",
		Name:        "宇奈月",
	},
	"0549_55": Observation{
		ID:          ObservationID("0549"),
		GroupNumber: "55",
		Name:        "嘉例沢",
	},
	"0552_55": Observation{
		ID:          ObservationID("0552"),
		GroupNumber: "55",
		Name:        "砺波",
	},
	"0553_55": Observation{
		ID:          ObservationID("0553"),
		GroupNumber: "55",
		Name:        "上市",
	},
	"0554_55": Observation{
		ID:          ObservationID("0554"),
		GroupNumber: "55",
		Name:        "南砺高宮",
	},
	"0555_55": Observation{
		ID:          ObservationID("0555"),
		GroupNumber: "55",
		Name:        "八尾",
	},
	"0556_55": Observation{
		ID:          ObservationID("0556"),
		GroupNumber: "55",
		Name:        "小谷",
	},
	"0557_55": Observation{
		ID:          ObservationID("0557"),
		GroupNumber: "55",
		Name:        "立山",
	},
	"0558_55": Observation{
		ID:          ObservationID("0558"),
		GroupNumber: "55",
		Name:        "細尾峠",
	},
	"0559_55": Observation{
		ID:          ObservationID("0559"),
		GroupNumber: "55",
		Name:        "谷折",
	},
	"1309_55": Observation{
		ID:          ObservationID("1309"),
		GroupNumber: "55",
		Name:        "氷見",
	},
	"1412_55": Observation{
		ID:          ObservationID("1412"),
		GroupNumber: "55",
		Name:        "宇奈月",
	},
	"1413_55": Observation{
		ID:          ObservationID("1413"),
		GroupNumber: "55",
		Name:        "平",
	},
	"1414_55": Observation{
		ID:          ObservationID("1414"),
		GroupNumber: "55",
		Name:        "大山",
	},
	"1415_55": Observation{
		ID:          ObservationID("1415"),
		GroupNumber: "55",
		Name:        "細入",
	},
	"1455_55": Observation{
		ID:          ObservationID("1455"),
		GroupNumber: "55",
		Name:        "猪谷",
	},
	"1533_55": Observation{
		ID:          ObservationID("1533"),
		GroupNumber: "55",
		Name:        "秋ヶ島",
	},
	"1553_55": Observation{
		ID:          ObservationID("1553"),
		GroupNumber: "55",
		Name:        "五箇山",
	},
	"1621_55": Observation{
		ID:          ObservationID("1621"),
		GroupNumber: "55",
		Name:        "立山芦峅",
	},
	"47606_55": Observation{
		ID:          ObservationID("47606"),
		GroupNumber: "55",
		Name:        "伏木",
	},
	"47607_55": Observation{
		ID:          ObservationID("47607"),
		GroupNumber: "55",
		Name:        "富山",
	},
	"0560_56": Observation{
		ID:          ObservationID("0560"),
		GroupNumber: "56",
		Name:        "珠洲",
	},
	"0562_56": Observation{
		ID:          ObservationID("0562"),
		GroupNumber: "56",
		Name:        "門前",
	},
	"0563_56": Observation{
		ID:          ObservationID("0563"),
		GroupNumber: "56",
		Name:        "木原岳",
	},
	"0564_56": Observation{
		ID:          ObservationID("0564"),
		GroupNumber: "56",
		Name:        "志賀",
	},
	"0565_56": Observation{
		ID:          ObservationID("0565"),
		GroupNumber: "56",
		Name:        "七尾",
	},
	"0566_56": Observation{
		ID:          ObservationID("0566"),
		GroupNumber: "56",
		Name:        "羽咋",
	},
	"0567_56": Observation{
		ID:          ObservationID("0567"),
		GroupNumber: "56",
		Name:        "かほく",
	},
	"0568_56": Observation{
		ID:          ObservationID("0568"),
		GroupNumber: "56",
		Name:        "宝達山",
	},
	"0570_56": Observation{
		ID:          ObservationID("0570"),
		GroupNumber: "56",
		Name:        "医王山",
	},
	"0571_56": Observation{
		ID:          ObservationID("0571"),
		GroupNumber: "56",
		Name:        "栢野",
	},
	"0973_56": Observation{
		ID:          ObservationID("0973"),
		GroupNumber: "56",
		Name:        "白山吉野",
	},
	"1171_56": Observation{
		ID:          ObservationID("1171"),
		GroupNumber: "56",
		Name:        "白山白峰",
	},
	"1324_56": Observation{
		ID:          ObservationID("1324"),
		GroupNumber: "56",
		Name:        "小松",
	},
	"1486_56": Observation{
		ID:          ObservationID("1486"),
		GroupNumber: "56",
		Name:        "穴水",
	},
	"1544_56": Observation{
		ID:          ObservationID("1544"),
		GroupNumber: "56",
		Name:        "三井",
	},
	"1545_56": Observation{
		ID:          ObservationID("1545"),
		GroupNumber: "56",
		Name:        "宝達志水",
	},
	"47600_56": Observation{
		ID:          ObservationID("47600"),
		GroupNumber: "56",
		Name:        "輪島",
	},
	"47605_56": Observation{
		ID:          ObservationID("47605"),
		GroupNumber: "56",
		Name:        "金沢",
	},
	"0573_57": Observation{
		ID:          ObservationID("0573"),
		GroupNumber: "57",
		Name:        "大野",
	},
	"0574_57": Observation{
		ID:          ObservationID("0574"),
		GroupNumber: "57",
		Name:        "タイラ山",
	},
	"0575_57": Observation{
		ID:          ObservationID("0575"),
		GroupNumber: "57",
		Name:        "春日野",
	},
	"0576_57": Observation{
		ID:          ObservationID("0576"),
		GroupNumber: "57",
		Name:        "板垣",
	},
	"0577_57": Observation{
		ID:          ObservationID("0577"),
		GroupNumber: "57",
		Name:        "今庄",
	},
	"0579_57": Observation{
		ID:          ObservationID("0579"),
		GroupNumber: "57",
		Name:        "小浜",
	},
	"0974_57": Observation{
		ID:          ObservationID("0974"),
		GroupNumber: "57",
		Name:        "美山",
	},
	"0975_57": Observation{
		ID:          ObservationID("0975"),
		GroupNumber: "57",
		Name:        "川上",
	},
	"1010_57": Observation{
		ID:          ObservationID("1010"),
		GroupNumber: "57",
		Name:        "美浜",
	},
	"1071_57": Observation{
		ID:          ObservationID("1071"),
		GroupNumber: "57",
		Name:        "三国",
	},
	"1226_57": Observation{
		ID:          ObservationID("1226"),
		GroupNumber: "57",
		Name:        "勝山",
	},
	"1316_57": Observation{
		ID:          ObservationID("1316"),
		GroupNumber: "57",
		Name:        "越廼",
	},
	"1382_57": Observation{
		ID:          ObservationID("1382"),
		GroupNumber: "57",
		Name:        "九頭竜",
	},
	"1454_57": Observation{
		ID:          ObservationID("1454"),
		GroupNumber: "57",
		Name:        "大飯",
	},
	"1537_57": Observation{
		ID:          ObservationID("1537"),
		GroupNumber: "57",
		Name:        "春江",
	},
	"1565_57": Observation{
		ID:          ObservationID("1565"),
		GroupNumber: "57",
		Name:        "武生",
	},
	"47616_57": Observation{
		ID:          ObservationID("47616"),
		GroupNumber: "57",
		Name:        "福井",
	},
	"47631_57": Observation{
		ID:          ObservationID("47631"),
		GroupNumber: "57",
		Name:        "敦賀",
	},
	"0580_60": Observation{
		ID:          ObservationID("0580"),
		GroupNumber: "60",
		Name:        "今津",
	},
	"0581_60": Observation{
		ID:          ObservationID("0581"),
		GroupNumber: "60",
		Name:        "荒川",
	},
	"0582_60": Observation{
		ID:          ObservationID("0582"),
		GroupNumber: "60",
		Name:        "春照",
	},
	"0584_60": Observation{
		ID:          ObservationID("0584"),
		GroupNumber: "60",
		Name:        "近江八幡",
	},
	"0585_60": Observation{
		ID:          ObservationID("0585"),
		GroupNumber: "60",
		Name:        "君ケ畑台地",
	},
	"0586_60": Observation{
		ID:          ObservationID("0586"),
		GroupNumber: "60",
		Name:        "大津",
	},
	"0587_60": Observation{
		ID:          ObservationID("0587"),
		GroupNumber: "60",
		Name:        "土山",
	},
	"0963_60": Observation{
		ID:          ObservationID("0963"),
		GroupNumber: "60",
		Name:        "東近江",
	},
	"0964_60": Observation{
		ID:          ObservationID("0964"),
		GroupNumber: "60",
		Name:        "信楽",
	},
	"0976_60": Observation{
		ID:          ObservationID("0976"),
		GroupNumber: "60",
		Name:        "長浜",
	},
	"1083_60": Observation{
		ID:          ObservationID("1083"),
		GroupNumber: "60",
		Name:        "南小松",
	},
	"1087_60": Observation{
		ID:          ObservationID("1087"),
		GroupNumber: "60",
		Name:        "柳ケ瀬",
	},
	"1361_60": Observation{
		ID:          ObservationID("1361"),
		GroupNumber: "60",
		Name:        "霜ケ原",
	},
	"1524_60": Observation{
		ID:          ObservationID("1524"),
		GroupNumber: "60",
		Name:        "米原",
	},
	"1578_60": Observation{
		ID:          ObservationID("1578"),
		GroupNumber: "60",
		Name:        "朽木平良",
	},
	"47751_60": Observation{
		ID:          ObservationID("47751"),
		GroupNumber: "60",
		Name:        "伊吹山",
	},
	"47761_60": Observation{
		ID:          ObservationID("47761"),
		GroupNumber: "60",
		Name:        "彦根",
	},
	"0588_61": Observation{
		ID:          ObservationID("0588"),
		GroupNumber: "61",
		Name:        "峰山",
	},
	"0589_61": Observation{
		ID:          ObservationID("0589"),
		GroupNumber: "61",
		Name:        "宮津",
	},
	"0591_61": Observation{
		ID:          ObservationID("0591"),
		GroupNumber: "61",
		Name:        "仏坂",
	},
	"0592_61": Observation{
		ID:          ObservationID("0592"),
		GroupNumber: "61",
		Name:        "浅原山",
	},
	"0593_61": Observation{
		ID:          ObservationID("0593"),
		GroupNumber: "61",
		Name:        "福知山",
	},
	"0594_61": Observation{
		ID:          ObservationID("0594"),
		GroupNumber: "61",
		Name:        "胡麻",
	},
	"0595_61": Observation{
		ID:          ObservationID("0595"),
		GroupNumber: "61",
		Name:        "妙高山",
	},
	"0596_61": Observation{
		ID:          ObservationID("0596"),
		GroupNumber: "61",
		Name:        "園部",
	},
	"0598_61": Observation{
		ID:          ObservationID("0598"),
		GroupNumber: "61",
		Name:        "京田辺",
	},
	"0599_61": Observation{
		ID:          ObservationID("0599"),
		GroupNumber: "61",
		Name:        "鷲峰山",
	},
	"0920_61": Observation{
		ID:          ObservationID("0920"),
		GroupNumber: "61",
		Name:        "花背峠",
	},
	"0927_61": Observation{
		ID:          ObservationID("0927"),
		GroupNumber: "61",
		Name:        "知井",
	},
	"0928_61": Observation{
		ID:          ObservationID("0928"),
		GroupNumber: "61",
		Name:        "京北",
	},
	"0965_61": Observation{
		ID:          ObservationID("0965"),
		GroupNumber: "61",
		Name:        "間人",
	},
	"0966_61": Observation{
		ID:          ObservationID("0966"),
		GroupNumber: "61",
		Name:        "本庄",
	},
	"1025_61": Observation{
		ID:          ObservationID("1025"),
		GroupNumber: "61",
		Name:        "長岡京",
	},
	"1160_61": Observation{
		ID:          ObservationID("1160"),
		GroupNumber: "61",
		Name:        "故屋岡",
	},
	"1303_61": Observation{
		ID:          ObservationID("1303"),
		GroupNumber: "61",
		Name:        "美山",
	},
	"1378_61": Observation{
		ID:          ObservationID("1378"),
		GroupNumber: "61",
		Name:        "須知",
	},
	"1379_61": Observation{
		ID:          ObservationID("1379"),
		GroupNumber: "61",
		Name:        "三岳",
	},
	"1380_61": Observation{
		ID:          ObservationID("1380"),
		GroupNumber: "61",
		Name:        "綾部",
	},
	"1383_61": Observation{
		ID:          ObservationID("1383"),
		GroupNumber: "61",
		Name:        "三和",
	},
	"47750_61": Observation{
		ID:          ObservationID("47750"),
		GroupNumber: "61",
		Name:        "舞鶴",
	},
	"47759_61": Observation{
		ID:          ObservationID("47759"),
		GroupNumber: "61",
		Name:        "京都",
	},
	"0600_62": Observation{
		ID:          ObservationID("0600"),
		GroupNumber: "62",
		Name:        "能勢",
	},
	"0601_62": Observation{
		ID:          ObservationID("0601"),
		GroupNumber: "62",
		Name:        "箕面",
	},
	"0602_62": Observation{
		ID:          ObservationID("0602"),
		GroupNumber: "62",
		Name:        "豊中",
	},
	"0604_62": Observation{
		ID:          ObservationID("0604"),
		GroupNumber: "62",
		Name:        "生駒山",
	},
	"0605_62": Observation{
		ID:          ObservationID("0605"),
		GroupNumber: "62",
		Name:        "河内長野",
	},
	"0606_62": Observation{
		ID:          ObservationID("0606"),
		GroupNumber: "62",
		Name:        "熊取",
	},
	"1062_62": Observation{
		ID:          ObservationID("1062"),
		GroupNumber: "62",
		Name:        "堺",
	},
	"1065_62": Observation{
		ID:          ObservationID("1065"),
		GroupNumber: "62",
		Name:        "枚方",
	},
	"1470_62": Observation{
		ID:          ObservationID("1470"),
		GroupNumber: "62",
		Name:        "八尾",
	},
	"1471_62": Observation{
		ID:          ObservationID("1471"),
		GroupNumber: "62",
		Name:        "関空島",
	},
	"1602_62": Observation{
		ID:          ObservationID("1602"),
		GroupNumber: "62",
		Name:        "茨木",
	},
	"47772_62": Observation{
		ID:          ObservationID("47772"),
		GroupNumber: "62",
		Name:        "大阪",
	},
	"0607_63": Observation{
		ID:          ObservationID("0607"),
		GroupNumber: "63",
		Name:        "香住",
	},
	"0608_63": Observation{
		ID:          ObservationID("0608"),
		GroupNumber: "63",
		Name:        "温泉",
	},
	"0609_63": Observation{
		ID:          ObservationID("0609"),
		GroupNumber: "63",
		Name:        "三川山",
	},
	"0611_63": Observation{
		ID:          ObservationID("0611"),
		GroupNumber: "63",
		Name:        "八鹿",
	},
	"0612_63": Observation{
		ID:          ObservationID("0612"),
		GroupNumber: "63",
		Name:        "和田山",
	},
	"0613_63": Observation{
		ID:          ObservationID("0613"),
		GroupNumber: "63",
		Name:        "生野",
	},
	"0614_63": Observation{
		ID:          ObservationID("0614"),
		GroupNumber: "63",
		Name:        "柏原",
	},
	"0615_63": Observation{
		ID:          ObservationID("0615"),
		GroupNumber: "63",
		Name:        "一宮",
	},
	"0616_63": Observation{
		ID:          ObservationID("0616"),
		GroupNumber: "63",
		Name:        "笠形山",
	},
	"0617_63": Observation{
		ID:          ObservationID("0617"),
		GroupNumber: "63",
		Name:        "佐用",
	},
	"0618_63": Observation{
		ID:          ObservationID("0618"),
		GroupNumber: "63",
		Name:        "後川",
	},
	"0619_63": Observation{
		ID:          ObservationID("0619"),
		GroupNumber: "63",
		Name:        "上郡",
	},
	"0620_63": Observation{
		ID:          ObservationID("0620"),
		GroupNumber: "63",
		Name:        "的場山",
	},
	"0622_63": Observation{
		ID:          ObservationID("0622"),
		GroupNumber: "63",
		Name:        "名塩",
	},
	"0623_63": Observation{
		ID:          ObservationID("0623"),
		GroupNumber: "63",
		Name:        "六甲山",
	},
	"0624_63": Observation{
		ID:          ObservationID("0624"),
		GroupNumber: "63",
		Name:        "家島",
	},
	"0625_63": Observation{
		ID:          ObservationID("0625"),
		GroupNumber: "63",
		Name:        "明石",
	},
	"0628_63": Observation{
		ID:          ObservationID("0628"),
		GroupNumber: "63",
		Name:        "福良",
	},
	"0956_63": Observation{
		ID:          ObservationID("0956"),
		GroupNumber: "63",
		Name:        "西脇",
	},
	"0967_63": Observation{
		ID:          ObservationID("0967"),
		GroupNumber: "63",
		Name:        "大屋",
	},
	"0968_63": Observation{
		ID:          ObservationID("0968"),
		GroupNumber: "63",
		Name:        "福崎",
	},
	"0969_63": Observation{
		ID:          ObservationID("0969"),
		GroupNumber: "63",
		Name:        "三田",
	},
	"0970_63": Observation{
		ID:          ObservationID("0970"),
		GroupNumber: "63",
		Name:        "郡家",
	},
	"0977_63": Observation{
		ID:          ObservationID("0977"),
		GroupNumber: "63",
		Name:        "兎和野高原",
	},
	"1227_63": Observation{
		ID:          ObservationID("1227"),
		GroupNumber: "63",
		Name:        "三木",
	},
	"1337_63": Observation{
		ID:          ObservationID("1337"),
		GroupNumber: "63",
		Name:        "南淡",
	},
	"1448_63": Observation{
		ID:          ObservationID("1448"),
		GroupNumber: "63",
		Name:        "淡路町",
	},
	"1449_63": Observation{
		ID:          ObservationID("1449"),
		GroupNumber: "63",
		Name:        "芦屋",
	},
	"1450_63": Observation{
		ID:          ObservationID("1450"),
		GroupNumber: "63",
		Name:        "神戸塩屋",
	},
	"1587_63": Observation{
		ID:          ObservationID("1587"),
		GroupNumber: "63",
		Name:        "神戸空港",
	},
	"1588_63": Observation{
		ID:          ObservationID("1588"),
		GroupNumber: "63",
		Name:        "西宮",
	},
	"47747_63": Observation{
		ID:          ObservationID("47747"),
		GroupNumber: "63",
		Name:        "豊岡",
	},
	"47769_63": Observation{
		ID:          ObservationID("47769"),
		GroupNumber: "63",
		Name:        "姫路",
	},
	"47770_63": Observation{
		ID:          ObservationID("47770"),
		GroupNumber: "63",
		Name:        "神戸",
	},
	"47776_63": Observation{
		ID:          ObservationID("47776"),
		GroupNumber: "63",
		Name:        "洲本",
	},
	"0630_64": Observation{
		ID:          ObservationID("0630"),
		GroupNumber: "64",
		Name:        "針",
	},
	"0631_64": Observation{
		ID:          ObservationID("0631"),
		GroupNumber: "64",
		Name:        "田原本",
	},
	"0632_64": Observation{
		ID:          ObservationID("0632"),
		GroupNumber: "64",
		Name:        "曽爾",
	},
	"0633_64": Observation{
		ID:          ObservationID("0633"),
		GroupNumber: "64",
		Name:        "大宇陀",
	},
	"0634_64": Observation{
		ID:          ObservationID("0634"),
		GroupNumber: "64",
		Name:        "高見山",
	},
	"0635_64": Observation{
		ID:          ObservationID("0635"),
		GroupNumber: "64",
		Name:        "五條",
	},
	"0636_64": Observation{
		ID:          ObservationID("0636"),
		GroupNumber: "64",
		Name:        "天辻",
	},
	"0637_64": Observation{
		ID:          ObservationID("0637"),
		GroupNumber: "64",
		Name:        "山上ケ岳",
	},
	"0638_64": Observation{
		ID:          ObservationID("0638"),
		GroupNumber: "64",
		Name:        "荒神岳",
	},
	"0639_64": Observation{
		ID:          ObservationID("0639"),
		GroupNumber: "64",
		Name:        "日出岳",
	},
	"0640_64": Observation{
		ID:          ObservationID("0640"),
		GroupNumber: "64",
		Name:        "上野地",
	},
	"0641_64": Observation{
		ID:          ObservationID("0641"),
		GroupNumber: "64",
		Name:        "玉置山",
	},
	"0957_64": Observation{
		ID:          ObservationID("0957"),
		GroupNumber: "64",
		Name:        "上北山",
	},
	"1228_64": Observation{
		ID:          ObservationID("1228"),
		GroupNumber: "64",
		Name:        "風屋",
	},
	"1366_64": Observation{
		ID:          ObservationID("1366"),
		GroupNumber: "64",
		Name:        "高見",
	},
	"1376_64": Observation{
		ID:          ObservationID("1376"),
		GroupNumber: "64",
		Name:        "葛城",
	},
	"1377_64": Observation{
		ID:          ObservationID("1377"),
		GroupNumber: "64",
		Name:        "壷坂",
	},
	"1439_64": Observation{
		ID:          ObservationID("1439"),
		GroupNumber: "64",
		Name:        "吉野",
	},
	"1619_64": Observation{
		ID:          ObservationID("1619"),
		GroupNumber: "64",
		Name:        "天川",
	},
	"47780_64": Observation{
		ID:          ObservationID("47780"),
		GroupNumber: "64",
		Name:        "奈良",
	},
	"0642_65": Observation{
		ID:          ObservationID("0642"),
		GroupNumber: "65",
		Name:        "葛城山",
	},
	"0644_65": Observation{
		ID:          ObservationID("0644"),
		GroupNumber: "65",
		Name:        "岩出",
	},
	"0645_65": Observation{
		ID:          ObservationID("0645"),
		GroupNumber: "65",
		Name:        "高野山",
	},
	"0646_65": Observation{
		ID:          ObservationID("0646"),
		GroupNumber: "65",
		Name:        "護摩壇山",
	},
	"0647_65": Observation{
		ID:          ObservationID("0647"),
		GroupNumber: "65",
		Name:        "御坊",
	},
	"0648_65": Observation{
		ID:          ObservationID("0648"),
		GroupNumber: "65",
		Name:        "虎ケ峰峠",
	},
	"0649_65": Observation{
		ID:          ObservationID("0649"),
		GroupNumber: "65",
		Name:        "新宮",
	},
	"0650_65": Observation{
		ID:          ObservationID("0650"),
		GroupNumber: "65",
		Name:        "白浜",
	},
	"0651_65": Observation{
		ID:          ObservationID("0651"),
		GroupNumber: "65",
		Name:        "日置川",
	},
	"0971_65": Observation{
		ID:          ObservationID("0971"),
		GroupNumber: "65",
		Name:        "栗栖川",
	},
	"0978_65": Observation{
		ID:          ObservationID("0978"),
		GroupNumber: "65",
		Name:        "湯浅",
	},
	"1063_65": Observation{
		ID:          ObservationID("1063"),
		GroupNumber: "65",
		Name:        "清水",
	},
	"1074_65": Observation{
		ID:          ObservationID("1074"),
		GroupNumber: "65",
		Name:        "本宮",
	},
	"1089_65": Observation{
		ID:          ObservationID("1089"),
		GroupNumber: "65",
		Name:        "龍神",
	},
	"1172_65": Observation{
		ID:          ObservationID("1172"),
		GroupNumber: "65",
		Name:        "色川",
	},
	"1342_65": Observation{
		ID:          ObservationID("1342"),
		GroupNumber: "65",
		Name:        "かつらぎ",
	},
	"1347_65": Observation{
		ID:          ObservationID("1347"),
		GroupNumber: "65",
		Name:        "西川",
	},
	"1457_65": Observation{
		ID:          ObservationID("1457"),
		GroupNumber: "65",
		Name:        "友ケ島",
	},
	"1485_65": Observation{
		ID:          ObservationID("1485"),
		GroupNumber: "65",
		Name:        "川辺",
	},
	"1589_65": Observation{
		ID:          ObservationID("1589"),
		GroupNumber: "65",
		Name:        "南紀白浜",
	},
	"47777_65": Observation{
		ID:          ObservationID("47777"),
		GroupNumber: "65",
		Name:        "和歌山",
	},
	"47778_65": Observation{
		ID:          ObservationID("47778"),
		GroupNumber: "65",
		Name:        "潮岬",
	},
	"0653_66": Observation{
		ID:          ObservationID("0653"),
		GroupNumber: "66",
		Name:        "大空山",
	},
	"0654_66": Observation{
		ID:          ObservationID("0654"),
		GroupNumber: "66",
		Name:        "那岐山",
	},
	"0655_66": Observation{
		ID:          ObservationID("0655"),
		GroupNumber: "66",
		Name:        "今岡",
	},
	"0656_66": Observation{
		ID:          ObservationID("0656"),
		GroupNumber: "66",
		Name:        "久世",
	},
	"0658_66": Observation{
		ID:          ObservationID("0658"),
		GroupNumber: "66",
		Name:        "新見",
	},
	"0659_66": Observation{
		ID:          ObservationID("0659"),
		GroupNumber: "66",
		Name:        "天子山",
	},
	"0660_66": Observation{
		ID:          ObservationID("0660"),
		GroupNumber: "66",
		Name:        "赤磐",
	},
	"0661_66": Observation{
		ID:          ObservationID("0661"),
		GroupNumber: "66",
		Name:        "陣山",
	},
	"0662_66": Observation{
		ID:          ObservationID("0662"),
		GroupNumber: "66",
		Name:        "大平山",
	},
	"0663_66": Observation{
		ID:          ObservationID("0663"),
		GroupNumber: "66",
		Name:        "福渡",
	},
	"0664_66": Observation{
		ID:          ObservationID("0664"),
		GroupNumber: "66",
		Name:        "和気",
	},
	"0665_66": Observation{
		ID:          ObservationID("0665"),
		GroupNumber: "66",
		Name:        "佐屋",
	},
	"0666_66": Observation{
		ID:          ObservationID("0666"),
		GroupNumber: "66",
		Name:        "矢掛",
	},
	"0668_66": Observation{
		ID:          ObservationID("0668"),
		GroupNumber: "66",
		Name:        "虫明",
	},
	"0669_66": Observation{
		ID:          ObservationID("0669"),
		GroupNumber: "66",
		Name:        "倉敷",
	},
	"0670_66": Observation{
		ID:          ObservationID("0670"),
		GroupNumber: "66",
		Name:        "玉野",
	},
	"0919_66": Observation{
		ID:          ObservationID("0919"),
		GroupNumber: "66",
		Name:        "笠岡",
	},
	"0929_66": Observation{
		ID:          ObservationID("0929"),
		GroupNumber: "66",
		Name:        "下呰部",
	},
	"0930_66": Observation{
		ID:          ObservationID("0930"),
		GroupNumber: "66",
		Name:        "高梁",
	},
	"1140_66": Observation{
		ID:          ObservationID("1140"),
		GroupNumber: "66",
		Name:        "奈義",
	},
	"1328_66": Observation{
		ID:          ObservationID("1328"),
		GroupNumber: "66",
		Name:        "上長田",
	},
	"1330_66": Observation{
		ID:          ObservationID("1330"),
		GroupNumber: "66",
		Name:        "千屋",
	},
	"1371_66": Observation{
		ID:          ObservationID("1371"),
		GroupNumber: "66",
		Name:        "恩原",
	},
	"1531_66": Observation{
		ID:          ObservationID("1531"),
		GroupNumber: "66",
		Name:        "日応寺",
	},
	"1584_66": Observation{
		ID:          ObservationID("1584"),
		GroupNumber: "66",
		Name:        "富",
	},
	"1585_66": Observation{
		ID:          ObservationID("1585"),
		GroupNumber: "66",
		Name:        "吉備中央",
	},
	"1618_66": Observation{
		ID:          ObservationID("1618"),
		GroupNumber: "66",
		Name:        "旭西",
	},
	"47756_66": Observation{
		ID:          ObservationID("47756"),
		GroupNumber: "66",
		Name:        "津山",
	},
	"47768_66": Observation{
		ID:          ObservationID("47768"),
		GroupNumber: "66",
		Name:        "岡山",
	},
	"0671_67": Observation{
		ID:          ObservationID("0671"),
		GroupNumber: "67",
		Name:        "道後山",
	},
	"0672_67": Observation{
		ID:          ObservationID("0672"),
		GroupNumber: "67",
		Name:        "安田山",
	},
	"0673_67": Observation{
		ID:          ObservationID("0673"),
		GroupNumber: "67",
		Name:        "犬伏山",
	},
	"0674_67": Observation{
		ID:          ObservationID("0674"),
		GroupNumber: "67",
		Name:        "三次",
	},
	"0675_67": Observation{
		ID:          ObservationID("0675"),
		GroupNumber: "67",
		Name:        "庄原",
	},
	"0676_67": Observation{
		ID:          ObservationID("0676"),
		GroupNumber: "67",
		Name:        "大朝",
	},
	"0677_67": Observation{
		ID:          ObservationID("0677"),
		GroupNumber: "67",
		Name:        "加計",
	},
	"0678_67": Observation{
		ID:          ObservationID("0678"),
		GroupNumber: "67",
		Name:        "海見山",
	},
	"0679_67": Observation{
		ID:          ObservationID("0679"),
		GroupNumber: "67",
		Name:        "上下",
	},
	"0680_67": Observation{
		ID:          ObservationID("0680"),
		GroupNumber: "67",
		Name:        "内黒山",
	},
	"0681_67": Observation{
		ID:          ObservationID("0681"),
		GroupNumber: "67",
		Name:        "世羅",
	},
	"0682_67": Observation{
		ID:          ObservationID("0682"),
		GroupNumber: "67",
		Name:        "恵下谷山",
	},
	"0683_67": Observation{
		ID:          ObservationID("0683"),
		GroupNumber: "67",
		Name:        "東広島",
	},
	"0686_67": Observation{
		ID:          ObservationID("0686"),
		GroupNumber: "67",
		Name:        "竹原",
	},
	"0687_67": Observation{
		ID:          ObservationID("0687"),
		GroupNumber: "67",
		Name:        "生口島",
	},
	"0688_67": Observation{
		ID:          ObservationID("0688"),
		GroupNumber: "67",
		Name:        "大竹",
	},
	"0690_67": Observation{
		ID:          ObservationID("0690"),
		GroupNumber: "67",
		Name:        "倉橋",
	},
	"0947_67": Observation{
		ID:          ObservationID("0947"),
		GroupNumber: "67",
		Name:        "高野",
	},
	"0948_67": Observation{
		ID:          ObservationID("0948"),
		GroupNumber: "67",
		Name:        "東城",
	},
	"0949_67": Observation{
		ID:          ObservationID("0949"),
		GroupNumber: "67",
		Name:        "甲田",
	},
	"0950_67": Observation{
		ID:          ObservationID("0950"),
		GroupNumber: "67",
		Name:        "三入",
	},
	"0951_67": Observation{
		ID:          ObservationID("0951"),
		GroupNumber: "67",
		Name:        "府中",
	},
	"0952_67": Observation{
		ID:          ObservationID("0952"),
		GroupNumber: "67",
		Name:        "河内",
	},
	"0953_67": Observation{
		ID:          ObservationID("0953"),
		GroupNumber: "67",
		Name:        "久比",
	},
	"1163_67": Observation{
		ID:          ObservationID("1163"),
		GroupNumber: "67",
		Name:        "志和",
	},
	"1257_67": Observation{
		ID:          ObservationID("1257"),
		GroupNumber: "67",
		Name:        "王泊",
	},
	"1262_67": Observation{
		ID:          ObservationID("1262"),
		GroupNumber: "67",
		Name:        "八幡",
	},
	"1325_67": Observation{
		ID:          ObservationID("1325"),
		GroupNumber: "67",
		Name:        "油木",
	},
	"1326_67": Observation{
		ID:          ObservationID("1326"),
		GroupNumber: "67",
		Name:        "廿日市津田",
	},
	"1472_67": Observation{
		ID:          ObservationID("1472"),
		GroupNumber: "67",
		Name:        "本郷",
	},
	"1575_67": Observation{
		ID:          ObservationID("1575"),
		GroupNumber: "67",
		Name:        "佐伯湯来",
	},
	"1579_67": Observation{
		ID:          ObservationID("1579"),
		GroupNumber: "67",
		Name:        "都志見",
	},
	"1582_67": Observation{
		ID:          ObservationID("1582"),
		GroupNumber: "67",
		Name:        "君田",
	},
	"1583_67": Observation{
		ID:          ObservationID("1583"),
		GroupNumber: "67",
		Name:        "美土里",
	},
	"1603_67": Observation{
		ID:          ObservationID("1603"),
		GroupNumber: "67",
		Name:        "安宿",
	},
	"1606_67": Observation{
		ID:          ObservationID("1606"),
		GroupNumber: "67",
		Name:        "呉市蒲刈",
	},
	"47765_67": Observation{
		ID:          ObservationID("47765"),
		GroupNumber: "67",
		Name:        "広島",
	},
	"47766_67": Observation{
		ID:          ObservationID("47766"),
		GroupNumber: "67",
		Name:        "呉",
	},
	"47767_67": Observation{
		ID:          ObservationID("47767"),
		GroupNumber: "67",
		Name:        "福山",
	},
	"0692_68": Observation{
		ID:          ObservationID("0692"),
		GroupNumber: "68",
		Name:        "西ノ島",
	},
	"0693_68": Observation{
		ID:          ObservationID("0693"),
		GroupNumber: "68",
		Name:        "鹿島",
	},
	"0694_68": Observation{
		ID:          ObservationID("0694"),
		GroupNumber: "68",
		Name:        "斐川",
	},
	"0696_68": Observation{
		ID:          ObservationID("0696"),
		GroupNumber: "68",
		Name:        "大東",
	},
	"0697_68": Observation{
		ID:          ObservationID("0697"),
		GroupNumber: "68",
		Name:        "大田",
	},
	"0698_68": Observation{
		ID:          ObservationID("0698"),
		GroupNumber: "68",
		Name:        "掛合",
	},
	"0699_68": Observation{
		ID:          ObservationID("0699"),
		GroupNumber: "68",
		Name:        "横田",
	},
	"0700_68": Observation{
		ID:          ObservationID("0700"),
		GroupNumber: "68",
		Name:        "吾妻山",
	},
	"0701_68": Observation{
		ID:          ObservationID("0701"),
		GroupNumber: "68",
		Name:        "川本",
	},
	"0703_68": Observation{
		ID:          ObservationID("0703"),
		GroupNumber: "68",
		Name:        "三隅",
	},
	"0704_68": Observation{
		ID:          ObservationID("0704"),
		GroupNumber: "68",
		Name:        "益田",
	},
	"0705_68": Observation{
		ID:          ObservationID("0705"),
		GroupNumber: "68",
		Name:        "聖山",
	},
	"0706_68": Observation{
		ID:          ObservationID("0706"),
		GroupNumber: "68",
		Name:        "十種峯",
	},
	"0707_68": Observation{
		ID:          ObservationID("0707"),
		GroupNumber: "68",
		Name:        "津和野",
	},
	"0954_68": Observation{
		ID:          ObservationID("0954"),
		GroupNumber: "68",
		Name:        "佐田",
	},
	"1093_68": Observation{
		ID:          ObservationID("1093"),
		GroupNumber: "68",
		Name:        "桜江",
	},
	"1094_68": Observation{
		ID:          ObservationID("1094"),
		GroupNumber: "68",
		Name:        "瑞穂",
	},
	"1095_68": Observation{
		ID:          ObservationID("1095"),
		GroupNumber: "68",
		Name:        "伯太",
	},
	"1175_68": Observation{
		ID:          ObservationID("1175"),
		GroupNumber: "68",
		Name:        "波佐",
	},
	"1176_68": Observation{
		ID:          ObservationID("1176"),
		GroupNumber: "68",
		Name:        "匹見",
	},
	"1307_68": Observation{
		ID:          ObservationID("1307"),
		GroupNumber: "68",
		Name:        "海士",
	},
	"1310_68": Observation{
		ID:          ObservationID("1310"),
		GroupNumber: "68",
		Name:        "出雲",
	},
	"1320_68": Observation{
		ID:          ObservationID("1320"),
		GroupNumber: "68",
		Name:        "赤名",
	},
	"1321_68": Observation{
		ID:          ObservationID("1321"),
		GroupNumber: "68",
		Name:        "弥栄",
	},
	"1322_68": Observation{
		ID:          ObservationID("1322"),
		GroupNumber: "68",
		Name:        "六日市",
	},
	"1336_68": Observation{
		ID:          ObservationID("1336"),
		GroupNumber: "68",
		Name:        "福光",
	},
	"1362_68": Observation{
		ID:          ObservationID("1362"),
		GroupNumber: "68",
		Name:        "原山",
	},
	"1473_68": Observation{
		ID:          ObservationID("1473"),
		GroupNumber: "68",
		Name:        "西郷岬",
	},
	"1474_68": Observation{
		ID:          ObservationID("1474"),
		GroupNumber: "68",
		Name:        "高津",
	},
	"1592_68": Observation{
		ID:          ObservationID("1592"),
		GroupNumber: "68",
		Name:        "吉賀",
	},
	"47740_68": Observation{
		ID:          ObservationID("47740"),
		GroupNumber: "68",
		Name:        "西郷",
	},
	"47741_68": Observation{
		ID:          ObservationID("47741"),
		GroupNumber: "68",
		Name:        "松江",
	},
	"47755_68": Observation{
		ID:          ObservationID("47755"),
		GroupNumber: "68",
		Name:        "浜田",
	},
	"0709_69": Observation{
		ID:          ObservationID("0709"),
		GroupNumber: "69",
		Name:        "青谷",
	},
	"0711_69": Observation{
		ID:          ObservationID("0711"),
		GroupNumber: "69",
		Name:        "岩井",
	},
	"0713_69": Observation{
		ID:          ObservationID("0713"),
		GroupNumber: "69",
		Name:        "倉吉",
	},
	"0714_69": Observation{
		ID:          ObservationID("0714"),
		GroupNumber: "69",
		Name:        "菅俵原",
	},
	"0715_69": Observation{
		ID:          ObservationID("0715"),
		GroupNumber: "69",
		Name:        "霊石山",
	},
	"0716_69": Observation{
		ID:          ObservationID("0716"),
		GroupNumber: "69",
		Name:        "大山",
	},
	"0717_69": Observation{
		ID:          ObservationID("0717"),
		GroupNumber: "69",
		Name:        "古峠山",
	},
	"0718_69": Observation{
		ID:          ObservationID("0718"),
		GroupNumber: "69",
		Name:        "智頭",
	},
	"0938_69": Observation{
		ID:          ObservationID("0938"),
		GroupNumber: "69",
		Name:        "日南",
	},
	"1177_69": Observation{
		ID:          ObservationID("1177"),
		GroupNumber: "69",
		Name:        "関金",
	},
	"1178_69": Observation{
		ID:          ObservationID("1178"),
		GroupNumber: "69",
		Name:        "若桜",
	},
	"1231_69": Observation{
		ID:          ObservationID("1231"),
		GroupNumber: "69",
		Name:        "塩津",
	},
	"1304_69": Observation{
		ID:          ObservationID("1304"),
		GroupNumber: "69",
		Name:        "茶屋",
	},
	"1385_69": Observation{
		ID:          ObservationID("1385"),
		GroupNumber: "69",
		Name:        "佐治",
	},
	"1389_69": Observation{
		ID:          ObservationID("1389"),
		GroupNumber: "69",
		Name:        "江尾",
	},
	"1391_69": Observation{
		ID:          ObservationID("1391"),
		GroupNumber: "69",
		Name:        "鹿野",
	},
	"1519_69": Observation{
		ID:          ObservationID("1519"),
		GroupNumber: "69",
		Name:        "湖山",
	},
	"47742_69": Observation{
		ID:          ObservationID("47742"),
		GroupNumber: "69",
		Name:        "境",
	},
	"47744_69": Observation{
		ID:          ObservationID("47744"),
		GroupNumber: "69",
		Name:        "米子",
	},
	"47746_69": Observation{
		ID:          ObservationID("47746"),
		GroupNumber: "69",
		Name:        "鳥取",
	},
	"0719_71": Observation{
		ID:          ObservationID("0719"),
		GroupNumber: "71",
		Name:        "大山寺",
	},
	"0720_71": Observation{
		ID:          ObservationID("0720"),
		GroupNumber: "71",
		Name:        "池田",
	},
	"0722_71": Observation{
		ID:          ObservationID("0722"),
		GroupNumber: "71",
		Name:        "江田山",
	},
	"0723_71": Observation{
		ID:          ObservationID("0723"),
		GroupNumber: "71",
		Name:        "太竜寺山",
	},
	"0724_71": Observation{
		ID:          ObservationID("0724"),
		GroupNumber: "71",
		Name:        "日和佐",
	},
	"0725_71": Observation{
		ID:          ObservationID("0725"),
		GroupNumber: "71",
		Name:        "宍喰",
	},
	"0972_71": Observation{
		ID:          ObservationID("0972"),
		GroupNumber: "71",
		Name:        "穴吹",
	},
	"1098_71": Observation{
		ID:          ObservationID("1098"),
		GroupNumber: "71",
		Name:        "福原旭",
	},
	"1242_71": Observation{
		ID:          ObservationID("1242"),
		GroupNumber: "71",
		Name:        "蒲生田",
	},
	"1252_71": Observation{
		ID:          ObservationID("1252"),
		GroupNumber: "71",
		Name:        "京上",
	},
	"1327_71": Observation{
		ID:          ObservationID("1327"),
		GroupNumber: "71",
		Name:        "半田",
	},
	"1352_71": Observation{
		ID:          ObservationID("1352"),
		GroupNumber: "71",
		Name:        "木頭",
	},
	"1358_71": Observation{
		ID:          ObservationID("1358"),
		GroupNumber: "71",
		Name:        "旭丸",
	},
	"1604_71": Observation{
		ID:          ObservationID("1604"),
		GroupNumber: "71",
		Name:        "海陽",
	},
	"47894_71": Observation{
		ID:          ObservationID("47894"),
		GroupNumber: "71",
		Name:        "剣山",
	},
	"47895_71": Observation{
		ID:          ObservationID("47895"),
		GroupNumber: "71",
		Name:        "徳島",
	},
	"0726_72": Observation{
		ID:          ObservationID("0726"),
		GroupNumber: "72",
		Name:        "内海",
	},
	"0729_72": Observation{
		ID:          ObservationID("0729"),
		GroupNumber: "72",
		Name:        "滝宮",
	},
	"0730_72": Observation{
		ID:          ObservationID("0730"),
		GroupNumber: "72",
		Name:        "引田",
	},
	"0731_72": Observation{
		ID:          ObservationID("0731"),
		GroupNumber: "72",
		Name:        "竜王山",
	},
	"1239_72": Observation{
		ID:          ObservationID("1239"),
		GroupNumber: "72",
		Name:        "財田",
	},
	"1475_72": Observation{
		ID:          ObservationID("1475"),
		GroupNumber: "72",
		Name:        "香南",
	},
	"47890_72": Observation{
		ID:          ObservationID("47890"),
		GroupNumber: "72",
		Name:        "多度津",
	},
	"47891_72": Observation{
		ID:          ObservationID("47891"),
		GroupNumber: "72",
		Name:        "高松",
	},
	"0732_73": Observation{
		ID:          ObservationID("0732"),
		GroupNumber: "73",
		Name:        "大三島",
	},
	"0733_73": Observation{
		ID:          ObservationID("0733"),
		GroupNumber: "73",
		Name:        "玉川",
	},
	"0734_73": Observation{
		ID:          ObservationID("0734"),
		GroupNumber: "73",
		Name:        "新居浜",
	},
	"0735_73": Observation{
		ID:          ObservationID("0735"),
		GroupNumber: "73",
		Name:        "四国中央",
	},
	"0737_73": Observation{
		ID:          ObservationID("0737"),
		GroupNumber: "73",
		Name:        "上林",
	},
	"0738_73": Observation{
		ID:          ObservationID("0738"),
		GroupNumber: "73",
		Name:        "成就社",
	},
	"0739_73": Observation{
		ID:          ObservationID("0739"),
		GroupNumber: "73",
		Name:        "長浜",
	},
	"0740_73": Observation{
		ID:          ObservationID("0740"),
		GroupNumber: "73",
		Name:        "中山",
	},
	"0741_73": Observation{
		ID:          ObservationID("0741"),
		GroupNumber: "73",
		Name:        "久万",
	},
	"0742_73": Observation{
		ID:          ObservationID("0742"),
		GroupNumber: "73",
		Name:        "獅子越峠",
	},
	"0743_73": Observation{
		ID:          ObservationID("0743"),
		GroupNumber: "73",
		Name:        "八幡浜",
	},
	"0745_73": Observation{
		ID:          ObservationID("0745"),
		GroupNumber: "73",
		Name:        "御荘",
	},
	"0958_73": Observation{
		ID:          ObservationID("0958"),
		GroupNumber: "73",
		Name:        "西条",
	},
	"0959_73": Observation{
		ID:          ObservationID("0959"),
		GroupNumber: "73",
		Name:        "大洲",
	},
	"0960_73": Observation{
		ID:          ObservationID("0960"),
		GroupNumber: "73",
		Name:        "三崎",
	},
	"0961_73": Observation{
		ID:          ObservationID("0961"),
		GroupNumber: "73",
		Name:        "宇和",
	},
	"1077_73": Observation{
		ID:          ObservationID("1077"),
		GroupNumber: "73",
		Name:        "今治",
	},
	"1260_73": Observation{
		ID:          ObservationID("1260"),
		GroupNumber: "73",
		Name:        "富郷",
	},
	"1350_73": Observation{
		ID:          ObservationID("1350"),
		GroupNumber: "73",
		Name:        "近永",
	},
	"1456_73": Observation{
		ID:          ObservationID("1456"),
		GroupNumber: "73",
		Name:        "瀬戸",
	},
	"1521_73": Observation{
		ID:          ObservationID("1521"),
		GroupNumber: "73",
		Name:        "松山南吉田",
	},
	"47887_73": Observation{
		ID:          ObservationID("47887"),
		GroupNumber: "73",
		Name:        "松山",
	},
	"47892_73": Observation{
		ID:          ObservationID("47892"),
		GroupNumber: "73",
		Name:        "宇和島",
	},
	"0746_74": Observation{
		ID:          ObservationID("0746"),
		GroupNumber: "74",
		Name:        "本山",
	},
	"0747_74": Observation{
		ID:          ObservationID("0747"),
		GroupNumber: "74",
		Name:        "繁藤",
	},
	"0748_74": Observation{
		ID:          ObservationID("0748"),
		GroupNumber: "74",
		Name:        "佐川",
	},
	"0749_74": Observation{
		ID:          ObservationID("0749"),
		GroupNumber: "74",
		Name:        "成山",
	},
	"0751_74": Observation{
		ID:          ObservationID("0751"),
		GroupNumber: "74",
		Name:        "南国",
	},
	"0752_74": Observation{
		ID:          ObservationID("0752"),
		GroupNumber: "74",
		Name:        "芸西",
	},
	"0753_74": Observation{
		ID:          ObservationID("0753"),
		GroupNumber: "74",
		Name:        "鳥形山",
	},
	"0754_74": Observation{
		ID:          ObservationID("0754"),
		GroupNumber: "74",
		Name:        "船戸",
	},
	"0755_74": Observation{
		ID:          ObservationID("0755"),
		GroupNumber: "74",
		Name:        "安芸",
	},
	"0756_74": Observation{
		ID:          ObservationID("0756"),
		GroupNumber: "74",
		Name:        "須崎",
	},
	"0757_74": Observation{
		ID:          ObservationID("0757"),
		GroupNumber: "74",
		Name:        "佐喜浜",
	},
	"0758_74": Observation{
		ID:          ObservationID("0758"),
		GroupNumber: "74",
		Name:        "窪川",
	},
	"0760_74": Observation{
		ID:          ObservationID("0760"),
		GroupNumber: "74",
		Name:        "堂ケ森",
	},
	"0761_74": Observation{
		ID:          ObservationID("0761"),
		GroupNumber: "74",
		Name:        "大正",
	},
	"0763_74": Observation{
		ID:          ObservationID("0763"),
		GroupNumber: "74",
		Name:        "中村",
	},
	"0921_74": Observation{
		ID:          ObservationID("0921"),
		GroupNumber: "74",
		Name:        "梼原",
	},
	"0955_74": Observation{
		ID:          ObservationID("0955"),
		GroupNumber: "74",
		Name:        "田野",
	},
	"1179_74": Observation{
		ID:          ObservationID("1179"),
		GroupNumber: "74",
		Name:        "池川",
	},
	"1229_74": Observation{
		ID:          ObservationID("1229"),
		GroupNumber: "74",
		Name:        "大栃",
	},
	"1233_74": Observation{
		ID:          ObservationID("1233"),
		GroupNumber: "74",
		Name:        "佐賀",
	},
	"1235_74": Observation{
		ID:          ObservationID("1235"),
		GroupNumber: "74",
		Name:        "江川崎",
	},
	"1249_74": Observation{
		ID:          ObservationID("1249"),
		GroupNumber: "74",
		Name:        "後免",
	},
	"1281_74": Observation{
		ID:          ObservationID("1281"),
		GroupNumber: "74",
		Name:        "魚梁瀬",
	},
	"1351_74": Observation{
		ID:          ObservationID("1351"),
		GroupNumber: "74",
		Name:        "本川",
	},
	"1476_74": Observation{
		ID:          ObservationID("1476"),
		GroupNumber: "74",
		Name:        "南国日章",
	},
	"1556_74": Observation{
		ID:          ObservationID("1556"),
		GroupNumber: "74",
		Name:        "三崎",
	},
	"47893_74": Observation{
		ID:          ObservationID("47893"),
		GroupNumber: "74",
		Name:        "高知",
	},
	"47897_74": Observation{
		ID:          ObservationID("47897"),
		GroupNumber: "74",
		Name:        "宿毛",
	},
	"47898_74": Observation{
		ID:          ObservationID("47898"),
		GroupNumber: "74",
		Name:        "清水",
	},
	"47899_74": Observation{
		ID:          ObservationID("47899"),
		GroupNumber: "74",
		Name:        "室戸岬",
	},
	"0765_81": Observation{
		ID:          ObservationID("0765"),
		GroupNumber: "81",
		Name:        "須佐",
	},
	"0767_81": Observation{
		ID:          ObservationID("0767"),
		GroupNumber: "81",
		Name:        "徳佐",
	},
	"0768_81": Observation{
		ID:          ObservationID("0768"),
		GroupNumber: "81",
		Name:        "鍋提峠",
	},
	"0769_81": Observation{
		ID:          ObservationID("0769"),
		GroupNumber: "81",
		Name:        "秋吉台",
	},
	"0770_81": Observation{
		ID:          ObservationID("0770"),
		GroupNumber: "81",
		Name:        "瀬戸原",
	},
	"0771_81": Observation{
		ID:          ObservationID("0771"),
		GroupNumber: "81",
		Name:        "広瀬",
	},
	"0772_81": Observation{
		ID:          ObservationID("0772"),
		GroupNumber: "81",
		Name:        "豊田",
	},
	"0773_81": Observation{
		ID:          ObservationID("0773"),
		GroupNumber: "81",
		Name:        "桜山",
	},
	"0775_81": Observation{
		ID:          ObservationID("0775"),
		GroupNumber: "81",
		Name:        "防府",
	},
	"0776_81": Observation{
		ID:          ObservationID("0776"),
		GroupNumber: "81",
		Name:        "下松",
	},
	"0778_81": Observation{
		ID:          ObservationID("0778"),
		GroupNumber: "81",
		Name:        "宇部",
	},
	"0779_81": Observation{
		ID:          ObservationID("0779"),
		GroupNumber: "81",
		Name:        "安下庄",
	},
	"0913_81": Observation{
		ID:          ObservationID("0913"),
		GroupNumber: "81",
		Name:        "玖珂",
	},
	"0939_81": Observation{
		ID:          ObservationID("0939"),
		GroupNumber: "81",
		Name:        "油谷",
	},
	"0940_81": Observation{
		ID:          ObservationID("0940"),
		GroupNumber: "81",
		Name:        "岩国",
	},
	"0941_81": Observation{
		ID:          ObservationID("0941"),
		GroupNumber: "81",
		Name:        "周東",
	},
	"0942_81": Observation{
		ID:          ObservationID("0942"),
		GroupNumber: "81",
		Name:        "柳井",
	},
	"1100_81": Observation{
		ID:          ObservationID("1100"),
		GroupNumber: "81",
		Name:        "羅漢山",
	},
	"1101_81": Observation{
		ID:          ObservationID("1101"),
		GroupNumber: "81",
		Name:        "長野山",
	},
	"1261_81": Observation{
		ID:          ObservationID("1261"),
		GroupNumber: "81",
		Name:        "和田",
	},
	"1367_81": Observation{
		ID:          ObservationID("1367"),
		GroupNumber: "81",
		Name:        "篠生",
	},
	"1614_81": Observation{
		ID:          ObservationID("1614"),
		GroupNumber: "81",
		Name:        "鹿野",
	},
	"1620_81": Observation{
		ID:          ObservationID("1620"),
		GroupNumber: "81",
		Name:        "東厚保",
	},
	"47754_81": Observation{
		ID:          ObservationID("47754"),
		GroupNumber: "81",
		Name:        "萩",
	},
	"47762_81": Observation{
		ID:          ObservationID("47762"),
		GroupNumber: "81",
		Name:        "下関",
	},
	"47784_81": Observation{
		ID:          ObservationID("47784"),
		GroupNumber: "81",
		Name:        "山口",
	},
	"0780_82": Observation{
		ID:          ObservationID("0780"),
		GroupNumber: "82",
		Name:        "八幡",
	},
	"0781_82": Observation{
		ID:          ObservationID("0781"),
		GroupNumber: "82",
		Name:        "小倉",
	},
	"0782_82": Observation{
		ID:          ObservationID("0782"),
		GroupNumber: "82",
		Name:        "行橋",
	},
	"0783_82": Observation{
		ID:          ObservationID("0783"),
		GroupNumber: "82",
		Name:        "篠栗",
	},
	"0785_82": Observation{
		ID:          ObservationID("0785"),
		GroupNumber: "82",
		Name:        "前原",
	},
	"0787_82": Observation{
		ID:          ObservationID("0787"),
		GroupNumber: "82",
		Name:        "九千部山",
	},
	"0788_82": Observation{
		ID:          ObservationID("0788"),
		GroupNumber: "82",
		Name:        "朝倉",
	},
	"0789_82": Observation{
		ID:          ObservationID("0789"),
		GroupNumber: "82",
		Name:        "英彦山",
	},
	"0790_82": Observation{
		ID:          ObservationID("0790"),
		GroupNumber: "82",
		Name:        "久留米",
	},
	"0791_82": Observation{
		ID:          ObservationID("0791"),
		GroupNumber: "82",
		Name:        "耳納山",
	},
	"0792_82": Observation{
		ID:          ObservationID("0792"),
		GroupNumber: "82",
		Name:        "黒木",
	},
	"0793_82": Observation{
		ID:          ObservationID("0793"),
		GroupNumber: "82",
		Name:        "大牟田",
	},
	"0943_82": Observation{
		ID:          ObservationID("0943"),
		GroupNumber: "82",
		Name:        "宗像",
	},
	"0944_82": Observation{
		ID:          ObservationID("0944"),
		GroupNumber: "82",
		Name:        "頂吉",
	},
	"0945_82": Observation{
		ID:          ObservationID("0945"),
		GroupNumber: "82",
		Name:        "柳川",
	},
	"1046_82": Observation{
		ID:          ObservationID("1046"),
		GroupNumber: "82",
		Name:        "添田",
	},
	"1141_82": Observation{
		ID:          ObservationID("1141"),
		GroupNumber: "82",
		Name:        "太宰府",
	},
	"1477_82": Observation{
		ID:          ObservationID("1477"),
		GroupNumber: "82",
		Name:        "博多",
	},
	"1527_82": Observation{
		ID:          ObservationID("1527"),
		GroupNumber: "82",
		Name:        "曽根",
	},
	"1590_82": Observation{
		ID:          ObservationID("1590"),
		GroupNumber: "82",
		Name:        "空港北町",
	},
	"1611_82": Observation{
		ID:          ObservationID("1611"),
		GroupNumber: "82",
		Name:        "早良脇山",
	},
	"47807_82": Observation{
		ID:          ObservationID("47807"),
		GroupNumber: "82",
		Name:        "福岡",
	},
	"47809_82": Observation{
		ID:          ObservationID("47809"),
		GroupNumber: "82",
		Name:        "飯塚",
	},
	"0794_83": Observation{
		ID:          ObservationID("0794"),
		GroupNumber: "83",
		Name:        "中津",
	},
	"0795_83": Observation{
		ID:          ObservationID("0795"),
		GroupNumber: "83",
		Name:        "豊後高田",
	},
	"0796_83": Observation{
		ID:          ObservationID("0796"),
		GroupNumber: "83",
		Name:        "武蔵",
	},
	"0797_83": Observation{
		ID:          ObservationID("0797"),
		GroupNumber: "83",
		Name:        "伏木",
	},
	"0799_83": Observation{
		ID:          ObservationID("0799"),
		GroupNumber: "83",
		Name:        "湯布院",
	},
	"0801_83": Observation{
		ID:          ObservationID("0801"),
		GroupNumber: "83",
		Name:        "佐賀関",
	},
	"0802_83": Observation{
		ID:          ObservationID("0802"),
		GroupNumber: "83",
		Name:        "釈迦岳",
	},
	"0803_83": Observation{
		ID:          ObservationID("0803"),
		GroupNumber: "83",
		Name:        "湯平",
	},
	"0804_83": Observation{
		ID:          ObservationID("0804"),
		GroupNumber: "83",
		Name:        "温見",
	},
	"0805_83": Observation{
		ID:          ObservationID("0805"),
		GroupNumber: "83",
		Name:        "犬飼",
	},
	"0806_83": Observation{
		ID:          ObservationID("0806"),
		GroupNumber: "83",
		Name:        "竹田",
	},
	"0807_83": Observation{
		ID:          ObservationID("0807"),
		GroupNumber: "83",
		Name:        "出羽",
	},
	"0808_83": Observation{
		ID:          ObservationID("0808"),
		GroupNumber: "83",
		Name:        "佐伯",
	},
	"0809_83": Observation{
		ID:          ObservationID("0809"),
		GroupNumber: "83",
		Name:        "倉木",
	},
	"0810_83": Observation{
		ID:          ObservationID("0810"),
		GroupNumber: "83",
		Name:        "蒲江",
	},
	"0931_83": Observation{
		ID:          ObservationID("0931"),
		GroupNumber: "83",
		Name:        "院内",
	},
	"0932_83": Observation{
		ID:          ObservationID("0932"),
		GroupNumber: "83",
		Name:        "玖珠",
	},
	"0933_83": Observation{
		ID:          ObservationID("0933"),
		GroupNumber: "83",
		Name:        "臼杵",
	},
	"0979_83": Observation{
		ID:          ObservationID("0979"),
		GroupNumber: "83",
		Name:        "耶馬渓",
	},
	"1137_83": Observation{
		ID:          ObservationID("1137"),
		GroupNumber: "83",
		Name:        "国見",
	},
	"1142_83": Observation{
		ID:          ObservationID("1142"),
		GroupNumber: "83",
		Name:        "宇目",
	},
	"1157_83": Observation{
		ID:          ObservationID("1157"),
		GroupNumber: "83",
		Name:        "別府",
	},
	"1237_83": Observation{
		ID:          ObservationID("1237"),
		GroupNumber: "83",
		Name:        "杵築",
	},
	"1554_83": Observation{
		ID:          ObservationID("1554"),
		GroupNumber: "83",
		Name:        "椿ヶ鼻",
	},
	"47814_83": Observation{
		ID:          ObservationID("47814"),
		GroupNumber: "83",
		Name:        "日田",
	},
	"47815_83": Observation{
		ID:          ObservationID("47815"),
		GroupNumber: "83",
		Name:        "大分",
	},
	"0811_84": Observation{
		ID:          ObservationID("0811"),
		GroupNumber: "84",
		Name:        "佐須奈",
	},
	"0814_84": Observation{
		ID:          ObservationID("0814"),
		GroupNumber: "84",
		Name:        "松浦",
	},
	"0815_84": Observation{
		ID:          ObservationID("0815"),
		GroupNumber: "84",
		Name:        "国見山",
	},
	"0817_84": Observation{
		ID:          ObservationID("0817"),
		GroupNumber: "84",
		Name:        "大瀬戸",
	},
	"0818_84": Observation{
		ID:          ObservationID("0818"),
		GroupNumber: "84",
		Name:        "長浦岳",
	},
	"0819_84": Observation{
		ID:          ObservationID("0819"),
		GroupNumber: "84",
		Name:        "五家原岳",
	},
	"0820_84": Observation{
		ID:          ObservationID("0820"),
		GroupNumber: "84",
		Name:        "諫早",
	},
	"0922_84": Observation{
		ID:          ObservationID("0922"),
		GroupNumber: "84",
		Name:        "口之津",
	},
	"0962_84": Observation{
		ID:          ObservationID("0962"),
		GroupNumber: "84",
		Name:        "島原",
	},
	"1080_84": Observation{
		ID:          ObservationID("1080"),
		GroupNumber: "84",
		Name:        "有川",
	},
	"1084_84": Observation{
		ID:          ObservationID("1084"),
		GroupNumber: "84",
		Name:        "大村",
	},
	"1138_84": Observation{
		ID:          ObservationID("1138"),
		GroupNumber: "84",
		Name:        "有川",
	},
	"1144_84": Observation{
		ID:          ObservationID("1144"),
		GroupNumber: "84",
		Name:        "芦辺",
	},
	"1441_84": Observation{
		ID:          ObservationID("1441"),
		GroupNumber: "84",
		Name:        "野母崎",
	},
	"1442_84": Observation{
		ID:          ObservationID("1442"),
		GroupNumber: "84",
		Name:        "深江",
	},
	"1446_84": Observation{
		ID:          ObservationID("1446"),
		GroupNumber: "84",
		Name:        "普賢岳",
	},
	"1451_84": Observation{
		ID:          ObservationID("1451"),
		GroupNumber: "84",
		Name:        "百花台",
	},
	"1452_84": Observation{
		ID:          ObservationID("1452"),
		GroupNumber: "84",
		Name:        "千々石",
	},
	"1453_84": Observation{
		ID:          ObservationID("1453"),
		GroupNumber: "84",
		Name:        "鰐浦",
	},
	"1478_84": Observation{
		ID:          ObservationID("1478"),
		GroupNumber: "84",
		Name:        "石田",
	},
	"1479_84": Observation{
		ID:          ObservationID("1479"),
		GroupNumber: "84",
		Name:        "上大津",
	},
	"1535_84": Observation{
		ID:          ObservationID("1535"),
		GroupNumber: "84",
		Name:        "小値賀",
	},
	"1536_84": Observation{
		ID:          ObservationID("1536"),
		GroupNumber: "84",
		Name:        "頭ヶ島",
	},
	"1538_84": Observation{
		ID:          ObservationID("1538"),
		GroupNumber: "84",
		Name:        "美津島",
	},
	"47800_84": Observation{
		ID:          ObservationID("47800"),
		GroupNumber: "84",
		Name:        "厳原",
	},
	"47805_84": Observation{
		ID:          ObservationID("47805"),
		GroupNumber: "84",
		Name:        "平戸",
	},
	"47812_84": Observation{
		ID:          ObservationID("47812"),
		GroupNumber: "84",
		Name:        "佐世保",
	},
	"47817_84": Observation{
		ID:          ObservationID("47817"),
		GroupNumber: "84",
		Name:        "長崎",
	},
	"47818_84": Observation{
		ID:          ObservationID("47818"),
		GroupNumber: "84",
		Name:        "雲仙岳",
	},
	"47843_84": Observation{
		ID:          ObservationID("47843"),
		GroupNumber: "84",
		Name:        "福江",
	},
	"0824_85": Observation{
		ID:          ObservationID("0824"),
		GroupNumber: "85",
		Name:        "枝去木",
	},
	"0825_85": Observation{
		ID:          ObservationID("0825"),
		GroupNumber: "85",
		Name:        "和多田",
	},
	"0826_85": Observation{
		ID:          ObservationID("0826"),
		GroupNumber: "85",
		Name:        "権現山",
	},
	"0827_85": Observation{
		ID:          ObservationID("0827"),
		GroupNumber: "85",
		Name:        "八幡岳",
	},
	"0829_85": Observation{
		ID:          ObservationID("0829"),
		GroupNumber: "85",
		Name:        "白石",
	},
	"0830_85": Observation{
		ID:          ObservationID("0830"),
		GroupNumber: "85",
		Name:        "嬉野",
	},
	"0831_85": Observation{
		ID:          ObservationID("0831"),
		GroupNumber: "85",
		Name:        "多良岳",
	},
	"1075_85": Observation{
		ID:          ObservationID("1075"),
		GroupNumber: "85",
		Name:        "伊万里",
	},
	"1480_85": Observation{
		ID:          ObservationID("1480"),
		GroupNumber: "85",
		Name:        "川副",
	},
	"1610_85": Observation{
		ID:          ObservationID("1610"),
		GroupNumber: "85",
		Name:        "唐津",
	},
	"1612_85": Observation{
		ID:          ObservationID("1612"),
		GroupNumber: "85",
		Name:        "鳥栖",
	},
	"1616_85": Observation{
		ID:          ObservationID("1616"),
		GroupNumber: "85",
		Name:        "北山",
	},
	"47813_85": Observation{
		ID:          ObservationID("47813"),
		GroupNumber: "85",
		Name:        "佐賀",
	},
	"0832_86": Observation{
		ID:          ObservationID("0832"),
		GroupNumber: "86",
		Name:        "鹿北",
	},
	"0833_86": Observation{
		ID:          ObservationID("0833"),
		GroupNumber: "86",
		Name:        "南小国",
	},
	"0834_86": Observation{
		ID:          ObservationID("0834"),
		GroupNumber: "86",
		Name:        "岱明",
	},
	"0835_86": Observation{
		ID:          ObservationID("0835"),
		GroupNumber: "86",
		Name:        "菊池",
	},
	"0836_86": Observation{
		ID:          ObservationID("0836"),
		GroupNumber: "86",
		Name:        "鞍岳",
	},
	"0838_86": Observation{
		ID:          ObservationID("0838"),
		GroupNumber: "86",
		Name:        "俵山",
	},
	"0840_86": Observation{
		ID:          ObservationID("0840"),
		GroupNumber: "86",
		Name:        "高森",
	},
	"0841_86": Observation{
		ID:          ObservationID("0841"),
		GroupNumber: "86",
		Name:        "間の谷山",
	},
	"0842_86": Observation{
		ID:          ObservationID("0842"),
		GroupNumber: "86",
		Name:        "甲佐",
	},
	"0843_86": Observation{
		ID:          ObservationID("0843"),
		GroupNumber: "86",
		Name:        "松島",
	},
	"0844_86": Observation{
		ID:          ObservationID("0844"),
		GroupNumber: "86",
		Name:        "大金峰",
	},
	"0845_86": Observation{
		ID:          ObservationID("0845"),
		GroupNumber: "86",
		Name:        "本渡",
	},
	"0846_86": Observation{
		ID:          ObservationID("0846"),
		GroupNumber: "86",
		Name:        "八代",
	},
	"0847_86": Observation{
		ID:          ObservationID("0847"),
		GroupNumber: "86",
		Name:        "戦山",
	},
	"0850_86": Observation{
		ID:          ObservationID("0850"),
		GroupNumber: "86",
		Name:        "一里山",
	},
	"0851_86": Observation{
		ID:          ObservationID("0851"),
		GroupNumber: "86",
		Name:        "白髪岳",
	},
	"0923_86": Observation{
		ID:          ObservationID("0923"),
		GroupNumber: "86",
		Name:        "益城",
	},
	"0924_86": Observation{
		ID:          ObservationID("0924"),
		GroupNumber: "86",
		Name:        "水俣",
	},
	"0925_86": Observation{
		ID:          ObservationID("0925"),
		GroupNumber: "86",
		Name:        "多良木",
	},
	"0926_86": Observation{
		ID:          ObservationID("0926"),
		GroupNumber: "86",
		Name:        "上",
	},
	"0946_86": Observation{
		ID:          ObservationID("0946"),
		GroupNumber: "86",
		Name:        "田浦",
	},
	"1081_86": Observation{
		ID:          ObservationID("1081"),
		GroupNumber: "86",
		Name:        "三角",
	},
	"1086_86": Observation{
		ID:          ObservationID("1086"),
		GroupNumber: "86",
		Name:        "湯前横谷",
	},
	"1088_86": Observation{
		ID:          ObservationID("1088"),
		GroupNumber: "86",
		Name:        "茶臼峠",
	},
	"1156_86": Observation{
		ID:          ObservationID("1156"),
		GroupNumber: "86",
		Name:        "屋形山",
	},
	"1158_86": Observation{
		ID:          ObservationID("1158"),
		GroupNumber: "86",
		Name:        "五家荘",
	},
	"1240_86": Observation{
		ID:          ObservationID("1240"),
		GroupNumber: "86",
		Name:        "阿蘇乙姫",
	},
	"1368_86": Observation{
		ID:          ObservationID("1368"),
		GroupNumber: "86",
		Name:        "五木",
	},
	"1435_86": Observation{
		ID:          ObservationID("1435"),
		GroupNumber: "86",
		Name:        "山江",
	},
	"1580_86": Observation{
		ID:          ObservationID("1580"),
		GroupNumber: "86",
		Name:        "宇土",
	},
	"1581_86": Observation{
		ID:          ObservationID("1581"),
		GroupNumber: "86",
		Name:        "一勝地",
	},
	"1613_86": Observation{
		ID:          ObservationID("1613"),
		GroupNumber: "86",
		Name:        "山都",
	},
	"47819_86": Observation{
		ID:          ObservationID("47819"),
		GroupNumber: "86",
		Name:        "熊本",
	},
	"47821_86": Observation{
		ID:          ObservationID("47821"),
		GroupNumber: "86",
		Name:        "阿蘇山",
	},
	"47824_86": Observation{
		ID:          ObservationID("47824"),
		GroupNumber: "86",
		Name:        "人吉",
	},
	"47838_86": Observation{
		ID:          ObservationID("47838"),
		GroupNumber: "86",
		Name:        "牛深",
	},
	"0852_87": Observation{
		ID:          ObservationID("0852"),
		GroupNumber: "87",
		Name:        "高千穂",
	},
	"0853_87": Observation{
		ID:          ObservationID("0853"),
		GroupNumber: "87",
		Name:        "大中尾",
	},
	"0854_87": Observation{
		ID:          ObservationID("0854"),
		GroupNumber: "87",
		Name:        "中小屋",
	},
	"0856_87": Observation{
		ID:          ObservationID("0856"),
		GroupNumber: "87",
		Name:        "猪原",
	},
	"0857_87": Observation{
		ID:          ObservationID("0857"),
		GroupNumber: "87",
		Name:        "日向",
	},
	"0858_87": Observation{
		ID:          ObservationID("0858"),
		GroupNumber: "87",
		Name:        "矢櫃岳",
	},
	"0859_87": Observation{
		ID:          ObservationID("0859"),
		GroupNumber: "87",
		Name:        "高鍋",
	},
	"0860_87": Observation{
		ID:          ObservationID("0860"),
		GroupNumber: "87",
		Name:        "加久藤",
	},
	"0861_87": Observation{
		ID:          ObservationID("0861"),
		GroupNumber: "87",
		Name:        "西都",
	},
	"0862_87": Observation{
		ID:          ObservationID("0862"),
		GroupNumber: "87",
		Name:        "えびの",
	},
	"0863_87": Observation{
		ID:          ObservationID("0863"),
		GroupNumber: "87",
		Name:        "小林",
	},
	"0864_87": Observation{
		ID:          ObservationID("0864"),
		GroupNumber: "87",
		Name:        "野尻",
	},
	"0865_87": Observation{
		ID:          ObservationID("0865"),
		GroupNumber: "87",
		Name:        "国富",
	},
	"0867_87": Observation{
		ID:          ObservationID("0867"),
		GroupNumber: "87",
		Name:        "霧島御池",
	},
	"0868_87": Observation{
		ID:          ObservationID("0868"),
		GroupNumber: "87",
		Name:        "青島",
	},
	"0870_87": Observation{
		ID:          ObservationID("0870"),
		GroupNumber: "87",
		Name:        "鰐塚山",
	},
	"0872_87": Observation{
		ID:          ObservationID("0872"),
		GroupNumber: "87",
		Name:        "串間",
	},
	"1135_87": Observation{
		ID:          ObservationID("1135"),
		GroupNumber: "87",
		Name:        "古江",
	},
	"1164_87": Observation{
		ID:          ObservationID("1164"),
		GroupNumber: "87",
		Name:        "深瀬",
	},
	"1275_87": Observation{
		ID:          ObservationID("1275"),
		GroupNumber: "87",
		Name:        "見立",
	},
	"1343_87": Observation{
		ID:          ObservationID("1343"),
		GroupNumber: "87",
		Name:        "諸塚",
	},
	"1345_87": Observation{
		ID:          ObservationID("1345"),
		GroupNumber: "87",
		Name:        "西米良",
	},
	"1348_87": Observation{
		ID:          ObservationID("1348"),
		GroupNumber: "87",
		Name:        "鞍岡",
	},
	"1353_87": Observation{
		ID:          ObservationID("1353"),
		GroupNumber: "87",
		Name:        "神門",
	},
	"1432_87": Observation{
		ID:          ObservationID("1432"),
		GroupNumber: "87",
		Name:        "北方",
	},
	"1433_87": Observation{
		ID:          ObservationID("1433"),
		GroupNumber: "87",
		Name:        "上椎葉",
	},
	"1481_87": Observation{
		ID:          ObservationID("1481"),
		GroupNumber: "87",
		Name:        "赤江",
	},
	"1608_87": Observation{
		ID:          ObservationID("1608"),
		GroupNumber: "87",
		Name:        "都農",
	},
	"1609_87": Observation{
		ID:          ObservationID("1609"),
		GroupNumber: "87",
		Name:        "日之影",
	},
	"47822_87": Observation{
		ID:          ObservationID("47822"),
		GroupNumber: "87",
		Name:        "延岡",
	},
	"47829_87": Observation{
		ID:          ObservationID("47829"),
		GroupNumber: "87",
		Name:        "都城",
	},
	"47830_87": Observation{
		ID:          ObservationID("47830"),
		GroupNumber: "87",
		Name:        "宮崎",
	},
	"47835_87": Observation{
		ID:          ObservationID("47835"),
		GroupNumber: "87",
		Name:        "油津",
	},
	"0874_88": Observation{
		ID:          ObservationID("0874"),
		GroupNumber: "88",
		Name:        "古原野",
	},
	"0875_88": Observation{
		ID:          ObservationID("0875"),
		GroupNumber: "88",
		Name:        "大口",
	},
	"0876_88": Observation{
		ID:          ObservationID("0876"),
		GroupNumber: "88",
		Name:        "魚野越",
	},
	"0877_88": Observation{
		ID:          ObservationID("0877"),
		GroupNumber: "88",
		Name:        "紫尾山",
	},
	"0878_88": Observation{
		ID:          ObservationID("0878"),
		GroupNumber: "88",
		Name:        "さつま柏原",
	},
	"0879_88": Observation{
		ID:          ObservationID("0879"),
		GroupNumber: "88",
		Name:        "中甑",
	},
	"0880_88": Observation{
		ID:          ObservationID("0880"),
		GroupNumber: "88",
		Name:        "川内",
	},
	"0881_88": Observation{
		ID:          ObservationID("0881"),
		GroupNumber: "88",
		Name:        "溝辺",
	},
	"0882_88": Observation{
		ID:          ObservationID("0882"),
		GroupNumber: "88",
		Name:        "八重山",
	},
	"0883_88": Observation{
		ID:          ObservationID("0883"),
		GroupNumber: "88",
		Name:        "東市来",
	},
	"0884_88": Observation{
		ID:          ObservationID("0884"),
		GroupNumber: "88",
		Name:        "権現ケ尾",
	},
	"0886_88": Observation{
		ID:          ObservationID("0886"),
		GroupNumber: "88",
		Name:        "高峠",
	},
	"0887_88": Observation{
		ID:          ObservationID("0887"),
		GroupNumber: "88",
		Name:        "加世田",
	},
	"0888_88": Observation{
		ID:          ObservationID("0888"),
		GroupNumber: "88",
		Name:        "吉ケ別府",
	},
	"0889_88": Observation{
		ID:          ObservationID("0889"),
		GroupNumber: "88",
		Name:        "大浦",
	},
	"0890_88": Observation{
		ID:          ObservationID("0890"),
		GroupNumber: "88",
		Name:        "喜入",
	},
	"0891_88": Observation{
		ID:          ObservationID("0891"),
		GroupNumber: "88",
		Name:        "肝属川",
	},
	"0893_88": Observation{
		ID:          ObservationID("0893"),
		GroupNumber: "88",
		Name:        "指宿",
	},
	"0894_88": Observation{
		ID:          ObservationID("0894"),
		GroupNumber: "88",
		Name:        "甫与志岳",
	},
	"0895_88": Observation{
		ID:          ObservationID("0895"),
		GroupNumber: "88",
		Name:        "内之浦",
	},
	"0897_88": Observation{
		ID:          ObservationID("0897"),
		GroupNumber: "88",
		Name:        "中種子",
	},
	"0898_88": Observation{
		ID:          ObservationID("0898"),
		GroupNumber: "88",
		Name:        "上中",
	},
	"0899_88": Observation{
		ID:          ObservationID("0899"),
		GroupNumber: "88",
		Name:        "尾之間",
	},
	"0934_88": Observation{
		ID:          ObservationID("0934"),
		GroupNumber: "88",
		Name:        "出水",
	},
	"0935_88": Observation{
		ID:          ObservationID("0935"),
		GroupNumber: "88",
		Name:        "牧之原",
	},
	"0936_88": Observation{
		ID:          ObservationID("0936"),
		GroupNumber: "88",
		Name:        "志布志",
	},
	"0937_88": Observation{
		ID:          ObservationID("0937"),
		GroupNumber: "88",
		Name:        "鹿屋",
	},
	"0980_88": Observation{
		ID:          ObservationID("0980"),
		GroupNumber: "88",
		Name:        "古仁屋",
	},
	"1104_88": Observation{
		ID:          ObservationID("1104"),
		GroupNumber: "88",
		Name:        "矢止岳",
	},
	"1136_88": Observation{
		ID:          ObservationID("1136"),
		GroupNumber: "88",
		Name:        "田代",
	},
	"1139_88": Observation{
		ID:          ObservationID("1139"),
		GroupNumber: "88",
		Name:        "伊仙",
	},
	"1146_88": Observation{
		ID:          ObservationID("1146"),
		GroupNumber: "88",
		Name:        "輝北",
	},
	"1149_88": Observation{
		ID:          ObservationID("1149"),
		GroupNumber: "88",
		Name:        "大隅",
	},
	"1161_88": Observation{
		ID:          ObservationID("1161"),
		GroupNumber: "88",
		Name:        "佐多",
	},
	"1181_88": Observation{
		ID:          ObservationID("1181"),
		GroupNumber: "88",
		Name:        "肝付前田",
	},
	"1277_88": Observation{
		ID:          ObservationID("1277"),
		GroupNumber: "88",
		Name:        "与論島",
	},
	"1296_88": Observation{
		ID:          ObservationID("1296"),
		GroupNumber: "88",
		Name:        "喜界島",
	},
	"1520_88": Observation{
		ID:          ObservationID("1520"),
		GroupNumber: "88",
		Name:        "笠利",
	},
	"1540_88": Observation{
		ID:          ObservationID("1540"),
		GroupNumber: "88",
		Name:        "中之島",
	},
	"1542_88": Observation{
		ID:          ObservationID("1542"),
		GroupNumber: "88",
		Name:        "天城",
	},
	"47823_88": Observation{
		ID:          ObservationID("47823"),
		GroupNumber: "88",
		Name:        "阿久根",
	},
	"47827_88": Observation{
		ID:          ObservationID("47827"),
		GroupNumber: "88",
		Name:        "鹿児島",
	},
	"47831_88": Observation{
		ID:          ObservationID("47831"),
		GroupNumber: "88",
		Name:        "枕崎",
	},
	"47836_88": Observation{
		ID:          ObservationID("47836"),
		GroupNumber: "88",
		Name:        "屋久島",
	},
	"47837_88": Observation{
		ID:          ObservationID("47837"),
		GroupNumber: "88",
		Name:        "種子島",
	},
	"47909_88": Observation{
		ID:          ObservationID("47909"),
		GroupNumber: "88",
		Name:        "名瀬",
	},
	"47942_88": Observation{
		ID:          ObservationID("47942"),
		GroupNumber: "88",
		Name:        "沖永良部",
	},
	"0901_91": Observation{
		ID:          ObservationID("0901"),
		GroupNumber: "91",
		Name:        "奥",
	},
	"0902_91": Observation{
		ID:          ObservationID("0902"),
		GroupNumber: "91",
		Name:        "与那覇岳",
	},
	"0903_91": Observation{
		ID:          ObservationID("0903"),
		GroupNumber: "91",
		Name:        "呉我山",
	},
	"0905_91": Observation{
		ID:          ObservationID("0905"),
		GroupNumber: "91",
		Name:        "名嘉真",
	},
	"0906_91": Observation{
		ID:          ObservationID("0906"),
		GroupNumber: "91",
		Name:        "読谷",
	},
	"0907_91": Observation{
		ID:          ObservationID("0907"),
		GroupNumber: "91",
		Name:        "胡屋",
	},
	"0909_91": Observation{
		ID:          ObservationID("0909"),
		GroupNumber: "91",
		Name:        "糸数",
	},
	"1064_91": Observation{
		ID:          ObservationID("1064"),
		GroupNumber: "91",
		Name:        "金武",
	},
	"1145_91": Observation{
		ID:          ObservationID("1145"),
		GroupNumber: "91",
		Name:        "伊原間",
	},
	"1147_91": Observation{
		ID:          ObservationID("1147"),
		GroupNumber: "91",
		Name:        "伊是名",
	},
	"1152_91": Observation{
		ID:          ObservationID("1152"),
		GroupNumber: "91",
		Name:        "渡嘉敷",
	},
	"1246_91": Observation{
		ID:          ObservationID("1246"),
		GroupNumber: "91",
		Name:        "伊良部",
	},
	"1247_91": Observation{
		ID:          ObservationID("1247"),
		GroupNumber: "91",
		Name:        "多良間",
	},
	"1248_91": Observation{
		ID:          ObservationID("1248"),
		GroupNumber: "91",
		Name:        "城辺",
	},
	"1250_91": Observation{
		ID:          ObservationID("1250"),
		GroupNumber: "91",
		Name:        "川平",
	},
	"1251_91": Observation{
		ID:          ObservationID("1251"),
		GroupNumber: "91",
		Name:        "大原",
	},
	"1354_91": Observation{
		ID:          ObservationID("1354"),
		GroupNumber: "91",
		Name:        "波照間",
	},
	"1373_91": Observation{
		ID:          ObservationID("1373"),
		GroupNumber: "91",
		Name:        "本部",
	},
	"1374_91": Observation{
		ID:          ObservationID("1374"),
		GroupNumber: "91",
		Name:        "東",
	},
	"1482_91": Observation{
		ID:          ObservationID("1482"),
		GroupNumber: "91",
		Name:        "安次嶺",
	},
	"1483_91": Observation{
		ID:          ObservationID("1483"),
		GroupNumber: "91",
		Name:        "鏡原",
	},
	"1484_91": Observation{
		ID:          ObservationID("1484"),
		GroupNumber: "91",
		Name:        "真栄里",
	},
	"1490_91": Observation{
		ID:          ObservationID("1490"),
		GroupNumber: "91",
		Name:        "下地",
	},
	"1491_91": Observation{
		ID:          ObservationID("1491"),
		GroupNumber: "91",
		Name:        "仲筋",
	},
	"1493_91": Observation{
		ID:          ObservationID("1493"),
		GroupNumber: "91",
		Name:        "所野",
	},
	"1494_91": Observation{
		ID:          ObservationID("1494"),
		GroupNumber: "91",
		Name:        "志多阿原",
	},
	"1515_91": Observation{
		ID:          ObservationID("1515"),
		GroupNumber: "91",
		Name:        "粟国",
	},
	"1516_91": Observation{
		ID:          ObservationID("1516"),
		GroupNumber: "91",
		Name:        "慶良間",
	},
	"1517_91": Observation{
		ID:          ObservationID("1517"),
		GroupNumber: "91",
		Name:        "北大東",
	},
	"1518_91": Observation{
		ID:          ObservationID("1518"),
		GroupNumber: "91",
		Name:        "旧東",
	},
	"1534_91": Observation{
		ID:          ObservationID("1534"),
		GroupNumber: "91",
		Name:        "北原",
	},
	"1586_91": Observation{
		ID:          ObservationID("1586"),
		GroupNumber: "91",
		Name:        "国頭",
	},
	"1596_91": Observation{
		ID:          ObservationID("1596"),
		GroupNumber: "91",
		Name:        "宮城島",
	},
	"47912_91": Observation{
		ID:          ObservationID("47912"),
		GroupNumber: "91",
		Name:        "与那国島",
	},
	"47917_91": Observation{
		ID:          ObservationID("47917"),
		GroupNumber: "91",
		Name:        "西表島",
	},
	"47918_91": Observation{
		ID:          ObservationID("47918"),
		GroupNumber: "91",
		Name:        "石垣島",
	},
	"47927_91": Observation{
		ID:          ObservationID("47927"),
		GroupNumber: "91",
		Name:        "宮古島",
	},
	"47929_91": Observation{
		ID:          ObservationID("47929"),
		GroupNumber: "91",
		Name:        "久米島",
	},
	"47936_91": Observation{
		ID:          ObservationID("47936"),
		GroupNumber: "91",
		Name:        "那覇",
	},
	"47940_91": Observation{
		ID:          ObservationID("47940"),
		GroupNumber: "91",
		Name:        "名護",
	},
	"47945_91": Observation{
		ID:          ObservationID("47945"),
		GroupNumber: "91",
		Name:        "南大東(南大東島)",
	},
	"89532_99": Observation{
		ID:          ObservationID("89532"),
		GroupNumber: "99",
		Name:        "昭和",
	},
}
