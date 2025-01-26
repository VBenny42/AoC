package day07

func (h *hand) fiveOfAKind() bool {
	return len(h.set) == 1
}

func (h *hand) fourOfAKind() bool {
	if len(h.set) != 2 {
		return false
	}
	for _, v := range h.set {
		if v == 4 || v == 1 {
			return true
		}
	}
	return false
}

func (h *hand) fullHouse() bool {
	if len(h.set) != 2 {
		return false
	}
	for _, v := range h.set {
		if v == 3 || v == 2 {
			return true
		}
	}
	return false
}

func (h *hand) threeOfAKind() bool {
	if len(h.set) != 3 {
		return false
	}
	for _, v := range h.set {
		if v == 3 {
			return true
		}
	}
	return false
}

func (h *hand) twoPairs() bool {
	if len(h.set) != 3 {
		return false
	}
	for _, v := range h.set {
		if v == 2 {
			return true
		}
	}
	return false
}

func (h *hand) onePair() bool {
	if len(h.set) != 4 {
		return false
	}
	for _, v := range h.set {
		if v == 2 {
			return true
		}
	}
	return false
}

func (h *hand) highCard() bool {
	return len(h.set) == 5
}

func (h *hand) determineHandType() {
	switch {
	case h.fiveOfAKind():
		h.hType = fiveOfAKind
	case h.fourOfAKind():
		h.hType = fourOfAKind
	case h.fullHouse():
		h.hType = fullHouse
	case h.threeOfAKind():
		h.hType = threeOfAKind
	case h.twoPairs():
		h.hType = twoPair
	case h.onePair():
		h.hType = onePair
	case h.highCard():
		h.hType = highCard
	default:
		panic("Invalid hand type")
	}
}
