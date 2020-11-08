// +build wasm
package git

func init() {
	// TODO: does was need this?
	// fillSystemInfo = func(e *index.Entry, sys interface{}) {
	// }
}

func isSymlinkWindowsNonAdmin(err error) bool {
	// TODO: is this false
	return false
}
