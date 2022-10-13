package fps

import "time"

func DeltaTime(fps int) time.Duration {
	return time.Nanosecond * time.Duration(float64(time.Second.Nanoseconds())/float64(fps))
}

func FramesPerSecond(deltaTime time.Duration) int {
	if deltaTime == 0 {
		return 0
	}
	return int(time.Second.Nanoseconds() / deltaTime.Nanoseconds())
}
