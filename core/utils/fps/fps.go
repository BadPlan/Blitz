package fps

import "time"

func DeltaTime(fps int) time.Duration {
	return time.Nanosecond * time.Duration(float64(time.Second.Nanoseconds())/float64(fps))
}

func FramesPerSecond(deltaTime time.Duration) int {
	return int(time.Second.Nanoseconds() / deltaTime.Nanoseconds())
}
