package scraper

import (
	"Rosobnadzor/logger"
	"Rosobnadzor/models"
	"Rosobnadzor/storage/excel"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

const PageDomain = "https://islod.obrnadzor.gov.ru/rlic/search/?page="
const VuzDomain = "https://islod.obrnadzor.gov.ru/rlic/details"

func GetVuzesId(pageNum int) (ids []string) {
	body, err := getBody(fmt.Sprintf("%s%d", PageDomain, pageNum))
	if err != nil{
		log.Println("error 10 seconds")
		time.Sleep(10*time.Second)
		body, err = getBody(fmt.Sprintf("%s%d", PageDomain, pageNum))
		if err != nil{
			log.Fatal(err)
		}
	}
	body.ForEach("table.search-result.table.table-hover tbody tr", func(i int, h *colly.HTMLElement) {
		ids = append(ids, h.Attr("data-guid"))
	})
	return
}

func SaveVuzes(vuzIds []string) {
	for _, id := range vuzIds{
		vuz := GetVuz(id)
		if vuz.FullName == "" {
			continue 
		}
		excel.AddVuz(&vuz)
		logger.Log.Printf("Saved vuz: %s\n", vuz.FullName)
	}
}

func GetVuz(vuzId string) (vuz models.Vuz) {
	url := fmt.Sprintf("%s/%s/", VuzDomain, vuzId)
	body, err := getBody(url)
	if err != nil{
		logger.Log.Printf("Failed to connect to the page the first time: %s\nTimeout 10 seconds\n", url)
		log.Println("error 10 seconds")
		time.Sleep(10*time.Second)
		body, err = getBody(url)
		if err != nil{
			logger.Log.Printf("Server error 502:%s\n", url)
			if err.Error() == "Server error 502" {
				return
			}
			log.Fatal(err)
		}
	}
	body.ForEach("table.table.table-bordered tr", func(i int, row *colly.HTMLElement) {
		tableRow := splitTableRow(row)
		switch tableRow.Title {
		case "Регистрационный номер лицензии": vuz.LicenceRegisterNumber = tableRow.Text
		case "Решение о предоставлении": vuz.DecisionOnGranting = tableRow.Text
		case "Текущий статус лицензии": vuz.LisenceStatus = tableRow.Text
		case "Полное наименование организации (ФИО индивидуального предпринимателя)": vuz.FullName = tableRow.Text
		case "Наименование органа, выдавшего лицензию": vuz.NameOfTheAuthorizyThatIssuedTheLicence = tableRow.Text
		case "Срок действия": vuz.ValidPeriod = tableRow.Text
		case "Субьект РФ": vuz.SubjectOfTheRussia = tableRow.Text
		case "Сокращенное наименование организации": vuz.ShortName = tableRow.Text
		case "Место нахождения организации": vuz.Address = tableRow.Text
		case "Решения лицензирующего органа о приостановлении действия": vuz.DecisionToSuspendTheLicence = tableRow.Text
		case "Решения лицензирующего органа о возобновлении действия": vuz.DecisionOnLicenceRenewal = tableRow.Text
		case "Основание и дата прекращения действия": vuz.BasisAndDateOfTermination = tableRow.Text
		case "Решения суда об аннулировании лицензии": vuz.CourtDecisionsOnLicenseCancellation = tableRow.Text
		case "Информация о должностном лице": vuz.InfoAboutTheOfficial = tableRow.Text
		case "Дата внесения изменений": vuz.DateOfChanges = tableRow.Text
		case "ОГРН": 
			digit, _ := strconv.Atoi(tableRow.Text)
			vuz.MainStateRegistrationNumber = digit
		case "ИНН": 
			digit, _ := strconv.Atoi(tableRow.Text)
			vuz.INN = digit
		case "КПП": 
			digit, _ := strconv.Atoi(tableRow.Text)
			vuz.KPP = digit
		}
	})
	return
}
