package demoinfocs_test

import (
	"fmt"
	"os"
	"testing"

	dem "github.com/visual42/demoinfocs-golang"
	events "github.com/visual42/demoinfocs-golang/events"
)

func TestReadmeExample(t *testing.T) {
	f, err := os.Open(defaultDemPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := dem.NewParser(f)

	// Register handler on kill events
	p.RegisterEventHandler(func(e events.Kill) {
		var hs string
		if e.IsHeadshot {
			hs = " (HS)"
		}
		var wallBang string
		if e.PenetratedObjects > 0 {
			wallBang = " (WB)"
		}
		fmt.Printf("%s <%v%s%s> %s\n", e.Killer.Name, e.Weapon.Weapon, hs, wallBang, e.Victim.Name)
	})

	// Parse to end
	err = p.ParseToEnd()
	if err != nil {
		panic(err)
	}
}
