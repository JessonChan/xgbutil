/*
   Defines event types and their associated methods automatically.

   Eventually, the no-op implementations of 'Bytes' will move out of here.

   This file is automatically generated using `scripts/write-events evtypes`.

   Edit it at your peril.
*/
package xevent

import "fmt"
import "burntsushi.net/go/x-go-binding/xgb"

type KeyPressEvent struct {
	*xgb.KeyPressEvent
}

const KeyPress = xgb.KeyPress

func (ev *KeyPressEvent) Bytes() []byte { return nil }

func (ev KeyPressEvent) String() string {
	return fmt.Sprintf("%v", ev.KeyPressEvent)
}

type KeyReleaseEvent struct {
	*xgb.KeyReleaseEvent
}

const KeyRelease = xgb.KeyRelease

func (ev *KeyReleaseEvent) Bytes() []byte { return nil }

func (ev KeyReleaseEvent) String() string {
	return fmt.Sprintf("%v", ev.KeyReleaseEvent)
}

type ButtonPressEvent struct {
	*xgb.ButtonPressEvent
}

const ButtonPress = xgb.ButtonPress

func (ev *ButtonPressEvent) Bytes() []byte { return nil }

func (ev ButtonPressEvent) String() string {
	return fmt.Sprintf("%v", ev.ButtonPressEvent)
}

type ButtonReleaseEvent struct {
	*xgb.ButtonReleaseEvent
}

const ButtonRelease = xgb.ButtonRelease

func (ev *ButtonReleaseEvent) Bytes() []byte { return nil }

func (ev ButtonReleaseEvent) String() string {
	return fmt.Sprintf("%v", ev.ButtonReleaseEvent)
}

type MotionNotifyEvent struct {
	*xgb.MotionNotifyEvent
}

const MotionNotify = xgb.MotionNotify

func (ev *MotionNotifyEvent) Bytes() []byte { return nil }

func (ev MotionNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.MotionNotifyEvent)
}

type EnterNotifyEvent struct {
	*xgb.EnterNotifyEvent
}

const EnterNotify = xgb.EnterNotify

func (ev *EnterNotifyEvent) Bytes() []byte { return nil }

func (ev EnterNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.EnterNotifyEvent)
}

type LeaveNotifyEvent struct {
	*xgb.LeaveNotifyEvent
}

const LeaveNotify = xgb.LeaveNotify

func (ev *LeaveNotifyEvent) Bytes() []byte { return nil }

func (ev LeaveNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.LeaveNotifyEvent)
}

type FocusInEvent struct {
	*xgb.FocusInEvent
}

const FocusIn = xgb.FocusIn

func (ev *FocusInEvent) Bytes() []byte { return nil }

func (ev FocusInEvent) String() string {
	return fmt.Sprintf("%v", ev.FocusInEvent)
}

type FocusOutEvent struct {
	*xgb.FocusOutEvent
}

const FocusOut = xgb.FocusOut

func (ev *FocusOutEvent) Bytes() []byte { return nil }

func (ev FocusOutEvent) String() string {
	return fmt.Sprintf("%v", ev.FocusOutEvent)
}

type KeymapNotifyEvent struct {
	*xgb.KeymapNotifyEvent
}

const KeymapNotify = xgb.KeymapNotify

func (ev *KeymapNotifyEvent) Bytes() []byte { return nil }

func (ev KeymapNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.KeymapNotifyEvent)
}

type ExposeEvent struct {
	*xgb.ExposeEvent
}

const Expose = xgb.Expose

func (ev *ExposeEvent) Bytes() []byte { return nil }

func (ev ExposeEvent) String() string {
	return fmt.Sprintf("%v", ev.ExposeEvent)
}

type GraphicsExposureEvent struct {
	*xgb.GraphicsExposureEvent
}

const GraphicsExposure = xgb.GraphicsExposure

func (ev *GraphicsExposureEvent) Bytes() []byte { return nil }

func (ev GraphicsExposureEvent) String() string {
	return fmt.Sprintf("%v", ev.GraphicsExposureEvent)
}

type NoExposureEvent struct {
	*xgb.NoExposureEvent
}

const NoExposure = xgb.NoExposure

func (ev *NoExposureEvent) Bytes() []byte { return nil }

func (ev NoExposureEvent) String() string {
	return fmt.Sprintf("%v", ev.NoExposureEvent)
}

type VisibilityNotifyEvent struct {
	*xgb.VisibilityNotifyEvent
}

const VisibilityNotify = xgb.VisibilityNotify

func (ev *VisibilityNotifyEvent) Bytes() []byte { return nil }

func (ev VisibilityNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.VisibilityNotifyEvent)
}

type CreateNotifyEvent struct {
	*xgb.CreateNotifyEvent
}

const CreateNotify = xgb.CreateNotify

func (ev *CreateNotifyEvent) Bytes() []byte { return nil }

func (ev CreateNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.CreateNotifyEvent)
}

type DestroyNotifyEvent struct {
	*xgb.DestroyNotifyEvent
}

const DestroyNotify = xgb.DestroyNotify

func (ev *DestroyNotifyEvent) Bytes() []byte { return nil }

func (ev DestroyNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.DestroyNotifyEvent)
}

type UnmapNotifyEvent struct {
	*xgb.UnmapNotifyEvent
}

const UnmapNotify = xgb.UnmapNotify

func (ev *UnmapNotifyEvent) Bytes() []byte { return nil }

func (ev UnmapNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.UnmapNotifyEvent)
}

type MapNotifyEvent struct {
	*xgb.MapNotifyEvent
}

const MapNotify = xgb.MapNotify

func (ev *MapNotifyEvent) Bytes() []byte { return nil }

func (ev MapNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.MapNotifyEvent)
}

type MapRequestEvent struct {
	*xgb.MapRequestEvent
}

const MapRequest = xgb.MapRequest

func (ev *MapRequestEvent) Bytes() []byte { return nil }

func (ev MapRequestEvent) String() string {
	return fmt.Sprintf("%v", ev.MapRequestEvent)
}

type ReparentNotifyEvent struct {
	*xgb.ReparentNotifyEvent
}

const ReparentNotify = xgb.ReparentNotify

func (ev *ReparentNotifyEvent) Bytes() []byte { return nil }

func (ev ReparentNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.ReparentNotifyEvent)
}

type ConfigureRequestEvent struct {
	*xgb.ConfigureRequestEvent
}

const ConfigureRequest = xgb.ConfigureRequest

func (ev *ConfigureRequestEvent) Bytes() []byte { return nil }

func (ev ConfigureRequestEvent) String() string {
	return fmt.Sprintf("%v", ev.ConfigureRequestEvent)
}

type GravityNotifyEvent struct {
	*xgb.GravityNotifyEvent
}

const GravityNotify = xgb.GravityNotify

func (ev *GravityNotifyEvent) Bytes() []byte { return nil }

func (ev GravityNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.GravityNotifyEvent)
}

type ResizeRequestEvent struct {
	*xgb.ResizeRequestEvent
}

const ResizeRequest = xgb.ResizeRequest

func (ev *ResizeRequestEvent) Bytes() []byte { return nil }

func (ev ResizeRequestEvent) String() string {
	return fmt.Sprintf("%v", ev.ResizeRequestEvent)
}

type CirculateNotifyEvent struct {
	*xgb.CirculateNotifyEvent
}

const CirculateNotify = xgb.CirculateNotify

func (ev *CirculateNotifyEvent) Bytes() []byte { return nil }

func (ev CirculateNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.CirculateNotifyEvent)
}

type CirculateRequestEvent struct {
	*xgb.CirculateRequestEvent
}

const CirculateRequest = xgb.CirculateRequest

func (ev *CirculateRequestEvent) Bytes() []byte { return nil }

func (ev CirculateRequestEvent) String() string {
	return fmt.Sprintf("%v", ev.CirculateRequestEvent)
}

type PropertyNotifyEvent struct {
	*xgb.PropertyNotifyEvent
}

const PropertyNotify = xgb.PropertyNotify

func (ev *PropertyNotifyEvent) Bytes() []byte { return nil }

func (ev PropertyNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.PropertyNotifyEvent)
}

type SelectionClearEvent struct {
	*xgb.SelectionClearEvent
}

const SelectionClear = xgb.SelectionClear

func (ev *SelectionClearEvent) Bytes() []byte { return nil }

func (ev SelectionClearEvent) String() string {
	return fmt.Sprintf("%v", ev.SelectionClearEvent)
}

type SelectionRequestEvent struct {
	*xgb.SelectionRequestEvent
}

const SelectionRequest = xgb.SelectionRequest

func (ev *SelectionRequestEvent) Bytes() []byte { return nil }

func (ev SelectionRequestEvent) String() string {
	return fmt.Sprintf("%v", ev.SelectionRequestEvent)
}

type SelectionNotifyEvent struct {
	*xgb.SelectionNotifyEvent
}

const SelectionNotify = xgb.SelectionNotify

func (ev *SelectionNotifyEvent) Bytes() []byte { return nil }

func (ev SelectionNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.SelectionNotifyEvent)
}

type ColormapNotifyEvent struct {
	*xgb.ColormapNotifyEvent
}

const ColormapNotify = xgb.ColormapNotify

func (ev *ColormapNotifyEvent) Bytes() []byte { return nil }

func (ev ColormapNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.ColormapNotifyEvent)
}

type MappingNotifyEvent struct {
	*xgb.MappingNotifyEvent
}

const MappingNotify = xgb.MappingNotify

func (ev *MappingNotifyEvent) Bytes() []byte { return nil }

func (ev MappingNotifyEvent) String() string {
	return fmt.Sprintf("%v", ev.MappingNotifyEvent)
}
