// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bubblegraph

// Below is standard struct for bubble graph related.
type (
	// Data includes List.
	Data struct {
		Title string
		List  []*Bubble
	}

	// Bubble .
	Bubble struct {
		X         *Axis       `json:"x"`
		Y         *Axis       `json:"y"`
		Size      *BubbleSize `json:"size"`
		Group     string      `json:"group"`
		Dimension string      `json:"dimension"`
	}

	// BubbleSize .
	BubbleSize struct {
		Value float64 `json:"value"`
	}

	// Axis .
	Axis struct {
		Value interface{} `json:"value"`
		Unit  string      `json:"unit"`
	}

	// BubbleBuilder .
	BubbleBuilder struct {
		bubble *Bubble
	}

	// DataBuilder .
	DataBuilder struct {
		data *Data
	}
)

// NewDataBuilder .
func NewDataBuilder() *DataBuilder {
	return &DataBuilder{data: &Data{}}
}

// WithTitle .
func (d *DataBuilder) WithTitle(title string) *DataBuilder {
	d.data.Title = title
	return d
}

// WithBubble .
func (d *DataBuilder) WithBubble(bubbles ...*Bubble) *DataBuilder {
	d.data.List = append(d.data.List, bubbles...)
	return d
}

// Build .
func (d *DataBuilder) Build() *Data {
	return d.data
}

// NewBubbleBuilder .
func NewBubbleBuilder() *BubbleBuilder {
	return &BubbleBuilder{bubble: &Bubble{}}
}

// WithValueX .
func (bb *BubbleBuilder) WithValueX(v interface{}) *BubbleBuilder {
	bb.bubble.X.Value = v
	return bb
}

// WithX .
func (bb *BubbleBuilder) WithX(x *Axis) *BubbleBuilder {
	bb.bubble.X = x
	return bb
}

// WithValueY .
func (bb *BubbleBuilder) WithValueY(v interface{}) *BubbleBuilder {
	bb.bubble.Y.Value = v
	return bb
}

// WithY .
func (bb *BubbleBuilder) WithY(y *Axis) *BubbleBuilder {
	bb.bubble.Y = y
	return bb
}

// WithSize .
func (bb *BubbleBuilder) WithSize(size *BubbleSize) *BubbleBuilder {
	bb.bubble.Size = size
	return bb
}

// WithValueSize .
func (bb *BubbleBuilder) WithValueSize(v float64) *BubbleBuilder {
	bb.bubble.Size.Value = v
	return bb
}

// WithGroup .
func (bb *BubbleBuilder) WithGroup(group string) *BubbleBuilder {
	bb.bubble.Group = group
	return bb
}

// WithDimension .
func (bb *BubbleBuilder) WithDimension(dimension string) *BubbleBuilder {
	bb.bubble.Dimension = dimension
	return bb
}

// Build .
func (bb *BubbleBuilder) Build() *Bubble {
	return bb.bubble
}
