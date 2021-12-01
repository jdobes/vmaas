package cache

import (
    "fmt"
    "strings"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var (
    db *sql.DB
    C  Cache = Cache{}
)

func init () {
    C.Packagename2Id = map[string]int{}
    C.Id2Packagename = map[int]string{}
    C.Evr2Id = map[Evr]int{}
    C.Id2Evr = map[int]Evr{}
    C.Arch2Id = map[string]int{}
    C.Id2Arch = map[int]string{}
    C.ArchCompat = map[int]*intSet{}
    C.PkgId2Nevra = map[int]Nevra{}
    C.Id2Repo = map[int]Repo{}
}

func openDb(filePath string) {
    tmpDb, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}
	db = tmpDb
}

func getRows(table string, cols []string) *sql.Rows {
    rows, err := db.Query(fmt.Sprintf("select %s from %s", strings.Join(cols, ","), table))
    if err != nil {
        panic(err)
    }
    return rows
}

func loadPkgNames() {
    rows := getRows("packagename", []string{"id", "packagename"})
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            panic(err)
        }
        C.Packagename2Id[name] = id
        C.Id2Packagename[id] = name
    }
    if err := rows.Err(); err != nil {
        panic(err)
    }
}

func loadEvrs() {
    rows := getRows("evr", []string{"id", "epoch", "version", "release"})
    defer rows.Close()

    for rows.Next() {
        var id int
        e := Evr{}
        if err := rows.Scan(&id, &e.Epoch, &e.Version, &e.Release); err != nil {
            panic(err)
        }
        C.Evr2Id[e] = id
        C.Id2Evr[id] = e
    }
    if err := rows.Err(); err != nil {
        panic(err)
    }
}

func loadArchs() {
    rows := getRows("arch", []string{"id", "arch"})
    defer rows.Close()

    for rows.Next() {
        var id int
        var arch string
        if err := rows.Scan(&id, &arch); err != nil {
            panic(err)
        }
        C.Arch2Id[arch] = id
        C.Id2Arch[id] = arch
    }
    if err := rows.Err(); err != nil {
        panic(err)
    }
}

func loadArchCompat() {
    rows := getRows("arch_compat", []string{"from_arch_id", "to_arch_id"})
    defer rows.Close()

    for rows.Next() {
        var from_arch_id, to_arch_id int
        if err := rows.Scan(&from_arch_id, &to_arch_id); err != nil {
            panic(err)
        }
        if _, ok := C.ArchCompat[from_arch_id]; !ok {
            C.ArchCompat[from_arch_id] = NewIntSet()
        }
        C.ArchCompat[from_arch_id].Add(to_arch_id)
    }
    if err := rows.Err(); err != nil {
        panic(err)
    }
}

func loadPkgNevras() {
    rows := getRows("package_detail", []string{"id", "name_id", "evr_id", "arch_id"})
    defer rows.Close()

    for rows.Next() {
        var id int
        n := Nevra{}
        if err := rows.Scan(&id, &n.NameId, &n.EvrId, &n.ArchId); err != nil {
            panic(err)
        }
        C.PkgId2Nevra[id] = n
    }
    if err := rows.Err(); err != nil {
        panic(err)
    }
}

func loadRepos() {
    rows := getRows("repo_detail", []string{"id", "label", "url", "basearch", "releasever"})
    defer rows.Close()

    for rows.Next() {
        var id int
        r := Repo{}
        if err := rows.Scan(&id, &r.Label, &r.Url, &r.BaseArch, &r.ReleaseVer); err != nil {
            panic(err)
        }
        C.Id2Repo[id] = r
    }
    if err := rows.Err(); err != nil {
        panic(err)
    }
}

func LoadCache(filePath string) {
    openDb(filePath)
    defer db.Close()
	loadPkgNames()
    loadEvrs()
    loadArchs()
    loadArchCompat()
    loadPkgNevras()
    loadRepos()
}
