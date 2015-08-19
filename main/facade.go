package main
import "fmt"

// 股票1
type Strock1 struct {

}
func (s *Strock1) sell() {
    fmt.Println("股票1卖出")
}
func (s *Strock1) buy() {
    fmt.Println("股票1买进")
}


// 股票2
type Strock2 struct {

}
func (s *Strock2) sell() {
    fmt.Println("股票2卖出")
}
func (s *Strock2) buy() {
    fmt.Println("股票2买进")
}


// 股票3
type Strock3 struct {

}
func (s *Strock3) sell() {
    fmt.Println("股票1卖出")
}
func (s *Strock3) buy() {
    fmt.Println("股票1买进")
}

// 国债1
type NationalDebt1 struct {

}

func (n *NationalDebt1) sell() {
    fmt.Println("国债1卖出")
}
func (n *NationalDebt1) buy() {
    fmt.Println("国债1买进")
}

// 房地产
type Realty1 struct {

}
func (r *Realty1) sell() {
    fmt.Println("房地产1卖出")
}
func (r *Realty1) buy() {
    fmt.Println("房地产1买进")
}

// 资金
type Fund struct {
    gu1 Strock1
    gu2 Strock2
    gu3 Strock3
    nd1 NationalDebt1
    rt1 Realty1
}

// 买进
func (f *Fund) buyFund() {
    f.gu1.buy()
    f.gu2.buy()
    f.gu3.buy()
    f.rt1.buy()
    f.nd1.buy()
}
// 卖出
func (f *Fund) sellFund() {
    f.gu1.sell()
    f.gu2.sell()
    f.gu3.sell()
    f.rt1.sell()
    f.nd1.sell()
}
func NewFund() *Fund {
    return &Fund{Strock1{}, Strock2{}, Strock3{}, NationalDebt1{}, Realty1{}}
}

func main() {
    money := NewFund()
    money.buyFund()
    fmt.Println("----------------------")
    money.sellFund()
}


