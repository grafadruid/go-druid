//go:build mage
// +build mage

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var ldflags = ""

// allow user to override go executable by running as GOEXE=xxx make ... on unix-like systems
var goexe = "go"

// Build is the default that fmt, vet, runs test and builds
var Default = Build

func init() {
	if exe := os.Getenv("GOEXE"); exe != "" {
		goexe = exe
	}

	// We want to use Go 1.11 modules even if the source lives inside GOPATH.
	// The default is "auto".
	os.Setenv("GO111MODULE", "on")
}

// Fmt run gofmt linter
func Fmt() error {
	if !isGoLatest() {
		return nil
	}
	pkgs, err := packages()
	if err != nil {
		return err
	}
	failed := false
	first := true
	for _, pkg := range pkgs {
		files, err := filepath.Glob(filepath.Join(pkg, "*.go"))
		if err != nil {
			return nil
		}
		for _, f := range files {
			// gofmt doesn't exit with non-zero when it finds unformatted code
			// so we have to explicitly look for output, and if we find any, we
			// should fail this target.
			s, err := sh.Output("gofmt", "-l", f)
			if err != nil {
				fmt.Printf("ERROR: running gofmt on %q: %v\n", f, err)
				failed = true
			}
			if s != "" {
				if first {
					fmt.Println("The following files are not gofmt'ed:")
					first = false
				}
				failed = true
				fmt.Println(s)
			}
		}
	}
	if failed {
		return errors.New("improperly formatted go files")
	}
	return nil
}

// Vet run go vet linter
func Vet() error {
	if err := sh.Run(goexe, "vet", "./..."); err != nil {
		return fmt.Errorf("error running go vet: %v", err)
	}
	return nil
}

// Run tests
func Test() error {
	env := map[string]string{"GOFLAGS": testGoFlags()}
	return runCmd(env, goexe, "test", "./...", buildFlags(), "-tags", buildTags())
}

// TestRace run tests with race detector
func TestRace() error {
	env := map[string]string{"GOFLAGS": testGoFlags()}
	return runCmd(env, goexe, "test", "-race", "./...", buildFlags(), "-tags", buildTags())
}

// TestCoverHTML generates test coverage report
func TestCoverHTML() error {
	const (
		coverAll = "coverage-all.out"
		cover    = "coverage.out"
	)
	f, err := os.Create(coverAll)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte("mode: count")); err != nil {
		return err
	}
	pkgs, err := packages()
	if err != nil {
		return err
	}
	for _, pkg := range pkgs {
		if err := sh.Run(goexe, "test", "-coverprofile="+cover, pkg); err != nil {
			return err
		}
		b, err := ioutil.ReadFile(cover)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		idx := bytes.Index(b, []byte{'\n'})
		b = b[idx+1:]
		if _, err := f.Write(b); err != nil {
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return sh.Run(goexe, "tool", "cover", "-html="+coverAll)
}

// Build run linters, vet and tests
func Build() error {
	if strings.Contains(runtime.Version(), "1.8") {
		// Go 1.8 doesn't play along with go test ./... and /vendor.
		// We could fix that, but that would take time.
		fmt.Printf("Skip Build on %s\n", runtime.Version())
		return nil
	}

	// TODO: Add lint after fixing errors
	mg.Deps(Fmt, Vet, TestRace)
	return sh.RunV(goexe, "build", "-ldflags="+ldflags, "./...")
}

var (
	pkgPrefixLen = len("github.com/h2oai/go-druid")
	pkgs         []string
	pkgsInit     sync.Once
)

// testGoFlags returns test flags that need to be set
func testGoFlags() string {
	return "-v"
}

func packages() ([]string, error) {
	var err error
	pkgsInit.Do(func() {
		var s string
		s, err = sh.Output(goexe, "list", "./...")
		if err != nil {
			return
		}
		pkgs = strings.Split(s, "\n")
		for i := range pkgs {
			pkgs[i] = "." + pkgs[i][pkgPrefixLen:]
		}
	})
	return pkgs, err
}

func buildFlags() []string {
	if runtime.GOOS == "windows" {
		return []string{"-buildmode", "exe"}
	}
	return nil
}

func buildTags() string {
	return "none"
}

func isGoLatest() bool {
	return strings.Contains(runtime.Version(), "1.21")
}

func runCmd(env map[string]string, cmd string, args ...interface{}) error {
	if mg.Verbose() {
		return runWith(env, cmd, args...)
	}
	output, err := sh.OutputWith(env, cmd, argsToStrings(args...)...)
	if err != nil {
		fmt.Fprint(os.Stderr, output)
	}

	return err
}

func runWith(env map[string]string, cmd string, inArgs ...interface{}) error {
	s := argsToStrings(inArgs...)
	return sh.RunWith(env, cmd, s...)
}

func argsToStrings(v ...interface{}) []string {
	var args []string
	for _, arg := range v {
		switch v := arg.(type) {
		case string:
			if v != "" {
				args = append(args, v)
			}
		case []string:
			if v != nil {
				args = append(args, v...)
			}
		default:
			panic("invalid type")
		}
	}

	return args
}
