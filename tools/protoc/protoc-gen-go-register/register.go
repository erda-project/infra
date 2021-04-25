// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	transportPackage = protogen.GoImportPath("github.com/erda-project/erda-infra/pkg/transport")
	transhttpPackage = protogen.GoImportPath("github.com/erda-project/erda-infra/pkg/transport/http")
	transgrpcPackage = protogen.GoImportPath("github.com/erda-project/erda-infra/pkg/transport/grpc")
	reflectPackage   = protogen.GoImportPath("reflect")
)

func generateFiles(gen *protogen.Plugin, flags flag.FlagSet, files []*protogen.File) (*protogen.GeneratedFile, error) {

	sort.Slice(files, func(i, j int) bool {
		return files[i].Desc.Name() < files[j].Desc.Name()
	})
	if len(files) <= 0 {
		return nil, nil
	}
	var file *protogen.File
	var count int
	var sources, services []string
	for _, f := range files {
		if len(f.Services) <= 0 {
			continue
		}
		count += len(f.Services)
		if file == nil {
			file = f
		}
		if f.GoImportPath != file.GoImportPath {
			return nil, fmt.Errorf("package path conflict between %s and %s", file.GoImportPath, f.GoImportPath)
		}
		for _, ser := range f.Services {
			services = append(services, strings.TrimRight(string(f.Desc.Package()), ".")+"."+ser.GoName)
		}
		sources = append(sources, f.Desc.Path())
	}
	if count <= 0 {
		return nil, nil
	}
	sort.Strings(services)

	const filename = "register.services.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by ", genName, ". DO NOT EDIT.")
	g.P("// Sources: ", strings.Join(sources, ", "))
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	if *genGRPC && *genHTTP {
		for _, file := range files {
			for _, ser := range file.Services {
				g.P("// Register", ser.GoName, "Imp ", file.Desc.Path())
				g.P("func Register", ser.GoName, "Imp(regester ", transportPackage.Ident("Register"), ", srv ", ser.GoName, "Server, opts ...", transportPackage.Ident("ServiceOption"), ") {")
				g.P("	_ops := ", transportPackage.Ident("DefaultServiceOptions"), "()")
				g.P("	for _, op := range opts {")
				g.P("		op(_ops)")
				g.P("	}")
				g.P("	Register", ser.GoName, "Handler(regester, ", ser.GoName, "Handler(srv), _ops.HTTP...)")
				g.P("	Register", ser.GoName, "Server(regester, srv, _ops.GRPC...)")
				g.P("}")
			}
			g.P()
		}
	} else if *genGRPC {
		for _, file := range files {
			for _, ser := range file.Services {
				g.P("// Register", ser.GoName, "Imp ", file.Desc.Path())
				g.P("func Register", ser.GoName, "Imp(regester ", transgrpcPackage.Ident("ServiceRegistrar"), ", srv ", ser.GoName, "Server, opts ...", transgrpcPackage.Ident("HandleOption"), ") {")
				g.P("	Register", ser.GoName, "Server(regester, srv, opts...)")
				g.P("}")
			}
			g.P()
		}
	} else if *genHTTP {
		for _, file := range files {
			for _, ser := range file.Services {
				g.P("// Register", ser.GoName, "Imp ", file.Desc.Path())
				g.P("func Register", ser.GoName, "Imp(regester ", transhttpPackage.Ident("Register"), ", srv ", ser.GoName, "Handler, opts ...", transhttpPackage.Ident("HandleOption"), ") {")
				g.P("	Register", ser.GoName, "Handler(regester, srv, opts...)")
				g.P("}")
			}
			g.P()
		}
	}

	g.P("// ServiceNames return all service names")
	g.P("func ServiceNames(svr ...string) []string {")
	g.P("	return append(svr,")
	for _, s := range services {
		g.P("	", strconv.Quote(s), ",")
	}
	g.P("	)")
	g.P("}")
	g.P()
	g.P("var (")
	for _, file := range files {
		for _, ser := range file.Services {
			if *genGRPC {
				g.P(lowerCaptain(ser.GoName+"ClientType"), " = ", reflectPackage.Ident("TypeOf"), "((*", file.GoImportPath.Ident(ser.GoName+"Client"), ")(nil)).Elem()")
				g.P(lowerCaptain(ser.GoName+"ServerType"), " = ", reflectPackage.Ident("TypeOf"), "((*", file.GoImportPath.Ident(ser.GoName+"Server"), ")(nil)).Elem()")
			}
			if *genHTTP {
				g.P(lowerCaptain(ser.GoName+"HandlerType"), " = ", reflectPackage.Ident("TypeOf"), "((*", file.GoImportPath.Ident(ser.GoName+"Handler"), ")(nil)).Elem()")
			}
		}
		g.P()
	}
	g.P(")")

	for _, file := range files {
		for _, ser := range file.Services {
			if *genGRPC {
				g.P("// ", ser.GoName+"ClientType .")
				g.P("func ", ser.GoName+"ClientType() ", reflectPackage.Ident("Type"), " { return ", lowerCaptain(ser.GoName+"ClientType"), "}")
				g.P("// ", ser.GoName+"ServerType .")
				g.P("func ", ser.GoName+"ServerType() ", reflectPackage.Ident("Type"), " { return ", lowerCaptain(ser.GoName+"ServerType"), "}")
			}
			if *genHTTP {
				g.P("// ", ser.GoName+"HandlerType .")
				g.P("func ", ser.GoName+"HandlerType() ", reflectPackage.Ident("Type"), " { return ", lowerCaptain(ser.GoName+"HandlerType"), "}")
			}
		}
		g.P()
	}
	g.P("func Types() []", reflectPackage.Ident("Type"), "{")
	g.P("	return []", reflectPackage.Ident("Type"), "{")
	if *genGRPC {
		g.P("// client types")
		for _, file := range files {
			for _, ser := range file.Services {
				g.P("	", lowerCaptain(ser.GoName+"ClientType"), ",")
			}
		}
		g.P("// server types")
		for _, file := range files {
			for _, ser := range file.Services {
				g.P("	", lowerCaptain(ser.GoName+"ServerType"), ",")
			}
		}
	}
	if *genHTTP {
		g.P("// handler types")
		for _, file := range files {
			for _, ser := range file.Services {
				g.P("	", lowerCaptain(ser.GoName+"HandlerType"), ",")
			}
		}
	}
	g.P("	}")
	g.P("}")
	return g, nil
}

func lowerCaptain(name string) string {
	if len(name) <= 0 {
		return name
	}
	chars := []rune(name)
	pre := chars[0]
	if unicode.IsLower(pre) {
		return name
	}
	for i, c := range chars {
		if unicode.IsUpper(c) != unicode.IsUpper(pre) {
			break
		}
		chars[i] = unicode.ToLower(c)
	}
	return string(chars)
}
