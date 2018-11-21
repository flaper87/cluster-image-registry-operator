package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	cVOOverridesFileName = "cvo-overrides.yaml.template"
)

var _ asset.WritableAsset = (*CVOOverrides)(nil)

// CVOOverrides is the constant to represent contents of cvo-override.yaml.template file
// This is a gate to prevent CVO from installing these operators which is conflicting
// with already owned resources by tectonic-operators.
// This files can be dropped when the overrides list becomes empty.
type CVOOverrides struct {
	fileName string
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *CVOOverrides) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *CVOOverrides) Name() string {
	return "CVOOverrides"
}

// Generate generates the actual files by this asset
func (t *CVOOverrides) Generate(parents asset.Parents) error {
	t.fileName = cVOOverridesFileName
	data, err := content.GetBootkubeTemplate(t.fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, t.fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *CVOOverrides) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *CVOOverrides) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, cVOOverridesFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
