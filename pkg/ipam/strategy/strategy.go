/*
  Copyright 2021 The Rama Authors.

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

package strategy

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/klog"

	ramav1 "github.com/oecp/rama/pkg/client/listers/networking/v1"
	"github.com/oecp/rama/pkg/ipam/types"
	"github.com/oecp/rama/pkg/utils/transform"
)

var (
	statefulWorkloadKindVar  = "StatefulSet"
	statelessWorkloadKindVar = ""
	StatefulWorkloadKind     map[string]bool
	StatelessWorkloadKind    map[string]bool
)

func init() {
	pflag.StringVar(&statefulWorkloadKindVar, "stateful-workload-kinds", statefulWorkloadKindVar, `stateful workload kinds to use strategic IP allocation,`+
		`eg: "StatefulSet,AdvancedStatefulSet", default: "StatefulSet"`)
	pflag.StringVar(&statelessWorkloadKindVar, "stateless-workload-kinds", statelessWorkloadKindVar, "stateless workload kinds to use strategic IP allocation,"+
		`eg: "ReplicaSet", default: ""`)

	StatefulWorkloadKind = make(map[string]bool)
	StatelessWorkloadKind = make(map[string]bool)

	for _, kind := range strings.Split(statefulWorkloadKindVar, ",") {
		if len(kind) > 0 {
			StatefulWorkloadKind[kind] = true
			klog.Infof("[strategy] Adding kind %s to known stateful workloads", kind)
		}
	}

	for _, kind := range strings.Split(statelessWorkloadKindVar, ",") {
		if len(kind) > 0 {
			StatelessWorkloadKind[kind] = true
			klog.Infof("[strategy] Adding kind %s to known stateless workloads", kind)
		}
	}
}

func OwnByStatefulWorkload(pod *v1.Pod) bool {
	ref := metav1.GetControllerOf(pod)
	if ref == nil {
		return false
	}

	return StatefulWorkloadKind[ref.Kind]
}

func OwnByStatelessWorkload(pod *v1.Pod) bool {
	ref := metav1.GetControllerOf(pod)
	if ref == nil {
		return false
	}

	return StatelessWorkloadKind[ref.Kind]
}

func GetKnownOwnReference(pod *v1.Pod) *metav1.OwnerReference {
	// only support stateful workloads
	if OwnByStatefulWorkload(pod) {
		return metav1.GetControllerOf(pod)
	}
	return nil
}

func GetIPbyPod(ipLister ramav1.IPInstanceLister, pod *v1.Pod) (string, error) {
	ips, err := ipLister.IPInstances(pod.Namespace).List(labels.Everything())
	if err != nil {
		return "", err
	}

	for _, ip := range ips {
		if ip.Status.PodName == pod.Name {
			ipStr, _ := toIPFormat(ip.Name)
			return ipStr, nil
		}
	}

	return "", nil
}

func GetIPsbyPod(ipLister ramav1.IPInstanceLister, pod *v1.Pod) ([]string, error) {
	ips, err := ipLister.IPInstances(pod.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var v4, v6 []string
	for _, ip := range ips {
		if ip.Status.PodName == pod.Name {
			ipStr, isIPv6 := toIPFormat(ip.Name)
			if isIPv6 {
				v6 = append(v6, ipStr)
			} else {
				v4 = append(v4, ipStr)
			}
		}
	}

	return append(v4, v6...), nil
}

func GetAllocatedIPsByPod(ipLister ramav1.IPInstanceLister, pod *v1.Pod) ([]*types.IP, error) {
	ips, err := ipLister.IPInstances(pod.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var allocatedIPs []*types.IP
	var networkName string
	for _, ip := range ips {
		if ip.Status.PodName == pod.Name {
			allocatedIPs = append(allocatedIPs, transform.TransferIPInstanceForIPAM(ip))
			if len(networkName) > 0 && networkName != ip.Spec.Network {
				return nil, fmt.Errorf("pod %s has allocated IPs from different networks", pod.Name)
			}
			networkName = ip.Spec.Network
		}
	}

	return allocatedIPs, nil
}

func GetIndexFromName(name string) int {
	nameSlice := strings.Split(name, "-")
	indexStr := nameSlice[len(nameSlice)-1]

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return math.MaxInt32
	}
	return index
}

func toIPFormat(name string) (string, bool) {
	const IPv6SeparatorCount = 7
	if isIPv6 := strings.Count(name, "-") == IPv6SeparatorCount; isIPv6 {
		return net.ParseIP(strings.ReplaceAll(name, "-", ":")).String(), true
	}
	return strings.ReplaceAll(name, "-", "."), false
}
