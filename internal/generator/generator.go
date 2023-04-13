package generator

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"time"

	"github.com/kwitsch/dnsroot/internal/rootfile"
	"github.com/kwitsch/dnsroot/internal/util"
)

//go:embed dnsroot.tmpl
var content embed.FS

func Run() error {
	cofv, cofvCheck := util.GetCurrentOutputFileVersion()

	rf, err := rootfile.Get()
	if err != nil {
		return err
	}

	if cofvCheck && rf.Version == cofv {
		fmt.Println("No changes necessary")
		return nil
	}

	if err := util.RemoveOutputFile(); err != nil {
		// can't remove old file
		return err
	}

	tmpl := template.Must(template.ParseFS(content, "*.tmpl"))

	outBuff := bytes.NewBuffer([]byte{})
	pkg, _ := util.GetGoPackage()
	err = tmpl.Execute(outBuff, map[string]interface{}{
		"package":  pkg,
		"dversion": util.ProgramVersion,
		"dupdate":  time.Now().Format("2006-01-02"),
		"rootfile": rf,
	})
	if err != nil {
		return err
	}

	return util.WriteOutputFile(outBuff.Bytes())
}
