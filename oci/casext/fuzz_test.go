// SPDX-License-Identifier: Apache-2.0
package casext

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/opencontainers/umoci/oci/cas/dir"
)

func FuzzCasext(f *testing.F) {
	f.Add([]byte(`{"mediaType":"application/vnd.oci.image.manifest.v1+json"}`))
	f.Fuzz(func(t *testing.T, data []byte) {
		ctx := context.Background()
		root, err := os.MkdirTemp("", "umoci-FuzzCasext")
		if err != nil {
			t.Skip()
		}
		defer os.RemoveAll(root)

		image := filepath.Join(root, "image")
		if err := dir.Create(image); err != nil {
			t.Skip()
		}

		engine, err := dir.Open(image)
		if err != nil {
			t.Skip()
		}
		engineExt := NewEngine(engine)
		defer engine.Close()

		digest, _, err := engineExt.PutBlobJSON(ctx, string(data))
		if err != nil {
			return
		}
		blobReader, err := engine.GetBlob(ctx, digest)
		if err != nil {
			return
		}
		defer blobReader.Close()
		_, _ = io.ReadAll(blobReader)
	})
}
