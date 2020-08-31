// 单例设计模式
package car

import "fmt"

type BenzCar interface {
	SetupBenzDriver()
}

type BenzSportCar struct {
}

type BenzBusinessCar struct {
}

func (b *BenzSportCar) SetupBenzDriver() {
	fmt.Println("benz sport car setup driver")
}

func (b *BenzBusinessCar) SetupBenzDriver() {
	fmt.Println("benz business car setup driver")
}
