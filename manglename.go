// MANGLENAME -- mangle file names (except their extensions)
// SvM 31-DEC-2020 - 15-JUL-2021
//
// -u: uppercase
// -l: lowercase
// -c: capitalise
// -e: also extension
// -L lang: select language
// -v: verbose
// -h: help
//
// if called as capitalise, run as if given -c

package main

import (
	"flag"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"log"
	"os"
	"path"
	"strings"
)

func b2i(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func main() {
	var u, l, c, e, v bool
	var lng string

	moi := path.Base(os.Args[0])
	lg := log.New(os.Stderr, moi+": ", 0)

	flag.Usage = func() {
		if moi == "capitalise" {
			lg.Printf("capitalise file names. See manglename(1) if you want more options.\n")
		} else {
			lg.Printf("change capitalisation of file names.\n\nOptions (exactly one of -u, -l or -c required):\n\n")
		}
		flag.PrintDefaults()
	}

	if moi == "capitalise" {
		u = false
		l = false
		c = true
		e = false
		v = false
		lng = "en"
	} else {
		flag.BoolVar(&u, "u", false, "convert to uppercase")
		flag.BoolVar(&l, "l", false, "convert to lowercase")
		flag.BoolVar(&c, "c", false, "capitalise")
		flag.BoolVar(&e, "e", false, "rename extension as well")
		flag.StringVar(&lng, "L", "en", "select language")
		flag.BoolVar(&v, "v", false, "verbose output")
	}
	flag.Parse()

	if x := b2i(u) + b2i(l) + b2i(c); x != 1 {
		lg.Printf("use one and only one of -u, -l or -c\n")
		os.Exit(1)
	}

	lngtag := language.Make(lng)
	if v {
		lg.Println("language:", display.English.Tags().Name(lngtag))
	}

	for _, old := range flag.Args() {
		old = path.Clean(old)

		if _, err := os.Stat(old); err == nil {
			// file old exists:
			dir := path.Dir(old)
			ext := path.Ext(old)
			base := strings.TrimSuffix(path.Base(old), ext)

			// set which caser function we need
			var cs cases.Caser
			if u {
				cs = cases.Upper(lngtag)
			} else if l {
				cs = cases.Lower(lngtag)
			} else if c {
				cs = cases.Title(lngtag)
			}

			// do the conversion and add all the bits back together
			base = cs.String(base)
			if e {
				ext = cs.String(ext)
			}

			new := path.Clean(dir + "/" + base + ext)

			if old == new {
				// no change from old to new: skip with warning (if being verbose)
				if v {
					lg.Printf("file %s already in desired format", old)
				}
			} else if _, err := os.Stat(new); err == nil {
				// file new alread exists: skip with warning
				lg.Printf("not overwriting existing file %s\n", new)
			} else {
				// file new doesn't exist: go for it
				if v {
					lg.Printf("rename %s to %s\n", old, new)
				}
				err := os.Rename(old, new)
				if err != nil {
					lg.Println(err)
				}
			}
		} else {
			// file old doesn't exist: nothing to do
			lg.Printf("file %s does not exist\n", old)
		}
	}
}
