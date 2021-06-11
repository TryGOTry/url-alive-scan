/*
* @Author: Try
* @Date:   2021/5/5 11:12
 */
package golimit

type Glimit struct {
	Num int
	C   chan struct{}
}
func NewG(num int) *Glimit {
	return &Glimit{
		Num: num,
		C : make(chan struct{}, num),
	}
}
func (g *Glimit) Run(f func()){
	g.C <- struct{}{}
	go func() {
		f()
		<-g.C
	}()
}

