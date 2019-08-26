// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xe3\u2502\xd4\xd7\xf1\xf3\x04\xe0\x81\xc7Sw\xa0\x89\x0f\x06xcN\xba]\x8c qY*%\xc0^\x9e\x83\xc0\x00\x00"
const testnetAllocData = "\xe3\u2502\xd4\xd7\xf1\xf3\x04\xe0\x81\xc7Sw\xa0\x89\x0f\x06xcN\xba]\x8c qY*%\xc0^\x9e\x83\xc0\x00\x00"
const rinkebyAllocData = "\xe3\u2502\xd4\xd7\xf1\xf3\x04\xe0\x81\xc7Sw\xa0\x89\x0f\x06xcN\xba]\x8c qY*%\xc0^\x9e\x83\xc0\x00\x00"
const goerliAllocData = "\xe3\u2502\xd4\xd7\xf1\xf3\x04\xe0\x81\xc7Sw\xa0\x89\x0f\x06xcN\xba]\x8c qY*%\xc0^\x9e\x83\xc0\x00\x00"
