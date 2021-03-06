/*
 * Licensed to the OpenSkywalking under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenSkywalking licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T)  {
	assert := assert.New(t)
	assert.Equal(30000,Conf.PendingSegmentLimit)
	assert.Equal(3000,Conf.Interval)
	assert.Equal(1000,Conf.PendingSegmentTimeout)
	assert.Equal(-1,Conf.SamplePer3Secs)
	assert.Equal(300,Conf.SpanLimitPerSegment)
}
