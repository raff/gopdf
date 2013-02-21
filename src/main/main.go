package main

import (
	"fmt"
	"gopdf"
	 iconv "github.com/djimenez/iconv-go"
	 "gopdf/fonts"
)

func main() {
	fmt.Println("start...")
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 A4
	pdf.AddFont("THSarabunPSK",new(fonts.THSarabun),"res/fonts/THSarabun.z")
	pdf.AddFont("Loma",new(fonts.Loma),"res/fonts/Loma.z")
	pdf.AddPage()
	pdf.SetFont("THSarabunPSK", "B", 14)
	pdf.Cell(nil,  ToCp874("กAโจตลาด  2 ล้อ พุ่ง 20% รับปีใหม่ คาดเอที toyota ยังแรงกุ้งตั้ว "))
	pdf.Ln(28)
	pdf.SetFont("THSarabunPSK", "B", 16)
	pdf.Cell(nil,  ToCp874("ด้วยการที่เราไม่ไปทำยังไง"))
	pdf.Cell(nil,  ToCp874("ศาลอาญา ไม่ชี้ชัด ใครยิง “ มานะ อาจราญ” เจ้าหน้าที่สวนสัตว์ดุสิตเสียชีวิต"))
	
	
	pdf.AddPage()
	pdf.SetFont("Loma", "B", 12)
	pdf.Cell(&gopdf.Rect{H: 100, W: 100}, ToCp874("การบ้านx"))
	pdf.SetFont("THSarabunPSK", "B", 12)
	pdf.Cell(&gopdf.Rect{H: 100, W: 100}, ToCp874("การบ้านx") )
	pdf.Ln(14)
	pdf.Cell(&gopdf.Rect{H: 100, W: 100}, ToCp874("การบ้านx"))
	pdf.Cell(&gopdf.Rect{H: 100, W: 100}, ToCp874("การบ้านx") )
	
	
	pdf.WritePdf("/var/www/fpdf17/output/x.pdf")
	fmt.Println("end...")
}

func ToCp874(str string) string{
	str, _ = iconv.ConvertString( str, "utf-8", "cp874") 
	return  str
}