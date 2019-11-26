// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output_test

import (
	"testing"

	"github.com/banzaicloud/logging-operator/pkg/sdk/model/output"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/render"
	"github.com/ghodss/yaml"
)

func TestSumologic(t *testing.T) {
	CONFIG := []byte(`
data_type: metrics
metric_data_format: carbon2
log_format: json
source_category: prod/someapp/logs
source_name: AppA
`)
	expected := `
  <match **>
    @type sumologic
    @id test_sumologic
    data_type metrics
    log_format json
    metric_data_format carbon2
    source_category prod/someapp/logs
    source_name AppA
  </match>
`
	s := &output.SumologicOutput{}
	yaml.Unmarshal(CONFIG, s)
	test := render.NewOutputPluginTest(t, s)
	test.DiffResult(expected)
}