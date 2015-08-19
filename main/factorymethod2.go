package main
import (

    "fmt"
)
type LeiFeng struct {

}
func (l *LeiFeng) sweep(){
    fmt.Println("扫地")
}
func (l *LeiFeng) wash(){
    fmt.Println("洗衣")
}
func (l *LeiFeng) buyRice(){
    fmt.Println("买米")
}
// 学雷锋者 雷锋子类
type Undergraduate struct {
    LeiFeng
}
type Volunteer struct {
    LeiFeng
}
// 工厂
type Ifactory interface {
    createLeifeng() LeiFeng
}
// 雷锋子类工厂
type UndergraduateFactory struct {

}

func (u *UndergraduateFactory) createLeifeng() LeiFeng{
    return new(Undergraduate).LeiFeng
}

// 志愿者工厂
type  VolunteerFactory struct {

}
func (v *VolunteerFactory) createLeifeng() LeiFeng{
    return new(Volunteer).LeiFeng
}

func main(){
    ifac :=new(UndergraduateFactory)
    student :=ifac.createLeifeng()
    student.wash()
    student.sweep()
    student.buyRice()
}


