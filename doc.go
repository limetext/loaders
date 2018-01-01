// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

// Loaders package is reponsible for loading plist and json files
//
// We have a specific peg file for each(e.g plist/plist.peg) which based on them
// we generate for example plist/plist_generated.go. The generator is a pegparser
// written by @quarnster which can be found here: github.com/quarnster/parser/pegparser.
// About the plist loading note that first we are converting it to json and then
// we pass it to LoadJSON function, after some parsing there it is passed to
// golangs "encoding/json" package for unmarshaling.
//
// If you want to change the peg rules in peg file, you should at least pass two
// tests. First add sth related to the new rule in testdata/*.plist files if the
// generated tests got pass it means probably there isn't any problem in reading,
// in second step the plist_test.go should pass which is related to parsing stuff.
// Also you may want to check plistconv() function for the special things we do on
// some nodes.
//
package loaders
