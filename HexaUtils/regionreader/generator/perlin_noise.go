package generator

import (
	"math"
	"math/rand"
)

type PerlinNoise struct {
	octaves     int
	persistence float64
	lacunarity  float64
	gradients   []float64
	permutation []uint8
}

// NewPerlinNoise creates a new PerlinNoise instance with default octaves(6)
func NewPerlinNoise() *PerlinNoise {
	return NewPerlinNoiseOctave(6, 0.5, 2.0)
}

// NewPerlinNoiseOctave creates a new PerlinNoise instance with the specified octaves, persistence and lacunarity.
func NewPerlinNoiseOctave(octaves int, persistence float64, lacunarity float64) *PerlinNoise {
	if octaves <= 0 {
		octaves = 6
	}
	if persistence <= 0 {
		persistence = 0.5
	}
	if lacunarity <= 0 {
		lacunarity = 2.0
	}

	p := &PerlinNoise{
		octaves:     octaves,
		persistence: persistence,
		lacunarity:  lacunarity,
		permutation: make([]uint8, 256),
		gradients:   make([]float64, 256),
	}
	p.init()
	return p
}

func (p *PerlinNoise) init() {
	r := rand.New(rand.NewSource(1))

	// Initialize permutation
	for i := 0; i < 256; i++ {
		p.permutation[i] = uint8(i)
	}
	for i := 255; i > 0; i-- {
		j := r.Intn(i + 1)
		p.permutation[i], p.permutation[j] = p.permutation[j], p.permutation[i]
	}

	// Initialize gradients (angles in radians)
	for i := 0; i < 256; i++ {
		p.gradients[i] = r.Float64() * 2 * math.Pi
	}
}

// Sample2D generates the noise value at a 2D position (x, z)
func (p *PerlinNoise) Sample2D(x, z float64) float64 {
	total := 0.0
	frequency := 1.0
	amplitude := 1.0
	maxValue := 0.0

	for i := 0; i < p.octaves; i++ {
		total += p.singleSample2D(x*frequency, z*frequency) * amplitude
		maxValue += amplitude
		amplitude *= p.persistence
		frequency *= p.lacunarity
	}
	return total / maxValue
}

func (p *PerlinNoise) singleSample2D(x, z float64) float64 {
	// 1. Find grid cell
	x0f := math.Floor(x)
	z0f := math.Floor(z)
	x0 := int(x0f)
	z0 := int(z0f)
	x1 := x0 + 1
	z1 := z0 + 1

	// 2. Calculate interpolation values between 0 and 1
	sx := x - x0f
	sz := z - z0f

	// 3. Calculate gradients at the 4 corners of the cell
	n00 := p.dotGridGradient(x0, z0, x, z)
	n10 := p.dotGridGradient(x1, z0, x, z)
	n01 := p.dotGridGradient(x0, z1, x, z)
	n11 := p.dotGridGradient(x1, z1, x, z)

	// 4. Interpolate results along X
	ix0 := p.interpolate(n00, n10, sx)

	// 5. Interpolate results along Z
	ix1 := p.interpolate(n01, n11, sx)

	// 6. Interpolate the final result
	return p.interpolate(ix0, ix1, sz)
}

// dotGridGradient calculates the dot product between the distance and the gradient vector
func (p *PerlinNoise) dotGridGradient(ix, iz int, x, z float64) float64 {
	// 1. Finds the gradient vector corresponding to the grid corner
	hash := p.permutation[(p.permutation[ix&255]+uint8(iz))&255]
	angle := p.gradients[hash]

	// 2. Calculates the distance vector between the point and the grid's origin
	dx := x - float64(ix)
	dz := z - float64(iz)

	// 3. Dot product
	return dx*math.Cos(angle) + dz*math.Sin(angle)
}

// interpolate performs a cosine interpolation
func (p *PerlinNoise) interpolate(a0, a1, w float64) float64 {

	// Smooth interpolation
	return (a1-a0)*(3.0-w*2.0)*w*w + a0
}

// ridge applies a ridge effect to the noise value
func Ridge(value float64) float64 {
	return 1.0 - math.Abs(value)
}
