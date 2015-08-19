package main

import (

    "fmt"
)
type cashSuper interface {
    AcceptCash(float64) float64
}
type cashNormal struct {

}
// "正常收费":
func (normal *cashNormal) AcceptCash(money float64) float64 {
    return money
}
// 打折
type cashRebate struct {
    moneyRebate float64
}
func (rebate *cashRebate) AcceptCash(money float64) float64 {
    return money * rebate.moneyRebate
}

type cashReturn struct {
    moneyConditon float64
    moneyReturn float64
}
// 满多少返现多少
func (returned *cashReturn) AcceptCash(money float64) float64 {
    if money>=returned.moneyConditon {
        return money -float64(int(money/returned.moneyConditon))*returned.moneyReturn
    }else {
        return money
    }
}
type CashContext struct {
    cashSuper
}

func NewCashContext(str string) *CashContext {
    cash := new(CashContext)
    switch str{
        case "正常消费":
        cash.cashSuper= &cashNormal{}
        case "满300返100":
        cash.cashSuper =&cashReturn{300, 100}
        case "打八折":
        cash.cashSuper =&cashRebate{0.8}
    }
    return cash
}
func main() {
    var total float64
    context := NewCashContext("满300返100")
    total+= context.AcceptCash(1 *1000)
    context = NewCashContext("正常消费")
    total +=context.AcceptCash(1*1000)
    context = NewCashContext("打八折")
    total+=context.AcceptCash(1 *1000)
    fmt.Println(total)
}