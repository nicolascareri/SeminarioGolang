package lib

import "fmt"

type Agencia struct {
	autos []Auto

}

func NewAgencia() Agencia{
	var aux []Auto
	return Agencia{
		autos: aux,
	}

}

func (ag *Agencia) AddAuto(auto ...Auto){

	ag.autos = append(ag.autos, auto...)

}

func (ag *Agencia) RemoveAuto(pos int){
	ag.autos = append(ag.autos[:pos], ag.autos[pos+1:]...)
}

func (ag *Agencia) GetAuto(pos int) Auto{
	return ag.autos[pos]
}

func (ag *Agencia) GetAutos() []Auto{
	return ag.autos
}

func (ag *Agencia) EditColor(Patente string, Color string){
	for index := 0; index < len(ag.autos); index++ {
		if ag.autos[index].Plate == Patente{
			ag.autos[index].Color = Color
		}
	}


}

func (ag *Agencia) BuscarAuto(Patente string) *Auto  {
	for _, auto := range ag.autos{
		fmt.Println(auto)
	}
	for index := 0; index < len(ag.autos); index++ {
		if ag.autos[index].Plate == Patente{
			return &ag.autos[index]
		}
	}
	return nil
}