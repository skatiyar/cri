package headlessexperimental

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type HeadlessExperimental struct {
	conn cri.Connector
}

func New(conn cri.Connector) *HeadlessExperimental {
	return &HeadlessExperimental{conn}
}
func (obj *HeadlessExperimental) Enable() (err error) {
	err = obj.conn.Send("HeadlessExperimental.enable", nil, nil)
	return
}
func (obj *HeadlessExperimental) Disable() (err error) {
	err = obj.conn.Send("HeadlessExperimental.disable", nil, nil)
	return
}

type BeginFrameRequest struct {
	FrameTime  *types.Runtime_Timestamp                     `json:"frameTime,omitempty"`
	Deadline   *types.Runtime_Timestamp                     `json:"deadline,omitempty"`
	Interval   *float32                                     `json:"interval,omitempty"`
	Screenshot *types.HeadlessExperimental_ScreenshotParams `json:"screenshot,omitempty"`
}

func (obj *HeadlessExperimental) BeginFrame(request *BeginFrameRequest) (response struct {
	HasDamage               bool    `json:"hasDamage"`
	MainFrameContentUpdated bool    `json:"mainFrameContentUpdated"`
	ScreenshotData          *string `json:"screenshotData,omitempty"`
}, err error) {
	err = obj.conn.Send("HeadlessExperimental.beginFrame", request, &response)
	return
}
