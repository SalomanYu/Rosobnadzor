package excel

import (
	"fmt"
	"Rosobnadzor/models"
	"github.com/xuri/excelize/v2"
)

const ExcelFile = "vuzes.xlsx"
const SheetName = "Sheet1"
var VuzBook *excelize.File
var row int

func init() {
	VuzBook = excelize.NewFile()
	defer VuzBook.Close()
	row++
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("A%d", row), "ОГРН")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("B%d", row), "Решение о предоставлении")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("C%d", row), "Текущий статус лицензии")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("D%d", row), "Полное наименование организации (ФИО индивидуального предпринимателя)")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("E%d", row), "Наименование органа, выдавшего лицензию")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("F%d", row), "Срок действия")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("G%d", row), "Субьект РФ")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("H%d", row), "Сокращенное наименование организации")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("I%d", row), "ИНН")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("J%d", row), "КПП")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("K%d", row), "Регистрационный номер лицензии")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("L%d", row), "Место нахождения организации")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("M%d", row), "Решения лицензирующего органа о приостановлении действия")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("N%d", row), "Решения лицензирующего органа о возобновлении действия")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("O%d", row), "Основание и дата прекращения действия")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("P%d", row), "Решения суда об аннулировании лицензии")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("Q%d", row), "Информация о должностном лице")
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("R%d", row), "Дата внесения изменений")
	if err := VuzBook.SaveAs(ExcelFile); err != nil{
		panic(err)
	}
}

func AddVuz(vuz *models.Vuz) {
	row++
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("A%d", row), vuz.MainStateRegistrationNumber)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("B%d", row), vuz.DecisionOnGranting)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("C%d", row), vuz.LisenceStatus)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("D%d", row), vuz.FullName)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("E%d", row), vuz.NameOfTheAuthorizyThatIssuedTheLicence)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("F%d", row), vuz.ValidPeriod)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("G%d", row), vuz.SubjectOfTheRussia)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("H%d", row), vuz.ShortName)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("I%d", row), vuz.INN)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("J%d", row), vuz.KPP)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("K%d", row), vuz.LicenceRegisterNumber)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("L%d", row), vuz.Address)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("M%d", row), vuz.DecisionToSuspendTheLicence)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("N%d", row), vuz.DecisionOnLicenceRenewal)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("O%d", row), vuz.BasisAndDateOfTermination)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("P%d", row), vuz.CourtDecisionsOnLicenseCancellation)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("Q%d", row), vuz.InfoAboutTheOfficial)
	VuzBook.SetCellValue(SheetName, fmt.Sprintf("R%d", row), vuz.DateOfChanges)

	if err := VuzBook.SaveAs(ExcelFile); err != nil{
		panic(err)
	}
}