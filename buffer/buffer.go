package buffer

type ClockReplacer struct {
	frames []int
}

func (c *ClockReplacer) Victim() bool {

	return false
}

func (c *ClockReplacer) Pin(frame int) {

}

func (c *ClockReplacer) Unpin(frame int) {

}

func (c *ClockReplacer) Size() int {

	return 0
}
