package main
import "fmt"

type WorkExperience struct {
    timeArea string
    compnany string
}

func (w *WorkExperience) getWorkDate() string {
    return w.timeArea
}
func (w *WorkExperience) setWorkDate(timeArea string) {
    w.timeArea =timeArea
}
func (w *WorkExperience) getCompany() string {
    return w.compnany
}
func (w *WorkExperience) setCompany(compnany string) {
    w.compnany =compnany
}

type Resume struct {
    name string
    sex string
    age string
    WorkExperience
}

func (r *Resume) setPersonalInfo(name, sex, age string) {
    r.sex =sex
    r.age =age
    r.name = name
}
func (r *Resume) setWorkExperience(timeArea, company string) {
    r.compnany =company
    r.timeArea =timeArea
}
func (r *Resume) dispaly() {
    fmt.Println(r.name, r.sex, r.age)
    fmt.Println("工作经历: ", r.timeArea, r.compnany)
}
func (r *Resume) clone() *Resume {
    new_obj := (*r)
    return &new_obj
}
func main(){
    a :=new(Resume)
    a.setPersonalInfo("员工1","男","21")
    a.setWorkExperience("1990-2014","Apple")


    b:=a.clone()
    b.setPersonalInfo("员工2","男","33")
    b.setWorkExperience("1988-2015","Google")


    c:=a.clone()
    c.setPersonalInfo("员工3","女","19")
    c.setWorkExperience("1996-2017","Logoocc")

    a.dispaly()
    b.dispaly()
    c.dispaly()

}