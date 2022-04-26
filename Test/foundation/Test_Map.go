package foundation

import "fmt"

func Map() {
	// 声明变量
	var countryCapitalMap map[string]string /*创建集合 */
	// 使用make函数创建映射对象实例
	countryCapitalMap = make(map[string]string)
	// 或者直接实例化
	//var m2 = map[string]string{}

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	// 使用range枚举map中的元素，输出地图值
	for country, city := range countryCapitalMap {
		fmt.Println(country, "首都是", city)
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}
