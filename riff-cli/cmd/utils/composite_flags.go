/*
 * Copyright 2018 the original author or authors.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package utils

import (
	"github.com/projectriff/riff/riff-cli/pkg/osutils"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func GetUseraccountWithOverride(name string, flagset pflag.FlagSet) string {
	userAcct := GetStringValueWithOverride("useraccount", flagset)
	if userAcct == DefaultValues.UserAccount {
		userAcct = osutils.GetCurrentUsername()
	}
	return userAcct
}

func GetStringValueWithOverride(name string, flagset pflag.FlagSet) string {
	viperVal := viper.GetString(name)
	val, _ := flagset.GetString(name)
	if flagset.Changed(name) || viperVal == "" {
		return val
	}
	return viperVal
}
