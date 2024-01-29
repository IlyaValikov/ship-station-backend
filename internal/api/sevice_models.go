package api

type Ship struct {
	ShipID            int     // Идентификатор судна
	ShipName          string  // Название судна
	ShipType          string  // Тип судна (контейнеровоз, танкер, пассажирское судно и т. д.)
	CargoCapacity     float64 // Грузоподъемность судна
	MaxDepth          float64 // Максимальная глубина судна (в метрах)
	MaxLength         float64 // Максимальная длина судна (в метрах)
	YearBuilt         int     // Год постройки судна
	Flag              string  // Флаг судна (страна регистрации)
	Classification    string  // Классификация судна (например, по Lloyd's Register)
	CrewCapacity      int     // Вместимость экипажа
	PassengerCapacity int     // Пассажирская вместимость (если судно пассажирское)
	Photo             string  // Фотография судна
	Href              string  // url отдельнеого судна
}