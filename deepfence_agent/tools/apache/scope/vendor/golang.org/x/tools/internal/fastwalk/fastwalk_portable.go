// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine !linux,!darwin,!freebsd,!openbsd,!netbsd

package fastwalk

import (
	"os"
)

// readDir calls fn for each directory entry in dirName.
// It does not descend into directories or follow symlinks.
// If fn returns a non-nil error, readDir returns with that error
// immediately.
func readDir(dirName string, fn func(dirName, entName string, typ os.FileMode) error) error {
	fis, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}
	skipFiles := false
	for _, fi := range fis {
		fInfo, _ := fi.Info()
		if fInfo.Mode().IsRegular() && skipFiles {
			continue
		}
		if err := fn(dirName, fi.Name(), fInfo.Mode()&os.ModeType); err != nil {
			if err == SkipFiles {
				skipFiles = true
				continue
			}
			return err
		}
	}
	return nil
}
