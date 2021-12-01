package cache

import (
	"time"
)

type Evr struct {
	Epoch   int
	Version string
	Release string
}

type Nevra struct {
	NameId int
	EvrId  int
	ArchId int
}

type Repo struct {
	Label      string
	Url        string
	BaseArch   *string
	ReleaseVer *string
}

type ErrataDetail struct {
	ID           int
	Synopsis     string
	Summary      *string
	Type         string
	Severity     *string
	Description  *string
	CVEs         []int
	PkgIds       []int
	ModulePkgIds []int
	Bugzillas    []string
	Refs         []string
	Modules      []Module
	Solution     *string
	Issued       *time.Time
	Updated      *time.Time
	Url          string
}

type PkgErrata struct {
	PkgId    int
	ErrataId int
}

type Module struct {
	Name              string
	Stream            string
	Version           string
	Context           string
	PackageList       []string
	SourcePackageList []string
}

type ModuleStream struct {
	Module string
	Stream string
}

type DbChange struct {
	ErrataChanges time.Time
	CveChanges    time.Time
	RepoChanges   time.Time
	LastChange    time.Time
	Exported      time.Time
}

type Cache struct {
	Packagename2Id      map[string]int
	Id2Packagename      map[int]string

	Evr2Id              map[Evr]int
	Id2Evr              map[int]Evr

	Arch2Id             map[string]int
	Id2Arch             map[int]string
	ArchCompat          map[int]*intSet

	PkgId2Nevra         map[int]Nevra

	Id2Repo             map[int]Repo
	RepoLabel2Ids       map[string][]int
	PkgId2RepoIds       map[int][]int

	ErrataDetail        map[string]ErrataDetail // reduce?
	ErrataId2Name       map[int]string
	PkgId2ErrataIds     map[int][]int
	ErrataId2RepoIds    map[int][]int

	ModuleName2Ids      map[ModuleStream][]int
	ModuleId2RequireIds map[int][]int
	PkgErrata2Module    map[PkgErrata][]int

	Updates             map[int][]int
	UpdatesIndex        map[int]map[int][]int

	DbChange            DbChange
}
