/*
Copyright 2019-2020 vChain, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache

import (
	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

const ROOT_FN = ".root"

type fileCache struct {
}

func NewFileCache() Cache {
	return &fileCache{}
}

func (w *fileCache) Get() (*schema.Root, error) {
	root := new(schema.Root)
	buf, err := ioutil.ReadFile(ROOT_FN)
	if err == nil {
		if err = proto.Unmarshal(buf, root); err != nil {
			return nil, err
		}
		return root, nil
	}
	return nil, err
}

func (w *fileCache) Set(root *schema.Root) error {
	raw, err := proto.Marshal(root)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ROOT_FN, raw, 0644)
	if err != nil {
		return err
	}
	return nil
}
