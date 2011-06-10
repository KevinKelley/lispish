# Copyright 2010 Petar Maymounkov. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

all:	install

install:
	cd stutter && make install

clean:
	cd stutter && make clean

nuke:
	cd stutter && make nuke
