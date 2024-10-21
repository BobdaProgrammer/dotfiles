package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"github.com/go-gl/mathgl/mgl32"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 500
	height = 500




	vertexShaderSource = `
    #version 410
    in vec3 vp;
    in vec3 vertexColor;
    out vec3 fragColor;

    uniform mat4 model;     // Model matrix for object transformations
    uniform mat4 view;      // View matrix for camera transformations
    uniform mat4 projection; // Projection matrix for perspective

    void main() {
        fragColor = vertexColor;
        gl_Position = projection * view * model * vec4(vp, 1.0);
    }
` + "\x00"





	fragmentShaderSource = `
	#version 410
	in vec3 fragColor;
	out vec4 frag_colour;
	void main() {
		frag_colour = vec4(fragColor, 1.0);  // use the interpolated color
	}
` + "\x00"
)

var (
square = []float32{
    // First triangle (top-left, bottom-left, bottom-right)
    -0.5,  0.5, 0.0, // top-left
    -0.5, -0.5, 0.0, // bottom-left
     0.5, -0.5, 0.0, // bottom-right

    // Second triangle (top-left, bottom-right, top-right)
    -0.5,  0.5, 0.0, // top-left
     0.5, -0.5, 0.0, // bottom-right
     0.5,  0.5, 0.0, // top-right
}
colors = []float32{
    // First triangle
    1.0, 0.0, 0.0, // red (top-left)
    0.0, 1.0, 0.0, // green (bottom-left)
    0.0, 0.0, 1.0, // blue (bottom-right)

    // Second triangle
    1.0, 0.0, 0.0, // red (top-left)
    0.0, 0.0, 1.0, // blue (bottom-right)
    1.0, 1.0, 0.0, // yellow (top-right)
}
)

// Create a function to generate the rotation matrix
func getRotationMatrix(angle float32) mgl32.Mat4 {
    return mgl32.HomogRotate3DY(angle) // Rotate around Z-axis
}

func getViewMatrix(angle float32) mgl32.Mat4 {
    // Camera position (orbiting around the square)
    radius := 2.0
    camX := float32(radius * math.Sin(float64(angle)))
    camZ := float32(radius * math.Cos(float64(angle)))

    // The camera is positioned at (camX, 0, camZ) looking at the origin (0, 0, 0)
    return mgl32.LookAt(
        camX, 0.0, camZ, // Camera position
        0.0, 0.0, 0.0,  // Look at the origin
        0.0, 1.0, 0.0,  // Up vector
    )
}

func getProjectionMatrix() mgl32.Mat4 {
    // Perspective projection with a 45-degree field of view
    fov := float32(mgl32.DegToRad(45.0))
    aspectRatio := float32(width) / float32(height)
    near := float32(0.1)
    far := float32(10.0)
    return mgl32.Perspective(fov, aspectRatio, near, far)
}
func main(){
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	vao := makeVao(square, colors)
	for !window.ShouldClose() {
		draw(vao, window, program)
	}
}



func draw(vao uint32, window *glfw.Window, program uint32) {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    gl.UseProgram(program)

    // Time-based angle for camera orbit
    angle := float32(glfw.GetTime())

    // Get view and projection matrices
    view := getViewMatrix(angle)
    projection := getProjectionMatrix()

    // Send the view and projection matrices to the shader
    viewLoc := gl.GetUniformLocation(program, gl.Str("view\x00"))
    projectionLoc := gl.GetUniformLocation(program, gl.Str("projection\x00"))
    modelLoc := gl.GetUniformLocation(program, gl.Str("model\x00"))

    // Set model to identity (no object transformation for now)
    model := mgl32.Ident4()

    // Send matrices to the shaders
    gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])
    gl.UniformMatrix4fv(projectionLoc, 1, false, &projection[0])
    gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

    // Draw the square
    gl.BindVertexArray(vao)
    gl.DrawArrays(gl.TRIANGLES, 0, 6)

    glfw.PollEvents()
    window.SwapBuffers()
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Go rendering engine", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.Enable(gl.DEPTH_TEST)
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

// makeVao initializes and returns a vertex array from the points provided.


func makeVao(vertices []float32, colors []float32) uint32 {
    var vboVertices, vboColors uint32
    
    // Vertex Buffer (VBO for positions)
    gl.GenBuffers(1, &vboVertices)
    gl.BindBuffer(gl.ARRAY_BUFFER, vboVertices)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

    var vao uint32
    gl.GenVertexArrays(1, &vao)
    gl.BindVertexArray(vao)

    // Set position attribute
    gl.EnableVertexAttribArray(0)
    gl.BindBuffer(gl.ARRAY_BUFFER, vboVertices)
    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

    // Color Buffer (VBO for colors)
    gl.GenBuffers(1, &vboColors)
    gl.BindBuffer(gl.ARRAY_BUFFER, vboColors)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(colors), gl.Ptr(colors), gl.STATIC_DRAW)

    // Set color attribute
    gl.EnableVertexAttribArray(1)
    gl.BindBuffer(gl.ARRAY_BUFFER, vboColors)
    gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 0, nil)

    return vao
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
