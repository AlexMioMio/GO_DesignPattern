package main
import "fmt"
type SchoolGirl struct {
    name string
}
func (s *SchoolGirl) setName(name string) {
    s.name = name
}
func (s *SchoolGirl) getName(name string) {
    s.name = name
}
type Pursuit struct {
    mm SchoolGirl
}
type GiveGift interface {
    giveDolls()
    giveFlowers()
    giveChocolate()
}

func (p *Pursuit) giveDolls() {
    fmt.Println(p.mm.name, "送你洋娃娃")
}
func (p *Pursuit) giveFlowers() {
    fmt.Println(p.mm.name, "送你鲜花")
}
func (p *Pursuit)  giveChocolate() {
    fmt.Println(p.mm.name, "送你巧克力")
}

// 代理人
type Proxy struct {
    gg Pursuit
}
// 创建代理人
func NewProxy(mm SchoolGirl) *Proxy {
    gg := Pursuit{mm}
    return &Proxy{gg}
}
func (p *Proxy) giveDolls() {
    p.gg.giveDolls()
}
func (p *Proxy) giveFlowers() {
    p.gg.giveFlowers()
}
func (p *Proxy)  giveChocolate() {
    p.gg.giveChocolate()
}

func main() {
    mm := SchoolGirl{}
    mm.setName("美女")

    proxyMen := NewProxy(mm)
    proxyMen.giveDolls()
    proxyMen.giveFlowers()
    proxyMen.giveChocolate()

}

