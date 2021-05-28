/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2021 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 *
 */

package internal

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	clusterclient "tkestack.io/tke/pkg/mesh/external/kubernetes"
)

type CommonService struct {
	clients clusterclient.Client
}

func New(clients clusterclient.Client) *CommonService {
	return &CommonService{
		clients: clients,
	}
}

// ListNamespaces list namespace with/without labels selector
func (m *CommonService) ListNamespaces(ctx context.Context, clusterName string, labelSelector labels.Selector) (
	[]corev1.Namespace, error) {

	clusterClient, err := m.clients.Cluster(clusterName)
	if err != nil {
		return nil, err
	}

	ns := &corev1.NamespaceList{}
	err = clusterClient.List(
		ctx, ns, &ctrlclient.ListOptions{
			LabelSelector: labelSelector,
		},
	)
	if err != nil {
		return nil, err
	}

	return ns.Items, nil
}
