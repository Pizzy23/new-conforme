package db

type Company struct {
	ID   uint64 `json:"id" xorm:"'id' notnull pk autoincr"`
	Name string `json:"name" xorm:"'name' VARCHAR(100) notnull"`
	Area string `json:"area" xorm:"'area' VARCHAR(100) notnull"`
	Type string `json:"type" xorm:"'type' VARCHAR(100) notnull"`
	Cnpj string `json:"cnpj" xorm:"'cnpj' VARCHAR(100) notnull"`
}

type Word struct {
	ID         uint64 `json:"id" xorm:"'id' notnull pk autoincr"`
	Word       string `json:"word" xorm:"'word' notnull"`
	Categories string `json:"category" xorm:"'category'"`
}

type NotConform struct {
	Id          uint64 `json:"id" xorm:"'id' notnull pk autoincr"`
	Number      int    `json:"number" xorm:"'number'"`
	Title       string `json:"title" xorm:"'title'"`
	Desc        string `json:"desc" xorm:"'desc'"`
	Tech        string `json:"tech" xorm:"'tech'"`
	Legal       string `json:"legal" xorm:"'legal'"`
	Recommended string `json:"recommended" xorm:"'recommended'"`
}

type Painel struct {
	Id          uint64 `json:"id" xorm:"'id' notnull pk autoincr"`
	SectorID    uint   `json:"sector_id" xorm:"'sector_id' notnull"`
	Name        string `json:"name" xorm:"'name' VARCHAR(100) notnull"`
	Number      string `json:"number" xorm:"'number' VARCHAR(50) notnull"`
	Review      string `json:"review" xorm:"'review' VARCHAR(255) notnull"`
	Description string `json:"description" xorm:"'description' VARCHAR(255) notnull"`
	PanelType   string `json:"panel_type" xorm:"'panel_type' VARCHAR(50) notnull"`
	ControlCopy string `json:"control_copy" xorm:"'control_copy' VARCHAR(50) notnull"`
}

type Pda struct {
	ID          uint64 `json:"id" xorm:"'id' notnull pk autoincr"`
	SectorID    uint64 `json:"sector_id" xorm:"'sector_id'"`
	CompanyID   uint64 `json:"company_id" xorm:"'company_id'"`
	Number      int    `json:"number" xorm:"'number'"`
	Date        string `json:"date" xorm:"'date'"`
	Responsible string `json:"responsible" xorm:"'responsible'"`
	Description string `json:"description" xorm:"'description'"`
	Process     string `json:"process" xorm:"'process'"`
	Pdf         uint64 `json:"pdf" xorm:"'idPDF'"`
}

type Pdf struct {
	IdPdf   uint64 `json:"id" xorm:"'idPDF' notnull pk autoincr"`
	Content []byte `json:"pdfContent" xorm:"'content' LONGBLOB"`
	Status  bool   `json:"status" xorm:"'status'"`
}

type User struct {
	ID        uint64 `json:"id" xorm:"'id' notnull pk autoincr"`
	CompanyID uint64 `json:"company_id" xorm:"'company_id'"`
	Name      string `json:"name" xorm:"'name'"`
	Email     string `json:"email" xorm:"'email'"`
	Password  string `json:"password" xorm:"'password'"`
	Type      string `json:"type" xorm:"'type'"`
	IsLogged  bool   `json:"isLogged" xorm:"'is_logged'"`
}

type Sector struct {
	ID          uint64 `json:"id" xorm:"'id' notnull pk autoincr"`
	CompanyID   uint64 `json:"company_id" xorm:"'company_id'"`
	Name        string `json:"name" xorm:"'name'"`
	Number      string `json:"number" xorm:"'number'"`
	Description string `json:"description" xorm:"'description'"`
	Content     []byte `json:"content" xorm:"'content' BLOB"`
}

type SectorTest struct {
	ID          uint64 `json:"id" xorm:"'id' notnull pk autoincr"`
	Name        string `json:"name" xorm:"'name'"`
	Number      string `json:"number" xorm:"'number'"`
	Description string `json:"description" xorm:"'description'"`
	Content     []byte `json:"content" xorm:"'content' BLOB"`
}

type PdfTest struct {
	ID       uint64 `json:"id" xorm:"'id' notnull pk autoincr"`
	FileName string `json:"fileName" xorm:"'file_name'"`
	Content  []byte `json:"content" xorm:"'content' BLOB"`
}
