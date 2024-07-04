package interfaces

type UserInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CompanyName string `json:"companyName" binding:"required"`
	Office      int    `json:"office" binding:"required"`
}

type CompanyInput struct {
	Company  string `json:"company" binding:"required"`
	Area     string `json:"area" binding:"required"`
	Type     string `json:"type" binding:"required"`
	CNPJ     string `json:"cnpj" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SectorInput struct {
	CompanyID   uint64 `json:"company_id" xorm:"'company_id'"`
	Name        string `json:"name" xorm:"'name'"`
	Number      string `json:"number" xorm:"'number'"`
	Description string `json:"description" xorm:"'description'"`
	Content     []byte `json:"content" xorm:"'content BLOB'"`
}
type SectorInputTest struct {
	Name        string `json:"name" xorm:"'name'"`
	Number      string `json:"number" xorm:"'number'"`
	Description string `json:"description" xorm:"'description'"`
	Content     []byte `json:"content" xorm:"'content BLOB'"`
}

type UserInputWithHashedPassword struct {
	UserInput
	HashedPassword string `json:"hashedpassword"`
}

type PainelInput struct {
	SectorID    uint   `json:"sector_id" xorm:"'sector_id' notnull"`
	Name        string `json:"name" xorm:"'name' VARCHAR(100) notnull"`
	Number      string `json:"number" xorm:"'number' VARCHAR(50) notnull"`
	Review      string `json:"review" xorm:"'review' VARCHAR(255) notnull"`
	Description string `json:"description" xorm:"'description' VARCHAR(255) notnull"`
	PanelType   string `json:"panel_type" xorm:"'panel_type' VARCHAR(50) notnull"`
	ControlCopy string `json:"control_copy" xorm:"'control_copy' VARCHAR(50) notnull"`
}

type NotConformInput struct {
	Number      int    `json:"Number"`
	Title       string `json:"Title"`
	Desc        string `json:"Desc"`
	Tech        string `json:"Tech"`
	Legal       string `json:"Legal"`
	Recommended string `json:"Recomended"`
}

type NotConformOutput struct {
	Id          uint64 `json:"id"`
	Number      int    `json:"Number"`
	Title       string `json:"Title"`
	Desc        string `json:"Desc"`
	Tech        string `json:"Tech"`
	Legal       string `json:"Legal"`
	Recommended string `json:"Recomended"`
}

type GetNC struct {
	Number int `json:"Number"`
}
