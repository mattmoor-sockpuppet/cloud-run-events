/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reconciler

import (
	"github.com/knative/pkg/controller"
	"go.uber.org/zap"
)

// MustNewStatsReporter creates a new instance of StatsReporter. Panics if creation fails.
func MustNewStatsReporter(reconciler string, logger *zap.SugaredLogger) controller.StatsReporter {
	stats, err := controller.NewStatsReporter(reconciler)
	if err != nil {
		logger.Fatalw("Failed to initialize the stats reporter", zap.Error(err))
	}
	return stats
}
