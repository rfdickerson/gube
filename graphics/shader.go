package graphics

import (
	"io/ioutil"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Shader used to draw stuff
type Shader struct {
	id   uint32
	Name string
}

// NewShader creates a new shader
func NewShader(name string, vertexFile string, fragmentFile string) (*Shader, error) {

	dat, err := ioutil.ReadFile(vertexFile)
	if err != nil {
		return nil, err
	}

	vs := gl.CreateShader(gl.VERTEX_SHADER)

	gl.ShaderBinary(1, &vs, gl.SHADER_BINARY_FORMAT_SPIR_V_ARB, gl.Ptr(dat), int32(len(dat)))

	ep := []uint8("main")

	gl.SpecializeShaderARB(vs, &ep[0], 0, nil, nil)

	return &Shader{id: vs, Name: name}, nil
}

// Destroy deconstructs a shader
func (s *Shader) Destroy() {
	gl.DeleteShader(s.id)
}

// Use binds the shader
func (s *Shader) Use() {
	gl.UseProgram(s.id)
}
