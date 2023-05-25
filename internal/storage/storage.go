// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storage

import (
	"context"
	"io"
	"io/fs"

	"github.com/hashicorp/boundary/internal/boundary"
	"github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/storagebuckets"
	plgpb "github.com/hashicorp/boundary/sdk/pbs/plugin"
)

// RecordingStorage can be used to create a FS usable for session recording.
type RecordingStorage interface {
	// NewSyncingFS returns a FS that will use local storage as a cache and sync files when they are closed.
	NewSyncingFS(ctx context.Context, bucket *storagebuckets.StorageBucket, _ ...Option) (FS, error)

	// NewRemoteFS returns a ReadOnly FS that can be used to retrieve files from a storage bucket.
	NewRemoteFS(ctx context.Context, bucket *storagebuckets.StorageBucket, _ ...Option) (FS, error)

	// PluginClients returns a map of storage plugin clients keyed on the plugin name.
	PluginClients() map[string]plgpb.StoragePluginServiceClient
}

// Bucket is a resource that represents a bucket in an external object store
type Bucket interface {
	boundary.Resource
	GetScopeId() string
	GetBucketName() string
	GetBucketPrefix() string
	GetWorkerFilter() string
}

// FS is a filesystem for creating or reading files and containers.
type FS interface {
	New(ctx context.Context, name string) (Container, error)
	Open(ctx context.Context, name string) (Container, error)
}

// A Container is a filesystem abstraction that can create files or other containers.
type Container interface {
	io.Closer
	Create(context.Context, string) (File, error)
	OpenFile(context.Context, string, ...Option) (File, error)
	SubContainer(context.Context, string, ...Option) (Container, error)
}

// File represents a storage File.
type File interface {
	fs.File
	io.Writer
	io.StringWriter
}
