package zh

import "github.com/olebedev/when/rules"

var All = []rules.Rule{
	Weekday(rules.Override),
	CasualDate(rules.Override),
	CasualTime(rules.Override),
	HourMinute(rules.Override),
	ExactMonthDate(rules.Override),
	TraditionHour(rules.Override),
	AfterTime(rules.Override),
}

var WEEKDAY_OFFSET = map[string]int{
	"天": 7,
	"一": 1,
	"二": 2,
	"三": 3,
	"四": 4,
	"五": 5,
	"六": 6,
	"1":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"日": 7,
}

var INTEGER_WORDS = map[string]int{
	"零一":   1,
	"零二":   2,
	"零三":   3,
	"零四":   4,
	"零五":   5,
	"零六":   6,
	"零七":   7,
	"零八":   8,
	"零九":   9,
	"一":     1,
	"二":     2,
	"两":     2,
	"三":     3,
	"四":     4,
	"五":     5,
	"六":     6,
	"七":     7,
	"八":     8,
	"九":     9,
	"十":     10,
	"十一":   11,
	"十二":   12,
	"十三":   13,
	"十四":   14,
	"十五":   15,
	"十六":   16,
	"十七":   17,
	"十八":   18,
	"十九":   19,
	"二十":   20,
	"二十一": 21,
	"二十二": 22,
	"二十三": 23,
	"二十五": 25,
	"二十六": 26,
	"二十七": 27,
	"二十八": 28,
	"二十九": 29,
	"三十":   30,
	"三十一": 31,
	"三十二": 32,
	"三十三": 33,
	"三十四": 34,
	"三十五": 35,
	"三十六": 36,
	"三十七": 37,
	"三十八": 38,
	"三十九": 39,
	"四十":   40,
	"四十一": 41,
	"四十二": 42,
	"四十三": 43,
	"四十四": 44,
	"四十五": 45,
	"四十六": 46,
	"四十七": 47,
	"四十八": 48,
	"四十九": 49,
	"五十":   50,
	"五十一": 51,
	"五十二": 52,
	"五十三": 53,
	"五十四": 54,
	"五十五": 55,
	"五十六": 56,
	"五十七": 57,
	"五十八": 58,
	"五十九": 59,
	"零":     0,
	"半":     30,
	"一刻":   15,
}
var INTEGER_WORDS_PATTERN = `(?:一刻|半|零一|零二|零三|零四|零五|零六|零七|零八|零九|一|两|二|三|四|五|六|七|八|九|十|十一|十二|十三|十四|十五|十六|十七|十八|十九|二十|二十一|二十二|二十三|二十五|二十六|二十七|二十八|二十九|三十|三十一|三十二|三十三|三十四|三十五|三十六|三十七|三十八|三十九|四十|四十一|四十二|四十三|四十四|四十五|四十六|四十七|四十八|四十九|五十|五十一|五十二|五十三|五十四|五十五|五十六|五十七|五十八|五十九|零)`

var TRADITION_HOUR_WORDS = map[string]int{
	"子时": 23,
	"丑时": 1,
	"寅时": 3,
	"卯时": 5,
	"辰时": 7,
	"巳时": 9,
	"午时": 11,
	"未时": 13,
	"申时": 15,
	"酉时": 17,
	"戌时": 19,
	"亥时": 21,
}
var TRADITION_MINUTE_WORDS = map[string]int{
	"一刻": 15,
	"两刻": 30,
	"二刻": 30,
	"三刻": 45,
	"四刻": 60,
	"五刻": 75,
	"六刻": 90,
	"七刻": 105,
	"八刻": 120,
	"1刻":  15,
	"2刻":  30,
	"3刻":  45,
	"4刻":  60,
	"5刻":  75,
	"6刻":  90,
	"7刻":  105,
	"8刻":  120,
}

var MON_WORDS = map[string]int{
	"一":   1,
	"二":   2,
	"三":   3,
	"四":   4,
	"五":   5,
	"六":   6,
	"七":   7,
	"八":   8,
	"九":   9,
	"十":   10,
	"十一": 11,
	"十二": 12,
}

var MON_WORDS_PATTERN = `十一|十二|一|二|三|四|五|六|七|八|九|十`

var DAY_WORDS = map[string]int{
	"一":     1,
	"二":     2,
	"四":     4,
	"五":     5,
	"六":     6,
	"七":     7,
	"八":     8,
	"九":     9,
	"十一":   11,
	"十二":   12,
	"十三":   13,
	"十四":   14,
	"十五":   15,
	"十六":   16,
	"十七":   17,
	"十八":   18,
	"十九":   19,
	"十":     10,
	"二十一": 21,
	"二十二": 22,
	"二十三": 23,
	"二十五": 25,
	"二十六": 26,
	"二十七": 27,
	"二十八": 28,
	"二十九": 29,
	"二十":   20,
	"三十一": 31,
	"三十":   30,
	"三":     3,
}
var DAY_WORDS_PATTERN = `一|二|四|五|六|七|八|九|十一|十二|十三|十四|十五|十六|十七|十八|十九|十|二十一|二十二|二十三|二十五|二十六|二十七|二十八|二十九|二十|三十一|三十|三`