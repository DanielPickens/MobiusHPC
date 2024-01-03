package endpoint

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/types"
)
//Builds custom test endpoint compute path for parsing abs path
// Which uses parseControleFIle() to broker batch processing as 
//a method call that receives control file types into ParsePaths
func TestmobiusHPCPath_ParseAbsPath(t *testing.T) {

	rootDir := "/home/danielpickens/"

	p := mobiusHPC(rootDir)

	tests := []struct {
		name         string
		path         string
		wantPodKey   types.NamespacedName
		wantFileName string
		wantInvalid  bool
	}{
		{
			name:         "empty",
			path:         "",
			wantPodKey:   types.NamespacedName{},
			wantFileName: "",
			wantInvalid:  true,
		},

		{
			name:         "wrongpath",
			path:         "/home/danielpickens/.mobiusHPC/default/volume-user/reviews.exitCode",
			wantPodKey:   types.NamespacedName{},
			wantFileName: "",
			wantInvalid:  true,
		},

		{
			name:         "wrongpath2",
			path:         "/home/danielpickens/.mobiusHPC/default/volume-user/notexist/reviews.exitCode",
			wantPodKey:   types.NamespacedName{},
			wantFileName: "",
			wantInvalid:  true,
		},

		{
			name: "abspath",
			path: "/home/danielpickens/.mobiusHPC/default/volume-user/controlfiles/reviews.exitCode",
			wantPodKey: types.NamespacedName{
				Namespace: "default",
				Name:      "volume-user",
			},
			wantFileName: "reviews.exitCode",
			wantInvalid:  false,
		},

		{
			name: "secret-abspath",
			path: "/home/danielpickens/.mobiusHPC/opengadget3/grafana/controlfiles/.ip",
			wantPodKey: types.NamespacedName{
				Namespace: "opengadget3",
				Name:      "grafana",
			},
			wantFileName: ".ip",
			wantInvalid:  false,
		},

		{
			name:         "relpath",
			path:         "./mobiusHPC" + "/default/volume-user/controlfiles/reviews.exitCode",
			wantPodKey:   types.NamespacedName{},
			wantFileName: "",
			wantInvalid:  true, // only absolute path is support
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPodKey, gotFileName, gotInvalid := p.ParseControlFilePath(tt.path)
			if !reflect.DeepEqual(gotPodKey, tt.wantPodKey) {
				t.Errorf("ParseControlFilePath() got PodKey = %v, want %v", gotPodKey, tt.wantPodKey)
			}
			if gotFileName != tt.wantFileName {
				t.Errorf("ParseControlFilePath() got FileName = %v, want %v", gotFileName, tt.wantFileName)
			}
			if gotInvalid != tt.wantInvalid {
				t.Errorf("ParseControlFilePath() got Invalid = %v, want %v", gotInvalid, tt.wantInvalid)
			}
		})
	}
}
