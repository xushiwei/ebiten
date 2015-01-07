// Copyright 2015 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !js

package ui

import (
	glfw "github.com/go-gl/glfw3"
)

var glfwKeyCodeToKey = map[glfw.Key]Key{
	glfw.Key0:           Key0,
	glfw.Key1:           Key1,
	glfw.Key2:           Key2,
	glfw.Key3:           Key3,
	glfw.Key4:           Key4,
	glfw.Key5:           Key5,
	glfw.Key6:           Key6,
	glfw.Key7:           Key7,
	glfw.Key8:           Key8,
	glfw.Key9:           Key9,
	glfw.KeyA:           KeyA,
	glfw.KeyB:           KeyB,
	glfw.KeyC:           KeyC,
	glfw.KeyD:           KeyD,
	glfw.KeyE:           KeyE,
	glfw.KeyF:           KeyF,
	glfw.KeyG:           KeyG,
	glfw.KeyH:           KeyH,
	glfw.KeyI:           KeyI,
	glfw.KeyJ:           KeyJ,
	glfw.KeyK:           KeyK,
	glfw.KeyL:           KeyL,
	glfw.KeyM:           KeyM,
	glfw.KeyN:           KeyN,
	glfw.KeyO:           KeyO,
	glfw.KeyP:           KeyP,
	glfw.KeyQ:           KeyQ,
	glfw.KeyR:           KeyR,
	glfw.KeyS:           KeyS,
	glfw.KeyT:           KeyT,
	glfw.KeyU:           KeyU,
	glfw.KeyV:           KeyV,
	glfw.KeyW:           KeyW,
	glfw.KeyX:           KeyX,
	glfw.KeyY:           KeyY,
	glfw.KeyZ:           KeyZ,
	glfw.KeyCapsLock:    KeyCapsLock,
	glfw.KeyComma:       KeyComma,
	glfw.KeyDelete:      KeyDelete,
	glfw.KeyDown:        KeyDown,
	glfw.KeyEnd:         KeyEnd,
	glfw.KeyEnter:       KeyEnter,
	glfw.KeyEscape:      KeyEscape,
	glfw.KeyF1:          KeyF1,
	glfw.KeyF2:          KeyF2,
	glfw.KeyF3:          KeyF3,
	glfw.KeyF4:          KeyF4,
	glfw.KeyF5:          KeyF5,
	glfw.KeyF6:          KeyF6,
	glfw.KeyF7:          KeyF7,
	glfw.KeyF8:          KeyF8,
	glfw.KeyF9:          KeyF9,
	glfw.KeyF10:         KeyF10,
	glfw.KeyF11:         KeyF11,
	glfw.KeyF12:         KeyF12,
	glfw.KeyHome:        KeyHome,
	glfw.KeyInsert:      KeyInsert,
	glfw.KeyLeft:        KeyLeft,
	glfw.KeyLeftAlt:     KeyLeftAlt,
	glfw.KeyLeftControl: KeyLeftControl,
	glfw.KeyLeftShift:   KeyLeftShift,
	glfw.KeyPageDown:    KeyPageDown,
	glfw.KeyPageUp:      KeyPageUp,
	glfw.KeyPeriod:      KeyPeriod,
	glfw.KeyRight:       KeyRight,
	glfw.KeySpace:       KeySpace,
	glfw.KeyTab:         KeyTab,
	glfw.KeyUp:          KeyUp,
}