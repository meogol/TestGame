package gamescene

type sceneState struct {
	tailItem   *SnakeNode
	headItem   *SnakeNode
	frameDelay int
}

func (s *sceneState) AddSnakeNode(sn *SnakeNode) {
	if s.tailItem == nil {
		s.tailItem = sn
	} else {
		sn.nextNode = s.tailItem
		s.tailItem = sn
	}
}
