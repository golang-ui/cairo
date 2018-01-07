package pugl

import "C"
import "unsafe"

func (e *Event) Type() EventType {
	return *(*EventType)(unsafe.Pointer(e))
}

func (e *Event) Any() *EventAny {
	return NewEventAnyRef(unsafe.Pointer(e))
}

func (e *Event) Button() *EventButton {
	return NewEventButtonRef(unsafe.Pointer(e))
}

func (e *Event) Configure() *EventConfigure {
	return NewEventConfigureRef(unsafe.Pointer(e))
}

func (e *Event) Expose() *EventExpose {
	return NewEventExposeRef(unsafe.Pointer(e))
}

func (e *Event) Close() *EventClose {
	return NewEventCloseRef(unsafe.Pointer(e))
}

func (e *Event) Key() *EventKey {
	return NewEventKeyRef(unsafe.Pointer(e))
}

func (e *Event) Crossing() *EventCrossing {
	return NewEventCrossingRef(unsafe.Pointer(e))
}

func (e *Event) Motion() *EventMotion {
	return NewEventMotionRef(unsafe.Pointer(e))
}

func (e *Event) Scroll() *EventScroll {
	return NewEventScrollRef(unsafe.Pointer(e))
}

func (e *Event) Focus() *EventFocus {
	return NewEventFocusRef(unsafe.Pointer(e))
}

func (e EventType) String() string {
	switch e {
	case Nothing:
		return "Nothing"
	case ButtonPress:
		return "ButtonPress"
	case ButtonRelease:
		return "ButtonRelease"
	case Configure:
		return "Configure"
	case Expose:
		return "Expose"
	case Close:
		return "Close"
	case KeyPress:
		return "KeyPress"
	case KeyRelease:
		return "KeyRelease"
	case EnterNotify:
		return "EnterNotify"
	case LeaveNotify:
		return "LeaveNotify"
	case MotionNotify:
		return "MotionNotify"
	case Scroll:
		return "Scroll"
	case FocusIn:
		return "FocusIn"
	case FocusOut:
		return "FocusOut"
	default:
		return "Unknown"
	}
}
