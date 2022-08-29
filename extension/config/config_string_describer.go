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
// +ioc:autowire:paramType=ConfigString
// +ioc:autowire:constructFunc=new
// +ioc:autowire:proxy:autoInjection=false

type ConfigString string

func (ci *ConfigString) Value() string {
	return string(*ci)
}

func (ci *ConfigString) new(impl *ConfigString) (*ConfigString, error) {
	*impl = *ci
	return impl, nil
}

func FromString(val string) *ConfigString {
	configInt := ConfigString(val)
	return &configInt
}
