package graphics

import "github.com/go-gl/gl/v3.3-core/gl"

type Quad struct {
	vao uint32
	vbo uint32
}

func NewQuad() *Quad {

	quadVerts := []float32{
		-1.0, 1.0, 0.0, 1.0,
		-1.0, -1.0, 0.0, 0.0,
		1.0, -1.0, 1.0, 0.0,

		-1.0, 1.0, 0.0, 1.0,
		1.0, -1.0, 1.0, 0.0,
		1.0, 1.0, 1.0, 1.0,
	}

	p := new(Quad)
	gl.GenVertexArrays(1, &p.vao)
	gl.GenBuffers(1, &p.vbo)
	gl.BindVertexArray(p.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, p.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(quadVerts), gl.Ptr(quadVerts), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 8, gl.PtrOffset(0))

	return p

}

// Draw renders the geometry
func (quad *Quad) Draw() {
	gl.BindVertexArray(quad.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}

// Destroy the Quad
func (quad *Quad) Destroy() {
	gl.DeleteBuffers(1, &quad.vbo)
	gl.DeleteVertexArrays(1, &quad.vao)
}
