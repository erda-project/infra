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

package i18n

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/erda-project/erda-infra/pkg/strutil"
)

type file struct {
	isDir    bool
	name     string
	fullPath string
	data     []byte
}

func newDir(fullPath string) *file {
	return &file{isDir: true, name: filepath.Base(fullPath), fullPath: fullPath, data: nil}
}
func newFile(fullPath string, data []byte) *file {
	return &file{isDir: false, name: filepath.Base(fullPath), fullPath: fullPath, data: data}
}

func (p *provider) RegisterFilesFromFS(fsPrefix string, rootFS embed.FS) error {
	var err error
	var filesItems []*file
	var commonItems []*file
	walkEmbedFS(rootFS, fmt.Sprintf("%s/common", fsPrefix), &commonItems)
	walkEmbedFS(rootFS, fmt.Sprintf("%s/files", fsPrefix), &filesItems)
	for _, file := range filesItems {
		if file.isDir {
			continue
		}
		if !strutil.HasSuffixes(file.name, ".yml", ".yaml") {
			continue
		}
		err = p.loadI18nFile(file.fullPath)
		if err != nil {
			return err
		}
	}
	for _, file := range commonItems {
		if file.isDir {
			continue
		}
		if !strutil.HasSuffixes(file.name, ".yml", ".yaml") {
			continue
		}
		err = p.loadToDic(file.fullPath, p.common)
		if err != nil {
			return err
		}
	}
	return nil
}

func walkEmbedFS(rootFS embed.FS, fullPath string, files *[]*file) {
	entries, err := fs.ReadDir(rootFS, fullPath)
	if err != nil {
		panic(fmt.Errorf("fullPath: %s, err: %v", fullPath, err))
	}
	for _, entry := range entries {
		entryPath := filepath.Join(fullPath, entry.Name())
		if !entry.IsDir() {
			data, err := rootFS.ReadFile(entryPath)
			if err != nil {
				panic(fmt.Errorf("failed to read file, filePath: %s, err: %v", entryPath, err))
			}
			*files = append(*files, newFile(entryPath, data))
			continue
		}
		*files = append(*files, newDir(entryPath))
		walkEmbedFS(rootFS, entryPath, files)
	}
}
