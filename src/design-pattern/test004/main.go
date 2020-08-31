// 抽象工厂设计模式
package main

import (
	"github.com/sahara-gopher/tour/src/design-pattern/test004/car"
)

// StoreFactory 工厂接口
type CarFactory interface {
	CreateBmwCar() car.BmwCar
	CreateBenzCar() car.BenzCar
}

// SportCarFactory 运动轿车具体工厂
type SportCarFactory struct {
}

func (f *SportCarFactory) CreateBmwCar() car.BmwCar {
	return &car.BmwSportCar{}
}

func (f *SportCarFactory) CreateBenzCar() car.BenzCar {
	return &car.BenzSportCar{}
}

// BusinessCarFactory 运动轿车具体工厂
type BusinessCarFactory struct {
}

func (f *BusinessCarFactory) CreateBmwCar() car.BmwCar {
	return &car.BmwBusinessCar{}
}

func (f *BusinessCarFactory) CreateBenzCar() car.BenzCar {
	return &car.BenzBusinessCar{}
}

// 测试抽象工厂设计模式
func main() {
	scFactory := &SportCarFactory{}
	bzCar := scFactory.CreateBenzCar()
	bzCar.SetupBenzDriver()

	bmCar := scFactory.CreateBmwCar()
	bmCar.SetupBmwDriver()

	bcFactory := &BusinessCarFactory{}
	bzCar = bcFactory.CreateBenzCar()
	bzCar.SetupBenzDriver()

	bmCar = bcFactory.CreateBmwCar()
	bmCar.SetupBmwDriver()
}
