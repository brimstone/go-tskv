package tskv

// Subscribe returns a new channel in which updates to the tskv are sent
func (t *Tskv) Subscribe() chan *Tskv {
	channel := make(chan *Tskv)
	t.subscribers = append(t.subscribers, channel)
	return channel
}
