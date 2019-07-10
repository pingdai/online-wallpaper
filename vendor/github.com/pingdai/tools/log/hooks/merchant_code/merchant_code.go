package merchant_code

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

var (
	mcs    = make(map[int64]string)
	locker = sync.RWMutex{}
)

func getGoID() int64 {
	return runtime.GoID()
}

func SetMerchantCode(mc string) {
	locker.Lock()
	defer locker.Unlock()

	mcs[getGoID()] = mc
}

func GetMerchantCode() string {
	locker.RLock()
	defer locker.RUnlock()

	goID := getGoID()

	if mc, ok := mcs[goID]; ok {
		return mc
	}

	return ""
}

type MerchantCodeHook struct {
}

func NewMerchantCodeHook() *MerchantCodeHook {
	return &MerchantCodeHook{}
}

func (hook *MerchantCodeHook) Fire(entry *logrus.Entry) error {
	entry.Data["merchant_code"] = GetMerchantCode()
	return nil
}

func (hook *MerchantCodeHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
