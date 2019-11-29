package lib

type Auto struct {
	Plate      string `json:"plate"`
	Color      string `json:"color"`
}
func NewAuto(Plate, Color string) Auto{
	return Auto{
		Plate: Plate,
		Color: Color,
	}
}