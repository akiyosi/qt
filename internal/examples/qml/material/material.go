package main

import (
	"os"

	"github.com/akiyosi/qt/core"
	"github.com/akiyosi/qt/gui"
	"github.com/akiyosi/qt/qml"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:/qml/material.qml", 0))

	gui.QGuiApplication_Exec()
}
