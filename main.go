package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type product struct {
	ID     int64  `json:"id"`
	Model  string  `json:"model"`
	Line  string  `json:"line"`
	ManufacturerCode  string  `json:"manufacturer_code"`
	Color  string  `json:"color"`
	Price  float64  `json:"price"`
	ProcessorModel  string  `json:"processor_model"`
	GpuModel  string  `json:"gpu_model"`
	Ram  string  `json:"ram"`
	ScreenType  string  `json:"screen_type"`
	DeltaE  float64  `json:"delta_e"`
	ColorCoverage  string  `json:"color_coverage"`
	ScreenResolution  string  `json:"screen_resolution"`
	ContrastRatio  string  `json:"contrast_ratio"`
}

var catalog = []product{
	{ID: 1, Model: "CreatorPro X18 HX A14VMG-415RU", Line: "Creator Pro", ManufacturerCode: "9S7-182253-415", Color: "серый", Price: 644490.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 5000", Ram: "64,00", ScreenType: "IPS (miniLED)", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "3840x2400", ContrastRatio: "1000000:1"},
	{ID: 2, Model: "CreatorPro X18 HX A14VKSG-455RU", Line: "Creator Pro", ManufacturerCode: "9S7-182253-455", Color: "серый", Price: 611990.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 3500", Ram: "64,00", ScreenType: "IPS (miniLED)", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "3840x2400", ContrastRatio: "1000000:1"},
	{ID: 3, Model: "CreatorPro X18 HX A14VMG-456RU", Line: "Creator Pro", ManufacturerCode: "9S7-182253-456", Color: "серый", Price: 641690.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 5000", Ram: "64,00", ScreenType: "IPS (miniLED)", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "3840x2400", ContrastRatio: "1000000:1"},
	{ID: 4, Model: "CreatorPro 16 AI Studio A1VKG-241RU", Line: "Creator Pro", ManufacturerCode: "S7-15F414-241", Color: "серый", Price: 619990.00, ProcessorModel: "Intel® Core™ Ultra 9 185H", GpuModel: "NVIDIA RTX™ 3000 Ada", Ram: "64,00", ScreenType: "IPS (miniLED)", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "3840x2400", ContrastRatio: "1200:1"},
	{ID: 5, Model: "CreatorPro 16 AI Studio A1VKG-299RU", Line: "Creator Pro", ManufacturerCode: "9S7-15F414-299", Color: "серый", Price: 614990.00, ProcessorModel: "Intel® Core™ Ultra 9 185H", GpuModel: "NVIDIA RTX™ 3000 Ada", Ram: "32,00", ScreenType: "IPS (miniLED)", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "3840x2400", ContrastRatio: "1200:1"},
	{ID: 6, Model: "CreatorPro 16 AI Studio A1VKG-406RU", Line: "Creator Pro", ManufacturerCode: "9S7-15F414-406", Color: "серый", Price: 529990.00, ProcessorModel: "Intel® Core™ Ultra 7 155H", GpuModel: "NVIDIA RTX™ 3000 Ada", Ram: "32,00", ScreenType: "IPS (miniLED)", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "3840x2400", ContrastRatio: "1200:1"},
	{ID: 7, Model: "CreatorPro 16 AI Studio A1VMG-227RU", Line: "Creator Pro", ManufacturerCode: "9S7-15F314-227", Color: "серый", Price: 658990.00, ProcessorModel: "Intel® Core™ Ultra 7 155H", GpuModel: "NVIDIA RTX™ 5000 Ada", Ram: "64,00", ScreenType: "IPS (miniLED)", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "3840x2400", ContrastRatio: "1200:1"},
	{ID: 8, Model: "CreatorPro M16 HX C14VIG-456RU", Line: "Creator Pro", ManufacturerCode: "9S7-15P215-456", Color: "серый", Price: 215590.00, ProcessorModel: "Intel® Core™ i7 14700HX", GpuModel: "NVIDIA RTX™ 1000 Ada", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 9, Model: "CreatorPro M16 HX C14VJG-457RU", Line: "Creator Pro", ManufacturerCode: "9S7-15P215-457", Color: "серый", Price: 283990.00, ProcessorModel: "Intel® Core™ i7 14700HX", GpuModel: "NVIDIA RTX™ 2000 Ada", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 10, Model: "CreatorPro M16 HX C14VJG-494RU", Line: "Creator Pro", ManufacturerCode: "9S7-15P215-494", Color: "серый", Price: 259990.00, ProcessorModel: "Intel® Core™ i7 14700HX", GpuModel: "NVIDIA RTX™ 2000 Ada", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 11, Model: "CreatorPro M16 HX C14VJG-647RU", Line: "Creator Pro", ManufacturerCode: "9S7-15P215-647", Color: "серый", Price: 269990.00, ProcessorModel: "Intel® Core™ i7 14700HX", GpuModel: "NVIDIA RTX™ 2000 Ada", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 12, Model: "CreatorPro M16 HX C14VIG-681RU", Line: "Creator Pro", ManufacturerCode: "9S7-15P215-681", Color: "серый", Price: 189990.00, ProcessorModel: "Intel® Core™ i7 14700HX", GpuModel: "NVIDIA RTX™ 1000 Ada", Ram: "16,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 13, Model: "CreatorPro M16 HX C14VIG-682XRU", Line: "Creator Pro", ManufacturerCode: "9S7-15P215-682", Color: "серый", Price: 185990.00, ProcessorModel: "Intel® Core™ i7 14700HX", GpuModel: "NVIDIA RTX™ 1000 Ada", Ram: "16,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 14, Model: "CreatorPro M16 HX C14VJG-680XRU", Line: "Creator Pro", ManufacturerCode: "9S7-15P215-680", Color: "серый", Price: 225990.00, ProcessorModel: "Intel® Core™ i7 14700HX", GpuModel: "NVIDIA RTX™ 2000 Ada", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 15, Model: "Creator Z16HXStudio B13VFTO-055RU", Line: "Creator Pro", ManufacturerCode: "9S7-15G231-055", Color: "серый", Price: 283999.00, ProcessorModel: "Intel® Core™ i7-13700HX", GpuModel: "NVIDIA® GeForce RTX™ 4060", Ram: "16,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 16, Model: "Creator Z16HXStudio B13VGTO-056RU", Line: "Creator Pro", ManufacturerCode: "9S7-15G231-056", Color: "серый", Price: 309999.00, ProcessorModel: "Intel® Core™ i7-13700HX", GpuModel: "NVIDIA® GeForce RTX™ 4060", Ram: "16,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 17, Model: "Creator A16 AI+ A3HVFG-279RU", Line: "Creator Pro", ManufacturerCode: "9S7-15FK14-279", Color: "серый", Price: 331990.00, ProcessorModel: "AMD Ryzen™ AI 9 365", GpuModel: "NVIDIA® GeForce RTX™ 4060", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 18, Model: "Creator A16 AI+ A3HVFG-282RU", Line: "Creator Pro", ManufacturerCode: "9S7-15FK14-282", Color: "серый", Price: 331990.00, ProcessorModel: "AMD Ryzen™ AI 9 365", GpuModel: "NVIDIA® GeForce RTX™ 4060", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 19, Model: "Creator 16 AI Studio A1VHG-051RU", Line: "Creator Pro", ManufacturerCode: "9S7-15F314-051", Color: "серый", Price: 279999.00, ProcessorModel: "Intel® Core™ Ultra 9 185H", GpuModel: "NVIDIA RTX™ 3000 Ada", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 20, Model: "Creator 16 AI Studio A1VIG-060RU", Line: "Creator Pro", ManufacturerCode: "9S7-15F314-060", Color: "серый", Price: 445990.00, ProcessorModel: "Intel® Core™ Ultra 9 185H", GpuModel: "NVIDIA RTX™ 3000 Ada", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 21, Model: "Creator Z17 HX Studio A14VGT-267RU", Line: "Creator Pro", ManufacturerCode: "9S7-17N212-267", Color: "серый", Price: 229999.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 5000", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 22, Model: "Creator Z17HXStudio A13VGT-039RU", Line: "Creator Pro", ManufacturerCode: "9S7-17N212-039", Color: "серый", Price: 279999.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 5000", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 23, Model: "Creator Z17HXStudio A13VFT-063RU", Line: "Creator Pro", ManufacturerCode: "9S7-17N212-063", Color: "серый", Price: 390999.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 5000", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 24, Model: "Creator Z17HXStudio A13VGT-230RU", Line: "Creator Pro", ManufacturerCode: "9S7-17N212-230", Color: "серый", Price: 390999.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 5000", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 25, Model: "Creator Z17HXStudio A13VGT-049RU", Line: "Creator Pro", ManufacturerCode: "9S7-17N212-049", Color: "серый", Price: 390999.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 5000", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 26, Model: "Creator Z17HXStudio A13VFT-220RU", Line: "Creator Pro", ManufacturerCode: "9S7-17N212-220", Color: "серый", Price: 390999.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 5000", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 27, Model: "Creator M14 A13VF-089RU", Line: "Creator Pro", ManufacturerCode: "9S7-14P112-089", Color: "серый", Price: 390999.00, ProcessorModel: "Intel® Core™ i9 14900HX", GpuModel: "NVIDIA RTX™ 5000", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 41, Model: "Creator A16 AI+ A3HVGG-268RU ", Line: "Creator", ManufacturerCode: "9S7-15FK14-268", Color: "серый", Price: 339999.00, ProcessorModel: "AMD Ryzen™ AI 9 365", GpuModel: "NVIDIA® GeForce RTX™ 4070", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 42, Model: "Creator A16 AI+ A3XVGG-095RU", Line: "Creator", ManufacturerCode: "9S7-15FK12-095", Color: "серый", Price: 343999.00, ProcessorModel: "AMD Ryzen™ AI 9 365", GpuModel: "NVIDIA® GeForce RTX™ 4070", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 43, Model: "Creator A16 AI+ A3XVGG-209RU", Line: "Creator", ManufacturerCode: "9S7-15FK13-209", Color: "серый", Price: 349999.00, ProcessorModel: "AMD Ryzen™ AI 9 365", GpuModel: "NVIDIA® GeForce RTX™ 4070", Ram: "32,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 44, Model: "Stealth 18 HX AI A2XWJG-022RU", Line: "Stealth", ManufacturerCode: "9S7-183341-022", Color: "черный", Price: 449900.00, ProcessorModel: "Intel® Core™ Ultra 9 275HX", GpuModel: "NVIDIA® GeForce RTX™ 5090", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 45, Model: "Stealth 18 HX AI A2XWJG-052RU", Line: "Stealth", ManufacturerCode: "9S7-183341-052", Color: "черный", Price: 449900.00, ProcessorModel: "Intel® Core™ Ultra 9 275HX", GpuModel: "NVIDIA® GeForce RTX™ 5090", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 46, Model: "Stealth 18 HX AI A2XWIG-023RU", Line: "Stealth", ManufacturerCode: "9S7-183341-023", Color: "черный", Price: 414650.00, ProcessorModel: "Intel® Core™ Ultra 9 275HX", GpuModel: "NVIDIA® GeForce RTX™ 5080", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 47, Model: "Stealth 18 HX AI A2XWIG-051RU", Line: "Stealth", ManufacturerCode: "9S7-183341-051", Color: "черный", Price: 414650.00, ProcessorModel: "Intel® Core™ Ultra 9 275HX", GpuModel: "NVIDIA® GeForce RTX™ 5080", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
	{ID: 48, Model: "Stealth 18 HX AI A2XWIG-066RU", Line: "Stealth", ManufacturerCode: "9S7-183341-066", Color: "черный", Price: 414650.00, ProcessorModel: "Intel® Core™ Ultra 9 275HX", GpuModel: "NVIDIA® GeForce RTX™ 5080", Ram: "64,00", ScreenType: "IPS", DeltaE: 0.68, ColorCoverage: "DCI-P3 (100%) в среднем", ScreenResolution: "2560x1600", ContrastRatio: "1200:1"},
}

func getCatalog(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, catalog)
}

func main() {
    router := gin.Default()
    router.GET("/catalog", getCatalog)

    router.Run("localhost:8080")
}