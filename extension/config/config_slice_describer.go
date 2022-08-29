/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

// +ioc:autowire=true
// +ioc:autowire:baseType=true
// +ioc:autowire:type=config
// +ioc:autowire:paramType=ConfigSlice
// +ioc:autowire:constructFunc=new
// +ioc:autowire:proxy:autoInjection=false

type ConfigSlice []interface{}

func (cs *ConfigSlice) Value() []interface{} {
	return *cs
}

func (ci *ConfigSlice) new(impl *ConfigSlice) (*ConfigSlice, error) {
	*impl = *ci
	return impl, nil
}

func FromSlice(val []interface{}) *ConfigSlice {
	configSlice := ConfigSlice(val)
	return &configSlice
}
