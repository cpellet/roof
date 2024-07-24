/*
 * Copyright (c) 2024. SPK Platforms - All rights reserved
 * Proprietary and confidential
 * Written by Cyrus Pellet <cyrus@spkplatforms.co> in July 2024
 */

package blob

// Config defines configuration for the blobstore.
type Config struct {
	// Type is the type of blobstore to use.
	Type    string `env:"BLOBSTORE, default=IN_MEMORY"`
	RootDir string `env:"BLOBSTORE_ROOT_DIR" default:"/tmp/blobstore"`
}
