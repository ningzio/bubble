package core

import (
	"errors"
)

type Pitch struct {
	note Note
	// 升降半音
	ud int8
	// 高低八度
	hl int8
}

var ErrInvalidPitch = errors.New("invalid Note")

func NewPitch(note Note) (*Pitch, error) {
	if _, ok := noteInterval[note]; !ok {
		return nil, ErrInvalidPitch
	}
	return &Pitch{note: note}, nil
}

func (p *Pitch) String() string {
	result := p.note
	switch {
	case p.ud > 0:
		var i int8
		for ; i < p.ud; i++ {
			result += sharp
		}
	case p.ud < 0:
		var i int8
		for ; i > p.ud; i-- {
			result += flat
		}
	}

	if p.hl > 0 {
		result = "{" + result + "}"
	}
	if p.hl < 0 {
		result = "(" + result + ")"
	}
	return string(result)
}

// TODO error
var ErrInvalid = errors.New("")

// Flat 降半音
func (p *Pitch) Flat() error {
	if p.ud <= -2 {
		return ErrInvalid
	}
	p.ud--
	return nil
}

// Sharp 升半音
func (p *Pitch) Sharp() error {
	if p.ud >= 2 {
		return ErrInvalid
	}
	p.ud++
	return nil
}

// High8 高8度
func (p *Pitch) High8() error {
	if p.hl >= 1 {
		return ErrInvalid
	}
	p.hl++
	return nil
}

// Low8 低8度
func (p *Pitch) Low8() error {
	if p.hl <= -1 {
		return ErrInvalid
	}
	p.hl--
	return nil
}

func (p *Pitch) Interval() int8 {
	i := noteInterval[p.note]
	i += p.ud
	i += p.hl * 11
	return i
}

func (p *Pitch) Index() int8 {
	index := noteIndex[p.note]
	index += p.hl * 6
	return index
}

// Distance 计算两个音的度数
// 先计算两个 Note 的距离得到宽泛的度数 二度 三度。。
// 在根据 ud 计算大小增减，hl 高低八度的关系
func (p *Pitch) DistanceEasy(pitch *Pitch) *Distance {
	return NewDistance(p, pitch)
}
