// Code generated by counterfeiter. DO NOT EDIT.
package actorfakes

import (
	"sync"

	"code.cloudfoundry.org/slack-attachment-resource/actor"
	"github.com/nlopes/slack"
)

type FakeInAPIClient struct {
	GetFileInfoStub        func(fileID string, count int, page int) (*slack.File, []slack.Comment, *slack.Paging, error)
	getFileInfoMutex       sync.RWMutex
	getFileInfoArgsForCall []struct {
		fileID string
		count  int
		page   int
	}
	getFileInfoReturns struct {
		result1 *slack.File
		result2 []slack.Comment
		result3 *slack.Paging
		result4 error
	}
	getFileInfoReturnsOnCall map[int]struct {
		result1 *slack.File
		result2 []slack.Comment
		result3 *slack.Paging
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInAPIClient) GetFileInfo(fileID string, count int, page int) (*slack.File, []slack.Comment, *slack.Paging, error) {
	fake.getFileInfoMutex.Lock()
	ret, specificReturn := fake.getFileInfoReturnsOnCall[len(fake.getFileInfoArgsForCall)]
	fake.getFileInfoArgsForCall = append(fake.getFileInfoArgsForCall, struct {
		fileID string
		count  int
		page   int
	}{fileID, count, page})
	fake.recordInvocation("GetFileInfo", []interface{}{fileID, count, page})
	fake.getFileInfoMutex.Unlock()
	if fake.GetFileInfoStub != nil {
		return fake.GetFileInfoStub(fileID, count, page)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.getFileInfoReturns.result1, fake.getFileInfoReturns.result2, fake.getFileInfoReturns.result3, fake.getFileInfoReturns.result4
}

func (fake *FakeInAPIClient) GetFileInfoCallCount() int {
	fake.getFileInfoMutex.RLock()
	defer fake.getFileInfoMutex.RUnlock()
	return len(fake.getFileInfoArgsForCall)
}

func (fake *FakeInAPIClient) GetFileInfoArgsForCall(i int) (string, int, int) {
	fake.getFileInfoMutex.RLock()
	defer fake.getFileInfoMutex.RUnlock()
	return fake.getFileInfoArgsForCall[i].fileID, fake.getFileInfoArgsForCall[i].count, fake.getFileInfoArgsForCall[i].page
}

func (fake *FakeInAPIClient) GetFileInfoReturns(result1 *slack.File, result2 []slack.Comment, result3 *slack.Paging, result4 error) {
	fake.GetFileInfoStub = nil
	fake.getFileInfoReturns = struct {
		result1 *slack.File
		result2 []slack.Comment
		result3 *slack.Paging
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeInAPIClient) GetFileInfoReturnsOnCall(i int, result1 *slack.File, result2 []slack.Comment, result3 *slack.Paging, result4 error) {
	fake.GetFileInfoStub = nil
	if fake.getFileInfoReturnsOnCall == nil {
		fake.getFileInfoReturnsOnCall = make(map[int]struct {
			result1 *slack.File
			result2 []slack.Comment
			result3 *slack.Paging
			result4 error
		})
	}
	fake.getFileInfoReturnsOnCall[i] = struct {
		result1 *slack.File
		result2 []slack.Comment
		result3 *slack.Paging
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeInAPIClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getFileInfoMutex.RLock()
	defer fake.getFileInfoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInAPIClient) recordInvocation(key string, args []interface{}) {
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

var _ actor.InAPIClient = new(FakeInAPIClient)