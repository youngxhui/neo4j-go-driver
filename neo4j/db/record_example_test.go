/*
 * Copyright (c) "Neo4j"
 * Neo4j Sweden AB [https://neo4j.com]
 *
 * This file is part of Neo4j.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package db

import (
	"encoding/json"
	"fmt"
)

func ExampleRecord_AsMap() {
	record := Record{
		Values: []any{true, 42, "yes"},
		Keys:   []string{"prop1", "prop2", "prop3"},
	}

	dictionary := record.AsMap()

	output, _ := json.Marshal(dictionary)
	fmt.Println(string(output))
	// output: {"prop1":true,"prop2":42,"prop3":"yes"}
}
