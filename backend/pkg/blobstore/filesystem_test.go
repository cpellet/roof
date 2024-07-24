/*
 * Copyright (c) 2024. SPK Platforms - All rights reserved
 * Proprietary and confidential
 * Written by Cyrus Pellet <cyrus@spkplatforms.co> in July 2024
 */

package blob

import (
	"bytes"
	"os"
	"path/filepath"
	"spot/pkg/utils"
	"testing"
)

func TestFilesystemStorage_CreateObject(t *testing.T) {
	t.Parallel()

	tmp, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(tmp) })

	cases := []struct {
		name     string
		folder   string
		filepath string
		contents []byte
		err      bool
	}{
		{
			name:     "default",
			folder:   "",
			filepath: "myfile",
			contents: []byte("contents"),
			err:      false,
		},
		{
			name:     "bad_path",
			folder:   "/path/that/definitely/doesnt/exist",
			filepath: "myfile",
			contents: []byte("contents"),
			err:      true,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := utils.TestContext(t)

			storage, err := NewFilesystem(ctx, &Config{
				RootDir: tmp,
			})
			if err != nil {
				t.Fatal(err)
			}

			err = storage.CreateObject(ctx, tc.folder, tc.filepath, tc.contents, false, ContentTypePDF)
			if (err != nil) != tc.err {
				t.Fatal(err)
			}

			if !tc.err {
				contents, err := os.ReadFile(filepath.Join(tmp, tc.folder, tc.filepath))
				if err != nil {
					t.Fatal(err)
				}

				if !bytes.Equal(contents, tc.contents) {
					t.Errorf("expected %q to be %q ", contents, tc.contents)
				}
			}
		})
	}
}

func TestFilesystemStorage_DeleteObject(t *testing.T) {
	t.Parallel()

	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name     string
		folder   string
		filepath string
	}{
		{
			name:     "default",
			folder:   "",
			filepath: filepath.Base(f.Name()),
		},
		{
			name:     "not_exist",
			folder:   "",
			filepath: "not-exist",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := utils.TestContext(t)

			storage, err := NewFilesystem(ctx, &Config{
				RootDir: filepath.Dir(f.Name()),
			})
			if err != nil {
				t.Fatal(err)
			}

			if err = storage.DeleteObject(ctx, tc.folder, tc.filepath); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestFilesystemStorage_GetObject(t *testing.T) {
	t.Parallel()

	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write([]byte("hello")); err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name     string
		folder   string
		filepath string
		contents []byte
		err      bool
	}{
		{
			name:     "default",
			folder:   "",
			filepath: filepath.Base(f.Name()),
			contents: []byte("hello"),
		},
		{
			name:     "not_exist",
			folder:   "",
			filepath: "not-exist",
			err:      true,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := utils.TestContext(t)

			storage, err := NewFilesystem(ctx, &Config{
				RootDir: filepath.Dir(f.Name()),
			})
			if err != nil {
				t.Fatal(err)
			}

			b, err := storage.GetObject(ctx, tc.folder, tc.filepath)
			if (err != nil) != tc.err {
				t.Fatal(err)
			}

			if got, want := b, tc.contents; !bytes.Equal(got, want) {
				t.Errorf("expected %v to be %v", got, want)
			}
		})
	}
}

func TestFilesystemStorage_GetPath(t *testing.T) {
	t.Parallel()

	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write([]byte("hello")); err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name     string
		folder   string
		filepath string
		contents []byte
		err      bool
	}{
		{
			name:     "default",
			folder:   "",
			filepath: filepath.Base(f.Name()),
			contents: []byte("hello"),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := utils.TestContext(t)

			storage, err := NewFilesystem(ctx, &Config{
				RootDir: filepath.Dir(f.Name()),
			})
			if err != nil {
				t.Fatal(err)
			}

			path, err := storage.GetPath(ctx, tc.folder, tc.filepath)
			if (err != nil) != tc.err {
				t.Fatal(err)
			}

			if !tc.err {
				contents, err := os.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}

				if got, want := contents, tc.contents; !bytes.Equal(got, want) {
					t.Errorf("expected %v to be %v", got, want)
				}
			}
		})
	}
}
