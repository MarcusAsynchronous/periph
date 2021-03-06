// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// +build !linux

package sysfs

import "errors"

type event struct{}

func (e event) makeEvent(f uintptr) error {
	return errors.New("sysfs-gpio: unreachable code")
}

func (e event) wait(timeoutms int) (int, error) {
	return 0, errors.New("sysfs-gpio: unreachable code")
}

func isErrBusy(err error) bool {
	// This function is not used on non-linux.
	return false
}
