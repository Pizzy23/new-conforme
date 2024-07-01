package db

type ComplianceStatus int
type UserRole int
type Category string

const (
	Conforme ComplianceStatus = iota
	NC
)

func (c ComplianceStatus) String() string {
	switch c {
	case Conforme:
		return "C"
	case NC:
		return "NC"
	default:
		return "Unknown"
	}
}

func ComplianceStatusFromString(s string) ComplianceStatus {
	switch s {
	case "C":
		return Conforme
	case "NC":
		return NC
	default:
		return -1
	}
}

const (
	Manager UserRole = iota
	ResponsibleProfessional
	ResponsiblePIE
	AuthorizedProfessional
	Owner
)

func (role UserRole) String() string {
	switch role {
	case Manager:
		return "Plant Manager"
	case ResponsibleProfessional:
		return "Responsible Professional"
	case ResponsiblePIE:
		return "Responsible for PIE"
	case AuthorizedProfessional:
		return "Authorized Professional"
	case Owner:
		return "Owner"
	default:
		return "Unknown"
	}
}

func ParseUserRole(id uint64) UserRole {
	switch id {
	case 10:
		return Manager
	case 11:
		return ResponsibleProfessional
	case 12:
		return ResponsiblePIE
	case 13:
		return AuthorizedProfessional
	case 14:
		return Owner
	default:
		return -1
	}
}

const (
	RiskManagement      Category = "RiskManagement"
	GeoElectricalSurvey Category = "GeoElectricalSurvey"
	SpdaMemorial        Category = "SpdaMemorial"
	MpsMemorial         Category = "MpsMemorial"
	MaterialsList       Category = "MaterialsList"
	Drawings            Category = "Drawings"
	InitialInspection   Category = "InitialInspection"
	FinalVerification   Category = "FinalVerification"
	PeriodicInspection  Category = "PeriodicInspection"
	Mythology           Category = "Mythology"
)

func (c Category) String() string {
	switch c {
	case RiskManagement:
		return "RiskManagement"
	case GeoElectricalSurvey:
		return "GeoElectricalSurvey"
	case SpdaMemorial:
		return "SpdaMemorial"
	case MpsMemorial:
		return "MpsMemorial"
	case MaterialsList:
		return "MaterialsList"
	case Drawings:
		return "Drawings"
	case InitialInspection:
		return "InitialInspection"
	case FinalVerification:
		return "FinalVerification"
	case PeriodicInspection:
		return "PeriodicInspection"
	case Mythology:
		return "Mythology"
	default:
		return "N/A"
	}
}
