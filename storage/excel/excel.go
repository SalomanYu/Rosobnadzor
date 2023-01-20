package excel

import (
	"Rosobnadzor/logger"
	"Rosobnadzor/models"
	"fmt"

	"github.com/xuri/excelize/v2"
)

const ExcelFile = "vuzes.xlsx"
const SheetName = "Sheet1"
var VuzBook *excelize.File
// var row = 12031

func init() {
	// VuzBook = excelize.NewFile()
	VuzBook, err := excelize.OpenFile(ExcelFile)

	if err != nil {
		VuzBook = excelize.NewFile()
	}
	defer VuzBook.Close()

	VuzBook.SetCellValue(SheetName, fmt.Sprintf("A%d", 0), "ОГРН")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("B%d", 0), "Решение о предоставлении")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("C%d", 0), "Текущий статус лицензии")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("D%d", 0), "Полное наименование организации (ФИО индивидуального предпринимателя)")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("E%d", 0), "Наименование органа, выдавшего лицензию")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("F%d", 0), "Срок действия")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("G%d", 0), "Субьект РФ")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("H%d", 0), "Сокращенное наименование организации")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("I%d", 0), "ИНН")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("J%d", 0), "КПП")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("K%d", 0), "Регистрационный номер лицензии")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("L%d", 0), "Место нахождения организации")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("M%d", 0), "Решения лицензирующего органа о приостановлении действия")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("N%d", 0), "Решения лицензирующего органа о возобновлении действия")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("O%d", 0), "Основание и дата прекращения действия")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("P%d", 0), "Решения суда об аннулировании лицензии")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("Q%d", 0), "Информация о должностном лице")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("R%d", 0), "Дата внесения изменений")
	if err := VuzBook.SaveAs(ExcelFile); err != nil{
		panic(err)
	}
}

func AddVuz(vuz *models.Vuz) {
	VuzBook, err := excelize.OpenFile(ExcelFile)
	rows, err := VuzBook.GetRows(VuzBook.GetSheetList()[0])
	if err != nil {
		panic(err)
	}
	rowsCount := len(rows) + 1
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("A%d", rowsCount), vuz.MainStateRegistrationNumber)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("B%d", rowsCount), vuz.DecisionOnGranting)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("C%d", rowsCount), vuz.LisenceStatus)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("D%d", rowsCount), vuz.FullName)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("E%d", rowsCount), vuz.NameOfTheAuthorizyThatIssuedTheLicence)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("F%d", rowsCount), vuz.ValidPeriod)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("G%d", rowsCount), vuz.SubjectOfTheRussia)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("H%d", rowsCount), vuz.ShortName)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("I%d", rowsCount), vuz.INN)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("J%d", rowsCount), vuz.KPP)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("K%d", rowsCount), vuz.LicenceRegisterNumber)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("L%d", rowsCount), vuz.Address)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("M%d", rowsCount), vuz.DecisionToSuspendTheLicence)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("N%d", rowsCount), vuz.DecisionOnLicenceRenewal)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("O%d", rowsCount), vuz.BasisAndDateOfTermination)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("P%d", rowsCount), vuz.CourtDecisionsOnLicenseCancellation)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("Q%d", rowsCount), vuz.InfoAboutTheOfficial)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("R%d", rowsCount), vuz.DateOfChanges)

	if err := VuzBook.SaveAs(ExcelFile); err != nil{
		panic(err)
	}
	logger.Log.Printf("[row:%d] Saved vuz: %s\n", rowsCount, vuz.FullName)

}