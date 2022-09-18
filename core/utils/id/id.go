package id

var (
	Channel chan int
)

func init() {
	Channel = make(chan int)
	go func() {
		i := 0
		for {
			Channel <- i
			i++
		}
	}()
}
