package models

type Vuz struct {
	MainStateRegistrationNumber				int
    DecisionOnGranting 						string
    LisenceStatus							string
    FullName								string
    NameOfTheAuthorizyThatIssuedTheLicence 	string
    ValidPeriod								string
    SubjectOfTheRussia						string
    ShortName								string
    INN										int
    KPP										int
    LicenceRegisterNumber					string
    Address									string
    DecisionToSuspendTheLicence				string
    DecisionOnLicenceRenewal				string
    BasisAndDateOfTermination				string
    CourtDecisionsOnLicenseCancellation		string
    InfoAboutTheOfficial					string
    DateOfChanges							string
}

type TableRow struct {
	Title									string
	Text									string
}