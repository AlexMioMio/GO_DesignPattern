package main

import (
    "testing"
)
// 测试加法
func Test_Add(t *testing.T) {
    var fac OperationFactory
    oper := fac.createOperation("+")
    oper.SetNumA(1.0)
    oper.SetNumB(2.0)
    if oper.getResult() ==3.0 {
        t.Log("加法测试通过")
    }else {
        t.Error("加法测试出错了")
    }
}
// 测试减法
func Test_Sub(t *testing.T) {
    var fac OperationFactory
    oper := fac.createOperation("-")
    oper.SetNumA(10.0)
    oper.SetNumB(2.0)
    if oper.getResult() ==8.0 {
        t.Log("减法测试通过")
    }else {
        t.Error("减法测试出错了")
    }
}
// 测试除法
func Test_Div(t *testing.T) {
    var fac OperationFactory
    oper := fac.createOperation("/")
    oper.SetNumA(1.0)
    oper.SetNumB(2.0)
    if oper.getResult() == 0.5 {
        t.Error("除法测试通过")
    }else {
        t.Error("除法测试没有通过")
    }

}
// 测试乘法
func Test_Mul(t *testing.T) {
    var fac OperationFactory
    oper := fac.createOperation("*")
    oper.SetNumA(9.0)
    oper.SetNumB(2.0)
    if oper.getResult() == 18 {
        t.Error("乘法测试通过")
    }else {
        t.Error("乘法测试没有通过")
    }
}

