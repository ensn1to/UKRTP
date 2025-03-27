package electricity  

import (  
	"time"  
)  

// UKElectricityCost 表示英国电费账单的完整结构  
type UKElectricityCost struct {  
	BasicInfo         BasicInfo          `json:"basic_info"`          // 文章基本信息  
	CostComponents    []CostComponent    `json:"cost_components"`     // 电费主要组成部分概览  
	WholesaleCost     WholesaleCost      `json:"wholesale_cost"`      // 批发成本详细信息  
	NetworkCosts      NetworkCosts       `json:"network_costs"`       // 网络成本详细信息  
	GovernmentCharges []GovernmentCharge `json:"government_charges"`  // 政府义务费用列表  
	SourceInfo        SourceInfo         `json:"source_info"`         // 数据来源信息  
	ContactInfo       ContactInfo        `json:"contact_info"`        // 项目联系信息  
}  

// BasicInfo 包含文章的基本元数据信息  
type BasicInfo struct {  
	Title       string    `json:"title"`        // 文章标题(如"英国的电费结构")  
	Translator  string    `json:"translator"`   // 译者信息(如"CLEANdata译")  
	PublishDate time.Time `json:"publish_date"` // 发布日期(如2025-03-27)  
	SourceURL   string    `json:"source_url"`   // 原文链接URL  
	Summary     string    `json:"summary"`      // 内容摘要(如"电费由不同费用组成...")  
}  

// CostComponent 描述电费的主要构成部分  
type CostComponent struct {  
	Name        string  `json:"name"`         // 组成部分名称(如"批发成本")  
	Description string  `json:"description"`  // 详细描述  
	Percentage  float64 `json:"percentage"`   // 占总费用的百分比(如35表示35%)  
}  

// WholesaleCost 描述电力批发市场成本信息  
type WholesaleCost struct {  
	Description string    `json:"description"` // 批发成本特点描述  
	PriceChart  PriceData `json:"price_chart"` // 历史价格数据图表  
	Factors     []string  `json:"factors"`     // 影响价格的因素列表(如["季节变化","天然气价格"])  
}  

// PriceData 表示电力价格的时间序列数据  
type PriceData struct {  
	TimePeriod string      `json:"time_period"` // 数据时间范围(如"2010-2020")  
	Unit       string      `json:"unit"`        // 价格单位(如"GBP/MWh")  
	DataPoints []DataPoint `json:"data_points"` // 具体数据点集合  
	Source     string      `json:"source"`      // 数据来源(如"Ofgem")  
}  

// DataPoint 表示单个时间点的价格数据  
type DataPoint struct {  
	Year  int     `json:"year"`  // 年份(如2020)  
	Month int     `json:"month"` // 月份(1-12)  
	Value float64 `json:"value"` // 该时间点的电价数值  
}  

// NetworkCosts 描述电网相关费用结构  
type NetworkCosts struct {  
	Description       string             `json:"description"`        // 网络成本总体描述  
	TransmissionCost  TransmissionCost   `json:"transmission_cost"`  // 输电系统费用详情  
	DistributionCost  DistributionCost   `json:"distribution_cost"`  // 配电系统费用详情  
	OtherCosts       []NetworkSubCharge `json:"other_costs"`        // 其他网络相关费用列表  
}  

// TransmissionCost 输电系统使用费(TNUoS)详情  
type TransmissionCost struct {  
	Name        string  `json:"name"`         // 费用名称(如"输电系统使用费")  
	Description string  `json:"description"`  // 费用用途描述  
	Percentage  float64 `json:"percentage"`   // 占总费用比例(如7-10%)  
	Calculation string  `json:"calculation"`  // 计算方法描述(如"基于高峰时段消耗计算")  
}  

// DistributionCost 配电系统使用费(DUoS)详情  
type DistributionCost struct {  
	Name        string           `json:"name"`        // 费用名称(如"配电系统使用费")  
	Description string           `json:"description"` // 费用用途描述  
	Percentage  float64          `json:"percentage"`  // 占总费用比例(如12-15%)  
	Components  []DUoSComponent  `json:"components"`  // DUoS费用子项列表  
	TimeBands   []TimeBandCharge `json:"time_bands"`  // 不同时间段费率信息  
}  

// DUoSComponent 配电系统使用费的子项  
type DUoSComponent struct {  
	Name        string  `json:"name"`        // 子项名称(如"单位收费")  
	Description string  `json:"description"` // 子项说明  
	Percentage  float64 `json:"percentage"`  // 占DUoS费用的比例  
}  

// TimeBandCharge 不同时间段的费率差异  
type TimeBandCharge struct {  
	DayType    string  `json:"day_type"`   // 日期类型(工作日/周末)  
	TimeBand   string  `json:"time_band"`  // 时间段标识(绿色/琥珀色/红色)  
	Rate       float64 `json:"rate"`       // 基础费率(单位:英镑/千瓦时)  
	Multiplier float64 `json:"multiplier"` // 相对于基础费率的倍数(如红色时段可能是4倍)  
}  

// NetworkSubCharge 其他网络相关费用  
type NetworkSubCharge struct {  
	Name        string  `json:"name"`        // 费用名称(如"平衡系统服务使用费")  
	Description string  `json:"description"` // 费用用途描述  
	Percentage  float64 `json:"percentage"`  // 占总费用比例  
}  

// GovernmentCharge 政府征收的环境和社会义务费用  
type GovernmentCharge struct {  
	Name         string  `json:"name"`          // 全称(如"可再生能源义务")  
	Abbreviation string  `json:"abbreviation"`  // 缩写(如"RO")  
	Description  string  `json:"description"`   // 政策目的描述  
	Percentage   float64 `json:"percentage"`    // 占总费用比例  
	Period       string  `json:"period"`        // 特殊征收时段(如"冬季工作日下午4-7点")  
}  
