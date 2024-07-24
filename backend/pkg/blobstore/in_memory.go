/*
 * Copyright (c) 2024. SPK Platforms - All rights reserved
 * Proprietary and confidential
 * Written by Cyrus Pellet <cyrus@spkplatforms.co> in July 2024
 */

package blob

import (
	"context"
	"os"
	"path"
	"sync"
)

func init() {
	RegisterBlobstore("IN_MEMORY", NewMemory)
}

// Compile time check to ensure Memory implements the Blobstore interface.
var _ Blobstore = (*Memory)(nil)

// Memory is a simple in-memory blobstore.
type Memory struct {
	lock sync.Mutex
	data map[string][]byte
}

// NewMemory creates a new Memory blobstore.
func NewMemory(_ context.Context, _ *Config) (Blobstore, error) {
	return &Memory{
		data: make(map[string][]byte),
	}, nil
}

// CreateObject creates a new object with the given name and contents. contentType is ignored.
func (s *Memory) CreateObject(_ context.Context, folder, filename string, contents []byte, cacheable bool, contentType string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	pth := path.Join(folder, filename)
	s.data[pth] = contents
	return nil
}

// DeleteObject deletes an object. It returns nil if the object was deleted or if the object no longer exists.
func (s *Memory) DeleteObject(_ context.Context, folder, filename string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	pth := path.Join(folder, filename)
	delete(s.data, pth)
	return nil
}

// GetObject returns the contents for the given object. If the object does not exist, it returns ErrNotFound.
func (s *Memory) GetObject(_ context.Context, folder, filename string) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	pth := path.Join(folder, filename)
	v, ok := s.data[pth]
	if !ok {
		return nil, ErrBlobNotFound
	}
	return v, nil
}

// GetPath returns the path of a local file that contains the object with the given name.
func (s *Memory) GetPath(_ context.Context, folder, filename string) (string, error) {
	filePath := path.Join(os.TempDir(), folder, filename)
	if err := os.MkdirAll(path.Dir(filePath), 0o700); err != nil {
		return "", err
	}
	if err := os.WriteFile(filePath, s.data[filename], 0o600); err != nil {
		return "", err
	}
	return filePath, nil
}
