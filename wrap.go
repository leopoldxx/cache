/*
Copyright 2020 leopoldxx@gmail.com.

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

import "time"

func Wrap(c Interface) Interface {
	if c != nil {
		return c
	}
	return &empty{}
}

type empty struct{}

func (e *empty) Put(key Key, value Value)                             {}
func (e *empty) PutWithTimeout(key Key, value Value, t time.Duration) {}
func (e *empty) Get(key Key) (Value, bool)                            { return nil, false }
func (e *empty) Del(key Key) Value                                    { return nil }
func (e *empty) Len() int                                             { return 0 }
func (e *empty) Close()                                               {}
