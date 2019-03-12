// Code generated; Do not regenerate the overwrite after editing.

package areacn

import (
	"github.com/wzshiming/areacn"
)

// AreacnService #path:"/areacn/"#
type AreacnService struct {
}

// NewAreacnService Create a new AreacnService
func NewAreacnService() (*AreacnService, error) {
	return &AreacnService{}, nil
}

// Get #route:"GET /{area_id}"# 获取行政区划分信息 总共5级 获取第一级省份传0
func (s *AreacnService) Get(areaID string /* #name:"area_id"# */) (areas []*areacn.Area, err error) {
	if areaID == "0" {
		areaID = ""
	}
	return areacn.Get(areaID), nil
}
