/*
 * Copyright (c) 2024. SPK Platforms - All rights reserved
 * Proprietary and confidential
 * Written by Cyrus Pellet <cyrus@spkplatforms.co> in July 2024
 */

package blob

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func init() {
	RegisterBlobstore("FILESYSTEM", NewFilesystem)
}

// Compile time check to ensure that Filesystem implements the Blobstore interface.
var _ Blobstore = (*Filesystem)(nil)

// Filesystem is a simple filesystem blobstore.
type Filesystem struct {
	lock sync.Mutex
	root string
}

// NewFilesystem creates a new Filesystem blobstore.
func NewFilesystem(ctx context.Context, cfg *Config) (Blobstore, error) {
	return &Filesystem{
		root: cfg.RootDir,
	}, nil
}

// CreateObject creates a new object with the given name and contents. contentType is ignored.
func (s *Filesystem) CreateObject(ctx context.Context, folder, filename string, contents []byte, cacheable bool, contentType string) error {
	pth := filepath.Join(s.root, folder, filename)
	if err := os.MkdirAll(filepath.Dir(pth), 0o700); err != nil {
		return fmt.Errorf("failed to create object: %w", err)
	}
	if err := os.WriteFile(pth, contents, 0o600); err != nil {
		return fmt.Errorf("failed to create object: %w", err)
	}
	return nil
}

// DeleteObject deletes an object. It returns nil if the object was deleted or if the object no longer exists.
func (s *Filesystem) DeleteObject(ctx context.Context, folder, filename string) error {
	pth := filepath.Join(s.root, folder, filename)
	if err := os.Remove(pth); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete object: %w", err)
	}
	return nil
}

// GetObject returns the contents for the given object. If the object does not exist, it returns ErrBlobNotFound.
func (s *Filesystem) GetObject(ctx context.Context, folder, filename string) ([]byte, error) {
	pth := filepath.Join(s.root, folder, filename)
	b, err := os.ReadFile(pth)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrBlobNotFound
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return b, nil
}

// GetPath returns the path of a local file that contains the object with the given name.
func (s *Filesystem) GetPath(ctx context.Context, folder, filename string) (string, error) {
	pth := filepath.Join(s.root, folder, filename)
	return pth, nil
}
