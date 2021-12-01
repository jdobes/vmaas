package main

import (
    "fmt"
    "vmaas-webapp/cache"
)

func main() {
    cache.LoadCache("/Users/jdobes/work/scripts/vmaas.db")
    if cache.C.DbChange.Exported.IsZero() {
        fmt.Println("Packagename2Id:", len(cache.C.Packagename2Id))
        fmt.Println("Id2Packagename:", len(cache.C.Id2Packagename))
        fmt.Println("Evr2Id:", len(cache.C.Evr2Id))
        fmt.Println("Id2Evr:", len(cache.C.Id2Evr))
        fmt.Println("Arch2Id:", len(cache.C.Arch2Id))
        fmt.Println("Id2Arch:", len(cache.C.Id2Arch))
        fmt.Println("ArchCompat:", len(cache.C.ArchCompat))
        fmt.Println("noarch compat:", cache.C.ArchCompat[cache.C.Arch2Id["noarch"]].Len())
        fmt.Println("x86_64 compat:", cache.C.ArchCompat[cache.C.Arch2Id["x86_64"]].Len())
        fmt.Println("PkgId2Nevra:", len(cache.C.PkgId2Nevra))
        fmt.Println("Id2Repo:", len(cache.C.Id2Repo))
        fmt.Println("repo basearch nil:", cache.C.Id2Repo[25400].BaseArch == nil)
    }
}
