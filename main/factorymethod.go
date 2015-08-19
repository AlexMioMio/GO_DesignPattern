package main

import (

    "fmt"
)
type RealOperation struct {
    numberA float64
    numberB float64
}

type OperateAdd struct {
    RealOperation
}
type OperateSub struct {
    RealOperation
}

type OperateDiv struct {
    RealOperation
}
type OperateMul struct {
    RealOperation
}

type Operation interface {
    GetResult() float64
    SetNumA(float64)
    SetNumB(float64)
}
// 工厂接口
type Ifactory interface {
    createOperation() Operation
}

// 加法工厂
type AddFactory struct {

}

func (a *AddFactory) createOperation() (operation Operation) {
    operation =new(OperateAdd)
    return
}

// 减法工厂
type SubFactory struct {

}
func (s *SubFactory) createOperation() (operation Operation) {
    operation = new(OperateSub)
    return
}

// 除法工厂
type DivFactory struct {

}
func (d *DivFactory) createOperation() (operation Operation) {
    operation =new(OperateDiv)
    return
}
// 乘法工厂
type MulFactory struct {

}
func (m *MulFactory) createOperation() (operation Operation) {
    operation =new(OperateMul)
    return
}

// 基本运算类方法
func (operation *RealOperation) SetNumA(numA float64) {
    operation.numberA = numA
}

func (operation *RealOperation) SetNumB(numB float64) {
    operation.numberB = numB
}

// 为四则运算类添加获取结果的方法

func (operation *OperateAdd) GetResult() float64{
 return operation.numberA + operation.numberB
}

func (operation *OperateSub) GetResult() float64{
    return  operation.numberA - operation.numberB
}
func (operation *OperateMul) GetResult() float64{
    return operation.numberA * operation.numberB
}
func (operation *OperateDiv) GetResult() float64{
    if operation.numberB == 0{
        panic("除数不能为 0")
    }
    return operation.numberA / operation.numberB
}

func main(){
    defer func (){
       if err :=recover();err !=nil{
           fmt.Println(err)
       }
    }()
    ifac :=new(AddFactory)
    oper:=ifac.createOperation();
    oper.SetNumA(1.0)
    oper.SetNumB(4.0)
    fmt.Println(oper.GetResult())
}