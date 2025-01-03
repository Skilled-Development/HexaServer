package entities

import "fmt"

type Location struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
}

func NewLocation(x float64, y float64, z float64, yaw float32, pitch float32) *Location {
	return &Location{
		X:     x,
		Y:     y,
		Z:     z,
		Yaw:   yaw,
		Pitch: pitch,
	}
}

func (l *Location) GetX() float64 {
	return l.X
}

func (l *Location) GetY() float64 {
	return l.Y
}

func (l *Location) GetZ() float64 {
	return l.Z
}

func (l *Location) GetYaw() float32 {
	return l.Yaw
}

func (l *Location) GetPitch() float32 {
	return l.Pitch
}

func (l *Location) SetX(x float64) {
	l.X = x
}

func (l *Location) SetY(y float64) {
	l.Y = y
}

func (l *Location) SetZ(z float64) {
	l.Z = z
}

func (l *Location) SetYaw(yaw float32) {
	l.Yaw = yaw
}

func (l *Location) SetPitch(pitch float32) {
	l.Pitch = pitch
}

func (l *Location) GetLocation() *Location {
	return l
}

func (l *Location) SetLocation(x float64, y float64, z float64, yaw float32, pitch float32) {
	l.X = x
	l.Y = y
	l.Z = z
	l.Yaw = yaw
	l.Pitch = pitch
}

func (l *Location) GetLocationAsString() string {
	return fmt.Sprintf("X: %f, Y: %f, Z: %f, Yaw: %f, Pitch: %f", l.X, l.Y, l.Z, l.Yaw, l.Pitch)
}

func (l *Location) GetLocationAsArray() []float64 {
	return []float64{l.X, l.Y, l.Z, float64(l.Yaw), float64(l.Pitch)}
}

func (l *Location) Add(x float64, y float64, z float64) {
	l.X += x
	l.Y += y
	l.Z += z
}

func (l *Location) Subtract(x float64, y float64, z float64) {
	l.X -= x
	l.Y -= y
	l.Z -= z
}

func (l *Location) Multiply(x float64, y float64, z float64) {
	l.X *= x
	l.Y *= y
	l.Z *= z
}

func (l *Location) Divide(x float64, y float64, z float64) {
	l.X /= x
	l.Y /= y
	l.Z /= z
}

func (l *Location) Clone() *Location {
	return &Location{
		X:     l.X,
		Y:     l.Y,
		Z:     l.Z,
		Yaw:   l.Yaw,
		Pitch: l.Pitch,
	}
}

func (l *Location) Equals(location *Location) bool {
	return l.X == location.X && l.Y == location.Y && l.Z == location.Z && l.Yaw == location.Yaw && l.Pitch == location.Pitch
}

func (l *Location) GetDistance(location *Location) float64 {
	distanceX := l.X - location.X
	distanceY := l.Y - location.Y
	distanceZ := l.Z - location.Z
	return distanceX*distanceX + distanceY*distanceY + distanceZ*distanceZ
}
