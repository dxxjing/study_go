type day struct {
    num       int
    shortName string
    longName  string
}
//与sort区别 本例采用结构体指针切片 且以指针作为接受者
type dayArray struct {
    data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
    Sunday    := day{0, "SUN", "Sunday"}
    Monday    := day{1, "MON", "Monday"}
    Tuesday   := day{2, "TUE", "Tuesday"}
    Wednesday := day{3, "WED", "Wednesday"}
    Thursday  := day{4, "THU", "Thursday"}
    Friday    := day{5, "FRI", "Friday"}
    Saturday  := day{6, "SAT", "Saturday"}
    data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
    a := dayArray{data}
    sort.Sort(&a)
    if !sort.IsSorted(&a) {
        panic("fail")
    }
    for _, d := range data {
        fmt.Printf("%s ", d.longName)
    }
    fmt.Printf("\n")
}

func main() {
    ints()
    strings()
    days()
}
