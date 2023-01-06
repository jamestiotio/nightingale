package conv

import (
	"fmt"
	"math"
	"strings"

	"github.com/prometheus/common/model"
)

type Vector struct {
	Key       string       `json:"key"`
	Labels    model.Metric `json:"labels"`
	Timestamp int64        `json:"timestamp"`
	Value     float64      `json:"value"`
}

func (v *Vector) ReadableValue() string {
	ret := fmt.Sprintf("%.5f", v.Value)
	ret = strings.TrimRight(ret, "0")
	return strings.TrimRight(ret, ".")
}

func ConvertVectors(value model.Value) (lst []Vector) {
	if value == nil {
		return
	}

	switch value.Type() {
	case model.ValVector:
		items, ok := value.(model.Vector)
		if !ok {
			return
		}

		for _, item := range items {
			if math.IsNaN(float64(item.Value)) {
				continue
			}

			lst = append(lst, Vector{
				Key:       item.Metric.String(),
				Timestamp: item.Timestamp.Unix(),
				Value:     float64(item.Value),
				Labels:    item.Metric,
			})
		}
	case model.ValMatrix:
		items, ok := value.(model.Matrix)
		if !ok {
			return
		}

		for _, item := range items {
			if len(item.Values) == 0 {
				return
			}

			last := item.Values[len(item.Values)-1]

			if math.IsNaN(float64(last.Value)) {
				continue
			}

			lst = append(lst, Vector{
				Key:       item.Metric.String(),
				Labels:    item.Metric,
				Timestamp: last.Timestamp.Unix(),
				Value:     float64(last.Value),
			})
		}
	case model.ValScalar:
		item, ok := value.(*model.Scalar)
		if !ok {
			return
		}

		if math.IsNaN(float64(item.Value)) {
			return
		}

		lst = append(lst, Vector{
			Key:       "{}",
			Timestamp: item.Timestamp.Unix(),
			Value:     float64(item.Value),
			Labels:    model.Metric{},
		})
	default:
		return
	}

	return
}
