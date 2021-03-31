// Author: recallsong
// Email: songruiguo@qq.com

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	transhttpPackage = protogen.GoImportPath("github.com/erda-project/erda-infra/pkg/transport/http")
	transgrpcPackage = protogen.GoImportPath("github.com/erda-project/erda-infra/pkg/transport/grpc")
	reflectPackage   = protogen.GoImportPath("reflect")
)

func generateFiles(gen *protogen.Plugin, files []*protogen.File) (*protogen.GeneratedFile, error) {
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
	g.P("// Source: ", strings.Join(sources, ", "))
	g.P()
	g.P("package ", file.GoPackageName)
	g.P("// RegisterServices register all services.")
	g.P("func RegisterServices(router_ ", transhttpPackage.Ident("Router"), ", server_ ", transgrpcPackage.Ident("ServiceRegistrar"), ",")
	for _, file := range files {
		g.P("// ", file.Desc.Path())
		for _, ser := range file.Services {
			g.P(lowerCaptain(ser.GoName), " ", ser.GoName, "Server,")
		}
	}
	g.P(") {")
	for _, file := range files {
		g.P("// ", file.Desc.Path())
		for _, ser := range file.Services {
			g.P("Register", ser.GoName, "Handler(router_, ", ser.GoName, "Handler(", lowerCaptain(ser.GoName), "))")
			g.P("Register", ser.GoName, "Server(server_, ", lowerCaptain(ser.GoName), ")")
		}
	}
	g.P("}")
	g.P()
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
			g.P(lowerCaptain(ser.GoName+"ClientType"), " = ", reflectPackage.Ident("TypeOf"), "((*", file.GoImportPath.Ident(ser.GoName+"Client"), ")(nil)).Elem()")
			g.P(lowerCaptain(ser.GoName+"ServerType"), " = ", reflectPackage.Ident("TypeOf"), "((*", file.GoImportPath.Ident(ser.GoName+"Server"), ")(nil)).Elem()")
		}
		g.P()
	}
	g.P(")")

	for _, file := range files {
		for _, ser := range file.Services {
			g.P("// ", ser.GoName+"ClientType .")
			g.P("func ", ser.GoName+"ClientType() ", reflectPackage.Ident("Type"), " { return ", lowerCaptain(ser.GoName+"ClientType"), "}")
			g.P("// ", ser.GoName+"ServerType .")
			g.P("func ", ser.GoName+"ServerType() ", reflectPackage.Ident("Type"), " { return ", lowerCaptain(ser.GoName+"ServerType"), "}")
		}
		g.P()
	}
	g.P("func Types() []", reflectPackage.Ident("Type"), "{")
	g.P("	return []", reflectPackage.Ident("Type"), "{")
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
