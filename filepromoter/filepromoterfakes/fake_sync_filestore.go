/*
Copyright The Kubernetes Authors.

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

// Code generated by counterfeiter. DO NOT EDIT.
package filepromoterfakes

import (
	"context"
	"io"
	"sync"

	"sigs.k8s.io/k8s-container-image-promoter/filepromoter"
)

type FakeSyncFilestore struct {
	ListFilesStub        func(context.Context) (map[string]*filepromoter.SyncFileInfo, error)
	listFilesMutex       sync.RWMutex
	listFilesArgsForCall []struct {
		arg1 context.Context
	}
	listFilesReturns struct {
		result1 map[string]*filepromoter.SyncFileInfo
		result2 error
	}
	listFilesReturnsOnCall map[int]struct {
		result1 map[string]*filepromoter.SyncFileInfo
		result2 error
	}
	OpenReaderStub        func(context.Context, string) (io.ReadCloser, error)
	openReaderMutex       sync.RWMutex
	openReaderArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	openReaderReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	openReaderReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	UploadFileStub        func(context.Context, string, string) error
	uploadFileMutex       sync.RWMutex
	uploadFileArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	uploadFileReturns struct {
		result1 error
	}
	uploadFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSyncFilestore) ListFiles(arg1 context.Context) (map[string]*filepromoter.SyncFileInfo, error) {
	fake.listFilesMutex.Lock()
	ret, specificReturn := fake.listFilesReturnsOnCall[len(fake.listFilesArgsForCall)]
	fake.listFilesArgsForCall = append(fake.listFilesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListFilesStub
	fakeReturns := fake.listFilesReturns
	fake.recordInvocation("ListFiles", []interface{}{arg1})
	fake.listFilesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSyncFilestore) ListFilesCallCount() int {
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	return len(fake.listFilesArgsForCall)
}

func (fake *FakeSyncFilestore) ListFilesCalls(stub func(context.Context) (map[string]*filepromoter.SyncFileInfo, error)) {
	fake.listFilesMutex.Lock()
	defer fake.listFilesMutex.Unlock()
	fake.ListFilesStub = stub
}

func (fake *FakeSyncFilestore) ListFilesArgsForCall(i int) context.Context {
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	argsForCall := fake.listFilesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSyncFilestore) ListFilesReturns(result1 map[string]*filepromoter.SyncFileInfo, result2 error) {
	fake.listFilesMutex.Lock()
	defer fake.listFilesMutex.Unlock()
	fake.ListFilesStub = nil
	fake.listFilesReturns = struct {
		result1 map[string]*filepromoter.SyncFileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeSyncFilestore) ListFilesReturnsOnCall(i int, result1 map[string]*filepromoter.SyncFileInfo, result2 error) {
	fake.listFilesMutex.Lock()
	defer fake.listFilesMutex.Unlock()
	fake.ListFilesStub = nil
	if fake.listFilesReturnsOnCall == nil {
		fake.listFilesReturnsOnCall = make(map[int]struct {
			result1 map[string]*filepromoter.SyncFileInfo
			result2 error
		})
	}
	fake.listFilesReturnsOnCall[i] = struct {
		result1 map[string]*filepromoter.SyncFileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeSyncFilestore) OpenReader(arg1 context.Context, arg2 string) (io.ReadCloser, error) {
	fake.openReaderMutex.Lock()
	ret, specificReturn := fake.openReaderReturnsOnCall[len(fake.openReaderArgsForCall)]
	fake.openReaderArgsForCall = append(fake.openReaderArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.OpenReaderStub
	fakeReturns := fake.openReaderReturns
	fake.recordInvocation("OpenReader", []interface{}{arg1, arg2})
	fake.openReaderMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSyncFilestore) OpenReaderCallCount() int {
	fake.openReaderMutex.RLock()
	defer fake.openReaderMutex.RUnlock()
	return len(fake.openReaderArgsForCall)
}

func (fake *FakeSyncFilestore) OpenReaderCalls(stub func(context.Context, string) (io.ReadCloser, error)) {
	fake.openReaderMutex.Lock()
	defer fake.openReaderMutex.Unlock()
	fake.OpenReaderStub = stub
}

func (fake *FakeSyncFilestore) OpenReaderArgsForCall(i int) (context.Context, string) {
	fake.openReaderMutex.RLock()
	defer fake.openReaderMutex.RUnlock()
	argsForCall := fake.openReaderArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSyncFilestore) OpenReaderReturns(result1 io.ReadCloser, result2 error) {
	fake.openReaderMutex.Lock()
	defer fake.openReaderMutex.Unlock()
	fake.OpenReaderStub = nil
	fake.openReaderReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeSyncFilestore) OpenReaderReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.openReaderMutex.Lock()
	defer fake.openReaderMutex.Unlock()
	fake.OpenReaderStub = nil
	if fake.openReaderReturnsOnCall == nil {
		fake.openReaderReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.openReaderReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeSyncFilestore) UploadFile(arg1 context.Context, arg2 string, arg3 string) error {
	fake.uploadFileMutex.Lock()
	ret, specificReturn := fake.uploadFileReturnsOnCall[len(fake.uploadFileArgsForCall)]
	fake.uploadFileArgsForCall = append(fake.uploadFileArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UploadFileStub
	fakeReturns := fake.uploadFileReturns
	fake.recordInvocation("UploadFile", []interface{}{arg1, arg2, arg3})
	fake.uploadFileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSyncFilestore) UploadFileCallCount() int {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return len(fake.uploadFileArgsForCall)
}

func (fake *FakeSyncFilestore) UploadFileCalls(stub func(context.Context, string, string) error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = stub
}

func (fake *FakeSyncFilestore) UploadFileArgsForCall(i int) (context.Context, string, string) {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	argsForCall := fake.uploadFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSyncFilestore) UploadFileReturns(result1 error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = nil
	fake.uploadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSyncFilestore) UploadFileReturnsOnCall(i int, result1 error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = nil
	if fake.uploadFileReturnsOnCall == nil {
		fake.uploadFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSyncFilestore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	fake.openReaderMutex.RLock()
	defer fake.openReaderMutex.RUnlock()
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSyncFilestore) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}