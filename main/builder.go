package main

type PersonBuilder interface {
    buildHead()
    buildBody()
    buildArmRight()
    buildArmLeft()
    buildLegRight()
    buildLegLeft()
    display_personBuilding() *PersonBuilding
}

type PersonBuilding struct {
    head,body,leftarm,rightarm,leftleg,rightleg string
}

type PersonThinBuilder struct {
    building *PersonBuilding
}

func (this *PersonThinBuilder) buildBody(){
    this.building.body ="draw thin body..."
}

func (this *PersonThinBuilder)  buildHead(){
    this.building.head ="draw thin head..."
}

func (this *PersonThinBuilder) buildArmRight() {
    this.building.rightarm = "draw thin arm right..."
}

func (this *PersonThinBuilder) buildArmLeft() {
    this.building.leftarm = "draw thin arm left..."
}

func (this *PersonThinBuilder) buildLegRight() {
    this.building.rightleg = "draw thin leg right..."
}

func (this *PersonThinBuilder) buildLegLeft() {
    this.building.leftleg = "draw thin leg left..."
}


