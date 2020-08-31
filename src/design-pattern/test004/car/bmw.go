// 单例设计模式
package car

import "fmt"

type BmwCar interface {
	SetupBmwDriver()
}

type BmwSportCar struct {
}

type BmwBusinessCar struct {
}

func (b *BmwSportCar) SetupBmwDriver() {
	fmt.Println("bmw sport car setup driver")
}

func (b *BmwBusinessCar) SetupBmwDriver() {
	fmt.Println("bmw business car setup driver")
}
