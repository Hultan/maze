package maze

import (
	"os"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
	mazeGen "github.com/hultan/maze/internal/maze-gen"

	"github.com/hultan/softteam/framework"
)

const applicationTitle = "maze"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	Window      *gtk.ApplicationWindow
	builder     *framework.GtkBuilder
	AboutDialog *gtk.AboutDialog
}

var maze *mazeGen.Maze

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	fw := framework.NewFramework()
	builder, err := fw.Gtk.CreateBuilder("main.glade")
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file
	m.Window = m.builder.GetObject("mainWindow").(*gtk.ApplicationWindow)

	// Set up main window
	m.Window.SetApplication(app)
	m.Window.SetTitle("maze main window")
	m.Window.Maximize()

	// Hook up the destroy event
	m.Window.Connect("destroy", m.Window.Close)

	// Quit button
	button := m.builder.GetObject("main_window_quit_button").(*gtk.ToolButton)
	button.Connect("clicked", m.Window.Close)

	// Status bar
	statusBar := m.builder.GetObject("main_window_status_bar").(*gtk.Statusbar)
	statusBar.Push(statusBar.GetContextId("maze"), "maze : version 0.1.0")

	// Menu
	m.setupMenu()

	maze = mazeGen.NewMaze()

	da := m.builder.GetObject("drawingArea").(*gtk.DrawingArea)
	da.SetVisible(true)
	da.Connect("draw", onDraw)

	// Show the main window
	m.Window.ShowAll()
}

func (m *MainForm) setupMenu() {
	menuQuit := m.builder.GetObject("menu_file_quit").(*gtk.MenuItem)
	menuQuit.Connect("activate", m.Window.Close)
}

func onDraw(da *gtk.DrawingArea, ctx *cairo.Context) {
	size := float64(da.GetAllocatedHeight() / 30)

	for x := float64(0); x < 25; x++ {
		for y := float64(0); y < 25; y++ {
			w := maze[int(y)][int(x)].Walls
			if w&mazeGen.North != 0 {
				ctx.MoveTo((x+1)*size, (y+1)*size)
				ctx.LineTo((x+2)*size, (y+1)*size)
			}
			if w&mazeGen.East != 0 {
				ctx.MoveTo((x+2)*size, (y+1)*size)
				ctx.LineTo((x+2)*size, (y+2)*size)
			}
			if w&mazeGen.South != 0 {
				ctx.MoveTo((x+1)*size, (y+2)*size)
				ctx.LineTo((x+2)*size, (y+2)*size)
			}
			if w&mazeGen.West != 0 {
				ctx.MoveTo((x+1)*size, (y+1)*size)
				ctx.LineTo((x+1)*size, (y+2)*size)
			}
			ctx.Stroke()
		}
	}

	// Draw green starting position
	ctx.SetSourceRGB(0, 255, 0)
	ctx.Rectangle(size+3, size+3, size-6, size-6)
	ctx.Fill()

	// Draw blue starting position
	ctx.SetSourceRGB(0, 0, 255)
	ctx.Rectangle(size*25+3, size*25+3, size-6, size-6)
	ctx.Fill()

}
