// 包cake提供了一个并发蛋糕商店的模拟程序
// 蛋糕店运作流程  烘焙 --> 加糖衣 --> 雕刻
package cake

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose        bool
	Cakes          int           // 需要烘焙的蛋糕数量
	BakeTime       time.Duration // 烘焙一个蛋糕的时间
	BakeStdDev     time.Duration // 烘焙时间的标准偏差
	BakeBuf        int           // 烘烤和加糖衣之间缓冲个数
	IceCookNum     int           // 加糖衣的厨师数
	IceTime        time.Duration // 对一个蛋糕加糖衣的时间
	IceStdDev      time.Duration // 加糖衣时间的标准偏差
	IceBuf         int           // 加糖衣和雕刻之间缓冲个数
	InscribeTime   time.Duration // 雕刻一个蛋糕的时间
	InscribeStdDev time.Duration // 雕刻时间的标准偏差
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("finished", c)
		}
	}
}

// Work 运行模拟程序runs次
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.IceCookNum; i++ {
			go s.icer(iced, baked)
		}
		s.inscriber(iced)
	}
}

func work(d, stdDev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stdDev))
	time.Sleep(delay)
}
