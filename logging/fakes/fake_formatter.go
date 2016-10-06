// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/SAP/aker/logging"
)

type FakeFormatter struct {
	FormatStub        func(*logging.AccessEntry) string
	formatMutex       sync.RWMutex
	formatArgsForCall []struct {
		arg1 *logging.AccessEntry
	}
	formatReturns struct {
		result1 string
	}
}

func (fake *FakeFormatter) Format(arg1 *logging.AccessEntry) string {
	fake.formatMutex.Lock()
	fake.formatArgsForCall = append(fake.formatArgsForCall, struct {
		arg1 *logging.AccessEntry
	}{arg1})
	fake.formatMutex.Unlock()
	if fake.FormatStub != nil {
		return fake.FormatStub(arg1)
	} else {
		return fake.formatReturns.result1
	}
}

func (fake *FakeFormatter) FormatCallCount() int {
	fake.formatMutex.RLock()
	defer fake.formatMutex.RUnlock()
	return len(fake.formatArgsForCall)
}

func (fake *FakeFormatter) FormatArgsForCall(i int) *logging.AccessEntry {
	fake.formatMutex.RLock()
	defer fake.formatMutex.RUnlock()
	return fake.formatArgsForCall[i].arg1
}

func (fake *FakeFormatter) FormatReturns(result1 string) {
	fake.FormatStub = nil
	fake.formatReturns = struct {
		result1 string
	}{result1}
}

var _ logging.Formatter = new(FakeFormatter)
