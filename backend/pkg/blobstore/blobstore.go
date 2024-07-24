/*
 * Copyright (c) 2024. SPK Platforms - All rights reserved
 * Proprietary and confidential
 * Written by Cyrus Pellet <cyrus@spkplatforms.co> in July 2024
 */

// Package blob provides an abstract interface for storing and retrieving blobs.
package blob

import (
	"context"
	"errors"
	"sort"
	"sync"
)

var ErrBlobNotFound = errors.New("storage object not found")

const (
	ContentTypeTextPlain = "text/plain"
	ContentTypePDF       = "application/pdf"
)

// Blobstore defines an abstract interface for storing and retrieving blobs.
type Blobstore interface {
	// CreateObject creates a new object with the given name and contents.
	CreateObject(ctx context.Context, parent, name string, contents []byte, cacheable bool, contentType string) error
	// DeleteObject deletes the object with the given name, if it exists.
	DeleteObject(ctx context.Context, parent, name string) error
	// GetObject retrieves the contents of the object with the given name.
	GetObject(ctx context.Context, parent, name string) ([]byte, error)
	// GetPath returns the path of a local file that contains the object with the given name.
	GetPath(ctx context.Context, parent, name string) (string, error)
}

// BlobstoreFunc is a function that creates a new Blobstore.
type BlobstoreFunc func(context.Context, *Config) (Blobstore, error)

// blobstores is a map of registered Blobstore constructors.
var (
	blobstores     = make(map[string]BlobstoreFunc)
	blobstoresLock sync.RWMutex
)

// RegisterBlobstore registers a new Blobstore constructor with the given name. If a constructor with the same name is already registered, it will panic.
func RegisterBlobstore(name string, f BlobstoreFunc) {
	blobstoresLock.Lock()
	defer blobstoresLock.Unlock()
	if _, ok := blobstores[name]; ok {
		panic("blobstore already registered: " + name)
	}
	blobstores[name] = f
}

// RegisteredBlobstores returns a sorted list of registered Blobstore names.
func RegisteredBlobstores() []string {
	blobstoresLock.RLock()
	defer blobstoresLock.RUnlock()
	names := make([]string, 0, len(blobstores))
	for name := range blobstores {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// BlobstoreFor returns the Blobstore with the given name, or an error if one does not exist.
func BlobstoreFor(ctx context.Context, cfg *Config) (Blobstore, error) {
	blobstoresLock.RLock()
	defer blobstoresLock.RUnlock()
	f, ok := blobstores[cfg.Type]
	if !ok {
		return nil, errors.New("unknown or unconfigured blobstore type: " + cfg.Type)
	}
	return f(ctx, cfg)
}
