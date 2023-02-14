package windicators

// Prefefined positions for quick setup
const (
	PositionCenterTop uint = iota
	PositionRightTop
	PositionRightCenter
	PositionRightBottom
	PositionCenterBottom
	PositionLeftBottom
	PositionLeftCenter
	PositionLeftTop
)

// PositionPoint ..
func PositionPoint(position, screenW, screenH, windowW, windowH uint) *Point {

	centerX := int((float32(screenW) / 2) - (float32(windowW) / 2))
	centerY := int((float32(screenH) / 2) - (float32(windowH) / 2))

	right := int(screenW - windowW)
	bottom := int(screenH - windowH)

	switch position {
	case PositionCenterTop:
		return NewPoint(centerX, 0)
	case PositionRightTop:
		return NewPoint(right, 0)
	case PositionRightCenter:
		return NewPoint(right, centerY)
	case PositionRightBottom:
		return NewPoint(right, bottom)
	// case PositionCenterBottom:
	// 	return NewPoint(centerX, bottom)
	case PositionLeftBottom:
		return NewPoint(0, bottom)
	case PositionLeftCenter:
		return NewPoint(0, centerY)
	case PositionLeftTop:
		return NewPoint(0, 0)
	}

	// same as `PositionCenterBottom`
	return NewPoint(centerX, bottom)
}
