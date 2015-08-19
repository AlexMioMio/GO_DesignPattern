package main
import "fmt"

type Person struct {
    Name string
}
func (p *Person) show() {
    fmt.Printf("装扮人 %s -> ", p.Name)

}

func (p *Person) hi() {
    fmt.Printf(" hi")

}

type AbstractPerson interface {
    show()
}
type Finery struct {
    AbstractPerson
}
func (f *Finery) Decorate(component AbstractPerson) {
    f.AbstractPerson =component
}

func (f *Finery) show() {
    if f.AbstractPerson !=nil {
        f.AbstractPerson.show()
    }
}
// T恤
type TShirts struct {
    Finery
}
func (t *TShirts) show() {
    t.Finery.show()
    fmt.Print(" 大T恤")
}
// 大裤衩
type BigTrouser struct {
    Finery
}

func (b *BigTrouser) show() {
    b.Finery.show()
    fmt.Print(" 大裤衩")
}
// 破球鞋
type Sneakers struct {
    Finery
}
func (s *Sneakers) show() {
    s.Finery.show()
    fmt.Print(" 破球鞋")
}
// 皮鞋
type LeatherShones struct {
    Finery
}
func (l *LeatherShones) show(){
    l.Finery.show()
    fmt.Print(" 皮鞋")
}
// 西装
type Suit struct {
    Finery
}
func (s *Suit) show(){
    s.Finery.show()
    fmt.Print(" 西装")
}
// 领带
type Tie struct {
    Finery
}
func (t *Tie) show(){
    t.Finery.show()
    fmt.Print(" 领带")
}

func main(){

    person :=&Person{"屌丝"}

    fmt.Println("第一种装扮: ")
    pqx :=new(Sneakers)
    kk := new(BigTrouser)
    dtk :=new(TShirts)
    pqx.Decorate(person)
    kk.Decorate(pqx)
    dtk.Decorate(kk)
    dtk.show()


   fmt.Println("")

    fmt.Println("第二种装扮")
    px :=new(LeatherShones)
    ld :=new(Tie)
    xz :=new(Suit)
    px.Decorate(person)
    ld.Decorate(px)
    xz.Decorate(ld)
    xz.show()

    fmt.Println("")

    fmt.Println("第三种装扮")
    pqx2 := new(Sneakers)
    px2 := new(LeatherShones)
    kk2 := new(BigTrouser)
    ld2 := new(Tie)
    pqx2.Decorate(person)
    px2.Decorate(pqx2)
    kk2.Decorate(px2)
    ld2.Decorate(kk2)
    ld2.show()





}
