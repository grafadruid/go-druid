package dimension

import (
	"testing"

	"github.com/grafadruid/go-druid/builder"
	"github.com/grafadruid/go-druid/builder/extractionfn"

	"github.com/stretchr/testify/assert"
)

func TestLoadUnsupportedType(t *testing.T) {
	assert := assert.New(t)

	f, err := Load([]byte("{\"type\": \"blahblahType\"}"))

	assert.Nil(f, "filter should be nil")
	assert.NotNil(err, "error should not be nil")
	assert.Error(err, "unsupported dimension type")
}

func TestLoad(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    func() builder.Dimension
		wantErr bool
	}{
		{
			name: "load extraction dimension success",
			args: args{data: []byte(`{
			  "type": "extraction",
			  "dimension": "lookupKey",
			  "outputName": "assetTier",
			  "extractionFn": {
				"type": "cascade",
				"extractionFns": [
				  {
					"type": "registeredLookup",
					"lookup": "asset_id_to_metadata",
					"retainMissingValue": true
				  },
				  {
					"type": "regex",
					"expr": "(?<=tier\":\")(.*?)(?=\")",
					"replaceMissingValue": false
				  }
				]
			  }
			}`)},
			want: func() builder.Dimension {
				eFunc1 := extractionfn.NewRegisteredLookup()
				eFunc1.Lookup = "asset_id_to_metadata"
				eFunc1RetainMissingValue := true
				eFunc1.RetainMissingValue = &eFunc1RetainMissingValue

				eFunc2 := extractionfn.NewRegex()
				eFunc2.Expr = "(?<=tier\":\")(.*?)(?=\")"
				eFunc2ReplaceMissingValue := false
				eFunc2.ReplaceMissingValue = &eFunc2ReplaceMissingValue

				eFunc := extractionfn.NewCascade()
				eFunc.SetExtractionFns([]builder.ExtractionFn{
					eFunc1,
					eFunc2,
				})

				d := NewExtraction()
				d.OutputName = "assetTier"
				d.Dimension = "lookupKey"
				d.ExtractionFn = eFunc

				return d
			},
			wantErr: false,
		},
		{
			name: "load extraction dimension empty extraction func",
			args: args{data: []byte(`{
			  "type": "extraction",
			  "dimension": "lookupKey",
			  "outputName": "assetTier",
			  "extractionFn": null
			}`)},
			want: func() builder.Dimension {
				d := NewExtraction()
				d.OutputName = "assetTier"
				d.Dimension = "lookupKey"
				d.ExtractionFn = nil

				return d
			},
			wantErr: false,
		},
		{
			name: "load extraction dimension missing extraction func",
			args: args{data: []byte(`{
			  "type": "extraction",
			  "dimension": "lookupKey",
			  "outputName": "assetTier"
			}`)},
			want: func() builder.Dimension {
				d := NewExtraction()
				d.OutputName = "assetTier"
				d.Dimension = "lookupKey"

				return d
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.NotNil(t, got)
			assert.Equalf(t, tt.want(), got, "Load(%v)", tt.args.data)
		})
	}
}
