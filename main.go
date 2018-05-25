// Copyright © 2017 Free Chess Club <help@freechess.club>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/Sirupsen/logrus"
)

var (
	log      = logrus.New()
)

func main() {
	user := "guest"
	pass := ""
	s, err := newSession(user, pass, "67.188.115.30")
	if err != nil {
		log.WithField("err", err).Println("Failed to create a new session")
		return
	}
	for {
	}
	s.end()
}
