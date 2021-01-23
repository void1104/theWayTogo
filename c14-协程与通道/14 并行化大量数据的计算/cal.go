package main

/**
1. 典型的用于解决按顺序执行每个步骤的顺序流水线算法
*/
//func SerialProcessData(in <-chan *Data, out chan<- *Data) {
//	for data := range in {
//		tmpA := PreprocessData(data)
//		tmpB := ProcessStepA(tmpA)
//		tmpC := ProcessStepB(tmpB)
//		out <- PostProcessData(tmpC)
//	}
//}

/**
2.更高效的方法：让每一个步骤作为一个协程独立工作。
每一个步骤从上一步的输出通道中获得输入数据。这种方式
仅有极少数时间会被浪费，而大部分时间所有的步骤都在一直
执行中
*/
//func ParallelProcessData (in <-chan *Data, out chan<- *Data) {
//	// make channels:
//	preOut := make(chan *Data, 100)
//	stepAOut := make(chan *Data, 100)
//	stepBOut := make(chan *Data, 100)
//	stepCOut := make(chan *Data, 100)
//	// start parallel computations:
//	go PreprocessData(in, preOut)
//	go ProcessStepA(preOut,StepAOut)
//	go ProcessStepB(StepAOut,StepBOut)
//	go ProcessStepC(StepBOut,StepCOut)
//	go PostProcessData(StepCOut,out)
//}
