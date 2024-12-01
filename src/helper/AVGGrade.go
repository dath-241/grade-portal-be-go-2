package helper

import (
	"Go2/models"
)

// AvgScore tính toán điểm trung bình tổng hợp từ các điểm BT, TN, BTL, GK và CK
func AvgScore(data models.InterfaceScore, HS []int) float32 {
	// Hàm tính trung bình
	calculateAvg := func(scores []float32) float32 {
		var total float32
		for _, score := range scores {
			total += score
		}
		if len(scores) > 0 {
			return total / float32(len(scores))
		}
		return 0
	}

	// Tính điểm trung bình cho từng loại
	avgBT := calculateAvg(data.BT)
	avgTN := calculateAvg(data.TN)
	avgBTL := calculateAvg(data.BTL)

	// Tính điểm trung bình tổng hợp
	return float32(HS[0])*avgBT/100 +
		float32(HS[1])*avgTN/100 +
		float32(HS[2])*avgBTL/100 +
		float32(HS[3])*data.GK/100 +
		float32(HS[4])*data.CK/100
}
