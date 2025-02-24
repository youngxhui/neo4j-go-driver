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
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testutil

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/internal/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/log"
)

type RouterFake struct {
	Invalidated            bool
	InvalidatedDb          string
	GetOrUpdateReadersRet  []string
	GetOrUpdateReadersHook func(bookmarks func(context.Context) ([]string, error), database string) ([]string, error)
	GetOrUpdateWritersRet  []string
	GetOrUpdateWritersHook func(bookmarks func(context.Context) ([]string, error), database string) ([]string, error)
	Err                    error
	CleanUpHook            func()
	GetNameOfDefaultDbHook func(user string) (string, error)
	InvalidatedServer      string
}

func (r *RouterFake) InvalidateReader(ctx context.Context, database string, server string) error {
	if err := r.Invalidate(ctx, database); err != nil {
		return err
	}
	r.InvalidatedServer = server
	return nil
}

func (r *RouterFake) InvalidateWriter(context.Context, string, string) error {
	return nil
}

func (r *RouterFake) Invalidate(ctx context.Context, database string) error {
	r.InvalidatedDb = database
	r.Invalidated = true
	return nil
}

func (r *RouterFake) GetOrUpdateReaders(_ context.Context, bookmarksFn func(context.Context) ([]string, error), database string, _ *db.ReAuthToken, _ log.BoltLogger) ([]string, error) {
	if r.GetOrUpdateReadersHook != nil {
		return r.GetOrUpdateReadersHook(bookmarksFn, database)
	}
	return r.GetOrUpdateReadersRet, r.Err
}

func (r *RouterFake) Readers(context.Context, string) ([]string, error) {
	return nil, nil
}

func (r *RouterFake) GetOrUpdateWriters(_ context.Context, bookmarksFn func(context.Context) ([]string, error), database string, _ *db.ReAuthToken, _ log.BoltLogger) ([]string, error) {
	if r.GetOrUpdateWritersHook != nil {
		return r.GetOrUpdateWritersHook(bookmarksFn, database)
	}
	return r.GetOrUpdateWritersRet, r.Err
}

func (r *RouterFake) Writers(context.Context, string) ([]string, error) {
	return nil, nil
}

func (r *RouterFake) GetNameOfDefaultDatabase(_ context.Context, _ []string, user string, _ *db.ReAuthToken, _ log.BoltLogger) (string, error) {
	if r.GetNameOfDefaultDbHook != nil {
		return r.GetNameOfDefaultDbHook(user)
	}
	return "", nil
}

func (r *RouterFake) CleanUp(ctx context.Context) error {
	if r.CleanUpHook != nil {
		r.CleanUpHook()
	}
	return nil
}
