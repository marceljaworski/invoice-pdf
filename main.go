package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/marceljaworski/invoice-pdf/data"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/invoice.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	// err = document.GetReport().Save("docs/assets/text/billingv2.txt")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber("Page {current} of {total}", props.RightBottom).
		WithMargins(10, 15, 10).
		Build()

	darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(getPageHeader())
	if err != nil {
		log.Fatal(err.Error())
	}

	err = m.RegisterFooter(getPageFooter())
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(text.NewRow(10, "Invoice ABC123456789", props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Align: align.Center,
	}))

	m.AddRow(7,
		text.NewCol(3, "Transactions", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: darkGrayColor})

	m.AddRows(getTransactions()...)

	m.AddRow(15,
		col.New(6).Add(
			code.NewBar("5123.151231.512314.1251251.123215", props.Barcode{
				Percent: 0,
				Proportion: props.Proportion{
					Width:  20,
					Height: 2,
				},
			}),
			text.New("5123.151231.512314.1251251.123215", props.Text{
				Top:    12,
				Family: "",
				Style:  fontstyle.Bold,
				Size:   9,
				Align:  align.Center,
			}),
		),
		col.New(6),
	)
	return m
}

func getTransactions() []core.Row {
	rows := []core.Row{
		row.New(5).Add(
			col.New(3),
			text.NewCol(1, "Product", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(2, "Description", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(3, "Price", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		),
	}

	var contentsRow []core.Row
	contents := data.FruitList(40)
	/*for i := 0; i < 8; i++ {
	    contents = append(contents, contents...)
	}*/

	for i, content := range contents {
		r := row.New(4).Add(
			col.New(3),
			text.NewCol(1, content[0], props.Text{Size: 8, Align: align.Center}),
			text.NewCol(2, content[1], props.Text{Size: 8, Align: align.Center}),
			text.NewCol(3, content[2], props.Text{Size: 8, Align: align.Center}),
		)
		if i%2 == 0 {
			gray := getGrayColor()
			r.WithStyle(&props.Cell{BackgroundColor: gray})
		}

		contentsRow = append(contentsRow, r)
	}

	rows = append(rows, contentsRow...)

	rows = append(rows, row.New(20).Add(
		col.New(7),
		text.NewCol(2, "Total:", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(3, "R$ 2.567,00", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Center,
		}),
	))

	return rows
}

func getPageHeader() core.Row {
	return row.New(20).Add(
		image.NewFromFileCol(3, "docs/assets/images/logo_div_rhino.jpg", props.Rect{
			Center:  true,
			Percent: 100,
		}),
		col.New(6),
		col.New(3).Add(
			text.New("AnyCompany Name Inc. 851 Any Street Name, Suite 120, Any City, CA 45123.", props.Text{
				Size:  8,
				Align: align.Right,
				Color: getRedColor(),
			}),
			text.New("Tel: 55 024 12345-1234", props.Text{
				Top:   12,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Right,
				Color: getDarkPurpleColor(),
			}),
			text.New("www.mycompany.com", props.Text{
				Top:   15,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Right,
				Color: getDarkPurpleColor(),
			}),
		),
	)
}

func getPageFooter() core.Row {
	return row.New(20).Add(
		col.New(12).Add(
			text.New("Tel: 55 024 12345-1234", props.Text{
				Top:   13,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getDarkPurpleColor(),
			}),
			text.New("www.mycompany.com", props.Text{
				Top:   16,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getDarkPurpleColor(),
			}),
		),
	)
}

func getDarkGrayColor() *props.Color {
	return &props.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getDarkPurpleColor() *props.Color {
	return &props.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getRedColor() *props.Color {
	return &props.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}
