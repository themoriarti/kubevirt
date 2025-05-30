/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright The KubeVirt Authors.
 *
 */

package vm_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"

	"kubevirt.io/kubevirt/pkg/virtctl/testing"
)

var _ = Describe("Guest os info command", func() {
	var vmiInterface *kubecli.MockVirtualMachineInstanceInterface
	var ctrl *gomock.Controller
	const vmName = "testvm"

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		kubecli.GetKubevirtClientFromClientConfig = kubecli.GetMockKubevirtClientFromClientConfig
		kubecli.MockKubevirtClientInstance = kubecli.NewMockKubevirtClient(ctrl)
		vmiInterface = kubecli.NewMockVirtualMachineInstanceInterface(ctrl)
	})

	It("should fail with missing input parameters", func() {
		cmd := testing.NewRepeatableVirtctlCommand("guestosinfo")
		err := cmd()
		Expect(err).To(HaveOccurred())
		Expect(err).Should(MatchError("accepts 1 arg(s), received 0"))
	})

	It("should fail with non existing vm", func() {
		kubecli.MockKubevirtClientInstance.
			EXPECT().
			VirtualMachineInstance(k8smetav1.NamespaceDefault).
			Return(vmiInterface).
			Times(1)

		vmiInterface.EXPECT().GuestOsInfo(context.Background(), vmName).Return(v1.VirtualMachineInstanceGuestAgentInfo{}, fmt.Errorf("an error on the server (\"virtualmachineinstance.kubevirt.io \"testvm\" not found\") has prevented the request from succeeding")).Times(1)

		cmd := testing.NewRepeatableVirtctlCommand("guestosinfo", vmName)
		err := cmd()
		Expect(err).To(HaveOccurred())
		Expect(err).Should(MatchError("Error getting guestosinfo of VirtualMachineInstance testvm, an error on the server (\"virtualmachineinstance.kubevirt.io \"testvm\" not found\") has prevented the request from succeeding"))
	})

	It("should fail when vm has no guest agent data", func() {
		vm := kubecli.NewMinimalVM(vmName)

		kubecli.MockKubevirtClientInstance.
			EXPECT().
			VirtualMachineInstance(k8smetav1.NamespaceDefault).
			Return(vmiInterface).
			Times(1)

		vmiInterface.EXPECT().GuestOsInfo(context.Background(), vm.Name).Return(v1.VirtualMachineInstanceGuestAgentInfo{}, fmt.Errorf("an error on the server (\"Operation cannot be fulfilled on virtualmachineinstance.kubevirt.io \"testvm\": VMI does not have guest agent connected\") has prevented the request from succeeding")).Times(1)

		cmd := testing.NewRepeatableVirtctlCommand("guestosinfo", vm.Name)
		err := cmd()
		Expect(err).To(HaveOccurred())
		Expect(err).Should(MatchError("Error getting guestosinfo of VirtualMachineInstance testvm, an error on the server (\"Operation cannot be fulfilled on virtualmachineinstance.kubevirt.io \"testvm\": VMI does not have guest agent connected\") has prevented the request from succeeding"))
	})

	It("should return guest agent data", func() {
		vm := kubecli.NewMinimalVM(vmName)
		guestOSInfo := v1.VirtualMachineInstanceGuestAgentInfo{
			GAVersion: "3.1.0",
		}

		kubecli.MockKubevirtClientInstance.
			EXPECT().
			VirtualMachineInstance(k8smetav1.NamespaceDefault).
			Return(vmiInterface).
			Times(1)

		vmiInterface.EXPECT().GuestOsInfo(context.Background(), vm.Name).Return(guestOSInfo, nil).Times(1)

		cmd := testing.NewRepeatableVirtctlCommand("guestosinfo", vm.Name)
		Expect(cmd()).To(Succeed())
	})
})
