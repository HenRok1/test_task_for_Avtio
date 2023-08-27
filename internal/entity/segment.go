package entity

type Segment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserSegment struct {
	UserID      int
	SegmentName []int
}
