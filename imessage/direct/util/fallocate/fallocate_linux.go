// beeper-imessage - A Matrix-iMessage puppeting bridge.
// Copyright (C) 2023 Beeper, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

//go:build linux

package fallocate

import (
	"os"

	"golang.org/x/sys/unix"
)

var ErrOutOfSpace error = unix.ENOSPC

func Fallocate(file *os.File, size int) error {
	if size <= 0 {
		return nil
	}
	return unix.Fallocate(int(file.Fd()), 0, 0, int64(size))
}
