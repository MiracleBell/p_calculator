package features

import (
	. "../../database/features"
	. "../../models/features"
	"time"
)

func AddFeature(feature *Feature) error {
	feature.CreatedAt = time.Now()
	feature.LastUpdatedAt = time.Now()
	return AddNewFeature(feature)
}

func UpdateFeature(feature *Feature) error {
	feature.LastUpdatedAt = time.Now()
	return ChangeFeature(feature)
}

func GetFeatures(projectId uint64) ([]Feature, error) {
	return GetProjectFeatures(projectId)
}
