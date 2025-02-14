/*
Copyright 2022 The OpenYurt Authors.

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
package kolecontroller

import (
	"github.com/eclipse/paho.golang/paho"
	"github.com/openyurtio/kole/pkg/message"
	"k8s.io/klog/v2"

	"github.com/openyurtio/kole/pkg/data"
	"github.com/openyurtio/kole/pkg/util"
)

func (c *KoleController) Mqtt5CreateSubscribes() []*message.SingleSubcribe {

	subs := make([]*message.SingleSubcribe, 0, 2)
	//HB
	subs = append(subs, &message.SingleSubcribe{
		Topic: util.TopicHeartBeat,
		Option: paho.SubscribeOptions{
			QoS: 1,
		},
		Handler: func(publish *paho.Publish) {
			hb, err := data.UnmarshalPayloadToHeartBeat(publish.Payload)
			if err != nil {
				klog.Errorf("UnmarshalPayloadToHeartBeat error %v", err)
				return
			}
			c.ConsumeHeartBeatDirect(hb)
			klog.V(5).Infof("sub heatbeat topic %s Name %s State %s", publish.Topic, hb.Name, hb.State)
		},
	})

	return subs
}
