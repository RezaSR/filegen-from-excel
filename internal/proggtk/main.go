package proggtk

import (
	"strconv"

	"github.com/gotk3/gotk3/gtk"
	"github.com/rezasr/filegen-from-excel/internal/prog"
)

var (
	bldr *gtk.Builder

	signals = map[string]interface{}{
		"Btn_HelpTemplateFile_clk": btn_HelpTemplateFile_Clicked,
		"Btn_HelpOutFileName_clk":  btn_HelpOutFileName_Clicked,
		"Btn_Generate_clk":         btn_Generate_Clicked,
		"Btn_DialogMain_clk":       btn_DialogMain_Clicked,
		"Menu_Quit_act":            menu_Quit_Activated,
		"Menu_About_act":           menu_About_Activated,
	}
)

func Main() error {
	var err error

	gtk.Init(nil)

	bldr, err = gtk.BuilderNew()
	if err != nil {
		return err
	}

	err = bldr.AddFromString(gladeStr)
	if err != nil {
		return err
	}

	err = initObjects()
	if err != nil {
		return err
	}

	initVars()

	bldr.ConnectSignals(signals)

	appWindow_Main.ShowAll()
	gtk.Main()

	return nil
}

func initVars() error {
	outDirPath := prog.NormalizePath(prog.DefaultOutDir)
	err := prog.DirExists(outDirPath, true)
	if err == nil {
		fileBtn_OutDir.SetFilename(outDirPath)
	}

	entry_OutFileName.SetText(prog.DefaultOutFileName)

	return nil
}

func showDialog(title, header, text string) {
	dialog_Main.SetTitle(title)
	lbl_DialogMainHeader.SetText(header)
	lbl_DialogMainText.SetText(text)
	dialog_Main.Show()
}

func btn_HelpTemplateFile_Clicked() {
	showDialog("Help", "Template File", prog.UsageTemplateFile)
}

func btn_HelpOutFileName_Clicked() {
	showDialog("Help", "Output File Name", prog.UsageOutFileName)
}

func btn_Generate_Clicked() {
	prog.DataFile.Set(fileBtn_DataFile.GetFilename())
	prog.TemplateFile.Set(fileBtn_TemplateFile.GetFilename())
	prog.OutDir.Set(fileBtn_OutDir.GetFilename())
	str, err := entry_OutFileName.GetText()
	if err == nil {
		prog.OutFileName.Set(str)
	} else {
		prog.OutFileName.Set(prog.DefaultOutFileName)
	}

	n, err := prog.Main()
	if err == nil {
		showDialog("Success", "File generation was successful", "Number of generated files: "+strconv.Itoa(n))
	} else {
		showDialog("Error", "Error generating files:", err.Error())
	}
}

func btn_DialogMain_Clicked() {
	dialog_Main.Hide()
}

func menu_Quit_Activated() {
	gtk.MainQuit()
}

func menu_About_Activated() {
	aboutDialog_Main.Show()
}
