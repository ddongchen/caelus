/*
 * Copyright (c) 2021 Tencent.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 *
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cgroupstore

const (
	KubernetesPodNameLabel       = "io.kubernetes.pod.name"
	KubernetesPodNamespaceLabel  = "io.kubernetes.pod.namespace"
	KubernetesPodUIDLabel        = "io.kubernetes.pod.uid"
	KubernetesContainerNameLabel = "io.kubernetes.container.name"

	PodInfraContainerName = "POD"
)

// GetContainerName get pod container name from docker container labels
func GetContainerName(labels map[string]string) string {
	return labels[KubernetesContainerNameLabel]
}

// GetPodName get pod name from docker container labels
func GetPodName(labels map[string]string) string {
	return labels[KubernetesPodNameLabel]
}

// GetPodUID get pod uid from docker container labels
func GetPodUID(labels map[string]string) string {
	return labels[KubernetesPodUIDLabel]
}

// GetPodNamespace get pod namespace from docker container labels
func GetPodNamespace(labels map[string]string) string {
	return labels[KubernetesPodNamespaceLabel]
}
