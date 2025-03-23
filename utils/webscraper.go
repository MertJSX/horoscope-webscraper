package utils

import (
	"github.com/gocolly/colly"
)

var url string = "https://www.edna.bg/horoskopi/" // https://www.edna.bg/horoskopi/deva/dneven

func GetHoroscopeComment(horoscope string, time string) string {
	c := colly.NewCollector()
	var element *colly.HTMLElement

	c.OnHTML("body > div.container-fluid > main > div.image-cover-layout > div > div > section > div.horoscope-inner.sticky-social > p:nth-child(4)", func(e *colly.HTMLElement) {
		element = e
	})

	err := c.Visit(url + "/" + horoscope + "/" + ConvertTimeName(time))

	if err != nil {
		return "Грешка..."
	} else {
		return element.Text
	}
}
