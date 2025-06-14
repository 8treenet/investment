package utility

import (
	"fmt"
	"testing"
)

var htmldata = `

04-14 22:28 星期一

2024年，面对复杂多变的市场环境和政策红利加速释放，海南机场锚定做强做优做大目标，抢抓重大发展机遇促转型、谋创新，持续拓宽空中“朋友圈”和全面提升服务效能。此外，公司进一步优化资源配置，在发挥多业态协同优势的同时，还大力推进重点项目建设、主动“去地产化”、完善分红机制等，以构建公司可持续发展的长效竞争力。

破浪启新局，海南机场2024年财报交出了一份创新提质、凸显韧性的答卷，机场主业营收占比大幅提升、重点建设项目如期推进、资产负债结构持续优化。2025年是海南全岛封关运作的重要节点，海南机场涉及的多个封关运作项目也已将悉数竣工验收。市场分析指出，海南机场拥有海南自贸港的关键空港枢纽码头优势和海口日月广场免税标杆商业资源等核心优势，中长线有望直接受益海南自贸港建设。

![](https://img.cls.cn/images/20250414/7F4pDKoygo.jpg)

海南机场在资本市场的长期投资价值获得多方肯定。作为海南机场间接控股股东，海南控股表示，坚定看好旗下上市公司未来发展前景，已于近日增持海南机场股票1.54亿元。多家券商也给予海南机场“优大于市”“买入”“增持”“推荐”等评级。

**业绩表现健康向好：主业营收占比跃升、关键经营指标创新高**

2024年，海南机场实现营业收入43.68亿元；其中，机场管理主业业务收入占比大幅提升至41.68%。短期受地产去化节奏及海南免税贡献回落等因素影响业绩承压，归母净利润为4.59亿元，同比下降51.88%。截至2024年底，海南机场总资产536.91亿元，同比下降3.77%，归母净资产235.19亿元，同比增长1.72%；资产负债率52.17%，较期初下降3.42个百分点。

近年来海南机场持续推进聚焦机场主业战略，加快存量地产项目去化，报告期内完成海口海航大厦、儋州3024亩农业发展观光用地等项目公司处置，资产结构进一步优化。

![](https://img.cls.cn/images/20250414/L58g3ge0nr.jpg)

按计划，海南将在2025年年底前适时启动封关运作。海南机场旗下三亚凤凰机场、琼海博鳌机场共涉及4个封关运作项目，包括三亚凤凰机场国际航站楼改扩建项目、“二线口岸”查验设施设备建设项目，琼海博鳌机场货运基建新建改造项目、国内航站楼基建改造及海关监管配套查验设备设施项目，目前已全部竣工验收。

**业绩短痛虽存但未改健康基本盘，海南机场2024年年度财报透露诸多向好信号。**

海南机场在年报中表示，消费增长放缓及行业周期等因素导致免税商业业务业绩同比有所下降，从而影响公司对联营单位的投资收益。由于公司在处置海南兴业国际联合实业发展有限公司产生的投资收益、部分债务重组收益等属于非经常性事项，导致归属于母公司所有者的扣除非经常性损益的净利润同比下降。

海南机场年内盈利承压或还与系列升级变革有关。转型是一场资金密集型的战役，过程中企业首先面临的可能是业绩下滑的困境，但变则通、通则久、久则达，因此用短期阵痛换护城河是不少业界公司的做法。

![](https://img.cls.cn/images/20250414/9329BevFFw.jpg)
`

func TestConvertMarkDown(t *testing.T) {

	// 删除特定的图片链接
	pattern := `!$$[^$$]*\]$https://img\.cls\.cn/images/[^)]*\.jpg$`
	output, err := ConvertMarkDown(htmldata, []string{pattern})
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
